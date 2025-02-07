package wire

import (
	"fmt"
	"strings"

	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/utils"
)

// LogFrame logs a frame, either sent or received
func LogFrame(logger utils.Logger, frame Frame, sent bool) {
	if !logger.Debug() {
		return
	}
	dir := "<-"
	if sent {
		dir = "->"
	}
	switch f := frame.(type) {
	case *CryptoFrame:
		dataLen := protocol.ByteCount(len(f.Data))
		logger.Debugf("\t%s &wire.CryptoFrame{Offset: 0x%x, Data length: 0x%x, Offset + Data length: 0x%x}", dir, f.Offset, dataLen, f.Offset+dataLen)
	case *StreamFrame:
		logger.Debugf("\t%s &wire.StreamFrame{StreamID: %d, FinBit: %t, Offset: 0x%x, Data length: 0x%x, Offset + Data length: 0x%x}", dir, f.StreamID, f.FinBit, f.Offset, f.DataLen(), f.Offset+f.DataLen())
	case *AckFrame:
		if len(f.AckRanges) > 1 {
			ackRanges := make([]string, len(f.AckRanges))
			for i, r := range f.AckRanges {
				ackRanges[i] = fmt.Sprintf("{Largest: %#x, Smallest: %#x}", r.Largest, r.Smallest)
			}
			logger.Debugf("\t%s &wire.AckFrame{LargestAcked: %#x, LowestAcked: %#x, AckRanges: {%s}, DelayTime: %s}", dir, f.LargestAcked(), f.LowestAcked(), strings.Join(ackRanges, ", "), f.DelayTime.String())
		} else {
			logger.Debugf("\t%s &wire.AckFrame{LargestAcked: %#x, LowestAcked: %#x, DelayTime: %s}", dir, f.LargestAcked(), f.LowestAcked(), f.DelayTime.String())
		}
	default:
		logger.Debugf("\t%s %#v", dir, frame)
	}
}
