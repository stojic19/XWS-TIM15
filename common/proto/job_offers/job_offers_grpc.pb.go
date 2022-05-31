// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: job_offers/job_offers.proto

package job_offers

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

// JobOffersServiceClient is the client API for JobOffersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobOffersServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Get(ctx context.Context, in *JobOfferId, opts ...grpc.CallOption) (*JobOffer, error)
	Create(ctx context.Context, in *NewJobOffer, opts ...grpc.CallOption) (*Response, error)
}

type jobOffersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobOffersServiceClient(cc grpc.ClientConnInterface) JobOffersServiceClient {
	return &jobOffersServiceClient{cc}
}

func (c *jobOffersServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/job_offers.JobOffersService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOffersServiceClient) Get(ctx context.Context, in *JobOfferId, opts ...grpc.CallOption) (*JobOffer, error) {
	out := new(JobOffer)
	err := c.cc.Invoke(ctx, "/job_offers.JobOffersService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOffersServiceClient) Create(ctx context.Context, in *NewJobOffer, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/job_offers.JobOffersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobOffersServiceServer is the server API for JobOffersService service.
// All implementations must embed UnimplementedJobOffersServiceServer
// for forward compatibility
type JobOffersServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Get(context.Context, *JobOfferId) (*JobOffer, error)
	Create(context.Context, *NewJobOffer) (*Response, error)
	mustEmbedUnimplementedJobOffersServiceServer()
}

// UnimplementedJobOffersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobOffersServiceServer struct {
}

func (UnimplementedJobOffersServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedJobOffersServiceServer) Get(context.Context, *JobOfferId) (*JobOffer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedJobOffersServiceServer) Create(context.Context, *NewJobOffer) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedJobOffersServiceServer) mustEmbedUnimplementedJobOffersServiceServer() {}

// UnsafeJobOffersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobOffersServiceServer will
// result in compilation errors.
type UnsafeJobOffersServiceServer interface {
	mustEmbedUnimplementedJobOffersServiceServer()
}

func RegisterJobOffersServiceServer(s grpc.ServiceRegistrar, srv JobOffersServiceServer) {
	s.RegisterService(&JobOffersService_ServiceDesc, srv)
}

func _JobOffersService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOffersServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_offers.JobOffersService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOffersServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOffersService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobOfferId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOffersServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_offers.JobOffersService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOffersServiceServer).Get(ctx, req.(*JobOfferId))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOffersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewJobOffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOffersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_offers.JobOffersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOffersServiceServer).Create(ctx, req.(*NewJobOffer))
	}
	return interceptor(ctx, in, info, handler)
}

// JobOffersService_ServiceDesc is the grpc.ServiceDesc for JobOffersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobOffersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job_offers.JobOffersService",
	HandlerType: (*JobOffersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _JobOffersService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _JobOffersService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _JobOffersService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_offers/job_offers.proto",
}
