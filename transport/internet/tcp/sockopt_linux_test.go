// +build linux

package tcp_test

import (
	"context"
	"strings"
	"testing"

	"github.com/clearcodecn/v2ray-core/common"
	"github.com/clearcodecn/v2ray-core/testing/servers/tcp"
	"github.com/clearcodecn/v2ray-core/transport/internet"
	. "github.com/clearcodecn/v2ray-core/transport/internet/tcp"
)

func TestGetOriginalDestination(t *testing.T) {
	tcpServer := tcp.Server{}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	config, err := internet.ToMemoryStreamConfig(nil)
	common.Must(err)
	conn, err := Dial(context.Background(), dest, config)
	common.Must(err)
	defer conn.Close()

	originalDest, err := GetOriginalDestination(conn)
	if !(dest == originalDest || strings.Contains(err.Error(), "failed to call getsockopt")) {
		t.Error("unexpected state")
	}
}
