package flowcontrol

import (
	"fmt"

	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/congestion"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/qerr"
	"github.com/clearcodecn/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/utils"
)

type streamFlowController struct {
	baseFlowController

	streamID protocol.StreamID

	queueWindowUpdate func()

	connection connectionFlowControllerI

	receivedFinalOffset bool
}

var _ StreamFlowController = &streamFlowController{}

// NewStreamFlowController gets a new flow controller for a stream
func NewStreamFlowController(
	streamID protocol.StreamID,
	cfc ConnectionFlowController,
	receiveWindow protocol.ByteCount,
	maxReceiveWindow protocol.ByteCount,
	initialSendWindow protocol.ByteCount,
	queueWindowUpdate func(protocol.StreamID),
	rttStats *congestion.RTTStats,
	logger utils.Logger,
) StreamFlowController {
	return &streamFlowController{
		streamID:          streamID,
		connection:        cfc.(connectionFlowControllerI),
		queueWindowUpdate: func() { queueWindowUpdate(streamID) },
		baseFlowController: baseFlowController{
			rttStats:             rttStats,
			receiveWindow:        receiveWindow,
			receiveWindowSize:    receiveWindow,
			maxReceiveWindowSize: maxReceiveWindow,
			sendWindow:           initialSendWindow,
			logger:               logger,
		},
	}
}

// UpdateHighestReceived updates the highestReceived value, if the byteOffset is higher
// it returns an ErrReceivedSmallerByteOffset if the received byteOffset is smaller than any byteOffset received before
func (c *streamFlowController) UpdateHighestReceived(byteOffset protocol.ByteCount, final bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// when receiving a final offset, check that this final offset is consistent with a final offset we might have received earlier
	if final && c.receivedFinalOffset && byteOffset != c.highestReceived {
		return qerr.Error(qerr.StreamDataAfterTermination, fmt.Sprintf("Received inconsistent final offset for stream %d (old: %d, new: %d bytes)", c.streamID, c.highestReceived, byteOffset))
	}
	// if we already received a final offset, check that the offset in the STREAM frames is below the final offset
	if c.receivedFinalOffset && byteOffset > c.highestReceived {
		return qerr.StreamDataAfterTermination
	}
	if final {
		c.receivedFinalOffset = true
	}
	if byteOffset == c.highestReceived {
		return nil
	}
	if byteOffset <= c.highestReceived {
		// a STREAM_FRAME with a higher offset was received before.
		if final {
			// If the current byteOffset is smaller than the offset in that STREAM_FRAME, this STREAM_FRAME contained data after the end of the stream
			return qerr.StreamDataAfterTermination
		}
		// this is a reordered STREAM_FRAME
		return nil
	}

	increment := byteOffset - c.highestReceived
	c.highestReceived = byteOffset
	if c.checkFlowControlViolation() {
		return qerr.Error(qerr.FlowControlReceivedTooMuchData, fmt.Sprintf("Received %d bytes on stream %d, allowed %d bytes", byteOffset, c.streamID, c.receiveWindow))
	}
	return c.connection.IncrementHighestReceived(increment)
}

func (c *streamFlowController) AddBytesRead(n protocol.ByteCount) {
	c.baseFlowController.AddBytesRead(n)
	c.connection.AddBytesRead(n)
}

func (c *streamFlowController) AddBytesSent(n protocol.ByteCount) {
	c.baseFlowController.AddBytesSent(n)
	c.connection.AddBytesSent(n)
}

func (c *streamFlowController) SendWindowSize() protocol.ByteCount {
	return utils.MinByteCount(c.baseFlowController.sendWindowSize(), c.connection.SendWindowSize())
}

func (c *streamFlowController) MaybeQueueWindowUpdate() {
	c.mutex.Lock()
	hasWindowUpdate := !c.receivedFinalOffset && c.hasWindowUpdate()
	c.mutex.Unlock()
	if hasWindowUpdate {
		c.queueWindowUpdate()
	}
	c.connection.MaybeQueueWindowUpdate()
}

func (c *streamFlowController) GetWindowUpdate() protocol.ByteCount {
	// don't use defer for unlocking the mutex here, GetWindowUpdate() is called frequently and defer shows up in the profiler
	c.mutex.Lock()
	// if we already received the final offset for this stream, the peer won't need any additional flow control credit
	if c.receivedFinalOffset {
		c.mutex.Unlock()
		return 0
	}

	oldWindowSize := c.receiveWindowSize
	offset := c.baseFlowController.getWindowUpdate()
	if c.receiveWindowSize > oldWindowSize { // auto-tuning enlarged the window size
		c.logger.Debugf("Increasing receive flow control window for stream %d to %d kB", c.streamID, c.receiveWindowSize/(1<<10))
		c.connection.EnsureMinimumWindowSize(protocol.ByteCount(float64(c.receiveWindowSize) * protocol.ConnectionFlowControlMultiplier))
	}
	c.mutex.Unlock()
	return offset
}
