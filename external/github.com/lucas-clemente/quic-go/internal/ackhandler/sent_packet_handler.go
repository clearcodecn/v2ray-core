package ackhandler

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/congestion"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/qerr"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/utils"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/wire"
)

const (
	// Maximum reordering in time space before time based loss detection considers a packet lost.
	// In fraction of an RTT.
	timeReorderingFraction = 1.0 / 8
	// Timer granularity. The timer will not be set to a value smaller than granularity.
	granularity = time.Millisecond
)

type sentPacketHandler struct {
	lastSentPacketNumber  protocol.PacketNumber
	packetNumberGenerator *packetNumberGenerator

	lastSentRetransmittablePacketTime time.Time
	lastSentCryptoPacketTime          time.Time

	nextPacketSendTime time.Time

	largestAcked                 protocol.PacketNumber
	largestReceivedPacketWithAck protocol.PacketNumber
	// lowestPacketNotConfirmedAcked is the lowest packet number that we sent an ACK for, but haven't received confirmation, that this ACK actually arrived
	// example: we send an ACK for packets 90-100 with packet number 20
	// once we receive an ACK from the peer for packet 20, the lowestPacketNotConfirmedAcked is 101
	lowestPacketNotConfirmedAcked protocol.PacketNumber

	packetHistory *sentPacketHistory

	retransmissionQueue []*Packet

	bytesInFlight protocol.ByteCount

	congestion congestion.SendAlgorithm
	rttStats   *congestion.RTTStats

	handshakeComplete bool

	// The number of times the crypto packets have been retransmitted without receiving an ack.
	cryptoCount uint32
	// The number of times a PTO has been sent without receiving an ack.
	ptoCount uint32
	// The number of PTO probe packets that should be sent.
	numProbesToSend int

	// The time at which the next packet will be considered lost based on early transmit or exceeding the reordering window in time.
	lossTime time.Time

	// The alarm timeout
	alarm time.Time

	logger utils.Logger
}

// NewSentPacketHandler creates a new sentPacketHandler
func NewSentPacketHandler(
	initialPacketNumber protocol.PacketNumber,
	rttStats *congestion.RTTStats,
	logger utils.Logger,
) SentPacketHandler {
	congestion := congestion.NewCubicSender(
		congestion.DefaultClock{},
		rttStats,
		false, /* don't use reno since chromium doesn't (why?) */
		protocol.InitialCongestionWindow,
		protocol.DefaultMaxCongestionWindow,
	)

	return &sentPacketHandler{
		packetNumberGenerator: newPacketNumberGenerator(initialPacketNumber, protocol.SkipPacketAveragePeriodLength),
		packetHistory:         newSentPacketHistory(),
		rttStats:              rttStats,
		congestion:            congestion,
		logger:                logger,
	}
}

func (h *sentPacketHandler) lowestUnacked() protocol.PacketNumber {
	if p := h.packetHistory.FirstOutstanding(); p != nil {
		return p.PacketNumber
	}
	return h.largestAcked + 1
}

func (h *sentPacketHandler) SetHandshakeComplete() {
	h.logger.Debugf("Handshake complete. Discarding all outstanding crypto packets.")
	var queue []*Packet
	for _, packet := range h.retransmissionQueue {
		if packet.EncryptionLevel == protocol.Encryption1RTT {
			queue = append(queue, packet)
		}
	}
	var cryptoPackets []*Packet
	h.packetHistory.Iterate(func(p *Packet) (bool, error) {
		if p.EncryptionLevel != protocol.Encryption1RTT {
			cryptoPackets = append(cryptoPackets, p)
		}
		return true, nil
	})
	for _, p := range cryptoPackets {
		h.packetHistory.Remove(p.PacketNumber)
	}
	h.retransmissionQueue = queue
	h.handshakeComplete = true
}

func (h *sentPacketHandler) SentPacket(packet *Packet) {
	if isRetransmittable := h.sentPacketImpl(packet); isRetransmittable {
		h.packetHistory.SentPacket(packet)
		h.updateLossDetectionAlarm()
	}
}

func (h *sentPacketHandler) SentPacketsAsRetransmission(packets []*Packet, retransmissionOf protocol.PacketNumber) {
	var p []*Packet
	for _, packet := range packets {
		if isRetransmittable := h.sentPacketImpl(packet); isRetransmittable {
			p = append(p, packet)
		}
	}
	h.packetHistory.SentPacketsAsRetransmission(p, retransmissionOf)
	h.updateLossDetectionAlarm()
}

