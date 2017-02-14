// Code generated by protoc-gen-go.
// source: proto/gen.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/gen.proto
	proto/word.proto

It has these top-level messages:
	GenRequest
	GenResponse
	WordRequest
	WordResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GenRequest struct {
	Text  string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Limit int32  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *GenRequest) Reset()                    { *m = GenRequest{} }
func (m *GenRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenRequest) ProtoMessage()               {}
func (*GenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GenRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *GenRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GenResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *GenResponse) Reset()                    { *m = GenResponse{} }
func (m *GenResponse) String() string            { return proto1.CompactTextString(m) }
func (*GenResponse) ProtoMessage()               {}
func (*GenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GenResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto1.RegisterType((*GenRequest)(nil), "proto.GenRequest")
	proto1.RegisterType((*GenResponse)(nil), "proto.GenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GenSvc service

type GenSvcClient interface {
	Gen(ctx context.Context, in *GenRequest, opts ...grpc.CallOption) (GenSvc_GenClient, error)
}

type genSvcClient struct {
	cc *grpc.ClientConn
}

func NewGenSvcClient(cc *grpc.ClientConn) GenSvcClient {
	return &genSvcClient{cc}
}

func (c *genSvcClient) Gen(ctx context.Context, in *GenRequest, opts ...grpc.CallOption) (GenSvc_GenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GenSvc_serviceDesc.Streams[0], c.cc, "/proto.GenSvc/Gen", opts...)
	if err != nil {
		return nil, err
	}
	x := &genSvcGenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GenSvc_GenClient interface {
	Recv() (*GenResponse, error)
	grpc.ClientStream
}

type genSvcGenClient struct {
	grpc.ClientStream
}

func (x *genSvcGenClient) Recv() (*GenResponse, error) {
	m := new(GenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GenSvc service

type GenSvcServer interface {
	Gen(*GenRequest, GenSvc_GenServer) error
}

func RegisterGenSvcServer(s *grpc.Server, srv GenSvcServer) {
	s.RegisterService(&_GenSvc_serviceDesc, srv)
}

func _GenSvc_Gen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GenSvcServer).Gen(m, &genSvcGenServer{stream})
}

type GenSvc_GenServer interface {
	Send(*GenResponse) error
	grpc.ServerStream
}

type genSvcGenServer struct {
	grpc.ServerStream
}

func (x *genSvcGenServer) Send(m *GenResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GenSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GenSvc",
	HandlerType: (*GenSvcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Gen",
			Handler:       _GenSvc_Gen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gen.proto",
}

func init() { proto1.RegisterFile("proto/gen.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4f, 0xcd, 0xd3, 0x03, 0xb3, 0x84, 0x58, 0xc1, 0x94, 0x92, 0x19, 0x17, 0x97,
	0x7b, 0x6a, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a,
	0x45, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x93,
	0x99, 0x9b, 0x59, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe1, 0x28, 0x29, 0x72, 0x71,
	0x83, 0xf5, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x62, 0xd3, 0x68, 0x64, 0xc5, 0xc5, 0xe6, 0x9e,
	0x9a, 0x17, 0x5c, 0x96, 0x2c, 0x64, 0xc0, 0xc5, 0xec, 0x9e, 0x9a, 0x27, 0x24, 0x08, 0xb1, 0x5a,
	0x0f, 0x61, 0xa1, 0x94, 0x10, 0xb2, 0x10, 0xc4, 0x2c, 0x25, 0x06, 0x03, 0xc6, 0x24, 0x36, 0xb0,
	0xb0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x65, 0x95, 0x3c, 0xf9, 0xb7, 0x00, 0x00, 0x00,
}