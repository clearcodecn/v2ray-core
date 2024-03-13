package udp

import (
	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
