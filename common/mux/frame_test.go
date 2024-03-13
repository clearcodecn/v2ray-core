package mux_test

import (
	"testing"

	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/common/buf"
	"github.com/clearcodecn/v2ray/common/mux"
	"github.com/clearcodecn/v2ray/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.v2ray.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
