package udp

import (
	"github.com/clearcodecn/v2ray/common/buf"
	"github.com/clearcodecn/v2ray/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
