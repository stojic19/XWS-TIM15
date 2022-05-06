// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: followers/followers.proto

package followers

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FollowersServiceClient is the client API for FollowersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowersServiceClient interface {
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	ConfirmFollow(ctx context.Context, in *ConfirmFollowRequest, opts ...grpc.CallOption) (*ConfirmFollowResponse, error)
	Unfollow(ctx context.Context, in *UnfollowRequest, opts ...grpc.CallOption) (*UnfollowResponse, error)
	RemoveFollowRequest(ctx context.Context, in *RemoveFollowRequestRequest, opts ...grpc.CallOption) (*RemoveFollowRequestResponse, error)
	GetFollows(ctx context.Context, in *GetFollowsRequest, opts ...grpc.CallOption) (*GetFollowsResponse, error)
	GetFollowers(ctx context.Context, in *GetFollowersRequest, opts ...grpc.CallOption) (*GetFollowersResponse, error)
	GetFollowRequests(ctx context.Context, in *GetFollowRequestsRequest, opts ...grpc.CallOption) (*GetFollowRequestsResponse, error)
	GetFollowerRequests(ctx context.Context, in *GetFollowerRequestsRequest, opts ...grpc.CallOption) (*GetFollowerRequestsResponse, error)
}

type followersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowersServiceClient(cc grpc.ClientConnInterface) FollowersServiceClient {
	return &followersServiceClient{cc}
}

func (c *followersServiceClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) ConfirmFollow(ctx context.Context, in *ConfirmFollowRequest, opts ...grpc.CallOption) (*ConfirmFollowResponse, error) {
	out := new(ConfirmFollowResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/ConfirmFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) Unfollow(ctx context.Context, in *UnfollowRequest, opts ...grpc.CallOption) (*UnfollowResponse, error) {
	out := new(UnfollowResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/Unfollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) RemoveFollowRequest(ctx context.Context, in *RemoveFollowRequestRequest, opts ...grpc.CallOption) (*RemoveFollowRequestResponse, error) {
	out := new(RemoveFollowRequestResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/RemoveFollowRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) GetFollows(ctx context.Context, in *GetFollowsRequest, opts ...grpc.CallOption) (*GetFollowsResponse, error) {
	out := new(GetFollowsResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/GetFollows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) GetFollowers(ctx context.Context, in *GetFollowersRequest, opts ...grpc.CallOption) (*GetFollowersResponse, error) {
	out := new(GetFollowersResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/GetFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) GetFollowRequests(ctx context.Context, in *GetFollowRequestsRequest, opts ...grpc.CallOption) (*GetFollowRequestsResponse, error) {
	out := new(GetFollowRequestsResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/GetFollowRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) GetFollowerRequests(ctx context.Context, in *GetFollowerRequestsRequest, opts ...grpc.CallOption) (*GetFollowerRequestsResponse, error) {
	out := new(GetFollowerRequestsResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/GetFollowerRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowersServiceServer is the server API for FollowersService service.
// All implementations must embed UnimplementedFollowersServiceServer
// for forward compatibility
type FollowersServiceServer interface {
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	ConfirmFollow(context.Context, *ConfirmFollowRequest) (*ConfirmFollowResponse, error)
	Unfollow(context.Context, *UnfollowRequest) (*UnfollowResponse, error)
	RemoveFollowRequest(context.Context, *RemoveFollowRequestRequest) (*RemoveFollowRequestResponse, error)
	GetFollows(context.Context, *GetFollowsRequest) (*GetFollowsResponse, error)
	GetFollowers(context.Context, *GetFollowersRequest) (*GetFollowersResponse, error)
	GetFollowRequests(context.Context, *GetFollowRequestsRequest) (*GetFollowRequestsResponse, error)
	GetFollowerRequests(context.Context, *GetFollowerRequestsRequest) (*GetFollowerRequestsResponse, error)
	mustEmbedUnimplementedFollowersServiceServer()
}

// UnimplementedFollowersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowersServiceServer struct {
}

func (UnimplementedFollowersServiceServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowersServiceServer) ConfirmFollow(context.Context, *ConfirmFollowRequest) (*ConfirmFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmFollow not implemented")
}
func (UnimplementedFollowersServiceServer) Unfollow(context.Context, *UnfollowRequest) (*UnfollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedFollowersServiceServer) RemoveFollowRequest(context.Context, *RemoveFollowRequestRequest) (*RemoveFollowRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFollowRequest not implemented")
}
func (UnimplementedFollowersServiceServer) GetFollows(context.Context, *GetFollowsRequest) (*GetFollowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollows not implemented")
}
func (UnimplementedFollowersServiceServer) GetFollowers(context.Context, *GetFollowersRequest) (*GetFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedFollowersServiceServer) GetFollowRequests(context.Context, *GetFollowRequestsRequest) (*GetFollowRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowRequests not implemented")
}
func (UnimplementedFollowersServiceServer) GetFollowerRequests(context.Context, *GetFollowerRequestsRequest) (*GetFollowerRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerRequests not implemented")
}
func (UnimplementedFollowersServiceServer) mustEmbedUnimplementedFollowersServiceServer() {}

// UnsafeFollowersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowersServiceServer will
// result in compilation errors.
type UnsafeFollowersServiceServer interface {
	mustEmbedUnimplementedFollowersServiceServer()
}

func RegisterFollowersServiceServer(s grpc.ServiceRegistrar, srv FollowersServiceServer) {
	s.RegisterService(&FollowersService_ServiceDesc, srv)
}

func _FollowersService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_ConfirmFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).ConfirmFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/ConfirmFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).ConfirmFollow(ctx, req.(*ConfirmFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/Unfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).Unfollow(ctx, req.(*UnfollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_RemoveFollowRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFollowRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).RemoveFollowRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/RemoveFollowRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).RemoveFollowRequest(ctx, req.(*RemoveFollowRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_GetFollows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).GetFollows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/GetFollows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).GetFollows(ctx, req.(*GetFollowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/GetFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).GetFollowers(ctx, req.(*GetFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_GetFollowRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).GetFollowRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/GetFollowRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).GetFollowRequests(ctx, req.(*GetFollowRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_GetFollowerRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowerRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).GetFollowerRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/GetFollowerRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).GetFollowerRequests(ctx, req.(*GetFollowerRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowersService_ServiceDesc is the grpc.ServiceDesc for FollowersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "followers.FollowersService",
	HandlerType: (*FollowersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Follow",
			Handler:    _FollowersService_Follow_Handler,
		},
		{
			MethodName: "ConfirmFollow",
			Handler:    _FollowersService_ConfirmFollow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _FollowersService_Unfollow_Handler,
		},
		{
			MethodName: "RemoveFollowRequest",
			Handler:    _FollowersService_RemoveFollowRequest_Handler,
		},
		{
			MethodName: "GetFollows",
			Handler:    _FollowersService_GetFollows_Handler,
		},
		{
			MethodName: "GetFollowers",
			Handler:    _FollowersService_GetFollowers_Handler,
		},
		{
			MethodName: "GetFollowRequests",
			Handler:    _FollowersService_GetFollowRequests_Handler,
		},
		{
			MethodName: "GetFollowerRequests",
			Handler:    _FollowersService_GetFollowerRequests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followers/followers.proto",
}
