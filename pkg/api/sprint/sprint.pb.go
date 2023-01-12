// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/sprint/sprint.proto

package sprint

import (
	status "github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	issue "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Project           string                 `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	StartedAt         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Status            status.Sprint          `protobuf:"varint,6,opt,name=status,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.Sprint" json:"status,omitempty"`
	StoriesBacklog    int64                  `protobuf:"varint,7,opt,name=stories_backlog,json=storiesBacklog,proto3" json:"stories_backlog,omitempty"`
	StoriesInProgress int64                  `protobuf:"varint,8,opt,name=stories_in_progress,json=storiesInProgress,proto3" json:"stories_in_progress,omitempty"`
	StoriesDone       int64                  `protobuf:"varint,9,opt,name=stories_done,json=storiesDone,proto3" json:"stories_done,omitempty"`
	Issues            []*issue.IssueInfo     `protobuf:"bytes,10,rep,name=issues,proto3" json:"issues,omitempty"`
}

func (x *Sprint) Reset() {
	*x = Sprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sprint) ProtoMessage() {}

func (x *Sprint) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sprint.ProtoReflect.Descriptor instead.
func (*Sprint) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{0}
}

func (x *Sprint) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sprint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sprint) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Sprint) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Sprint) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Sprint) GetStatus() status.Sprint {
	if x != nil {
		return x.Status
	}
	return status.Sprint(0)
}

func (x *Sprint) GetStoriesBacklog() int64 {
	if x != nil {
		return x.StoriesBacklog
	}
	return 0
}

func (x *Sprint) GetStoriesInProgress() int64 {
	if x != nil {
		return x.StoriesInProgress
	}
	return 0
}

func (x *Sprint) GetStoriesDone() int64 {
	if x != nil {
		return x.StoriesDone
	}
	return 0
}

func (x *Sprint) GetIssues() []*issue.IssueInfo {
	if x != nil {
		return x.Issues
	}
	return nil
}

type SprintCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sprint *Sprint `protobuf:"bytes,1,opt,name=sprint,proto3" json:"sprint,omitempty"`
}

func (x *SprintCreateRequest) Reset() {
	*x = SprintCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintCreateRequest) ProtoMessage() {}

func (x *SprintCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintCreateRequest.ProtoReflect.Descriptor instead.
func (*SprintCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{1}
}

func (x *SprintCreateRequest) GetSprint() *Sprint {
	if x != nil {
		return x.Sprint
	}
	return nil
}

type SprintCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SprintCreateResponse) Reset() {
	*x = SprintCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintCreateResponse) ProtoMessage() {}

func (x *SprintCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintCreateResponse.ProtoReflect.Descriptor instead.
func (*SprintCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{2}
}

func (x *SprintCreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SprintUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sprint *Sprint `protobuf:"bytes,2,opt,name=sprint,proto3" json:"sprint,omitempty"`
}

func (x *SprintUpdateRequest) Reset() {
	*x = SprintUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintUpdateRequest) ProtoMessage() {}

func (x *SprintUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintUpdateRequest.ProtoReflect.Descriptor instead.
func (*SprintUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{3}
}

func (x *SprintUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SprintUpdateRequest) GetSprint() *Sprint {
	if x != nil {
		return x.Sprint
	}
	return nil
}

type SprintUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SprintUpdateResponse) Reset() {
	*x = SprintUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintUpdateResponse) ProtoMessage() {}

func (x *SprintUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintUpdateResponse.ProtoReflect.Descriptor instead.
func (*SprintUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{4}
}

func (x *SprintUpdateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SprintGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids              []int64                `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	StartedAtAfter   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=started_at_after,json=startedAtAfter,proto3" json:"started_at_after,omitempty"`
	FinishedAtBefore *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=finished_at_before,json=finishedAtBefore,proto3" json:"finished_at_before,omitempty"`
	Status           status.Sprint          `protobuf:"varint,5,opt,name=status,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.Sprint" json:"status,omitempty"`
}

func (x *SprintGetRequest) Reset() {
	*x = SprintGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintGetRequest) ProtoMessage() {}

