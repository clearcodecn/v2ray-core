package inbound

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	protocol "github.com/clearcodecn/v2ray-core/common/protocol"
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

type DetourConfig struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetourConfig) Reset()         { *m = DetourConfig{} }
func (m *DetourConfig) String() string { return proto.CompactTextString(m) }
func (*DetourConfig) ProtoMessage()    {}
func (*DetourConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a47d4a41f33382d2, []int{0}
}

func (m *DetourConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetourConfig.Unmarshal(m, b)
}
func (m *DetourConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetourConfig.Marshal(b, m, deterministic)
}
func (m *DetourConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetourConfig.Merge(m, src)
}
func (m *DetourConfig) XXX_Size() int {
	return xxx_messageInfo_DetourConfig.Size(m)
}
func (m *DetourConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DetourConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DetourConfig proto.InternalMessageInfo

func (m *DetourConfig) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type DefaultConfig struct {
	AlterId              uint32   `protobuf:"varint,1,opt,name=alter_id,json=alterId,proto3" json:"alter_id,omitempty"`
	Level                uint32   `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultConfig) Reset()         { *m = DefaultConfig{} }
func (m *DefaultConfig) String() string { return proto.CompactTextString(m) }
func (*DefaultConfig) ProtoMessage()    {}
func (*DefaultConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a47d4a41f33382d2, []int{1}
}

func (m *DefaultConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultConfig.Unmarshal(m, b)
}
func (m *DefaultConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultConfig.Marshal(b, m, deterministic)
}
func (m *DefaultConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultConfig.Merge(m, src)
}
func (m *DefaultConfig) XXX_Size() int {
	return xxx_messageInfo_DefaultConfig.Size(m)
}
func (m *DefaultConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultConfig proto.InternalMessageInfo

func (m *DefaultConfig) GetAlterId() uint32 {
	if m != nil {
		return m.AlterId
	}
	return 0
}

func (m *DefaultConfig) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type Config struct {
	User                 []*protocol.User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Default              *DefaultConfig   `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Detour               *DetourConfig    `protobuf:"bytes,3,opt,name=detour,proto3" json:"detour,omitempty"`
	SecureEncryptionOnly bool             `protobuf:"varint,4,opt,name=secure_encryption_only,json=secureEncryptionOnly,proto3" json:"secure_encryption_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_a47d4a41f33382d2, []int{2}
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

func (m *Config) GetUser() []*protocol.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Config) GetDefault() *DefaultConfig {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *Config) GetDetour() *DetourConfig {
	if m != nil {
		return m.Detour
	}
	return nil
}

func (m *Config) GetSecureEncryptionOnly() bool {
	if m != nil {
		return m.SecureEncryptionOnly
	}
	return false
}

func init() {
	proto.RegisterType((*DetourConfig)(nil), "v2ray.core.proxy.vmess.inbound.DetourConfig")
	proto.RegisterType((*DefaultConfig)(nil), "v2ray.core.proxy.vmess.inbound.DefaultConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.vmess.inbound.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/proxy/vmess/inbound/config.proto", fileDescriptor_a47d4a41f33382d2)
}

var fileDescriptor_a47d4a41f33382d2 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xd3, 0xc2, 0x0b, 0xbc, 0x8b, 0x78, 0x68, 0x88, 0xa9, 0x1e, 0x48, 0xd3, 0x13, 0x26,
	0xba, 0x9b, 0x54, 0x3e, 0x80, 0x11, 0x8c, 0xe1, 0x24, 0x69, 0x22, 0x07, 0x2f, 0xa4, 0x6c, 0x07,
	0xd3, 0x64, 0xbb, 0x43, 0xb6, 0x5b, 0x62, 0xcf, 0x7e, 0x1b, 0x3f, 0xa5, 0x61, 0x5a, 0xfc, 0x77,
	0x90, 0xdb, 0xce, 0xce, 0xef, 0x79, 0xe6, 0x99, 0x61, 0x62, 0x17, 0x99, 0xa4, 0xe2, 0x12, 0x73,
	0x21, 0xd1, 0x80, 0xd8, 0x1a, 0x7c, 0xad, 0xc4, 0x2e, 0x87, 0xa2, 0x10, 0x99, 0x5e, 0x63, 0xa9,
	0x53, 0x21, 0x51, 0x6f, 0xb2, 0x17, 0xbe, 0x35, 0x68, 0xd1, 0x1b, 0x1d, 0x04, 0x06, 0x38, 0xc1,
	0x9c, 0x60, 0xde, 0xc0, 0x17, 0x97, 0xbf, 0x0c, 0x25, 0xe6, 0x39, 0x6a, 0x41, 0x62, 0x89, 0x4a,
	0x94, 0x05, 0x98, 0xda, 0x2a, 0x1c, 0xb1, 0x93, 0x19, 0x58, 0x2c, 0xcd, 0x94, 0x06, 0x78, 0xa7,
	0xcc, 0xb5, 0xe8, 0x3b, 0x81, 0x33, 0xfe, 0x1f, 0xbb, 0x16, 0xc3, 0x5b, 0x36, 0x98, 0xc1, 0x26,
	0x29, 0x95, 0x6d, 0x80, 0x73, 0xd6, 0x4b, 0x94, 0x05, 0xb3, 0xca, 0x52, 0xc2, 0x06, 0x71, 0x97,
	0xea, 0x79, 0xea, 0x0d, 0xd9, 0x3f, 0x05, 0x3b, 0x50, 0xbe, 0x4b, 0xff, 0x75, 0x11, 0xbe, 0xb9,
	0xac, 0xd3, 0x68, 0x27, 0xac, 0xbd, 0x1f, 0xed, 0x3b, 0x41, 0x6b, 0xdc, 0x8f, 0x02, 0xfe, 0x6d,
	0x8d, 0x3a, 0x22, 0x3f, 0x44, 0xe4, 0x4f, 0x05, 0x98, 0x98, 0x68, 0xef, 0x81, 0x75, 0xd3, 0x3a,
	0x02, 0x19, 0xf7, 0xa3, 0x6b, 0xfe, 0xf7, 0xfe, 0xfc, 0x47, 0xe2, 0xf8, 0xa0, 0xf6, 0x66, 0xac,
	0x93, 0xd2, 0xae, 0x7e, 0x8b, 0x7c, 0xae, 0x8e, 0xfb, 0x7c, 0x5d, 0x26, 0x6e, 0xb4, 0xde, 0x84,
	0x9d, 0x15, 0x20, 0x4b, 0x03, 0x2b, 0xd0, 0xd2, 0x54, 0x5b, 0x9b, 0xa1, 0x5e, 0xa1, 0x56, 0x95,
	0xdf, 0x0e, 0x9c, 0x71, 0x2f, 0x1e, 0xd6, 0xdd, 0xfb, 0xcf, 0xe6, 0xa3, 0x56, 0xd5, 0xdd, 0x82,
	0x85, 0x12, 0xf3, 0x23, 0x03, 0x17, 0xce, 0x73, 0xb7, 0x79, 0xbe, 0xbb, 0xa3, 0x65, 0x14, 0x27,
	0x15, 0x9f, 0xee, 0xd9, 0x05, 0xb1, 0x4b, 0x62, 0xe7, 0x35, 0xb0, 0xee, 0xd0, 0xad, 0x6e, 0x3e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x49, 0x84, 0xe1, 0x3e, 0x02, 0x00, 0x00,
}