func (h *sentPacketHandler) sentPacketImpl(packet *Packet) bool /* isRetransmittable */ {
	if h.logger.Debug() && h.lastSentPacketNumber != 0 {
		for p := h.lastSentPacketNumber + 1; p < packet.PacketNumber; p++ {
			h.logger.Debugf("Skipping packet number %#x", p)
		}
	}

	h.lastSentPacketNumber = packet.PacketNumber

	if len(packet.Frames) > 0 {
		if ackFrame, ok := packet.Frames[0].(*wire.AckFrame); ok {
			packet.largestAcked = ackFrame.LargestAcked()
		}
	}

	packet.Frames = stripNonRetransmittableFrames(packet.Frames)
	isRetransmittable := len(packet.Frames) != 0

	if isRetransmittable {
		if packet.EncryptionLevel != protocol.Encryption1RTT {
			h.lastSentCryptoPacketTime = packet.SendTime
		}
		h.lastSentRetransmittablePacketTime = packet.SendTime
		packet.includedInBytesInFlight = true
		h.bytesInFlight += packet.Length
		packet.canBeRetransmitted = true
		if h.numProbesToSend > 0 {
			h.numProbesToSend--
		}
	}
	h.congestion.OnPacketSent(packet.SendTime, h.bytesInFlight, packet.PacketNumber, packet.Length, isRetransmittable)

	h.nextPacketSendTime = utils.MaxTime(h.nextPacketSendTime, packet.SendTime).Add(h.congestion.TimeUntilSend(h.bytesInFlight))
	return isRetransmittable
}

func (h *sentPacketHandler) ReceivedAck(ackFrame *wire.AckFrame, withPacketNumber protocol.PacketNumber, encLevel protocol.EncryptionLevel, rcvTime time.Time) error {
	largestAcked := ackFrame.LargestAcked()
	if largestAcked > h.lastSentPacketNumber {
		return qerr.Error(qerr.InvalidAckData, "Received ACK for an unsent package")
	}

	// duplicate or out of order ACK
	if withPacketNumber != 0 && withPacketNumber < h.largestReceivedPacketWithAck {
		h.logger.Debugf("Ignoring ACK frame (duplicate or out of order).")
		return nil
	}
	h.largestReceivedPacketWithAck = withPacketNumber
	h.largestAcked = utils.MaxPacketNumber(h.largestAcked, largestAcked)

	if !h.packetNumberGenerator.Validate(ackFrame) {
		return qerr.Error(qerr.InvalidAckData, "Received an ACK for a skipped packet number")
	}

	if rttUpdated := h.maybeUpdateRTT(largestAcked, ackFrame.DelayTime, rcvTime); rttUpdated {
		h.congestion.MaybeExitSlowStart()
	}

	ackedPackets, err := h.determineNewlyAckedPackets(ackFrame)
	if err != nil {
		return err
	}
	if len(ackedPackets) == 0 {
		return nil
	}

	priorInFlight := h.bytesInFlight
	for _, p := range ackedPackets {
		// TODO(#1534): check the encryption level
		// if encLevel < p.EncryptionLevel {
		// 	return fmt.Errorf("Received ACK with encryption level %s that acks a packet %d (encryption level %s)", encLevel, p.PacketNumber, p.EncryptionLevel)
		// }

		// largestAcked == 0 either means that the packet didn't contain an ACK, or it just acked packet 0
		// It is safe to ignore the corner case of packets that just acked packet 0, because
		// the lowestPacketNotConfirmedAcked is only used to limit the number of ACK ranges we will send.
		if p.largestAcked != 0 {
			h.lowestPacketNotConfirmedAcked = utils.MaxPacketNumber(h.lowestPacketNotConfirmedAcked, p.largestAcked+1)
		}
		if err := h.onPacketAcked(p, rcvTime); err != nil {
			return err
		}
		if p.includedInBytesInFlight {
			h.congestion.OnPacketAcked(p.PacketNumber, p.Length, priorInFlight, rcvTime)
		}
	}

	if err := h.detectLostPackets(rcvTime, priorInFlight); err != nil {
		return err
	}

	h.ptoCount = 0
	h.cryptoCount = 0

	h.updateLossDetectionAlarm()
	return nil
}

func (h *sentPacketHandler) GetLowestPacketNotConfirmedAcked() protocol.PacketNumber {
	return h.lowestPacketNotConfirmedAcked
}

