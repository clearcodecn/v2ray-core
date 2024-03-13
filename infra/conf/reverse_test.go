package conf_test

import (
	"testing"

	"github.com/clearcodecn/v2ray-core/app/reverse"
	"github.com/clearcodecn/v2ray-core/infra/conf"
)

func TestReverseConfig(t *testing.T) {
	creator := func() conf.Buildable {
		return new(conf.ReverseConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"bridges": [{
					"tag": "test",
					"domain": "test.v2ray.com"
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &reverse.Config{
				BridgeConfig: []*reverse.BridgeConfig{
					{Tag: "test", Domain: "test.v2ray.com"},
				},
			},
		},
		{
			Input: `{
				"portals": [{
					"tag": "test",
					"domain": "test.v2ray.com"
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &reverse.Config{
				PortalConfig: []*reverse.PortalConfig{
					{Tag: "test", Domain: "test.v2ray.com"},
				},
			},
		},
	})
}