func (x *SprintGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintGetRequest.ProtoReflect.Descriptor instead.
func (*SprintGetRequest) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{5}
}

func (x *SprintGetRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *SprintGetRequest) GetStartedAtAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAtAfter
	}
	return nil
}

func (x *SprintGetRequest) GetFinishedAtBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAtBefore
	}
	return nil
}

func (x *SprintGetRequest) GetStatus() status.Sprint {
	if x != nil {
		return x.Status
	}
	return status.Sprint(0)
}

type SprintGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sprint []*Sprint `protobuf:"bytes,1,rep,name=sprint,proto3" json:"sprint,omitempty"`
}

func (x *SprintGetResponse) Reset() {
	*x = SprintGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintGetResponse) ProtoMessage() {}

func (x *SprintGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintGetResponse.ProtoReflect.Descriptor instead.
func (*SprintGetResponse) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{6}
}

func (x *SprintGetResponse) GetSprint() []*Sprint {
	if x != nil {
		return x.Sprint
	}
	return nil
}

type SprintGetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SprintGetByIDRequest) Reset() {
	*x = SprintGetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintGetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintGetByIDRequest) ProtoMessage() {}

func (x *SprintGetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintGetByIDRequest.ProtoReflect.Descriptor instead.
func (*SprintGetByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{7}
}

func (x *SprintGetByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SprintGetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sprint *Sprint `protobuf:"bytes,1,opt,name=sprint,proto3" json:"sprint,omitempty"`
}

func (x *SprintGetByIDResponse) Reset() {
	*x = SprintGetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sprint_sprint_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SprintGetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SprintGetByIDResponse) ProtoMessage() {}

func (x *SprintGetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sprint_sprint_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SprintGetByIDResponse.ProtoReflect.Descriptor instead.
func (*SprintGetByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_sprint_sprint_proto_rawDescGZIP(), []int{8}
}

func (x *SprintGetByIDResponse) GetSprint() *Sprint {
	if x != nil {
		return x.Sprint
	}
	return nil
}

var File_api_sprint_sprint_proto protoreflect.FileDescriptor

var file_api_sprint_sprint_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b,
	0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2f, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x06, 0x53, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x58, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37,
	0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x12, 0x2e, 0x0a, 0x13,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x5a, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x13, 0x53,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x79,
	0x0a, 0x13, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x8e, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x48,
	0x0a, 0x12, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b,
	0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x67, 0x0a, 0x11, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x73, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e,
	0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x15, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x06,
	0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x32, 0xfb, 0x05, 0x0a, 0x0e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0xb8, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x12, 0x47, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a,
	0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0xbd,
	0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12,
	0x47, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b,
	0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xad,
	0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x44, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x73, 0x12, 0xbd,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32,
	0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e,
	0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x7c,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x4b, 0x2f, 0x63, 0x72, 0x6e, 0x74,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x3b, 0x73, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x92, 0x41, 0x36, 0x12, 0x0d, 0x0a, 0x06, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x32, 0x03, 0x30, 0x2e, 0x31, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_sprint_sprint_proto_rawDescOnce sync.Once
	file_api_sprint_sprint_proto_rawDescData = file_api_sprint_sprint_proto_rawDesc
)

func file_api_sprint_sprint_proto_rawDescGZIP() []byte {
	file_api_sprint_sprint_proto_rawDescOnce.Do(func() {
		file_api_sprint_sprint_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sprint_sprint_proto_rawDescData)
	})
	return file_api_sprint_sprint_proto_rawDescData
}