func (h *sentPacketHandler) determineNewlyAckedPackets(ackFrame *wire.AckFrame) ([]*Packet, error) {
	var ackedPackets []*Packet
	ackRangeIndex := 0
	lowestAcked := ackFrame.LowestAcked()
	largestAcked := ackFrame.LargestAcked()
	err := h.packetHistory.Iterate(func(p *Packet) (bool, error) {
		// Ignore packets below the lowest acked
		if p.PacketNumber < lowestAcked {
			return true, nil
		}
		// Break after largest acked is reached
		if p.PacketNumber > largestAcked {
			return false, nil
		}

		if ackFrame.HasMissingRanges() {
			ackRange := ackFrame.AckRanges[len(ackFrame.AckRanges)-1-ackRangeIndex]

			for p.PacketNumber > ackRange.Largest && ackRangeIndex < len(ackFrame.AckRanges)-1 {
				ackRangeIndex++
				ackRange = ackFrame.AckRanges[len(ackFrame.AckRanges)-1-ackRangeIndex]
			}

			if p.PacketNumber >= ackRange.Smallest { // packet i contained in ACK range
				if p.PacketNumber > ackRange.Largest {
					return false, fmt.Errorf("BUG: ackhandler would have acked wrong packet 0x%x, while evaluating range 0x%x -> 0x%x", p.PacketNumber, ackRange.Smallest, ackRange.Largest)
				}
				ackedPackets = append(ackedPackets, p)
			}
		} else {
			ackedPackets = append(ackedPackets, p)
		}
		return true, nil
	})
	if h.logger.Debug() && len(ackedPackets) > 0 {
		pns := make([]protocol.PacketNumber, len(ackedPackets))
		for i, p := range ackedPackets {
			pns[i] = p.PacketNumber
		}
		h.logger.Debugf("\tnewly acked packets (%d): %#x", len(pns), pns)
	}
	return ackedPackets, err
}

func (h *sentPacketHandler) maybeUpdateRTT(largestAcked protocol.PacketNumber, ackDelay time.Duration, rcvTime time.Time) bool {
	if p := h.packetHistory.GetPacket(largestAcked); p != nil {
		h.rttStats.UpdateRTT(rcvTime.Sub(p.SendTime), ackDelay, rcvTime)
		if h.logger.Debug() {
			h.logger.Debugf("\tupdated RTT: %s (σ: %s)", h.rttStats.SmoothedRTT(), h.rttStats.MeanDeviation())
		}
		return true
	}
	return false
}

func (h *sentPacketHandler) updateLossDetectionAlarm() {
	// Cancel the alarm if no packets are outstanding
	if !h.packetHistory.HasOutstandingPackets() {
		h.alarm = time.Time{}
		return
	}

	if h.packetHistory.HasOutstandingCryptoPackets() {
		h.alarm = h.lastSentCryptoPacketTime.Add(h.computeCryptoTimeout())
	} else if !h.lossTime.IsZero() {
		// Early retransmit timer or time loss detection.
		h.alarm = h.lossTime
	} else { // PTO alarm
		h.alarm = h.lastSentRetransmittablePacketTime.Add(h.computePTOTimeout())
	}
}

func (h *sentPacketHandler) detectLostPackets(now time.Time, priorInFlight protocol.ByteCount) error {
	h.lossTime = time.Time{}

	maxRTT := float64(utils.MaxDuration(h.rttStats.LatestRTT(), h.rttStats.SmoothedRTT()))
	delayUntilLost := time.Duration((1.0 + timeReorderingFraction) * maxRTT)

	var lostPackets []*Packet
	h.packetHistory.Iterate(func(packet *Packet) (bool, error) {
		if packet.PacketNumber > h.largestAcked {
			return false, nil
		}

		timeSinceSent := now.Sub(packet.SendTime)
		if timeSinceSent > delayUntilLost {
			lostPackets = append(lostPackets, packet)
		} else if h.lossTime.IsZero() {
			if h.logger.Debug() {
				h.logger.Debugf("\tsetting loss timer for packet %#x to %s (in %s)", packet.PacketNumber, delayUntilLost, delayUntilLost-timeSinceSent)
			}
			// Note: This conditional is only entered once per call
			h.lossTime = now.Add(delayUntilLost - timeSinceSent)
		}
		return true, nil
	})

	if h.logger.Debug() && len(lostPackets) > 0 {
		pns := make([]protocol.PacketNumber, len(lostPackets))
		for i, p := range lostPackets {
			pns[i] = p.PacketNumber
		}
		h.logger.Debugf("\tlost packets (%d): %#x", len(pns), pns)
	}

	for _, p := range lostPackets {
		// the bytes in flight need to be reduced no matter if this packet will be retransmitted
		if p.includedInBytesInFlight {
			h.bytesInFlight -= p.Length
			h.congestion.OnPacketLost(p.PacketNumber, p.Length, priorInFlight)
		}
		if p.canBeRetransmitted {
			// queue the packet for retransmission, and report the loss to the congestion controller
			if err := h.queuePacketForRetransmission(p); err != nil {
				return err
			}
		}
		h.packetHistory.Remove(p.PacketNumber)
	}
	return nil
}

