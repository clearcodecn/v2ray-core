package udp

import (
	"github.com/clearcodecn/v2ray-core/common"
	"github.com/clearcodecn/v2ray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
