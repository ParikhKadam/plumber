// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_mqtt.proto

package args

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type MQTTQoSLevel int32

const (
	MQTTQoSLevel_MQTT_QOS_LEVEL_AT_MOST_ONCE  MQTTQoSLevel = 0
	MQTTQoSLevel_MQTT_QOS_LEVEL_AT_LEAST_ONCE MQTTQoSLevel = 1
	MQTTQoSLevel_MQTT_QOS_LEVEL_EXACTLY_ONCE  MQTTQoSLevel = 2
)

var MQTTQoSLevel_name = map[int32]string{
	0: "MQTT_QOS_LEVEL_AT_MOST_ONCE",
	1: "MQTT_QOS_LEVEL_AT_LEAST_ONCE",
	2: "MQTT_QOS_LEVEL_EXACTLY_ONCE",
}

var MQTTQoSLevel_value = map[string]int32{
	"MQTT_QOS_LEVEL_AT_MOST_ONCE":  0,
	"MQTT_QOS_LEVEL_AT_LEAST_ONCE": 1,
	"MQTT_QOS_LEVEL_EXACTLY_ONCE":  2,
}

func (x MQTTQoSLevel) String() string {
	return proto.EnumName(MQTTQoSLevel_name, int32(x))
}

func (MQTTQoSLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{0}
}