func (h *sentPacketHandler) OnAlarm() error {
	// When all outstanding are acknowledged, the alarm is canceled in
	// updateLossDetectionAlarm. This doesn't reset the timer in the session though.
	// When OnAlarm is called, we therefore need to make sure that there are
	// actually packets outstanding.
	if h.packetHistory.HasOutstandingPackets() {
		if err := h.onVerifiedAlarm(); err != nil {
			return err
		}
	}
	h.updateLossDetectionAlarm()
	return nil
}

func (h *sentPacketHandler) onVerifiedAlarm() error {
	var err error
	if h.packetHistory.HasOutstandingCryptoPackets() {
		if h.logger.Debug() {
			h.logger.Debugf("Loss detection alarm fired in crypto mode. Crypto count: %d", h.cryptoCount)
		}
		h.cryptoCount++
		err = h.queueCryptoPacketsForRetransmission()
	} else if !h.lossTime.IsZero() {
		if h.logger.Debug() {
			h.logger.Debugf("Loss detection alarm fired in loss timer mode. Loss time: %s", h.lossTime)
		}
		// Early retransmit or time loss detection
		err = h.detectLostPackets(time.Now(), h.bytesInFlight)
	} else { // PTO
		if h.logger.Debug() {
			h.logger.Debugf("Loss detection alarm fired in PTO mode. PTO count: %d", h.ptoCount)
		}
		h.ptoCount++
		h.numProbesToSend += 2
	}
	return err
}

func (h *sentPacketHandler) GetAlarmTimeout() time.Time {
	return h.alarm
}

func (h *sentPacketHandler) onPacketAcked(p *Packet, rcvTime time.Time) error {
	// This happens if a packet and its retransmissions is acked in the same ACK.
	// As soon as we process the first one, this will remove all the retransmissions,
	// so we won't find the retransmitted packet number later.
	if packet := h.packetHistory.GetPacket(p.PacketNumber); packet == nil {
		return nil
	}

	// only report the acking of this packet to the congestion controller if:
	// * it is a retransmittable packet
	// * this packet wasn't retransmitted yet
	if p.isRetransmission {
		// that the parent doesn't exist is expected to happen every time the original packet was already acked
		if parent := h.packetHistory.GetPacket(p.retransmissionOf); parent != nil {
			if len(parent.retransmittedAs) == 1 {
				parent.retransmittedAs = nil
			} else {
				// remove this packet from the slice of retransmission
				retransmittedAs := make([]protocol.PacketNumber, 0, len(parent.retransmittedAs)-1)
				for _, pn := range parent.retransmittedAs {
					if pn != p.PacketNumber {
						retransmittedAs = append(retransmittedAs, pn)
					}
				}
				parent.retransmittedAs = retransmittedAs
			}
		}
	}
	// this also applies to packets that have been retransmitted as probe packets
	if p.includedInBytesInFlight {
		h.bytesInFlight -= p.Length
	}
	if err := h.stopRetransmissionsFor(p); err != nil {
		return err
	}
	return h.packetHistory.Remove(p.PacketNumber)
}

func (h *sentPacketHandler) stopRetransmissionsFor(p *Packet) error {
	if err := h.packetHistory.MarkCannotBeRetransmitted(p.PacketNumber); err != nil {
		return err
	}
	for _, r := range p.retransmittedAs {
		packet := h.packetHistory.GetPacket(r)
		if packet == nil {
			return fmt.Errorf("sent packet handler BUG: marking packet as not retransmittable %d (retransmission of %d) not found in history", r, p.PacketNumber)
		}
		h.stopRetransmissionsFor(packet)
	}
	return nil
}

func (h *sentPacketHandler) DequeuePacketForRetransmission() *Packet {
	if len(h.retransmissionQueue) == 0 {
		return nil
	}
	packet := h.retransmissionQueue[0]
	// Shift the slice and don't retain anything that isn't needed.
	copy(h.retransmissionQueue, h.retransmissionQueue[1:])
	h.retransmissionQueue[len(h.retransmissionQueue)-1] = nil
	h.retransmissionQueue = h.retransmissionQueue[:len(h.retransmissionQueue)-1]
	return packet
}

