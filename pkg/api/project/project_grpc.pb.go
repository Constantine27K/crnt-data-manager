// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/project/project.proto

package project

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

// ProjectRegistryClient is the client API for ProjectRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectRegistryClient interface {
	CreateProject(ctx context.Context, in *ProjectCreateRequest, opts ...grpc.CallOption) (*ProjectCreateResponse, error)
	AddResponsibleTeam(ctx context.Context, in *ProjectAddTeamRequest, opts ...grpc.CallOption) (*ProjectAddTeamResponse, error)
	RemoveResponsibleTeam(ctx context.Context, in *ProjectRemoveTeamRequest, opts ...grpc.CallOption) (*ProjectRemoveTeamResponse, error)
	UpdateProject(ctx context.Context, in *ProjectUpdateRequest, opts ...grpc.CallOption) (*ProjectUpdateResponse, error)
	GetProjects(ctx context.Context, in *ProjectGetRequest, opts ...grpc.CallOption) (*ProjectGetResponse, error)
	GetProjectByID(ctx context.Context, in *ProjectGetByIDRequest, opts ...grpc.CallOption) (*ProjectGetByIDResponse, error)
}

type projectRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectRegistryClient(cc grpc.ClientConnInterface) ProjectRegistryClient {
	return &projectRegistryClient{cc}
}

func (c *projectRegistryClient) CreateProject(ctx context.Context, in *ProjectCreateRequest, opts ...grpc.CallOption) (*ProjectCreateResponse, error) {
	out := new(ProjectCreateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectRegistryClient) AddResponsibleTeam(ctx context.Context, in *ProjectAddTeamRequest, opts ...grpc.CallOption) (*ProjectAddTeamResponse, error) {
	out := new(ProjectAddTeamResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/AddResponsibleTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectRegistryClient) RemoveResponsibleTeam(ctx context.Context, in *ProjectRemoveTeamRequest, opts ...grpc.CallOption) (*ProjectRemoveTeamResponse, error) {
	out := new(ProjectRemoveTeamResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/RemoveResponsibleTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectRegistryClient) UpdateProject(ctx context.Context, in *ProjectUpdateRequest, opts ...grpc.CallOption) (*ProjectUpdateResponse, error) {
	out := new(ProjectUpdateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectRegistryClient) GetProjects(ctx context.Context, in *ProjectGetRequest, opts ...grpc.CallOption) (*ProjectGetResponse, error) {
	out := new(ProjectGetResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/GetProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectRegistryClient) GetProjectByID(ctx context.Context, in *ProjectGetByIDRequest, opts ...grpc.CallOption) (*ProjectGetByIDResponse, error) {
	out := new(ProjectGetByIDResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/GetProjectByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectRegistryServer is the server API for ProjectRegistry service.
// All implementations should embed UnimplementedProjectRegistryServer
// for forward compatibility
type ProjectRegistryServer interface {
	CreateProject(context.Context, *ProjectCreateRequest) (*ProjectCreateResponse, error)
	AddResponsibleTeam(context.Context, *ProjectAddTeamRequest) (*ProjectAddTeamResponse, error)
	RemoveResponsibleTeam(context.Context, *ProjectRemoveTeamRequest) (*ProjectRemoveTeamResponse, error)
	UpdateProject(context.Context, *ProjectUpdateRequest) (*ProjectUpdateResponse, error)
	GetProjects(context.Context, *ProjectGetRequest) (*ProjectGetResponse, error)
	GetProjectByID(context.Context, *ProjectGetByIDRequest) (*ProjectGetByIDResponse, error)
}

// UnimplementedProjectRegistryServer should be embedded to have forward compatible implementations.
type UnimplementedProjectRegistryServer struct {
}

func (UnimplementedProjectRegistryServer) CreateProject(context.Context, *ProjectCreateRequest) (*ProjectCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectRegistryServer) AddResponsibleTeam(context.Context, *ProjectAddTeamRequest) (*ProjectAddTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddResponsibleTeam not implemented")
}
func (UnimplementedProjectRegistryServer) RemoveResponsibleTeam(context.Context, *ProjectRemoveTeamRequest) (*ProjectRemoveTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveResponsibleTeam not implemented")
}
func (UnimplementedProjectRegistryServer) UpdateProject(context.Context, *ProjectUpdateRequest) (*ProjectUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectRegistryServer) GetProjects(context.Context, *ProjectGetRequest) (*ProjectGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedProjectRegistryServer) GetProjectByID(context.Context, *ProjectGetByIDRequest) (*ProjectGetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectByID not implemented")
}

// UnsafeProjectRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectRegistryServer will
// result in compilation errors.
type UnsafeProjectRegistryServer interface {
	mustEmbedUnimplementedProjectRegistryServer()
}

func RegisterProjectRegistryServer(s grpc.ServiceRegistrar, srv ProjectRegistryServer) {
	s.RegisterService(&ProjectRegistry_ServiceDesc, srv)
}

func _ProjectRegistry_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectRegistryServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectRegistryServer).CreateProject(ctx, req.(*ProjectCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectRegistry_AddResponsibleTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAddTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectRegistryServer).AddResponsibleTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/AddResponsibleTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectRegistryServer).AddResponsibleTeam(ctx, req.(*ProjectAddTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectRegistry_RemoveResponsibleTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRemoveTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectRegistryServer).RemoveResponsibleTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/RemoveResponsibleTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectRegistryServer).RemoveResponsibleTeam(ctx, req.(*ProjectRemoveTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectRegistry_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectRegistryServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectRegistryServer).UpdateProject(ctx, req.(*ProjectUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectRegistry_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectRegistryServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/GetProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectRegistryServer).GetProjects(ctx, req.(*ProjectGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectRegistry_GetProjectByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectGetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectRegistryServer).GetProjectByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_data_manager.api.project.ProjectRegistry/GetProjectByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectRegistryServer).GetProjectByID(ctx, req.(*ProjectGetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectRegistry_ServiceDesc is the grpc.ServiceDesc for ProjectRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.constantine27k.crnt_data_manager.api.project.ProjectRegistry",
	HandlerType: (*ProjectRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectRegistry_CreateProject_Handler,
		},
		{
			MethodName: "AddResponsibleTeam",
			Handler:    _ProjectRegistry_AddResponsibleTeam_Handler,
		},
		{
			MethodName: "RemoveResponsibleTeam",
			Handler:    _ProjectRegistry_RemoveResponsibleTeam_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _ProjectRegistry_UpdateProject_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _ProjectRegistry_GetProjects_Handler,
		},
		{
			MethodName: "GetProjectByID",
			Handler:    _ProjectRegistry_GetProjectByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/project/project.proto",
}
