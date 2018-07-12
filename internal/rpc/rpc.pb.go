// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/rpc/rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	internal/rpc/rpc.proto

It has these top-level messages:
	Message
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"

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

type Message struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "rpc.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GrafoRPC service

type GrafoRPCClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type grafoRPCClient struct {
	cc *grpc.ClientConn
}

func NewGrafoRPCClient(cc *grpc.ClientConn) GrafoRPCClient {
	return &grafoRPCClient{cc}
}

func (c *grafoRPCClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/rpc.GrafoRPC/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GrafoRPC service

type GrafoRPCServer interface {
	SendMessage(context.Context, *Message) (*Message, error)
}

func RegisterGrafoRPCServer(s *grpc.Server, srv GrafoRPCServer) {
	s.RegisterService(&_GrafoRPC_serviceDesc, srv)
}

func _GrafoRPC_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafoRPCServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.GrafoRPC/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafoRPCServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrafoRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.GrafoRPC",
	HandlerType: (*GrafoRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _GrafoRPC_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/rpc/rpc.proto",
}

func init() { proto.RegisterFile("internal/rpc/rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2a, 0x48, 0x06, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0x21, 0xe6, 0xa2, 0x82, 0x64, 0x25, 0x65, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54,
	0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x9e, 0x20, 0x18, 0xd7, 0xc8, 0x9c, 0x8b, 0xc3, 0xbd, 0x28, 0x31, 0x2d, 0x3f, 0x28, 0xc0, 0x59,
	0x48, 0x9b, 0x8b, 0x3b, 0x38, 0x35, 0x2f, 0x05, 0xa6, 0x89, 0x47, 0x0f, 0x64, 0x20, 0x94, 0x27,
	0x85, 0xc2, 0x53, 0x62, 0x48, 0x62, 0x03, 0xdb, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x30, 0xc8, 0x8f, 0x83, 0x00, 0x00, 0x00,
}
