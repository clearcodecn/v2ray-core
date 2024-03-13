package congestion

import "github.com/clearcodecn/v2ray/external/github.com/lucas-clemente/quic-go/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
