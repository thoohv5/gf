// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/service.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("user/service.proto", fileDescriptor_d9e7c516be0a4556) }

var fileDescriptor_d9e7c516be0a4556 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x49, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27,
	0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0xd4, 0x48, 0x41, 0xf4, 0xe5,
	0xa6, 0x16, 0x17, 0x27, 0xa6, 0x43, 0xf5, 0x19, 0xd9, 0x70, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09,
	0x99, 0x70, 0xb1, 0x78, 0xe6, 0xa5, 0xe5, 0x0b, 0xf1, 0xea, 0x81, 0x14, 0xe9, 0x81, 0xd8, 0x41,
	0xa9, 0x85, 0x52, 0x7c, 0xc8, 0xdc, 0xe2, 0x02, 0x25, 0xde, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xb1,
	0x0b, 0xb1, 0xea, 0x67, 0xe6, 0xa5, 0xe5, 0x3b, 0xc9, 0x45, 0xc9, 0xa4, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x97, 0x64, 0xe4, 0xe7, 0x67, 0x94, 0x99, 0xea, 0x57, 0x55,
	0xe4, 0x14, 0xe9, 0x83, 0xf4, 0x25, 0xb1, 0x81, 0x2d, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0xa3, 0xa3, 0xe7, 0xb2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	// 详情
	Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	out := new(InfoResp)
	err := c.cc.Invoke(ctx, "/user.User/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// 详情
	Info(context.Context, *InfoReq) (*InfoResp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Info(ctx context.Context, req *InfoReq) (*InfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Info(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _User_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/service.proto",
}