type MQTTTLSOptions struct {
	// @gotags: kong:"help='CA file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CA_FILE,type=existingfile"
	CaFile string `protobuf:"bytes,1,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty" kong:"help='CA file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CA_FILE,type=existingfile"`
	// @gotags: kong:"help='Client cert file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CERT_FILE,type=existingfile"
	CertFile string `protobuf:"bytes,2,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty" kong:"help='Client cert file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_CERT_FILE,type=existingfile"`
	// @gotags: kong:"help='Client key file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_KEY_FILE,type=existingfile"
	KeyFile string `protobuf:"bytes,3,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty" kong:"help='Client key file (only needed if addr is ssl://)',env=PLUMBER_RELAY_MQTT_TLS_KEY_FILE,type=existingfile"`
	// @gotags: kong:"help='Whether to verify server certificate',env=PLUMBER_RELAY_MQTT_SKIP_VERIFY_TLS"
	SkipVerify           bool     `protobuf:"varint,4,opt,name=skip_verify,json=skipVerify,proto3" json:"skip_verify,omitempty" kong:"help='Whether to verify server certificate',env=PLUMBER_RELAY_MQTT_SKIP_VERIFY_TLS"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTTLSOptions) Reset()         { *m = MQTTTLSOptions{} }
func (m *MQTTTLSOptions) String() string { return proto.CompactTextString(m) }
func (*MQTTTLSOptions) ProtoMessage()    {}
func (*MQTTTLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{0}
}

func (m *MQTTTLSOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTTLSOptions.Unmarshal(m, b)
}
func (m *MQTTTLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTTLSOptions.Marshal(b, m, deterministic)
}
func (m *MQTTTLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTTLSOptions.Merge(m, src)
}
func (m *MQTTTLSOptions) XXX_Size() int {
	return xxx_messageInfo_MQTTTLSOptions.Size(m)
}
func (m *MQTTTLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTTLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTTLSOptions proto.InternalMessageInfo

func (m *MQTTTLSOptions) GetCaFile() string {
	if m != nil {
		return m.CaFile
	}
	return ""
}

func (m *MQTTTLSOptions) GetCertFile() string {
	if m != nil {
		return m.CertFile
	}
	return ""
}

func (m *MQTTTLSOptions) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

func (m *MQTTTLSOptions) GetSkipVerify() bool {
	if m != nil {
		return m.SkipVerify
	}
	return false
}

type MQTTConn struct {
	// @gotags: kong:"help='MQTT address',default='tcp://localhost:1883',env='PLUMBER_RELAY_MQTT_ADDRESS',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='MQTT address',default='tcp://localhost:1883',env='PLUMBER_RELAY_MQTT_ADDRESS',required"`
	// @gotags: kong:"help='How long to attempt to connect for',env='PLUMBER_RELAY_MQTT_CONNECT_TIMEOUT',default=5"
	ConnTimeoutSeconds uint32 `protobuf:"varint,3,opt,name=conn_timeout_seconds,json=connTimeoutSeconds,proto3" json:"conn_timeout_seconds,omitempty" kong:"help='How long to attempt to connect for',env='PLUMBER_RELAY_MQTT_CONNECT_TIMEOUT',default=5"`
	// @gotags: kong:"help='Client id presented to MQTT broker',env='PLUMBER_RELAY_MQTT_CLIENT_ID',default=plumber"
	ClientId string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" kong:"help='Client id presented to MQTT broker',env='PLUMBER_RELAY_MQTT_CLIENT_ID',default=plumber"`
	// @gotags: kong:"help='QoS level to use for pub/sub (options: at_most_once, at_least_once, exactly_once)',env=PLUMBER_RELAY_MQTT_QOS,type=pbenum,pbenum_strip_prefix=MQTT_QOS_LEVEL_,pbenum_lowercase,default=at_most_once"
	QosLevel MQTTQoSLevel `protobuf:"varint,5,opt,name=qos_level,json=qosLevel,proto3,enum=protos.args.MQTTQoSLevel" json:"qos_level,omitempty" kong:"help='QoS level to use for pub/sub (options: at_most_once, at_least_once, exactly_once)',env=PLUMBER_RELAY_MQTT_QOS,type=pbenum,pbenum_strip_prefix=MQTT_QOS_LEVEL_,pbenum_lowercase,default=at_most_once"`
	// @gotags: kong:"embed"
	TlsOptions           *MQTTTLSOptions `protobuf:"bytes,6,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty" kong:"embed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MQTTConn) Reset()         { *m = MQTTConn{} }
func (m *MQTTConn) String() string { return proto.CompactTextString(m) }
func (*MQTTConn) ProtoMessage()    {}
func (*MQTTConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{1}
}

func (m *MQTTConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTConn.Unmarshal(m, b)
}
func (m *MQTTConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTConn.Marshal(b, m, deterministic)
}
func (m *MQTTConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTConn.Merge(m, src)
}
func (m *MQTTConn) XXX_Size() int {
	return xxx_messageInfo_MQTTConn.Size(m)
}
func (m *MQTTConn) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTConn.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTConn proto.InternalMessageInfo

func (m *MQTTConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MQTTConn) GetConnTimeoutSeconds() uint32 {
	if m != nil {
		return m.ConnTimeoutSeconds
	}
	return 0
}

func (m *MQTTConn) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *MQTTConn) GetQosLevel() MQTTQoSLevel {
	if m != nil {
		return m.QosLevel
	}
	return MQTTQoSLevel_MQTT_QOS_LEVEL_AT_MOST_ONCE
}

func (m *MQTTConn) GetTlsOptions() *MQTTTLSOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

type MQTTReadArgs struct {
	// @gotags: kong:"help='Topic to read message(s) from',env='PLUMBER_RELAY_MQTT_TOPIC',required"
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic to read message(s) from',env='PLUMBER_RELAY_MQTT_TOPIC',required"`
	// @gotags: kong:"help='How long to attempt to read message(s)',default=0,env='PLUMBER_RELAY_MQTT_READ_TIMEOUT_SECONDS'"
	ReadTimeoutSeconds   uint32   `protobuf:"varint,2,opt,name=read_timeout_seconds,json=readTimeoutSeconds,proto3" json:"read_timeout_seconds,omitempty" kong:"help='How long to attempt to read message(s)',default=0,env='PLUMBER_RELAY_MQTT_READ_TIMEOUT_SECONDS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTReadArgs) Reset()         { *m = MQTTReadArgs{} }
func (m *MQTTReadArgs) String() string { return proto.CompactTextString(m) }
func (*MQTTReadArgs) ProtoMessage()    {}
func (*MQTTReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{2}
}

func (m *MQTTReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTReadArgs.Unmarshal(m, b)
}
func (m *MQTTReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTReadArgs.Marshal(b, m, deterministic)
}
func (m *MQTTReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTReadArgs.Merge(m, src)
}
func (m *MQTTReadArgs) XXX_Size() int {
	return xxx_messageInfo_MQTTReadArgs.Size(m)
}
func (m *MQTTReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTReadArgs proto.InternalMessageInfo

func (m *MQTTReadArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MQTTReadArgs) GetReadTimeoutSeconds() uint32 {
	if m != nil {
		return m.ReadTimeoutSeconds
	}
	return 0
}

type MQTTWriteArgs struct {
	// @gotags: kong:"help='Topic to write message(s) to',required"
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic to write message(s) to',required"`
	// @gotags: kong:"help='How long to attempt to publish message(s)',default=5"
	WriteTimeoutSeconds  uint32   `protobuf:"varint,2,opt,name=write_timeout_seconds,json=writeTimeoutSeconds,proto3" json:"write_timeout_seconds,omitempty" kong:"help='How long to attempt to publish message(s)',default=5"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MQTTWriteArgs) Reset()         { *m = MQTTWriteArgs{} }
func (m *MQTTWriteArgs) String() string { return proto.CompactTextString(m) }
func (*MQTTWriteArgs) ProtoMessage()    {}
func (*MQTTWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_11ff87d0ae52f3d2, []int{3}
}

func (m *MQTTWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MQTTWriteArgs.Unmarshal(m, b)
}
func (m *MQTTWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MQTTWriteArgs.Marshal(b, m, deterministic)
}
func (m *MQTTWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MQTTWriteArgs.Merge(m, src)
}
func (m *MQTTWriteArgs) XXX_Size() int {
	return xxx_messageInfo_MQTTWriteArgs.Size(m)
}
func (m *MQTTWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MQTTWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MQTTWriteArgs proto.InternalMessageInfo

func (m *MQTTWriteArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MQTTWriteArgs) GetWriteTimeoutSeconds() uint32 {
	if m != nil {
		return m.WriteTimeoutSeconds
	}
	return 0
}

func init() {
	proto.RegisterEnum("protos.args.MQTTQoSLevel", MQTTQoSLevel_name, MQTTQoSLevel_value)
	proto.RegisterType((*MQTTTLSOptions)(nil), "protos.args.MQTTTLSOptions")
	proto.RegisterType((*MQTTConn)(nil), "protos.args.MQTTConn")
	proto.RegisterType((*MQTTReadArgs)(nil), "protos.args.MQTTReadArgs")
	proto.RegisterType((*MQTTWriteArgs)(nil), "protos.args.MQTTWriteArgs")
}

func init() { proto.RegisterFile("ps_args_mqtt.proto", fileDescriptor_11ff87d0ae52f3d2) }

var fileDescriptor_11ff87d0ae52f3d2 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x4d, 0x75, 0xfb, 0xe7, 0xd6, 0x5d, 0x64, 0x5c, 0x31, 0x4b, 0x85, 0x0d, 0x7d, 0x2a,
	0x82, 0x89, 0x54, 0x10, 0x44, 0x5f, 0x6a, 0x89, 0x20, 0x64, 0x2d, 0x4d, 0x42, 0x75, 0x7d, 0x19,
	0x92, 0xc9, 0x6c, 0x3b, 0x34, 0xc9, 0xa4, 0x33, 0xd3, 0x95, 0xbe, 0xfa, 0x6d, 0xfd, 0x16, 0x32,
	0x33, 0x2d, 0xba, 0x54, 0xf7, 0x29, 0x73, 0xef, 0xef, 0xe6, 0xde, 0x73, 0x0e, 0xa0, 0x46, 0xe2,
	0x4c, 0x2c, 0x25, 0xae, 0x36, 0x4a, 0xf9, 0x8d, 0xe0, 0x8a, 0xa3, 0xbe, 0xf9, 0x48, 0x5f, 0xf7,
	0x87, 0x3f, 0x1d, 0x38, 0xbb, 0x9a, 0xa7, 0x69, 0x1a, 0x25, 0xb3, 0x46, 0x31, 0x5e, 0x4b, 0xf4,
	0x1c, 0x3a, 0x24, 0xc3, 0x37, 0xac, 0xa4, 0xae, 0xe3, 0x39, 0xa3, 0x5e, 0xdc, 0x26, 0xd9, 0x27,
	0x56, 0x52, 0x34, 0x80, 0x1e, 0xa1, 0x42, 0x59, 0xd4, 0x32, 0xa8, 0xab, 0x1b, 0x06, 0x5e, 0x40,
	0x77, 0x4d, 0x77, 0x96, 0x3d, 0x34, 0xac, 0xb3, 0xa6, 0x3b, 0x83, 0x2e, 0xa1, 0x2f, 0xd7, 0xac,
	0xc1, 0xb7, 0x54, 0xb0, 0x9b, 0x9d, 0xfb, 0xc8, 0x73, 0x46, 0xdd, 0x18, 0x74, 0x6b, 0x61, 0x3a,
	0xc3, 0x5f, 0x0e, 0x74, 0xb5, 0x88, 0x29, 0xaf, 0x6b, 0xe4, 0x42, 0x27, 0x2b, 0x0a, 0x41, 0xa5,
	0xdc, 0x9f, 0x3f, 0x94, 0xe8, 0x35, 0x9c, 0x13, 0x5e, 0xd7, 0x58, 0xb1, 0x8a, 0xf2, 0xad, 0xc2,
	0x92, 0x12, 0x5e, 0x17, 0xd2, 0x9c, 0x3b, 0x8d, 0x91, 0x66, 0xa9, 0x45, 0x89, 0x25, 0x46, 0x71,
	0xc9, 0x68, 0xad, 0x30, 0x2b, 0xcc, 0x5d, 0xad, 0xd8, 0x34, 0x3e, 0x17, 0xe8, 0x2d, 0xf4, 0x36,
	0x5c, 0xe2, 0x92, 0xde, 0xd2, 0xd2, 0x3d, 0xf1, 0x9c, 0xd1, 0xd9, 0xf8, 0xc2, 0xff, 0x2b, 0x1b,
	0x5f, 0x4b, 0x9a, 0xf3, 0x24, 0xd2, 0x03, 0x71, 0x77, 0xc3, 0xa5, 0x79, 0xa1, 0x0f, 0xd0, 0x57,
	0xa5, 0xc4, 0xdc, 0xc6, 0xe5, 0xb6, 0x3d, 0x67, 0xd4, 0x1f, 0x0f, 0x8e, 0xfe, 0xfc, 0x93, 0x68,
	0x0c, 0xaa, 0x94, 0xfb, 0xf7, 0x70, 0x01, 0x8f, 0x35, 0x8d, 0x69, 0x56, 0x4c, 0xc4, 0x52, 0xa2,
	0x73, 0x38, 0x51, 0xbc, 0x61, 0x64, 0x6f, 0xd6, 0x16, 0xda, 0xaa, 0xa0, 0x59, 0x71, 0x64, 0xb5,
	0x65, 0xad, 0x6a, 0x76, 0xd7, 0xea, 0xf0, 0x1a, 0x4e, 0xf5, 0xde, 0xaf, 0x82, 0x29, 0x7a, 0xcf,
	0xe2, 0x31, 0x3c, 0xfb, 0xa1, 0x47, 0xfe, 0xb3, 0xf9, 0xa9, 0x81, 0x77, 0x57, 0xbf, 0x14, 0x56,
	0xf2, 0x21, 0x0a, 0x74, 0x09, 0x03, 0x5d, 0xe3, 0xf9, 0x2c, 0xc1, 0x51, 0xb8, 0x08, 0x23, 0x3c,
	0x49, 0xf1, 0xd5, 0x2c, 0x49, 0xf1, 0xec, 0xcb, 0x34, 0x7c, 0xf2, 0x00, 0x79, 0xf0, 0xe2, 0x78,
	0x20, 0x0a, 0x27, 0x87, 0x09, 0xe7, 0x1f, 0x2b, 0xc2, 0x6f, 0x93, 0x69, 0x1a, 0x5d, 0xdb, 0x81,
	0xd6, 0xc7, 0xf7, 0xdf, 0xdf, 0x2d, 0x99, 0x5a, 0x6d, 0x73, 0x9f, 0xf0, 0x2a, 0xc8, 0x33, 0x45,
	0x56, 0x84, 0x8b, 0x26, 0x68, 0xca, 0x6d, 0x95, 0x53, 0xf1, 0x4a, 0x92, 0x15, 0xad, 0x32, 0x19,
	0xe4, 0x5b, 0x56, 0x16, 0xc1, 0x92, 0x07, 0x36, 0xfe, 0x40, 0xc7, 0x9f, 0xb7, 0x4d, 0xf1, 0xe6,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xc4, 0x17, 0xfe, 0xfe, 0x02, 0x00, 0x00,
}
