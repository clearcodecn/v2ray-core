// +build !confonly

package core

import (
	"bytes"
	"context"

	"github.com/clearcodecn/v2ray-core/common"
	"github.com/clearcodecn/v2ray-core/common/net"
	"github.com/clearcodecn/v2ray-core/features/routing"
	"github.com/clearcodecn/v2ray-core/transport/internet/udp"
)

// CreateObject creates a new object based on the given V2Ray instance and config. The V2Ray instance may be nil.
func CreateObject(v *Instance, config interface{}) (interface{}, error) {
	ctx := context.Background()
	if v != nil {
		ctx = context.WithValue(ctx, v2rayKey, v)
	}
	return common.CreateObject(ctx, config)
}

// StartInstance starts a new V2Ray instance with given serialized config.
// By default V2Ray only support config in protobuf format, i.e., configFormat = "protobuf". Caller need to load other packages to add JSON support.
//
// v2ray:api:stable
func StartInstance(configFormat string, configBytes []byte) (*Instance, error) {
	config, err := LoadConfig(configFormat, "", bytes.NewReader(configBytes))
	if err != nil {
		return nil, err
	}
	instance, err := New(config)
	if err != nil {
		return nil, err
	}
	if err := instance.Start(); err != nil {
		return nil, err
	}
	return instance, nil
}

// Dial provides an easy way for upstream caller to create net.Conn through V2Ray.
// It dispatches the request to the given destination by the given V2Ray instance.
// Since it is under a proxy context, the LocalAddr() and RemoteAddr() in returned net.Conn
// will not show real addresses being used for communication.
//
// v2ray:api:stable
func Dial(ctx context.Context, v *Instance, dest net.Destination) (net.Conn, error) {
	dispatcher := v.GetFeature(routing.DispatcherType())
	if dispatcher == nil {
		return nil, newError("routing.Dispatcher is not registered in V2Ray core")
	}
	r, err := dispatcher.(routing.Dispatcher).Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt net.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = net.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = net.ConnectionOutputMultiUDP(r.Reader)
	}
	return net.NewConnection(net.ConnectionInputMulti(r.Writer), readerOpt), nil
}

// DialUDP provides a way to exchange UDP packets through V2Ray instance to remote servers.
// Since it is under a proxy context, the LocalAddr() in returned PacketConn will not show the real address.
//
// TODO: SetDeadline() / SetReadDeadline() / SetWriteDeadline() are not implemented.
//
// v2ray:api:beta
func DialUDP(ctx context.Context, v *Instance) (net.PacketConn, error) {
	dispatcher := v.GetFeature(routing.DispatcherType())
	if dispatcher == nil {
		return nil, newError("routing.Dispatcher is not registered in V2Ray core")
	}
	return udp.DialDispatcher(ctx, dispatcher.(routing.Dispatcher))
}
