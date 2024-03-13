package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	net "github.com/clearcodecn/v2ray-core/common/net"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServerEndpoint struct {
	Address              *net.IPOrDomain `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32          `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	User                 []*User         `protobuf:"bytes,3,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ServerEndpoint) Reset()         { *m = ServerEndpoint{} }
func (m *ServerEndpoint) String() string { return proto.CompactTextString(m) }
func (*ServerEndpoint) ProtoMessage()    {}
func (*ServerEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_8741b2fa976e72fc, []int{0}
}

func (m *ServerEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerEndpoint.Unmarshal(m, b)
}
func (m *ServerEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerEndpoint.Marshal(b, m, deterministic)
}
func (m *ServerEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerEndpoint.Merge(m, src)
}
func (m *ServerEndpoint) XXX_Size() int {
	return xxx_messageInfo_ServerEndpoint.Size(m)
}
func (m *ServerEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ServerEndpoint proto.InternalMessageInfo

func (m *ServerEndpoint) GetAddress() *net.IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ServerEndpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServerEndpoint) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerEndpoint)(nil), "v2ray.core.common.protocol.ServerEndpoint")
}

func init() {
	proto.RegisterFile("v2ray.com/core/common/protocol/server_spec.proto", fileDescriptor_8741b2fa976e72fc)
}

var fileDescriptor_8741b2fa976e72fc = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcf, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x07, 0x70, 0xb9, 0xad, 0x00, 0xb9, 0x82, 0xc1, 0x53, 0x94, 0x01, 0x05, 0x16, 0xc2, 0x72,
	0x46, 0x81, 0x8d, 0x8d, 0xc2, 0xd0, 0x89, 0x28, 0x08, 0x06, 0x16, 0x14, 0x9c, 0x1b, 0x2a, 0x61,
	0x5f, 0x74, 0x36, 0x95, 0xfa, 0x24, 0xbc, 0x03, 0x4f, 0x89, 0x6a, 0xd7, 0x13, 0x5f, 0xdb, 0xc9,
	0xfe, 0xdd, 0xdd, 0xff, 0xe4, 0xc5, 0xba, 0xe1, 0x7e, 0x03, 0x86, 0xac, 0x36, 0xc4, 0xa8, 0x0d,
	0x59, 0x4b, 0x4e, 0x8f, 0x4c, 0x81, 0x0c, 0xbd, 0x69, 0x8f, 0xbc, 0x46, 0x7e, 0xf1, 0x23, 0x1a,
	0x88, 0x8f, 0xaa, 0xcc, 0x1d, 0x8c, 0x90, 0x34, 0x64, 0x5d, 0x9e, 0xfd, 0x3c, 0xcd, 0x61, 0xd0,
	0xfd, 0x30, 0x30, 0x7a, 0x9f, 0x6c, 0x79, 0xfe, 0xcf, 0xda, 0x77, 0x8f, 0x9c, 0xe8, 0xe9, 0x87,
	0x90, 0x47, 0x0f, 0x31, 0xc5, 0x9d, 0x1b, 0x46, 0x5a, 0xb9, 0xa0, 0xae, 0xe5, 0xfe, 0x6e, 0x5c,
	0x21, 0x2a, 0x51, 0xcf, 0x9b, 0x13, 0xf8, 0x1e, 0xca, 0x61, 0x80, 0x65, 0x7b, 0xcf, 0xb7, 0x64,
	0xfb, 0x95, 0xeb, 0x72, 0x87, 0x52, 0x72, 0x36, 0x12, 0x87, 0x62, 0x52, 0x89, 0xfa, 0xb0, 0x8b,
	0xb5, 0xba, 0x92, 0xb3, 0xed, 0xc6, 0x62, 0x5a, 0x4d, 0xeb, 0x79, 0x53, 0xc1, 0xef, 0x27, 0xc2,
	0xa3, 0x47, 0xee, 0xa2, 0xbe, 0x59, 0xca, 0x63, 0x43, 0xf6, 0x0f, 0xdc, 0x8a, 0xe7, 0x83, 0x5c,
	0x7f, 0x4e, 0xca, 0xa7, 0xa6, 0xeb, 0x37, 0xb0, 0xd8, 0xc2, 0x45, 0x82, 0xed, 0xee, 0xf3, 0x75,
	0x2f, 0xb2, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x61, 0xa2, 0x3c, 0x8f, 0x01, 0x00,
	0x00,
}
