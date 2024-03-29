// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/team/team.proto

package team

import (
	_ "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Members       []string `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Projects      []int64  `protobuf:"varint,4,rep,packed,name=projects,proto3" json:"projects,omitempty"`
	TechOwner     string   `protobuf:"bytes,5,opt,name=tech_owner,json=techOwner,proto3" json:"tech_owner,omitempty"`
	BusinessOwner string   `protobuf:"bytes,6,opt,name=business_owner,json=businessOwner,proto3" json:"business_owner,omitempty"`
	Department    string   `protobuf:"bytes,7,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Team) GetProjects() []int64 {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *Team) GetTechOwner() string {
	if x != nil {
		return x.TechOwner
	}
	return ""
}

func (x *Team) GetBusinessOwner() string {
	if x != nil {
		return x.BusinessOwner
	}
	return ""
}

func (x *Team) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

type TeamCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *TeamCreateRequest) Reset() {
	*x = TeamCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamCreateRequest) ProtoMessage() {}

func (x *TeamCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamCreateRequest.ProtoReflect.Descriptor instead.
func (*TeamCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{1}
}

func (x *TeamCreateRequest) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type TeamCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamCreateResponse) Reset() {
	*x = TeamCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamCreateResponse) ProtoMessage() {}

func (x *TeamCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamCreateResponse.ProtoReflect.Descriptor instead.
func (*TeamCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{2}
}

func (x *TeamCreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TeamUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Team *Team `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *TeamUpdateRequest) Reset() {
	*x = TeamUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamUpdateRequest) ProtoMessage() {}

func (x *TeamUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamUpdateRequest.ProtoReflect.Descriptor instead.
func (*TeamUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{3}
}

func (x *TeamUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TeamUpdateRequest) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type TeamUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamUpdateResponse) Reset() {
	*x = TeamUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamUpdateResponse) ProtoMessage() {}

func (x *TeamUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamUpdateResponse.ProtoReflect.Descriptor instead.
func (*TeamUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{4}
}

func (x *TeamUpdateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TeamGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids         []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Names       []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	Departments []string `protobuf:"bytes,3,rep,name=departments,proto3" json:"departments,omitempty"`
}

func (x *TeamGetRequest) Reset() {
	*x = TeamGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamGetRequest) ProtoMessage() {}

func (x *TeamGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamGetRequest.ProtoReflect.Descriptor instead.
func (*TeamGetRequest) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{5}
}

func (x *TeamGetRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *TeamGetRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *TeamGetRequest) GetDepartments() []string {
	if x != nil {
		return x.Departments
	}
	return nil
}

type TeamGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *TeamGetResponse) Reset() {
	*x = TeamGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamGetResponse) ProtoMessage() {}

func (x *TeamGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamGetResponse.ProtoReflect.Descriptor instead.
func (*TeamGetResponse) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{6}
}

func (x *TeamGetResponse) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

type TeamGetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamGetByIDRequest) Reset() {
	*x = TeamGetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamGetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamGetByIDRequest) ProtoMessage() {}

func (x *TeamGetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamGetByIDRequest.ProtoReflect.Descriptor instead.
func (*TeamGetByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{7}
}

func (x *TeamGetByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TeamGetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *TeamGetByIDResponse) Reset() {
	*x = TeamGetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_team_team_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamGetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamGetByIDResponse) ProtoMessage() {}

func (x *TeamGetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_team_team_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamGetByIDResponse.ProtoReflect.Descriptor instead.
func (*TeamGetByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_team_team_proto_rawDescGZIP(), []int{8}
}

func (x *TeamGetByIDResponse) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

var File_api_team_team_proto protoreflect.FileDescriptor

var file_api_team_team_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x63, 0x68, 0x5f, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x63, 0x68, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x11, 0x54, 0x65,
	0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4a, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x24, 0x0a, 0x12, 0x54,
	0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x6f, 0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x22, 0x24, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x5f, 0x0a, 0x0f, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63,
	0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x13, 0x54,
	0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x32, 0xca,
	0x05, 0x0a, 0x0c, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0xac, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x43,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x12, 0xb1,
	0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x43, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x3a, 0x01, 0x2a, 0x1a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0xa2, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12,
	0x40, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x41, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x12, 0xb1, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x12, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e,
	0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xb3, 0x01, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x4b, 0x2f, 0x63, 0x72, 0x6e, 0x74, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x92, 0x41, 0x71,
	0x12, 0x0b, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a, 0x01, 0x01,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x26, 0x0a, 0x24, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x13, 0x0a, 0x11,
	0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_team_team_proto_rawDescOnce sync.Once
	file_api_team_team_proto_rawDescData = file_api_team_team_proto_rawDesc
)

func file_api_team_team_proto_rawDescGZIP() []byte {
	file_api_team_team_proto_rawDescOnce.Do(func() {
		file_api_team_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_team_team_proto_rawDescData)
	})
	return file_api_team_team_proto_rawDescData
}

