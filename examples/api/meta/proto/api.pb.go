// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/xxdawn/micro/examples/api/meta/proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/xxdawn/micro/examples/api/meta/proto/api.proto

It has these top-level messages:
	CallRequest
	CallResponse
	EmptyRequest
	EmptyResponse
*/
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

type CallRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CallRequest) Reset()                    { *m = CallRequest{} }
func (m *CallRequest) String() string            { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()               {}
func (*CallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CallRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallResponse struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *CallResponse) Reset()                    { *m = CallResponse{} }
func (m *CallResponse) String() string            { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()               {}
func (*CallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CallResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CallRequest)(nil), "CallRequest")
	proto.RegisterType((*CallResponse)(nil), "CallResponse")
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := grpc.Invoke(ctx, "/Example/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleServer interface {
	Call(context.Context, *CallRequest) (*CallResponse, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Example/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Example_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/xxdawn/micro/examples/api/meta/proto/api.proto",
}

// Client API for Foo service

type FooClient interface {
	Bar(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type fooClient struct {
	cc *grpc.ClientConn
}

func NewFooClient(cc *grpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Bar(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/Foo/Bar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Foo service

type FooServer interface {
	Bar(context.Context, *EmptyRequest) (*EmptyResponse, error)
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Foo/Bar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Bar(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Foo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bar",
			Handler:    _Foo_Bar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/xxdawn/micro/examples/api/meta/proto/api.proto",
}

func init() { proto.RegisterFile("github.com/xxdawn/micro/examples/api/meta/proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xd1, 0x6a, 0x83, 0x30,
	0x14, 0x86, 0x75, 0xca, 0x64, 0x67, 0xea, 0x20, 0x57, 0xe2, 0xd5, 0x16, 0xd8, 0xf0, 0x66, 0xc9,
	0x70, 0x6f, 0xd0, 0x62, 0x1f, 0xc0, 0x37, 0x88, 0x72, 0xb0, 0x82, 0x31, 0xa9, 0x89, 0xd0, 0xbe,
	0x7d, 0x31, 0x5a, 0xb0, 0x77, 0xff, 0x07, 0x7f, 0xbe, 0x3f, 0x07, 0xca, 0xae, 0xb7, 0xe7, 0xb9,
	0x61, 0xad, 0x92, 0x5c, 0xf6, 0xed, 0xa4, 0x38, 0x5e, 0x85, 0xd4, 0x03, 0x1a, 0x2e, 0x74, 0xcf,
	0x25, 0x5a, 0xc1, 0xf5, 0xa4, 0xac, 0x5a, 0x90, 0xb9, 0x44, 0xbf, 0xe0, 0xfd, 0x28, 0x86, 0xa1,
	0xc6, 0xcb, 0x8c, 0xc6, 0x12, 0x02, 0xe1, 0x28, 0x24, 0x66, 0xfe, 0xa7, 0x5f, 0xbc, 0xd5, 0x2e,
	0xd3, 0x02, 0xe2, 0xb5, 0x62, 0xb4, 0x1a, 0x0d, 0x92, 0x0c, 0x22, 0x89, 0xc6, 0x88, 0x0e, 0xb3,
	0x17, 0x57, 0x7b, 0x20, 0x4d, 0x21, 0xae, 0xa4, 0xb6, 0xb7, 0xcd, 0x46, 0x3f, 0x20, 0xd9, 0x78,
	0x7d, 0x5a, 0xfe, 0x41, 0x54, 0xad, 0x5f, 0x22, 0xdf, 0x10, 0x2e, 0x56, 0x12, 0xb3, 0xdd, 0x7e,
	0x9e, 0xb0, 0xfd, 0x14, 0xf5, 0xca, 0x5f, 0x08, 0x4e, 0x4a, 0x91, 0x1f, 0x08, 0x0e, 0x62, 0x22,
	0x09, 0xdb, 0xfb, 0xf3, 0x94, 0x3d, 0xe9, 0xa9, 0xd7, 0xbc, 0xba, 0xab, 0xfe, 0xef, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x14, 0x42, 0xb8, 0x1d, 0x0b, 0x01, 0x00, 0x00,
}
