// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// proto spec for the heartbeat
type HeartbeatMsg struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsRunning            bool     `protobuf:"varint,3,opt,name=isRunning,proto3" json:"isRunning,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatMsg) Reset()         { *m = HeartbeatMsg{} }
func (m *HeartbeatMsg) String() string { return proto.CompactTextString(m) }
func (*HeartbeatMsg) ProtoMessage()    {}
func (*HeartbeatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_7a8b8c1eff21deb6, []int{0}
}
func (m *HeartbeatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatMsg.Unmarshal(m, b)
}
func (m *HeartbeatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatMsg.Marshal(b, m, deterministic)
}
func (dst *HeartbeatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatMsg.Merge(dst, src)
}
func (m *HeartbeatMsg) XXX_Size() int {
	return xxx_messageInfo_HeartbeatMsg.Size(m)
}
func (m *HeartbeatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatMsg proto.InternalMessageInfo

func (m *HeartbeatMsg) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HeartbeatMsg) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HeartbeatMsg) GetIsRunning() bool {
	if m != nil {
		return m.IsRunning
	}
	return false
}

func (m *HeartbeatMsg) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A generic empty message that you can re-use to avoid defining duplicated empty messages in your APIs.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_7a8b8c1eff21deb6, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HeartbeatMsg)(nil), "api.HeartbeatMsg")
	proto.RegisterType((*Empty)(nil), "api.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	GetHeartbeat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatMsg, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) GetHeartbeat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatMsg, error) {
	out := new(HeartbeatMsg)
	err := c.cc.Invoke(ctx, "/api.Heartbeat/getHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	GetHeartbeat(context.Context, *Empty) (*HeartbeatMsg, error)
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_GetHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).GetHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Heartbeat/GetHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).GetHeartbeat(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getHeartbeat",
			Handler:    _Heartbeat_GetHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_heartbeat_7a8b8c1eff21deb6) }

var fileDescriptor_heartbeat_7a8b8c1eff21deb6 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0x8d, 0x55, 0x6b, 0xcf, 0x82, 0x98, 0x29, 0x88, 0x43, 0xe8, 0xd4, 0xa9, 0x82, 0xae,
	0xae, 0x82, 0x8b, 0x0e, 0x79, 0x83, 0x54, 0x8e, 0x98, 0x21, 0x6d, 0x68, 0x4e, 0xc1, 0x17, 0xf0,
	0xb9, 0xc5, 0x40, 0x6d, 0xb7, 0xfb, 0xbf, 0xff, 0x38, 0xbe, 0x83, 0xf5, 0x03, 0x75, 0x47, 0x35,
	0x6a, 0xaa, 0x7c, 0xd7, 0x52, 0xcb, 0x13, 0xed, 0x6d, 0xf1, 0x61, 0x90, 0x5f, 0xfa, 0xe2, 0x1a,
	0x0c, 0x97, 0xb0, 0x0a, 0xd8, 0xbd, 0xec, 0x1d, 0x6f, 0xda, 0xa1, 0x60, 0x92, 0x95, 0x99, 0x1a,
	0x23, 0xbe, 0x83, 0x8c, 0xac, 0xc3, 0x40, 0xda, 0x79, 0x31, 0x95, 0xac, 0x4c, 0xd4, 0x00, 0x7e,
	0xad, 0x0d, 0xea, 0xd9, 0x34, 0xb6, 0x31, 0x22, 0x91, 0xac, 0x5c, 0xaa, 0x01, 0x70, 0x01, 0xa9,
	0xc3, 0x10, 0xb4, 0x41, 0x31, 0x8b, 0x97, 0xfb, 0x58, 0xa4, 0x30, 0x3f, 0x3b, 0x4f, 0xef, 0xc3,
	0x09, 0xb2, 0xbf, 0x10, 0xdf, 0x43, 0x6e, 0x90, 0x86, 0x0c, 0x95, 0xf6, 0xb6, 0x8a, 0x8b, 0xdb,
	0x4d, 0x9c, 0xc7, 0xf2, 0xc5, 0xa4, 0x5e, 0xc4, 0xdf, 0x8e, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0x6c, 0xfe, 0x01, 0xee, 0x00, 0x00, 0x00,
}