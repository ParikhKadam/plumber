// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_azure_event_hub.proto

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

type AzureEventHubConn struct {
	// @gotags: kong:"help='Connection string',env='EVENTHUB_CONNECTION_STRING',required"
	ConnectionString     string   `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty" kong:"help='Connection string',env='EVENTHUB_CONNECTION_STRING',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureEventHubConn) Reset()         { *m = AzureEventHubConn{} }
func (m *AzureEventHubConn) String() string { return proto.CompactTextString(m) }
func (*AzureEventHubConn) ProtoMessage()    {}
func (*AzureEventHubConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b95bd39efe3a29, []int{0}
}

func (m *AzureEventHubConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureEventHubConn.Unmarshal(m, b)
}
func (m *AzureEventHubConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureEventHubConn.Marshal(b, m, deterministic)
}
func (m *AzureEventHubConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureEventHubConn.Merge(m, src)
}
func (m *AzureEventHubConn) XXX_Size() int {
	return xxx_messageInfo_AzureEventHubConn.Size(m)
}
func (m *AzureEventHubConn) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureEventHubConn.DiscardUnknown(m)
}

var xxx_messageInfo_AzureEventHubConn proto.InternalMessageInfo

func (m *AzureEventHubConn) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

type AzureEventHubReadArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureEventHubReadArgs) Reset()         { *m = AzureEventHubReadArgs{} }
func (m *AzureEventHubReadArgs) String() string { return proto.CompactTextString(m) }
func (*AzureEventHubReadArgs) ProtoMessage()    {}
func (*AzureEventHubReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b95bd39efe3a29, []int{1}
}

func (m *AzureEventHubReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureEventHubReadArgs.Unmarshal(m, b)
}
func (m *AzureEventHubReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureEventHubReadArgs.Marshal(b, m, deterministic)
}
func (m *AzureEventHubReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureEventHubReadArgs.Merge(m, src)
}
func (m *AzureEventHubReadArgs) XXX_Size() int {
	return xxx_messageInfo_AzureEventHubReadArgs.Size(m)
}
func (m *AzureEventHubReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureEventHubReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AzureEventHubReadArgs proto.InternalMessageInfo

type AzureEventHubWriteArgs struct {
	// @gotags: kong:"help='Send message with this ID'"
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty" kong:"help='Send message with this ID'"`
	// @gotags: kong:"help='Send message with this partition key'"
	PartitionKey         string   `protobuf:"bytes,2,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty" kong:"help='Send message with this partition key'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureEventHubWriteArgs) Reset()         { *m = AzureEventHubWriteArgs{} }
func (m *AzureEventHubWriteArgs) String() string { return proto.CompactTextString(m) }
func (*AzureEventHubWriteArgs) ProtoMessage()    {}
func (*AzureEventHubWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b95bd39efe3a29, []int{2}
}

func (m *AzureEventHubWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureEventHubWriteArgs.Unmarshal(m, b)
}
func (m *AzureEventHubWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureEventHubWriteArgs.Marshal(b, m, deterministic)
}
func (m *AzureEventHubWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureEventHubWriteArgs.Merge(m, src)
}
func (m *AzureEventHubWriteArgs) XXX_Size() int {
	return xxx_messageInfo_AzureEventHubWriteArgs.Size(m)
}
func (m *AzureEventHubWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureEventHubWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AzureEventHubWriteArgs proto.InternalMessageInfo

func (m *AzureEventHubWriteArgs) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *AzureEventHubWriteArgs) GetPartitionKey() string {
	if m != nil {
		return m.PartitionKey
	}
	return ""
}

func init() {
	proto.RegisterType((*AzureEventHubConn)(nil), "protos.args.AzureEventHubConn")
	proto.RegisterType((*AzureEventHubReadArgs)(nil), "protos.args.AzureEventHubReadArgs")
	proto.RegisterType((*AzureEventHubWriteArgs)(nil), "protos.args.AzureEventHubWriteArgs")
}

func init() { proto.RegisterFile("ps_args_azure_event_hub.proto", fileDescriptor_16b95bd39efe3a29) }

var fileDescriptor_16b95bd39efe3a29 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xa9, 0x07, 0xa1, 0x51, 0xc1, 0x2e, 0xa8, 0xbd, 0x14, 0xa4, 0x5e, 0x04, 0xb1, 0x39,
	0x78, 0x12, 0x2f, 0x56, 0x11, 0x14, 0x6f, 0xf5, 0x20, 0x88, 0x10, 0x92, 0xec, 0x90, 0x0d, 0x76,
	0x93, 0x30, 0x33, 0x11, 0xea, 0xaf, 0x97, 0xc4, 0x62, 0xd9, 0x53, 0xc8, 0x7b, 0x1f, 0x6f, 0xf8,
	0xc4, 0x2c, 0x91, 0xd2, 0xe8, 0x48, 0xe9, 0x9f, 0x8c, 0xa0, 0xe0, 0x1b, 0x02, 0xab, 0x2e, 0x9b,
	0x45, 0xc2, 0xc8, 0xb1, 0x39, 0xa8, 0x0f, 0x2d, 0x0a, 0x32, 0xbf, 0x17, 0x93, 0x65, 0xa1, 0x9e,
	0x0a, 0xf4, 0x9c, 0xcd, 0x63, 0x0c, 0xa1, 0xb9, 0x12, 0x13, 0x1b, 0x43, 0x00, 0xcb, 0x3e, 0x06,
	0x45, 0x8c, 0x3e, 0xb8, 0xe9, 0xe8, 0x7c, 0x74, 0x39, 0x5e, 0x1d, 0xef, 0x8a, 0xb7, 0x9a, 0xcf,
	0xcf, 0xc4, 0xc9, 0x60, 0x61, 0x05, 0xba, 0x5d, 0x96, 0xe9, 0x4f, 0x71, 0x3a, 0x28, 0xde, 0xd1,
	0x33, 0x94, 0xa6, 0x99, 0x09, 0xd1, 0x03, 0x91, 0x76, 0xa0, 0x7c, 0xbb, 0x1d, 0x1e, 0x6f, 0x93,
	0x97, 0xb6, 0xb9, 0x10, 0x47, 0x49, 0x23, 0xfb, 0x7a, 0xfd, 0x0b, 0x36, 0xd3, 0xbd, 0x4a, 0x1c,
	0xfe, 0x87, 0xaf, 0xb0, 0x79, 0xb8, 0xfb, 0xb8, 0x75, 0x9e, 0x8b, 0x95, 0x8d, 0xbd, 0x34, 0x9a,
	0x6d, 0x67, 0x23, 0x26, 0x99, 0xd6, 0xb9, 0x37, 0x80, 0xd7, 0x64, 0x3b, 0xe8, 0x35, 0x49, 0x93,
	0xfd, 0xba, 0x95, 0x2e, 0xca, 0x3f, 0x6b, 0x59, 0xac, 0xcd, 0x7e, 0xfd, 0xdc, 0xfc, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x54, 0x1b, 0x42, 0x69, 0x2a, 0x01, 0x00, 0x00,
}