// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/project/project.proto

package project

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsArchived         bool    `protobuf:"varint,3,opt,name=is_archived,json=isArchived,proto3" json:"is_archived,omitempty"`
	ResponsibleTeamIds []int64 `protobuf:"varint,4,rep,packed,name=responsible_team_ids,json=responsibleTeamIds,proto3" json:"responsible_team_ids,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *Project) GetResponsibleTeamIds() []int64 {
	if x != nil {
		return x.ResponsibleTeamIds
	}
	return nil
}

type ProjectCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectCreateRequest) Reset() {
	*x = ProjectCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreateRequest) ProtoMessage() {}

func (x *ProjectCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreateRequest.ProtoReflect.Descriptor instead.
func (*ProjectCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectCreateRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type ProjectCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectCreateResponse) Reset() {
	*x = ProjectCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreateResponse) ProtoMessage() {}

func (x *ProjectCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreateResponse.ProtoReflect.Descriptor instead.
func (*ProjectCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectCreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Project *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectUpdateRequest) Reset() {
	*x = ProjectUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUpdateRequest) ProtoMessage() {}

func (x *ProjectUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUpdateRequest.ProtoReflect.Descriptor instead.
func (*ProjectUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectUpdateRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type ProjectUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectUpdateResponse) Reset() {
	*x = ProjectUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUpdateResponse) ProtoMessage() {}

func (x *ProjectUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUpdateResponse.ProtoReflect.Descriptor instead.
func (*ProjectUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectUpdateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids                []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Names              []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	IsArchived         bool     `protobuf:"varint,3,opt,name=is_archived,json=isArchived,proto3" json:"is_archived,omitempty"`
	ResponsibleTeamIds []int64  `protobuf:"varint,4,rep,packed,name=responsible_team_ids,json=responsibleTeamIds,proto3" json:"responsible_team_ids,omitempty"`
}

func (x *ProjectGetRequest) Reset() {
	*x = ProjectGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectGetRequest) ProtoMessage() {}

func (x *ProjectGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectGetRequest.ProtoReflect.Descriptor instead.
func (*ProjectGetRequest) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectGetRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProjectGetRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ProjectGetRequest) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *ProjectGetRequest) GetResponsibleTeamIds() []int64 {
	if x != nil {
		return x.ResponsibleTeamIds
	}
	return nil
}

type ProjectGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ProjectGetResponse) Reset() {
	*x = ProjectGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectGetResponse) ProtoMessage() {}

func (x *ProjectGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectGetResponse.ProtoReflect.Descriptor instead.
func (*ProjectGetResponse) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectGetResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ProjectGetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectGetByIDRequest) Reset() {
	*x = ProjectGetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectGetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectGetByIDRequest) ProtoMessage() {}

func (x *ProjectGetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectGetByIDRequest.ProtoReflect.Descriptor instead.
func (*ProjectGetByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectGetByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectGetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectGetByIDResponse) Reset() {
	*x = ProjectGetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_project_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectGetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectGetByIDResponse) ProtoMessage() {}

func (x *ProjectGetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_project_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectGetByIDResponse.ProtoReflect.Descriptor instead.
func (*ProjectGetByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_project_project_proto_rawDescGZIP(), []int{8}
}

func (x *ProjectGetByIDResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

var File_api_project_project_proto protoreflect.FileDescriptor

var file_api_project_project_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32,
	0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x80, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64,
	0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x73, 0x22, 0x6e, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65,
	0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x56, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x73, 0x22, 0x6e, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x70,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37,
	0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x32, 0x95, 0x06, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x12, 0xbe, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x49, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63,
	0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xc3, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x49, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e,
	0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xb4, 0x01, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x46, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65,
	0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x12, 0xc3, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x4a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x7f, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x32, 0x37, 0x4b, 0x2f, 0x63, 0x72, 0x6e, 0x74, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x92,
	0x41, 0x37, 0x12, 0x0e, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x03, 0x30,
	0x2e, 0x31, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_project_project_proto_rawDescOnce sync.Once
	file_api_project_project_proto_rawDescData = file_api_project_project_proto_rawDesc
)

func file_api_project_project_proto_rawDescGZIP() []byte {
	file_api_project_project_proto_rawDescOnce.Do(func() {
		file_api_project_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_project_project_proto_rawDescData)
	})
	return file_api_project_project_proto_rawDescData
}

var file_api_project_project_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_project_project_proto_goTypes = []interface{}{
	(*Project)(nil),                // 0: github.constantine27k.crnt_data_manager.api.project.Project
	(*ProjectCreateRequest)(nil),   // 1: github.constantine27k.crnt_data_manager.api.project.ProjectCreateRequest
	(*ProjectCreateResponse)(nil),  // 2: github.constantine27k.crnt_data_manager.api.project.ProjectCreateResponse
	(*ProjectUpdateRequest)(nil),   // 3: github.constantine27k.crnt_data_manager.api.project.ProjectUpdateRequest
	(*ProjectUpdateResponse)(nil),  // 4: github.constantine27k.crnt_data_manager.api.project.ProjectUpdateResponse
	(*ProjectGetRequest)(nil),      // 5: github.constantine27k.crnt_data_manager.api.project.ProjectGetRequest
	(*ProjectGetResponse)(nil),     // 6: github.constantine27k.crnt_data_manager.api.project.ProjectGetResponse
	(*ProjectGetByIDRequest)(nil),  // 7: github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDRequest
	(*ProjectGetByIDResponse)(nil), // 8: github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDResponse
}
var file_api_project_project_proto_depIdxs = []int32{
	0, // 0: github.constantine27k.crnt_data_manager.api.project.ProjectCreateRequest.project:type_name -> github.constantine27k.crnt_data_manager.api.project.Project
	0, // 1: github.constantine27k.crnt_data_manager.api.project.ProjectUpdateRequest.project:type_name -> github.constantine27k.crnt_data_manager.api.project.Project
	0, // 2: github.constantine27k.crnt_data_manager.api.project.ProjectGetResponse.projects:type_name -> github.constantine27k.crnt_data_manager.api.project.Project
	0, // 3: github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDResponse.project:type_name -> github.constantine27k.crnt_data_manager.api.project.Project
	1, // 4: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.CreateProject:input_type -> github.constantine27k.crnt_data_manager.api.project.ProjectCreateRequest
	3, // 5: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.UpdateProject:input_type -> github.constantine27k.crnt_data_manager.api.project.ProjectUpdateRequest
	5, // 6: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.GetProjects:input_type -> github.constantine27k.crnt_data_manager.api.project.ProjectGetRequest
	7, // 7: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.GetProjectByID:input_type -> github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDRequest
	2, // 8: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.CreateProject:output_type -> github.constantine27k.crnt_data_manager.api.project.ProjectCreateResponse
	4, // 9: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.UpdateProject:output_type -> github.constantine27k.crnt_data_manager.api.project.ProjectUpdateResponse
	6, // 10: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.GetProjects:output_type -> github.constantine27k.crnt_data_manager.api.project.ProjectGetResponse
	8, // 11: github.constantine27k.crnt_data_manager.api.project.ProjectRegistry.GetProjectByID:output_type -> github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_project_project_proto_init() }
func file_api_project_project_proto_init() {
	if File_api_project_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_project_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectUpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectGetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectGetByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_project_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectGetByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_project_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_project_project_proto_goTypes,
		DependencyIndexes: file_api_project_project_proto_depIdxs,
		MessageInfos:      file_api_project_project_proto_msgTypes,
	}.Build()
	File_api_project_project_proto = out.File
	file_api_project_project_proto_rawDesc = nil
	file_api_project_project_proto_goTypes = nil
	file_api_project_project_proto_depIdxs = nil
}
