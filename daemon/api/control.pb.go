// Code generated by protoc-gen-go.
// source: control.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	control.proto
	objects.proto
	specs.proto
	types.proto

It has these top-level messages:
	ListRoomsRequest
	ListRoomsResponse
	Room
	RoomSpec
	UpdateStatus
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

type ListRoomsRequest struct {
}

func (m *ListRoomsRequest) Reset()                    { *m = ListRoomsRequest{} }
func (m *ListRoomsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRoomsRequest) ProtoMessage()               {}
func (*ListRoomsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListRoomsResponse struct {
	Room []*Room `protobuf:"bytes,1,rep,name=room" json:"room,omitempty"`
}

func (m *ListRoomsResponse) Reset()                    { *m = ListRoomsResponse{} }
func (m *ListRoomsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRoomsResponse) ProtoMessage()               {}
func (*ListRoomsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListRoomsResponse) GetRoom() []*Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRoomsRequest)(nil), "api.ListRoomsRequest")
	proto.RegisterType((*ListRoomsResponse)(nil), "api.ListRoomsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GipfeliRoomService service

type GipfeliRoomServiceClient interface {
	ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error)
}

type gipfeliRoomServiceClient struct {
	cc *grpc.ClientConn
}

func NewGipfeliRoomServiceClient(cc *grpc.ClientConn) GipfeliRoomServiceClient {
	return &gipfeliRoomServiceClient{cc}
}

func (c *gipfeliRoomServiceClient) ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error) {
	out := new(ListRoomsResponse)
	err := grpc.Invoke(ctx, "/api.GipfeliRoomService/ListRooms", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GipfeliRoomService service

type GipfeliRoomServiceServer interface {
	ListRooms(context.Context, *ListRoomsRequest) (*ListRoomsResponse, error)
}

func RegisterGipfeliRoomServiceServer(s *grpc.Server, srv GipfeliRoomServiceServer) {
	s.RegisterService(&_GipfeliRoomService_serviceDesc, srv)
}

func _GipfeliRoomService_ListRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GipfeliRoomServiceServer).ListRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GipfeliRoomService/ListRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GipfeliRoomServiceServer).ListRooms(ctx, req.(*ListRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GipfeliRoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GipfeliRoomService",
	HandlerType: (*GipfeliRoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRooms",
			Handler:    _GipfeliRoomService_ListRooms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("control.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xe2,
	0xcd, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0x88, 0x29, 0x09, 0x71, 0x09, 0xf8, 0x64, 0x16,
	0x97, 0x04, 0xe5, 0xe7, 0xe7, 0x16, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x19, 0x71,
	0x09, 0x22, 0x89, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x72, 0xb1, 0x14, 0x01, 0x05,
	0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x38, 0xf5, 0x80, 0x66, 0xe9, 0x81, 0x54, 0x04, 0x81,
	0x85, 0x8d, 0x02, 0xb8, 0x84, 0xdc, 0x33, 0x0b, 0xd2, 0x52, 0x73, 0x32, 0x41, 0x82, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x56, 0x5c, 0x9c, 0x70, 0x93, 0x84, 0x44, 0xc1, 0x7a, 0xd0,
	0x6d, 0x93, 0x12, 0x43, 0x17, 0x86, 0x58, 0x98, 0xc4, 0x06, 0x76, 0xa0, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x12, 0x0d, 0x47, 0xc5, 0x00, 0x00, 0x00,
}
