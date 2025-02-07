package dns

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

type Config struct {
	// Server is the DNS server address. If specified, this address overrides the original one.
	Server               *net.Endpoint `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_c49bb2d51e576d57, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetServer() *net.Endpoint {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.dns.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/proxy/dns/config.proto", fileDescriptor_c49bb2d51e576d57)
}

var fileDescriptor_c49bb2d51e576d57 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x31, 0xab, 0xc2, 0x30,
	0x14, 0x46, 0xe9, 0x7b, 0xd0, 0x21, 0x6f, 0x2b, 0x1d, 0xca, 0x5b, 0xde, 0x43, 0x10, 0x04, 0xe1,
	0x06, 0xea, 0xa0, 0xab, 0x56, 0xf7, 0xd2, 0xc1, 0xc1, 0xad, 0x26, 0x51, 0x32, 0xe4, 0xde, 0x72,
	0x13, 0x8a, 0xfd, 0x4b, 0xfe, 0x4a, 0x31, 0x55, 0x10, 0x71, 0xfe, 0xce, 0x77, 0x38, 0x62, 0xda,
	0x97, 0xdc, 0x0e, 0xa0, 0xc8, 0x49, 0x45, 0x6c, 0x64, 0xc7, 0x74, 0x19, 0xa4, 0x46, 0x2f, 0x15,
	0xe1, 0xc9, 0x9e, 0xa1, 0x63, 0x0a, 0x94, 0xe5, 0x4f, 0x8c, 0x0d, 0x44, 0x04, 0x34, 0xfa, 0xdf,
	0xf9, 0xdb, 0x59, 0x91, 0x73, 0x84, 0x12, 0x4d, 0x90, 0xda, 0xf8, 0x60, 0xb1, 0x0d, 0x96, 0x70,
	0x54, 0x4c, 0xd6, 0x22, 0xad, 0xa2, 0x32, 0x5b, 0x8a, 0xd4, 0x1b, 0xee, 0x0d, 0x17, 0xc9, 0x7f,
	0x32, 0xfb, 0x29, 0xff, 0xe0, 0xc5, 0x3e, 0x3a, 0x00, 0x4d, 0x80, 0x1d, 0xea, 0x8e, 0x2c, 0x86,
	0xe6, 0x81, 0x6f, 0x56, 0xa2, 0x50, 0xe4, 0xe0, 0x53, 0x4b, 0x9d, 0x1c, 0xbe, 0x35, 0xfa, 0xeb,
	0x57, 0xbe, 0x2f, 0x9b, 0x76, 0x80, 0xea, 0xbe, 0xd6, 0x71, 0xdd, 0xa2, 0x3f, 0xa6, 0xb1, 0x61,
	0x71, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x84, 0x9c, 0x79, 0xef, 0x00, 0x00, 0x00,
}
