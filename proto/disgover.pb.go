// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/disgover.proto

/*
Package disgover is a generated protocol buffer package.

It is generated from these files:
	proto/disgover.proto

It has these top-level messages:
	Endpoint
	Contact
	FindRequest
	Empty
*/
package disgover

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

type Endpoint struct {
	Host string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
	Port int64  `protobuf:"varint,2,opt,name=Port" json:"Port,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Contact struct {
	Id       []byte    `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Endpoint *Endpoint `protobuf:"bytes,2,opt,name=Endpoint" json:"Endpoint,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Contact) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Contact) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type FindRequest struct {
	ContactId []byte   `protobuf:"bytes,1,opt,name=ContactId,proto3" json:"ContactId,omitempty"`
	Sender    *Contact `protobuf:"bytes,2,opt,name=Sender" json:"Sender,omitempty"`
}

func (m *FindRequest) Reset()                    { *m = FindRequest{} }
func (m *FindRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()               {}
func (*FindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FindRequest) GetContactId() []byte {
	if m != nil {
		return m.ContactId
	}
	return nil
}

func (m *FindRequest) GetSender() *Contact {
	if m != nil {
		return m.Sender
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Endpoint)(nil), "disgover.Endpoint")
	proto.RegisterType((*Contact)(nil), "disgover.Contact")
	proto.RegisterType((*FindRequest)(nil), "disgover.FindRequest")
	proto.RegisterType((*Empty)(nil), "disgover.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DisgoverRPC service

type DisgoverRPCClient interface {
	PeerPingGrpc(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error)
	PeerFindGrpc(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Contact, error)
}

type disgoverRPCClient struct {
	cc *grpc.ClientConn
}

func NewDisgoverRPCClient(cc *grpc.ClientConn) DisgoverRPCClient {
	return &disgoverRPCClient{cc}
}

func (c *disgoverRPCClient) PeerPingGrpc(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := grpc.Invoke(ctx, "/disgover.DisgoverRPC/PeerPingGrpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disgoverRPCClient) PeerFindGrpc(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := grpc.Invoke(ctx, "/disgover.DisgoverRPC/PeerFindGrpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DisgoverRPC service

type DisgoverRPCServer interface {
	PeerPingGrpc(context.Context, *Contact) (*Contact, error)
	PeerFindGrpc(context.Context, *FindRequest) (*Contact, error)
}

func RegisterDisgoverRPCServer(s *grpc.Server, srv DisgoverRPCServer) {
	s.RegisterService(&_DisgoverRPC_serviceDesc, srv)
}

func _DisgoverRPC_PeerPingGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisgoverRPCServer).PeerPingGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/disgover.DisgoverRPC/PeerPingGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisgoverRPCServer).PeerPingGrpc(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisgoverRPC_PeerFindGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisgoverRPCServer).PeerFindGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/disgover.DisgoverRPC/PeerFindGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisgoverRPCServer).PeerFindGrpc(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DisgoverRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "disgover.DisgoverRPC",
	HandlerType: (*DisgoverRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PeerPingGrpc",
			Handler:    _DisgoverRPC_PeerPingGrpc_Handler,
		},
		{
			MethodName: "PeerFindGrpc",
			Handler:    _DisgoverRPC_PeerFindGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/disgover.proto",
}

func init() { proto.RegisterFile("proto/disgover.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x4d, 0xd5, 0xfd, 0x78, 0x1d, 0x82, 0x0f, 0x85, 0x21, 0x1e, 0x46, 0x4e, 0xf3, 0x52,
	0xa1, 0x82, 0x07, 0xaf, 0x73, 0x6a, 0x6f, 0x21, 0x82, 0x77, 0x5d, 0xc2, 0xc8, 0xc1, 0x24, 0xa6,
	0x4f, 0xc1, 0xa3, 0xff, 0xb9, 0x2c, 0xa6, 0x8d, 0xd0, 0xdd, 0xde, 0xfb, 0xf6, 0xfb, 0xf9, 0x34,
	0x09, 0x9c, 0xf9, 0xe0, 0xc8, 0x5d, 0x2b, 0xd3, 0x6e, 0xdd, 0x97, 0x0e, 0x55, 0x5c, 0x71, 0xd2,
	0xed, 0xbc, 0x86, 0xc9, 0xda, 0x2a, 0xef, 0x8c, 0x25, 0x44, 0x38, 0x7a, 0x72, 0x2d, 0xcd, 0xd9,
	0x82, 0x2d, 0xa7, 0x32, 0xce, 0xbb, 0x4c, 0xb8, 0x40, 0xf3, 0x62, 0xc1, 0x96, 0x87, 0x32, 0xce,
	0xbc, 0x81, 0xf1, 0xca, 0x59, 0x7a, 0xdd, 0x10, 0x9e, 0x40, 0xd1, 0xa8, 0x08, 0xcc, 0x64, 0xd1,
	0x28, 0xac, 0xb2, 0x2e, 0x22, 0x65, 0x8d, 0x55, 0xff, 0xef, 0xee, 0x8b, 0xec, 0x3b, 0xfc, 0x05,
	0xca, 0x07, 0x63, 0x95, 0xd4, 0x1f, 0x9f, 0xba, 0x25, 0xbc, 0x84, 0x69, 0x32, 0xf7, 0xd6, 0x1c,
	0xe0, 0x15, 0x8c, 0x9e, 0xb5, 0x55, 0x3a, 0x24, 0xf5, 0x69, 0x56, 0xa7, 0x92, 0x4c, 0x05, 0x3e,
	0x86, 0xe3, 0xf5, 0xbb, 0xa7, 0xef, 0xfa, 0x87, 0x41, 0x79, 0x9f, 0x5a, 0x52, 0xac, 0xf0, 0x16,
	0x66, 0x42, 0xeb, 0x20, 0x8c, 0xdd, 0x3e, 0x06, 0xbf, 0xc1, 0xa1, 0xe3, 0x62, 0x18, 0xf1, 0x03,
	0xbc, 0xfb, 0xe3, 0x76, 0x87, 0x8d, 0xdc, 0x79, 0x2e, 0xfd, 0xbb, 0xc0, 0x5e, 0xf6, 0x6d, 0x14,
	0x1f, 0xfd, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x46, 0x5a, 0x7c, 0x8c, 0x01, 0x00, 0x00,
}
