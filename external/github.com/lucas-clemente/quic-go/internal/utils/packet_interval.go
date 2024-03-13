package utils

import "github.com/clearcodecn/v2ray/external/github.com/lucas-clemente/quic-go/internal/protocol"

// PacketInterval is an interval from one PacketNumber to the other
type PacketInterval struct {
	Start protocol.PacketNumber
	End   protocol.PacketNumber
}
