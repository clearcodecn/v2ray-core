//go:build !linux && !freebsd && !confonly
// +build !linux,!freebsd,!confonly

package tcp

import (
	"github.com/clearcodecn/v2ray-core/common/net"
	"github.com/clearcodecn/v2ray-core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
