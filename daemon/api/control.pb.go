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
	CreateRoomRequest
	CreateRoomResponse
	UpdateRoomRequest
	UpdateRoomResponse
	DeleteRoomRequest
	DeleteRoomResponse
	Room
	RoomSpec
	UpdateStatus
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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
	Filters *ListRoomsRequest_Filters `protobuf:"bytes,1,opt,name=filters" json:"filters,omitempty"`
}

func (m *ListRoomsRequest) Reset()                    { *m = ListRoomsRequest{} }
func (m *ListRoomsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRoomsRequest) ProtoMessage()               {}
func (*ListRoomsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListRoomsRequest) GetFilters() *ListRoomsRequest_Filters {
	if m != nil {
		return m.Filters
	}
	return nil
}

type ListRoomsRequest_Filters struct {
	Names    []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	Ids      []string `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
	Catalogs []string `protobuf:"bytes,3,rep,name=catalogs" json:"catalogs,omitempty"`
}

func (m *ListRoomsRequest_Filters) Reset()                    { *m = ListRoomsRequest_Filters{} }
func (m *ListRoomsRequest_Filters) String() string            { return proto.CompactTextString(m) }
func (*ListRoomsRequest_Filters) ProtoMessage()               {}
func (*ListRoomsRequest_Filters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ListRoomsRequest_Filters) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListRoomsRequest_Filters) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ListRoomsRequest_Filters) GetCatalogs() []string {
	if m != nil {
		return m.Catalogs
	}
	return nil
}

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

type CreateRoomRequest struct {
	RoomSpec *RoomSpec `protobuf:"bytes,1,opt,name=room_spec,json=roomSpec" json:"room_spec,omitempty"`
}

func (m *CreateRoomRequest) Reset()                    { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()               {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRoomRequest) GetRoomSpec() *RoomSpec {
	if m != nil {
		return m.RoomSpec
	}
	return nil
}

type CreateRoomResponse struct {
	Room *Room `protobuf:"bytes,1,opt,name=room" json:"room,omitempty"`
}

func (m *CreateRoomResponse) Reset()                    { *m = CreateRoomResponse{} }
func (m *CreateRoomResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRoomResponse) ProtoMessage()               {}
func (*CreateRoomResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type UpdateRoomRequest struct {
	RoomId   string    `protobuf:"bytes,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	RoomSpec *RoomSpec `protobuf:"bytes,2,opt,name=room_spec,json=roomSpec" json:"room_spec,omitempty"`
}

func (m *UpdateRoomRequest) Reset()                    { *m = UpdateRoomRequest{} }
func (m *UpdateRoomRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRoomRequest) ProtoMessage()               {}
func (*UpdateRoomRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateRoomRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *UpdateRoomRequest) GetRoomSpec() *RoomSpec {
	if m != nil {
		return m.RoomSpec
	}
	return nil
}

type UpdateRoomResponse struct {
	Room *Room `protobuf:"bytes,1,opt,name=room" json:"room,omitempty"`
}

func (m *UpdateRoomResponse) Reset()                    { *m = UpdateRoomResponse{} }
func (m *UpdateRoomResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateRoomResponse) ProtoMessage()               {}
func (*UpdateRoomResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type DeleteRoomRequest struct {
	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
}

func (m *DeleteRoomRequest) Reset()                    { *m = DeleteRoomRequest{} }
func (m *DeleteRoomRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoomRequest) ProtoMessage()               {}
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRoomRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type DeleteRoomResponse struct {
	Room *Room `protobuf:"bytes,1,opt,name=room" json:"room,omitempty"`
}

func (m *DeleteRoomResponse) Reset()                    { *m = DeleteRoomResponse{} }
func (m *DeleteRoomResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoomResponse) ProtoMessage()               {}
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRoomsRequest)(nil), "api.ListRoomsRequest")
	proto.RegisterType((*ListRoomsRequest_Filters)(nil), "api.ListRoomsRequest.Filters")
	proto.RegisterType((*ListRoomsResponse)(nil), "api.ListRoomsResponse")
	proto.RegisterType((*CreateRoomRequest)(nil), "api.CreateRoomRequest")
	proto.RegisterType((*CreateRoomResponse)(nil), "api.CreateRoomResponse")
	proto.RegisterType((*UpdateRoomRequest)(nil), "api.UpdateRoomRequest")
	proto.RegisterType((*UpdateRoomResponse)(nil), "api.UpdateRoomResponse")
	proto.RegisterType((*DeleteRoomRequest)(nil), "api.DeleteRoomRequest")
	proto.RegisterType((*DeleteRoomResponse)(nil), "api.DeleteRoomResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GipfeliRoomService service

type GipfeliRoomServiceClient interface {
	ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error)
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error)
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

