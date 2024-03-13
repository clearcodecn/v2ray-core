package command_test

import (
	"context"
	"testing"

	"github.com/clearcodecn/v2ray"
	"github.com/clearcodecn/v2ray/app/dispatcher"
	"github.com/clearcodecn/v2ray/app/log"
	. "github.com/clearcodecn/v2ray/app/log/command"
	"github.com/clearcodecn/v2ray/app/proxyman"
	_ "github.com/clearcodecn/v2ray/app/proxyman/inbound"
	_ "github.com/clearcodecn/v2ray/app/proxyman/outbound"
	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/common/serial"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