func (h *sentPacketHandler) DequeueProbePacket() (*Packet, error) {
	if len(h.retransmissionQueue) == 0 {
		p := h.packetHistory.FirstOutstanding()
		if p == nil {
			return nil, errors.New("cannot dequeue a probe packet. No outstanding packets")
		}
		if err := h.queuePacketForRetransmission(p); err != nil {
			return nil, err
		}
	}
	return h.DequeuePacketForRetransmission(), nil
}

func (h *sentPacketHandler) PeekPacketNumber() (protocol.PacketNumber, protocol.PacketNumberLen) {
	pn := h.packetNumberGenerator.Peek()
	return pn, protocol.GetPacketNumberLengthForHeader(pn, h.lowestUnacked())
}

func (h *sentPacketHandler) PopPacketNumber() protocol.PacketNumber {
	return h.packetNumberGenerator.Pop()
}

func (h *sentPacketHandler) SendMode() SendMode {
	numTrackedPackets := len(h.retransmissionQueue) + h.packetHistory.Len()

	// Don't send any packets if we're keeping track of the maximum number of packets.
	// Note that since MaxOutstandingSentPackets is smaller than MaxTrackedSentPackets,
	// we will stop sending out new data when reaching MaxOutstandingSentPackets,
	// but still allow sending of retransmissions and ACKs.
	if numTrackedPackets >= protocol.MaxTrackedSentPackets {
		if h.logger.Debug() {
			h.logger.Debugf("Limited by the number of tracked packets: tracking %d packets, maximum %d", numTrackedPackets, protocol.MaxTrackedSentPackets)
		}
		return SendNone
	}
	if h.numProbesToSend > 0 {
		return SendPTO
	}
	// Only send ACKs if we're congestion limited.
	if cwnd := h.congestion.GetCongestionWindow(); h.bytesInFlight > cwnd {
		if h.logger.Debug() {
			h.logger.Debugf("Congestion limited: bytes in flight %d, window %d", h.bytesInFlight, cwnd)
		}
		return SendAck
	}
	// Send retransmissions first, if there are any.
	if len(h.retransmissionQueue) > 0 {
		return SendRetransmission
	}
	if numTrackedPackets >= protocol.MaxOutstandingSentPackets {
		if h.logger.Debug() {
			h.logger.Debugf("Max outstanding limited: tracking %d packets, maximum: %d", numTrackedPackets, protocol.MaxOutstandingSentPackets)
		}
		return SendAck
	}
	return SendAny
}

func (h *sentPacketHandler) TimeUntilSend() time.Time {
	return h.nextPacketSendTime
}

func (h *sentPacketHandler) ShouldSendNumPackets() int {
	if h.numProbesToSend > 0 {
		// RTO probes should not be paced, but must be sent immediately.
		return h.numProbesToSend
	}
	delay := h.congestion.TimeUntilSend(h.bytesInFlight)
	if delay == 0 || delay > protocol.MinPacingDelay {
		return 1
	}
	return int(math.Ceil(float64(protocol.MinPacingDelay) / float64(delay)))
}

func (h *sentPacketHandler) queueCryptoPacketsForRetransmission() error {
	var cryptoPackets []*Packet
	h.packetHistory.Iterate(func(p *Packet) (bool, error) {
		if p.canBeRetransmitted && p.EncryptionLevel != protocol.Encryption1RTT {
			cryptoPackets = append(cryptoPackets, p)
		}
		return true, nil
	})
	for _, p := range cryptoPackets {
		h.logger.Debugf("Queueing packet %#x as a crypto retransmission", p.PacketNumber)
		if err := h.queuePacketForRetransmission(p); err != nil {
			return err
		}
	}
	return nil
}

func (h *sentPacketHandler) queuePacketForRetransmission(p *Packet) error {
	if !p.canBeRetransmitted {
		return fmt.Errorf("sent packet handler BUG: packet %d already queued for retransmission", p.PacketNumber)
	}
	if err := h.packetHistory.MarkCannotBeRetransmitted(p.PacketNumber); err != nil {
		return err
	}
	h.retransmissionQueue = append(h.retransmissionQueue, p)
	return nil
}

func (h *sentPacketHandler) computeCryptoTimeout() time.Duration {
	duration := utils.MaxDuration(2*h.rttStats.SmoothedOrInitialRTT(), granularity)
	// exponential backoff
	// There's an implicit limit to this set by the crypto timeout.
	return duration << h.cryptoCount
}

func (h *sentPacketHandler) computePTOTimeout() time.Duration {
	// TODO(#1236): include the max_ack_delay
	duration := utils.MaxDuration(h.rttStats.SmoothedOrInitialRTT()+4*h.rttStats.MeanDeviation(), granularity)
	return duration << h.ptoCount
}
