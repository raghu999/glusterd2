// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/peers/peer-rpc.proto

/*
Package peercommands is a generated protocol buffer package.

It is generated from these files:
	commands/peers/peer-rpc.proto

It has these top-level messages:
	StoreConfig
	JoinReq
	JoinRsp
	LeaveReq
	LeaveRsp
*/
package peercommands

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

type StoreConfig struct {
	Endpoints []string `protobuf:"bytes,1,rep,name=Endpoints" json:"Endpoints,omitempty"`
}

func (m *StoreConfig) Reset()                    { *m = StoreConfig{} }
func (m *StoreConfig) String() string            { return proto.CompactTextString(m) }
func (*StoreConfig) ProtoMessage()               {}
func (*StoreConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreConfig) GetEndpoints() []string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type JoinReq struct {
	PeerID string       `protobuf:"bytes,1,opt,name=PeerID" json:"PeerID,omitempty"`
	Config *StoreConfig `protobuf:"bytes,2,opt,name=Config" json:"Config,omitempty"`
}

func (m *JoinReq) Reset()                    { *m = JoinReq{} }
func (m *JoinReq) String() string            { return proto.CompactTextString(m) }
func (*JoinReq) ProtoMessage()               {}
func (*JoinReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JoinReq) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *JoinReq) GetConfig() *StoreConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type JoinRsp struct {
	PeerID string `protobuf:"bytes,1,opt,name=PeerID" json:"PeerID,omitempty"`
	Err    int32  `protobuf:"varint,2,opt,name=Err" json:"Err,omitempty"`
}

func (m *JoinRsp) Reset()                    { *m = JoinRsp{} }
func (m *JoinRsp) String() string            { return proto.CompactTextString(m) }
func (*JoinRsp) ProtoMessage()               {}
func (*JoinRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JoinRsp) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *JoinRsp) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

type LeaveReq struct {
	PeerID string `protobuf:"bytes,1,opt,name=PeerID" json:"PeerID,omitempty"`
}

func (m *LeaveReq) Reset()                    { *m = LeaveReq{} }
func (m *LeaveReq) String() string            { return proto.CompactTextString(m) }
func (*LeaveReq) ProtoMessage()               {}
func (*LeaveReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LeaveReq) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type LeaveRsp struct {
	Err int32 `protobuf:"varint,1,opt,name=Err" json:"Err,omitempty"`
}

func (m *LeaveRsp) Reset()                    { *m = LeaveRsp{} }
func (m *LeaveRsp) String() string            { return proto.CompactTextString(m) }
func (*LeaveRsp) ProtoMessage()               {}
func (*LeaveRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LeaveRsp) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreConfig)(nil), "peercommands.StoreConfig")
	proto.RegisterType((*JoinReq)(nil), "peercommands.JoinReq")
	proto.RegisterType((*JoinRsp)(nil), "peercommands.JoinRsp")
	proto.RegisterType((*LeaveReq)(nil), "peercommands.LeaveReq")
	proto.RegisterType((*LeaveRsp)(nil), "peercommands.LeaveRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PeerService service

type PeerServiceClient interface {
	Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (*JoinRsp, error)
	Leave(ctx context.Context, in *LeaveReq, opts ...grpc.CallOption) (*LeaveRsp, error)
}

type peerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPeerServiceClient(cc *grpc.ClientConn) PeerServiceClient {
	return &peerServiceClient{cc}
}

func (c *peerServiceClient) Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (*JoinRsp, error) {
	out := new(JoinRsp)
	err := grpc.Invoke(ctx, "/peercommands.PeerService/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerServiceClient) Leave(ctx context.Context, in *LeaveReq, opts ...grpc.CallOption) (*LeaveRsp, error) {
	out := new(LeaveRsp)
	err := grpc.Invoke(ctx, "/peercommands.PeerService/Leave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PeerService service

type PeerServiceServer interface {
	Join(context.Context, *JoinReq) (*JoinRsp, error)
	Leave(context.Context, *LeaveReq) (*LeaveRsp, error)
}

func RegisterPeerServiceServer(s *grpc.Server, srv PeerServiceServer) {
	s.RegisterService(&_PeerService_serviceDesc, srv)
}

func _PeerService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peercommands.PeerService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).Join(ctx, req.(*JoinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerService_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peercommands.PeerService/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).Leave(ctx, req.(*LeaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peercommands.PeerService",
	HandlerType: (*PeerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _PeerService_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _PeerService_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands/peers/peer-rpc.proto",
}

func init() { proto.RegisterFile("commands/peers/peer-rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x29, 0xd6, 0x2f, 0x48, 0x4d, 0x2d, 0x82, 0x90, 0xba, 0x45, 0x05, 0xc9, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x3c, 0x20, 0x3e, 0x4c, 0x89, 0x92, 0x36, 0x17, 0x77, 0x70,
	0x49, 0x7e, 0x51, 0xaa, 0x73, 0x7e, 0x5e, 0x5a, 0x66, 0xba, 0x90, 0x0c, 0x17, 0xa7, 0x6b, 0x5e,
	0x4a, 0x41, 0x7e, 0x66, 0x5e, 0x49, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x42, 0x40,
	0x29, 0x84, 0x8b, 0xdd, 0x2b, 0x3f, 0x33, 0x2f, 0x28, 0xb5, 0x50, 0x48, 0x8c, 0x8b, 0x2d, 0x20,
	0x35, 0xb5, 0xc8, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0x32, 0xe4,
	0x62, 0x83, 0x18, 0x25, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x87, 0x6c, 0x9d, 0x1e,
	0x92, 0x5d, 0x41, 0x50, 0x85, 0x4a, 0xc6, 0x50, 0x53, 0x8b, 0x0b, 0x70, 0x9a, 0x2a, 0xc0, 0xc5,
	0xec, 0x5a, 0x54, 0x04, 0x36, 0x92, 0x35, 0x08, 0xc4, 0x54, 0x52, 0xe2, 0xe2, 0xf0, 0x49, 0x4d,
	0x2c, 0x4b, 0xc5, 0xe3, 0x16, 0x25, 0x19, 0x98, 0x9a, 0xe2, 0x02, 0x98, 0x09, 0x8c, 0x70, 0x13,
	0x8c, 0x1a, 0x18, 0xb9, 0xb8, 0x41, 0x0a, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xcc,
	0xb8, 0x58, 0x40, 0xce, 0x10, 0x12, 0x45, 0x75, 0x31, 0xd4, 0xc3, 0x52, 0xd8, 0x84, 0x8b, 0x0b,
	0x94, 0x18, 0x84, 0x2c, 0xb9, 0x58, 0xc1, 0xb6, 0x08, 0x89, 0xa1, 0xaa, 0x80, 0x39, 0x4f, 0x0a,
	0xab, 0x38, 0x48, 0x6b, 0x12, 0x1b, 0x38, 0x46, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88,
	0x1d, 0x73, 0x8f, 0xb2, 0x01, 0x00, 0x00,
}
