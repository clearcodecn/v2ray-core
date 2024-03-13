package log

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	log "github.com/clearcodecn/v2ray/common/log"
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

type LogType int32

const (
	LogType_None    LogType = 0
	LogType_Console LogType = 1
	LogType_File    LogType = 2
	LogType_Event   LogType = 3
)

var LogType_name = map[int32]string{
	0: "None",
	1: "Console",
	2: "File",
	3: "Event",
}

var LogType_value = map[string]int32{
	"None":    0,
	"Console": 1,
	"File":    2,
	"Event":   3,
}

func (x LogType) String() string {
	return proto.EnumName(LogType_name, int32(x))
}

func (LogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92dfeade43d9e989, []int{0}
}

type Config struct {
	ErrorLogType         LogType      `protobuf:"varint,1,opt,name=error_log_type,json=errorLogType,proto3,enum=v2ray.core.app.log.LogType" json:"error_log_type,omitempty"`
	ErrorLogLevel        log.Severity `protobuf:"varint,2,opt,name=error_log_level,json=errorLogLevel,proto3,enum=v2ray.core.common.log.Severity" json:"error_log_level,omitempty"`
	ErrorLogPath         string       `protobuf:"bytes,3,opt,name=error_log_path,json=errorLogPath,proto3" json:"error_log_path,omitempty"`
	AccessLogType        LogType      `protobuf:"varint,4,opt,name=access_log_type,json=accessLogType,proto3,enum=v2ray.core.app.log.LogType" json:"access_log_type,omitempty"`
	AccessLogPath        string       `protobuf:"bytes,5,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_92dfeade43d9e989, []int{0}
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

func (m *Config) GetErrorLogType() LogType {
	if m != nil {
		return m.ErrorLogType
	}
	return LogType_None
}

func (m *Config) GetErrorLogLevel() log.Severity {
	if m != nil {
		return m.ErrorLogLevel
	}
	return log.Severity_Unknown
}

func (m *Config) GetErrorLogPath() string {
	if m != nil {
		return m.ErrorLogPath
	}
	return ""
}

func (m *Config) GetAccessLogType() LogType {
	if m != nil {
		return m.AccessLogType
	}
	return LogType_None
}

func (m *Config) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func init() {
	proto.RegisterEnum("v2ray.core.app.log.LogType", LogType_name, LogType_value)
	proto.RegisterType((*Config)(nil), "v2ray.core.app.log.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/log/config.proto", fileDescriptor_92dfeade43d9e989)
}

var fileDescriptor_92dfeade43d9e989 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0x47, 0xff, 0x49, 0xbf, 0xfe, 0x1d, 0x6d, 0x1b, 0x66, 0x21, 0x45, 0x17, 0x16, 0x15, 0x29,
	0x2e, 0x26, 0x50, 0xf5, 0x01, 0x6a, 0x50, 0x37, 0x45, 0x4a, 0x14, 0x17, 0x6e, 0xca, 0x38, 0x5c,
	0xa7, 0x81, 0x49, 0xee, 0x30, 0x19, 0x02, 0x79, 0x22, 0xc1, 0xa7, 0x94, 0x4c, 0x1b, 0x52, 0x3f,
	0xc0, 0x65, 0x92, 0x7b, 0xce, 0xef, 0x40, 0xc8, 0x69, 0x31, 0x33, 0xbc, 0x64, 0x02, 0xd3, 0x50,
	0xa0, 0x81, 0x90, 0x6b, 0x1d, 0x2a, 0x94, 0xa1, 0xc0, 0xec, 0x2d, 0x91, 0x4c, 0x1b, 0xb4, 0x48,
	0x69, 0x7d, 0x64, 0x80, 0x71, 0xad, 0x99, 0x42, 0x79, 0xf8, 0x1d, 0x14, 0x98, 0xa6, 0x98, 0x39,
	0x56, 0xe1, 0x16, 0x3c, 0x79, 0xf7, 0x49, 0x37, 0x72, 0x26, 0x3a, 0x27, 0x43, 0x30, 0x06, 0xcd,
	0x4a, 0xa1, 0x5c, 0xd9, 0x52, 0xc3, 0xd8, 0x9b, 0x78, 0xd3, 0xe1, 0xec, 0x88, 0xfd, 0x94, 0xb3,
	0x05, 0xca, 0xa7, 0x52, 0x43, 0xbc, 0xef, 0x90, 0xed, 0x13, 0xbd, 0x27, 0xa3, 0x46, 0xa1, 0xa0,
	0x00, 0x35, 0xf6, 0x9d, 0xe3, 0x78, 0xd7, 0xb1, 0x09, 0x71, 0x9a, 0x47, 0x28, 0xc0, 0x24, 0xb6,
	0x8c, 0x07, 0xb5, 0x67, 0x51, 0x51, 0xf4, 0x6c, 0xb7, 0x45, 0x73, 0xbb, 0x1e, 0xb7, 0x26, 0xde,
	0xb4, 0xdf, 0xcc, 0x2d, 0xb9, 0x5d, 0xd3, 0x88, 0x8c, 0xb8, 0x10, 0x90, 0xe7, 0x4d, 0x72, 0xfb,
	0xef, 0xe4, 0xc1, 0x86, 0xa9, 0x9b, 0xcf, 0xbf, 0x48, 0xdc, 0x56, 0xc7, 0x6d, 0x35, 0x77, 0xd5,
	0xd8, 0xc5, 0x35, 0xe9, 0xd5, 0xc8, 0x7f, 0xd2, 0x7e, 0xc0, 0x0c, 0x82, 0x7f, 0x74, 0x8f, 0xf4,
	0x22, 0xcc, 0x72, 0x54, 0x10, 0x78, 0xd5, 0xeb, 0xbb, 0x44, 0x41, 0xe0, 0xd3, 0x3e, 0xe9, 0xdc,
	0x16, 0x90, 0xd9, 0xa0, 0x75, 0x73, 0x45, 0x0e, 0x04, 0xa6, 0xbf, 0xf4, 0x2c, 0xbd, 0x97, 0x96,
	0x42, 0xf9, 0xe1, 0xd3, 0xe7, 0x59, 0xcc, 0x4b, 0x16, 0x55, 0xdf, 0xe6, 0x5a, 0x57, 0x9d, 0xaf,
	0x5d, 0xf7, 0x77, 0x2e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x33, 0xd5, 0x99, 0x3f, 0xfd, 0x01,
	0x00, 0x00,
}
