package conf_test

import (
	"testing"

	"github.com/clearcodecn/v2ray-core/common/protocol"
	"github.com/clearcodecn/v2ray-core/common/serial"
	. "github.com/clearcodecn/v2ray-core/infra/conf"
	"github.com/clearcodecn/v2ray-core/proxy/mtproto"
)

func TestMTProtoServerConfig(t *testing.T) {
	creator := func() Buildable {
		return new(MTProtoServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"users": [{
					"email": "love@v2ray.com",
					"level": 1,
					"secret": "b0cbcef5a486d9636472ac27f8e11a9d"
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &mtproto.ServerConfig{
				User: []*protocol.User{
					{
						Email: "love@v2ray.com",
						Level: 1,
						Account: serial.ToTypedMessage(&mtproto.Account{
							Secret: []byte{176, 203, 206, 245, 164, 134, 217, 99, 100, 114, 172, 39, 248, 225, 26, 157},
						}),
					},
				},
			},
		},
	})
}
