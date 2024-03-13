// +build !confonly

package tcp

import (
	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
