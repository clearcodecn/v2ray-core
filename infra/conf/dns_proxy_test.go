package conf_test

import (
	"testing"

	"github.com/clearcodecn/v2ray-core/common/net"
	. "github.com/clearcodecn/v2ray-core/infra/conf"
	"github.com/clearcodecn/v2ray-core/proxy/dns"
)

func TestDnsProxyConfig(t *testing.T) {
	creator := func() Buildable {
		return new(DnsOutboundConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"address": "8.8.8.8",
				"port": 53,
				"network": "tcp"
			}`,
			Parser: loadJSON(creator),
			Output: &dns.Config{
				Server: &net.Endpoint{
					Network: net.Network_TCP,
					Address: net.NewIPOrDomain(net.IPAddress([]byte{8, 8, 8, 8})),
					Port:    53,
				},
			},
		},
	})
}