func (c *gipfeliRoomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := grpc.Invoke(ctx, "/api.GipfeliRoomService/CreateRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gipfeliRoomServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error) {
	out := new(UpdateRoomResponse)
	err := grpc.Invoke(ctx, "/api.GipfeliRoomService/UpdateRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gipfeliRoomServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error) {
	out := new(DeleteRoomResponse)
	err := grpc.Invoke(ctx, "/api.GipfeliRoomService/DeleteRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GipfeliRoomService service

type GipfeliRoomServiceServer interface {
	ListRooms(context.Context, *ListRoomsRequest) (*ListRoomsResponse, error)
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error)
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

func _GipfeliRoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GipfeliRoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GipfeliRoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GipfeliRoomServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GipfeliRoomService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GipfeliRoomServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GipfeliRoomService/UpdateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GipfeliRoomServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GipfeliRoomService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GipfeliRoomServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GipfeliRoomService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GipfeliRoomServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
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
		{
			MethodName: "CreateRoom",
			Handler:    _GipfeliRoomService_CreateRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _GipfeliRoomService_UpdateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _GipfeliRoomService_DeleteRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

func init() { proto.RegisterFile("control.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x06, 0xb6, 0x9b, 0xa9, 0x2a, 0x5a, 0x6b, 0x61, 0x4b, 0x60, 0x11, 0x58, 0x1c,
	0x56, 0x15, 0x6a, 0x44, 0xf7, 0x80, 0xc4, 0x85, 0x03, 0x08, 0x84, 0x04, 0x97, 0x00, 0x12, 0x9c,
	0xc0, 0x9b, 0xcc, 0x56, 0x46, 0x69, 0xc6, 0xd8, 0x66, 0x2f, 0x88, 0x0b, 0x67, 0x6e, 0xdc, 0x78,
	0x2d, 0x5e, 0x81, 0x07, 0x41, 0x76, 0xbc, 0x6d, 0x49, 0x8a, 0xe8, 0x2d, 0xf3, 0x7b, 0xfc, 0x7f,
	0xe3, 0xf9, 0x5b, 0x18, 0x16, 0x54, 0x5b, 0x4d, 0xd5, 0x4c, 0x69, 0xb2, 0xc4, 0x62, 0xa1, 0x64,
	0x3a, 0xa4, 0xd3, 0x8f, 0x58, 0x58, 0xd3, 0x68, 0xe9, 0xc0, 0x28, 0x2c, 0x2e, 0x8a, 0x9b, 0x0b,
	0xa2, 0x45, 0x85, 0x99, 0x50, 0x32, 0x13, 0x75, 0x4d, 0x56, 0x58, 0x49, 0x75, 0x38, 0xe5, 0x3f,
	0x23, 0x18, 0xbd, 0x90, 0xc6, 0xe6, 0x44, 0x4b, 0x93, 0xe3, 0xa7, 0xcf, 0x68, 0x2c, 0x7b, 0x00,
	0xfd, 0x33, 0x59, 0x59, 0xd4, 0x66, 0x12, 0xdd, 0x8e, 0x8e, 0x07, 0xf3, 0xa3, 0x99, 0x50, 0x72,
	0xd6, 0xee, 0x9b, 0x3d, 0x6d, 0x9a, 0xf2, 0x8b, 0xee, 0xf4, 0x25, 0xf4, 0x83, 0xc6, 0x0e, 0xe0,
	0x72, 0x2d, 0x96, 0xe8, 0x1c, 0xe2, 0xe3, 0x24, 0x6f, 0x0a, 0x36, 0x82, 0x58, 0x96, 0x66, 0xd2,
	0xf3, 0x9a, 0xfb, 0x64, 0x29, 0xec, 0x17, 0xc2, 0x8a, 0x8a, 0x16, 0x66, 0x12, 0x7b, 0x79, 0x55,
	0xf3, 0x39, 0x8c, 0x37, 0x98, 0x46, 0x51, 0x6d, 0x90, 0x1d, 0xc1, 0x25, 0x4d, 0xb4, 0xf4, 0xbe,
	0x83, 0x79, 0xe2, 0x27, 0x73, 0x1d, 0xb9, 0x97, 0xf9, 0x23, 0x18, 0x3f, 0xd6, 0x28, 0x2c, 0x7a,
	0x2d, 0x3c, 0x68, 0x0a, 0x89, 0x3b, 0x7c, 0xef, 0xf6, 0x12, 0x9e, 0x34, 0x5c, 0x5d, 0x7c, 0xa5,
	0xb0, 0xc8, 0xf7, 0x75, 0xf8, 0xe2, 0x27, 0xc0, 0x36, 0x0d, 0x3a, 0xd4, 0x68, 0x1b, 0xf5, 0x2d,
	0x8c, 0xdf, 0xa8, 0xb2, 0x45, 0x3d, 0x84, 0xbe, 0xa7, 0xca, 0xd2, 0x5f, 0x4b, 0xf2, 0x3d, 0x57,
	0x3e, 0x2f, 0xff, 0x1e, 0xa7, 0xf7, 0xdf, 0x71, 0x36, 0x9d, 0x77, 0x1b, 0xe7, 0x1e, 0x8c, 0x9f,
	0x60, 0x85, 0xbb, 0x8d, 0xe3, 0x10, 0x9b, 0xdd, 0x3b, 0x21, 0xe6, 0xdf, 0x63, 0x60, 0xcf, 0xa4,
	0x3a, 0xc3, 0x4a, 0xfa, 0xa9, 0x51, 0x9f, 0xcb, 0x02, 0xd9, 0x6b, 0x48, 0x56, 0x91, 0xb1, 0xab,
	0x5b, 0x7f, 0x36, 0xe9, 0xb5, 0xb6, 0xdc, 0x10, 0xf9, 0xe4, 0xdb, 0xaf, 0xdf, 0x3f, 0x7a, 0x8c,
	0x8d, 0xb2, 0xf3, 0xfb, 0x59, 0x29, 0x70, 0x49, 0x75, 0xa6, 0xbd, 0xd1, 0x3b, 0x80, 0x75, 0x26,
	0xac, 0xb9, 0xdf, 0x49, 0x39, 0x3d, 0xec, 0xe8, 0xc1, 0x38, 0xf5, 0xc6, 0x07, 0xfc, 0x4a, 0xcb,
	0xf8, 0x61, 0x34, 0x65, 0x05, 0xc0, 0x7a, 0xbf, 0xc1, 0xba, 0x13, 0x65, 0xb0, 0xee, 0x06, 0xc1,
	0xef, 0x7a, 0xeb, 0x5b, 0xe9, 0xf5, 0x96, 0x75, 0xf6, 0x25, 0xec, 0xfa, 0xab, 0x83, 0x7c, 0x00,
	0x58, 0x6f, 0x38, 0x40, 0x3a, 0x01, 0x05, 0x48, 0x37, 0x0a, 0x7e, 0xc7, 0x43, 0x6e, 0x4c, 0xff,
	0x0d, 0x39, 0xdd, 0xf3, 0x7f, 0xe7, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xce, 0x7e, 0x90,
	0x1f, 0x1e, 0x04, 0x00, 0x00,
}
