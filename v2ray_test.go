package core_test

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	. "github.com/clearcodecn/v2ray"
	"github.com/clearcodecn/v2ray/app/dispatcher"
	"github.com/clearcodecn/v2ray/app/proxyman"
	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/common/net"
	"github.com/clearcodecn/v2ray/common/protocol"
	"github.com/clearcodecn/v2ray/common/serial"
	"github.com/clearcodecn/v2ray/common/uuid"
	"github.com/clearcodecn/v2ray/features/dns"
	"github.com/clearcodecn/v2ray/features/dns/localdns"
	_ "github.com/clearcodecn/v2ray/main/distro/all"
	"github.com/clearcodecn/v2ray/proxy/dokodemo"
	"github.com/clearcodecn/v2ray/proxy/vmess"
	"github.com/clearcodecn/v2ray/proxy/vmess/outbound"
	"github.com/clearcodecn/v2ray/testing/servers/tcp"
)

func TestV2RayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestV2RayClose(t *testing.T) {
	port := tcp.PickPort()

	userId := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(port),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP, net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userId.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
