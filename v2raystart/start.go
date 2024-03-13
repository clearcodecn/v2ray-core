package v2raystart

import (
	"github.com/clearcodecn/v2ray"
	"github.com/clearcodecn/v2ray/common/errors"
	"github.com/clearcodecn/v2ray/main/confloader/external"

	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "github.com/clearcodecn/v2ray/app/dispatcher"
	_ "github.com/clearcodecn/v2ray/app/proxyman/inbound"
	_ "github.com/clearcodecn/v2ray/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/clearcodecn/v2ray/app/commander"
	_ "github.com/clearcodecn/v2ray/app/log/command"
	_ "github.com/clearcodecn/v2ray/app/proxyman/command"
	_ "github.com/clearcodecn/v2ray/app/stats/command"

	// Other optional features.
	_ "github.com/clearcodecn/v2ray/app/dns"
	_ "github.com/clearcodecn/v2ray/app/log"
	_ "github.com/clearcodecn/v2ray/app/policy"
	_ "github.com/clearcodecn/v2ray/app/reverse"
	_ "github.com/clearcodecn/v2ray/app/router"
	_ "github.com/clearcodecn/v2ray/app/stats"

	// Inbound and outbound proxies.
	_ "github.com/clearcodecn/v2ray/proxy/blackhole"
	_ "github.com/clearcodecn/v2ray/proxy/dns"
	_ "github.com/clearcodecn/v2ray/proxy/dokodemo"
	_ "github.com/clearcodecn/v2ray/proxy/freedom"
	_ "github.com/clearcodecn/v2ray/proxy/http"
	_ "github.com/clearcodecn/v2ray/proxy/mtproto"
	_ "github.com/clearcodecn/v2ray/proxy/shadowsocks"
	_ "github.com/clearcodecn/v2ray/proxy/socks"
	_ "github.com/clearcodecn/v2ray/proxy/vmess/inbound"
	_ "github.com/clearcodecn/v2ray/proxy/vmess/outbound"

	// Transports
	_ "github.com/clearcodecn/v2ray/transport/internet/domainsocket"
	_ "github.com/clearcodecn/v2ray/transport/internet/http"
	_ "github.com/clearcodecn/v2ray/transport/internet/kcp"
	_ "github.com/clearcodecn/v2ray/transport/internet/quic"
	_ "github.com/clearcodecn/v2ray/transport/internet/tcp"
	_ "github.com/clearcodecn/v2ray/transport/internet/tls"
	_ "github.com/clearcodecn/v2ray/transport/internet/udp"
	_ "github.com/clearcodecn/v2ray/transport/internet/websocket"

	// Transport headers
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/http"
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/noop"
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/srtp"
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/tls"
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/utp"
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/wechat"
	_ "github.com/clearcodecn/v2ray/transport/internet/headers/wireguard"

	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from v2ctl
	_ "github.com/clearcodecn/v2ray/main/json"
	// The following line loads JSON internally
	// _ "v2ray.com/core/main/jsonem"

	// Load config from file or http(s)
	_ "github.com/clearcodecn/v2ray/main/confloader/external"
)

type errPathObjHolder struct{}

func newError(values ...interface{}) *errors.Error {
	return errors.New(values...).WithPathObj(errPathObjHolder{})
}

func Start(uri string, stopChan chan struct{}) (core.Server, error) {
	out, err := external.ConfigLoader(uri)
	if err != nil {
		return nil, err
	}
	config, err := core.LoadConfig("json", uri, out)
	if err != nil {
		return nil, newError("failed to read config files: [", uri, "]").Base(err)
	}

	server, err := core.New(config)
	if err != nil {
		return nil, newError("failed to create server").Base(err)
	}

	go func() {
		if err = server.Start(); err != nil {
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	go func() {
		<-stopChan
		server.Close()
	}()

	return server, nil
}
