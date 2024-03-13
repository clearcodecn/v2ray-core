// +build !linux,!freebsd
// +build !confonly

package tcp

import (
	"github.com/clearcodecn/v2ray/common/net"
	"github.com/clearcodecn/v2ray/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
