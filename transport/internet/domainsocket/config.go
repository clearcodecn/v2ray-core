// +build !confonly

package domainsocket

import (
	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/common/net"
	"github.com/clearcodecn/v2ray/transport/internet"
)

const protocolName = "domainsocket"

func (c *Config) GetUnixAddr() (*net.UnixAddr, error) {
	path := c.Path
	if path == "" {
		return nil, newError("empty domain socket path")
	}
	if c.Abstract && path[0] != '\x00' {
		path = "\x00" + path
	}
	return &net.UnixAddr{
		Name: path,
		Net:  "unix",
	}, nil
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
