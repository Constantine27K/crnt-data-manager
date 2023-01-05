// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: tasks/story/story.proto

package story

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

// StoryRegistryClient is the client API for StoryRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoryRegistryClient interface {
	Create(ctx context.Context, in *StoryCreateRequest, opts ...grpc.CallOption) (*StoryCreateResponse, error)
	Update(ctx context.Context, in *StoryUpdateRequest, opts ...grpc.CallOption) (*StoryUpdateResponse, error)
	Get(ctx context.Context, in *StoryGetRequest, opts ...grpc.CallOption) (*StoryGetResponse, error)
	GetByID(ctx context.Context, in *StoryGetByIDRequest, opts ...grpc.CallOption) (*StoryGetByIDResponse, error)
}

type storyRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewStoryRegistryClient(cc grpc.ClientConnInterface) StoryRegistryClient {
	return &storyRegistryClient{cc}
}

func (c *storyRegistryClient) Create(ctx context.Context, in *StoryCreateRequest, opts ...grpc.CallOption) (*StoryCreateResponse, error) {
	out := new(StoryCreateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyRegistryClient) Update(ctx context.Context, in *StoryUpdateRequest, opts ...grpc.CallOption) (*StoryUpdateResponse, error) {
	out := new(StoryUpdateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyRegistryClient) Get(ctx context.Context, in *StoryGetRequest, opts ...grpc.CallOption) (*StoryGetResponse, error) {
	out := new(StoryGetResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyRegistryClient) GetByID(ctx context.Context, in *StoryGetByIDRequest, opts ...grpc.CallOption) (*StoryGetByIDResponse, error) {
	out := new(StoryGetByIDResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoryRegistryServer is the server API for StoryRegistry service.
// All implementations should embed UnimplementedStoryRegistryServer
// for forward compatibility
type StoryRegistryServer interface {
	Create(context.Context, *StoryCreateRequest) (*StoryCreateResponse, error)
	Update(context.Context, *StoryUpdateRequest) (*StoryUpdateResponse, error)
	Get(context.Context, *StoryGetRequest) (*StoryGetResponse, error)
	GetByID(context.Context, *StoryGetByIDRequest) (*StoryGetByIDResponse, error)
}

// UnimplementedStoryRegistryServer should be embedded to have forward compatible implementations.
type UnimplementedStoryRegistryServer struct {
}

func (UnimplementedStoryRegistryServer) Create(context.Context, *StoryCreateRequest) (*StoryCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStoryRegistryServer) Update(context.Context, *StoryUpdateRequest) (*StoryUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStoryRegistryServer) Get(context.Context, *StoryGetRequest) (*StoryGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStoryRegistryServer) GetByID(context.Context, *StoryGetByIDRequest) (*StoryGetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

// UnsafeStoryRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoryRegistryServer will
// result in compilation errors.
type UnsafeStoryRegistryServer interface {
	mustEmbedUnimplementedStoryRegistryServer()
}

func RegisterStoryRegistryServer(s grpc.ServiceRegistrar, srv StoryRegistryServer) {
	s.RegisterService(&StoryRegistry_ServiceDesc, srv)
}

func _StoryRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryRegistryServer).Create(ctx, req.(*StoryCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryRegistryServer).Update(ctx, req.(*StoryUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryRegistryServer).Get(ctx, req.(*StoryGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryRegistry_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryGetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryRegistryServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryRegistryServer).GetByID(ctx, req.(*StoryGetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoryRegistry_ServiceDesc is the grpc.ServiceDesc for StoryRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoryRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.constantine27k.crnt_data_manager.api.tasks.story.StoryRegistry",
	HandlerType: (*StoryRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StoryRegistry_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StoryRegistry_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StoryRegistry_Get_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _StoryRegistry_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks/story/story.proto",
}
