// +build !confonly

package kcp

import (
	"context"
	"crypto/tls"
	"io"
	"sync/atomic"

	"github.com/clearcodecn/v2ray-core/common"
	"github.com/clearcodecn/v2ray-core/common/buf"
	"github.com/clearcodecn/v2ray-core/common/dice"
	"github.com/clearcodecn/v2ray-core/common/net"
	"github.com/clearcodecn/v2ray-core/transport/internet"
	v2tls "github.com/clearcodecn/v2ray-core/transport/internet/tls"
)

var (
	globalConv = uint32(dice.RollUint16())
)

func fetchInput(ctx context.Context, input io.Reader, reader PacketReader, conn *Connection) {
	cache := make(chan *buf.Buffer, 1024)
	go func() {
		for {
			payload := buf.New()
			if _, err := payload.ReadFrom(input); err != nil {
				payload.Release()
				close(cache)
				return
			}
			select {
			case cache <- payload:
			default:
				payload.Release()
			}
		}
	}()

	for payload := range cache {
		segments := reader.Read(payload.Bytes())
		payload.Release()
		if len(segments) > 0 {
			conn.Input(segments)
		}
	}
}

// DialKCP dials a new KCP connections to the specific destination.
func DialKCP(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (internet.Connection, error) {
	dest.Network = net.Network_UDP
	newError("dialing mKCP to ", dest).WriteToLog()

	rawConn, err := internet.DialSystem(ctx, dest, streamSettings.SocketSettings)
	if err != nil {
		return nil, newError("failed to dial to dest: ", err).AtWarning().Base(err)
	}

	kcpSettings := streamSettings.ProtocolSettings.(*Config)

	header, err := kcpSettings.GetPackerHeader()
	if err != nil {
		return nil, newError("failed to create packet header").Base(err)
	}
	security, err := kcpSettings.GetSecurity()
	if err != nil {
		return nil, newError("failed to create security").Base(err)
	}
	reader := &KCPPacketReader{
		Header:   header,
		Security: security,
	}
	writer := &KCPPacketWriter{
		Header:   header,
		Security: security,
		Writer:   rawConn,
	}

	conv := uint16(atomic.AddUint32(&globalConv, 1))
	session := NewConnection(ConnMetadata{
		LocalAddr:    rawConn.LocalAddr(),
		RemoteAddr:   rawConn.RemoteAddr(),
		Conversation: conv,
	}, writer, rawConn, kcpSettings)

	go fetchInput(ctx, rawConn, reader, session)

	var iConn internet.Connection = session

	if config := v2tls.ConfigFromStreamSettings(streamSettings); config != nil {
		tlsConn := tls.Client(iConn, config.GetTLSConfig(v2tls.WithDestination(dest)))
		iConn = tlsConn
	}

	return iConn, nil
}

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName, DialKCP))
}