var file_api_team_team_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_team_team_proto_goTypes = []interface{}{
	(*Team)(nil),                // 0: github.constantine27k.crnt_data_manager.api.team.Team
	(*TeamCreateRequest)(nil),   // 1: github.constantine27k.crnt_data_manager.api.team.TeamCreateRequest
	(*TeamCreateResponse)(nil),  // 2: github.constantine27k.crnt_data_manager.api.team.TeamCreateResponse
	(*TeamUpdateRequest)(nil),   // 3: github.constantine27k.crnt_data_manager.api.team.TeamUpdateRequest
	(*TeamUpdateResponse)(nil),  // 4: github.constantine27k.crnt_data_manager.api.team.TeamUpdateResponse
	(*TeamGetRequest)(nil),      // 5: github.constantine27k.crnt_data_manager.api.team.TeamGetRequest
	(*TeamGetResponse)(nil),     // 6: github.constantine27k.crnt_data_manager.api.team.TeamGetResponse
	(*TeamGetByIDRequest)(nil),  // 7: github.constantine27k.crnt_data_manager.api.team.TeamGetByIDRequest
	(*TeamGetByIDResponse)(nil), // 8: github.constantine27k.crnt_data_manager.api.team.TeamGetByIDResponse
}
var file_api_team_team_proto_depIdxs = []int32{
	0, // 0: github.constantine27k.crnt_data_manager.api.team.TeamCreateRequest.team:type_name -> github.constantine27k.crnt_data_manager.api.team.Team
	0, // 1: github.constantine27k.crnt_data_manager.api.team.TeamUpdateRequest.team:type_name -> github.constantine27k.crnt_data_manager.api.team.Team
	0, // 2: github.constantine27k.crnt_data_manager.api.team.TeamGetResponse.teams:type_name -> github.constantine27k.crnt_data_manager.api.team.Team
	0, // 3: github.constantine27k.crnt_data_manager.api.team.TeamGetByIDResponse.team:type_name -> github.constantine27k.crnt_data_manager.api.team.Team
	1, // 4: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.CreateTeam:input_type -> github.constantine27k.crnt_data_manager.api.team.TeamCreateRequest
	3, // 5: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.UpdateTeam:input_type -> github.constantine27k.crnt_data_manager.api.team.TeamUpdateRequest
	5, // 6: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.GetTeams:input_type -> github.constantine27k.crnt_data_manager.api.team.TeamGetRequest
	7, // 7: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.GetTeamByID:input_type -> github.constantine27k.crnt_data_manager.api.team.TeamGetByIDRequest
	2, // 8: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.CreateTeam:output_type -> github.constantine27k.crnt_data_manager.api.team.TeamCreateResponse
	4, // 9: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.UpdateTeam:output_type -> github.constantine27k.crnt_data_manager.api.team.TeamUpdateResponse
	6, // 10: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.GetTeams:output_type -> github.constantine27k.crnt_data_manager.api.team.TeamGetResponse
	8, // 11: github.constantine27k.crnt_data_manager.api.team.TeamRegistry.GetTeamByID:output_type -> github.constantine27k.crnt_data_manager.api.team.TeamGetByIDResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_team_team_proto_init() }
func file_api_team_team_proto_init() {
	if File_api_team_team_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_team_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_api_team_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamCreateRequest); i {
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
		file_api_team_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamCreateResponse); i {
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
		file_api_team_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamUpdateRequest); i {
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
		file_api_team_team_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamUpdateResponse); i {
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
		file_api_team_team_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamGetRequest); i {
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
		file_api_team_team_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamGetResponse); i {
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
		file_api_team_team_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamGetByIDRequest); i {
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
		file_api_team_team_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamGetByIDResponse); i {
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
			RawDescriptor: file_api_team_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_team_team_proto_goTypes,
		DependencyIndexes: file_api_team_team_proto_depIdxs,
		MessageInfos:      file_api_team_team_proto_msgTypes,
	}.Build()
	File_api_team_team_proto = out.File
	file_api_team_team_proto_rawDesc = nil
	file_api_team_team_proto_goTypes = nil
	file_api_team_team_proto_depIdxs = nil
}