var file_api_sprint_sprint_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_sprint_sprint_proto_goTypes = []interface{}{
	(*Sprint)(nil),                // 0: github.constantine27k.crnt_data_manager.api.sprint.Sprint
	(*SprintCreateRequest)(nil),   // 1: github.constantine27k.crnt_data_manager.api.sprint.SprintCreateRequest
	(*SprintCreateResponse)(nil),  // 2: github.constantine27k.crnt_data_manager.api.sprint.SprintCreateResponse
	(*SprintUpdateRequest)(nil),   // 3: github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateRequest
	(*SprintUpdateResponse)(nil),  // 4: github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateResponse
	(*SprintGetRequest)(nil),      // 5: github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest
	(*SprintGetResponse)(nil),     // 6: github.constantine27k.crnt_data_manager.api.sprint.SprintGetResponse
	(*SprintGetByIDRequest)(nil),  // 7: github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDRequest
	(*SprintGetByIDResponse)(nil), // 8: github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(status.Sprint)(0),            // 10: github.constantine27k.crnt_data_manager.api.state.status.Sprint
	(*issue.IssueInfo)(nil),       // 11: github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfo
}
var file_api_sprint_sprint_proto_depIdxs = []int32{
	9,  // 0: github.constantine27k.crnt_data_manager.api.sprint.Sprint.started_at:type_name -> google.protobuf.Timestamp
	9,  // 1: github.constantine27k.crnt_data_manager.api.sprint.Sprint.finished_at:type_name -> google.protobuf.Timestamp
	10, // 2: github.constantine27k.crnt_data_manager.api.sprint.Sprint.status:type_name -> github.constantine27k.crnt_data_manager.api.state.status.Sprint
	11, // 3: github.constantine27k.crnt_data_manager.api.sprint.Sprint.issues:type_name -> github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfo
	0,  // 4: github.constantine27k.crnt_data_manager.api.sprint.SprintCreateRequest.sprint:type_name -> github.constantine27k.crnt_data_manager.api.sprint.Sprint
	0,  // 5: github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateRequest.sprint:type_name -> github.constantine27k.crnt_data_manager.api.sprint.Sprint
	9,  // 6: github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest.started_at_after:type_name -> google.protobuf.Timestamp
	9,  // 7: github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest.finished_at_before:type_name -> google.protobuf.Timestamp
	10, // 8: github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest.status:type_name -> github.constantine27k.crnt_data_manager.api.state.status.Sprint
	0,  // 9: github.constantine27k.crnt_data_manager.api.sprint.SprintGetResponse.sprint:type_name -> github.constantine27k.crnt_data_manager.api.sprint.Sprint
	0,  // 10: github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDResponse.sprint:type_name -> github.constantine27k.crnt_data_manager.api.sprint.Sprint
	1,  // 11: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.CreateSprint:input_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintCreateRequest
	3,  // 12: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.UpdateSprint:input_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateRequest
	5,  // 13: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.GetSprint:input_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest
	7,  // 14: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.GetSprintByID:input_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDRequest
	2,  // 15: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.CreateSprint:output_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintCreateResponse
	4,  // 16: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.UpdateSprint:output_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateResponse
	6,  // 17: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.GetSprint:output_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintGetResponse
	8,  // 18: github.constantine27k.crnt_data_manager.api.sprint.SprintRegistry.GetSprintByID:output_type -> github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_sprint_sprint_proto_init() }
func file_api_sprint_sprint_proto_init() {
	if File_api_sprint_sprint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sprint_sprint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sprint); i {
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
		file_api_sprint_sprint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintCreateRequest); i {
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
		file_api_sprint_sprint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintCreateResponse); i {
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
		file_api_sprint_sprint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintUpdateRequest); i {
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
		file_api_sprint_sprint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintUpdateResponse); i {
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
		file_api_sprint_sprint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintGetRequest); i {
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
		file_api_sprint_sprint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintGetResponse); i {
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
		file_api_sprint_sprint_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintGetByIDRequest); i {
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
		file_api_sprint_sprint_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SprintGetByIDResponse); i {
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
			RawDescriptor: file_api_sprint_sprint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sprint_sprint_proto_goTypes,
		DependencyIndexes: file_api_sprint_sprint_proto_depIdxs,
		MessageInfos:      file_api_sprint_sprint_proto_msgTypes,
	}.Build()
	File_api_sprint_sprint_proto = out.File
	file_api_sprint_sprint_proto_rawDesc = nil
	file_api_sprint_sprint_proto_goTypes = nil
	file_api_sprint_sprint_proto_depIdxs = nil
}
