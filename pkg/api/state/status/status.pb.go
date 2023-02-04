// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/state/status/status.proto

package status

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Common int32

const (
	Common_STATUS_COMMON_UNKNOWN           Common = 0
	Common_STATUS_COMMON_BACKLOG           Common = 1
	Common_STATUS_COMMON_IN_PROGRESS       Common = 2
	Common_STATUS_COMMON_DONE              Common = 3
	Common_STATUS_COMMON_READY_FOR_REVIEW  Common = 4
	Common_STATUS_COMMON_IN_REVIEW         Common = 5
	Common_STATUS_COMMON_REVIEW_FAILED     Common = 6
	Common_STATUS_COMMON_REVIEW_PASSED     Common = 7
	Common_STATUS_COMMON_GIVEN_TO_CUSTOMER Common = 8
	Common_STATUS_COMMON_CLOSED            Common = 9
	Common_STATUS_COMMON_ON_HOLD           Common = 10
)

// Enum value maps for Common.
var (
	Common_name = map[int32]string{
		0:  "STATUS_COMMON_UNKNOWN",
		1:  "STATUS_COMMON_BACKLOG",
		2:  "STATUS_COMMON_IN_PROGRESS",
		3:  "STATUS_COMMON_DONE",
		4:  "STATUS_COMMON_READY_FOR_REVIEW",
		5:  "STATUS_COMMON_IN_REVIEW",
		6:  "STATUS_COMMON_REVIEW_FAILED",
		7:  "STATUS_COMMON_REVIEW_PASSED",
		8:  "STATUS_COMMON_GIVEN_TO_CUSTOMER",
		9:  "STATUS_COMMON_CLOSED",
		10: "STATUS_COMMON_ON_HOLD",
	}
	Common_value = map[string]int32{
		"STATUS_COMMON_UNKNOWN":           0,
		"STATUS_COMMON_BACKLOG":           1,
		"STATUS_COMMON_IN_PROGRESS":       2,
		"STATUS_COMMON_DONE":              3,
		"STATUS_COMMON_READY_FOR_REVIEW":  4,
		"STATUS_COMMON_IN_REVIEW":         5,
		"STATUS_COMMON_REVIEW_FAILED":     6,
		"STATUS_COMMON_REVIEW_PASSED":     7,
		"STATUS_COMMON_GIVEN_TO_CUSTOMER": 8,
		"STATUS_COMMON_CLOSED":            9,
		"STATUS_COMMON_ON_HOLD":           10,
	}
)

func (x Common) Enum() *Common {
	p := new(Common)
	*p = x
	return p
}

func (x Common) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Common) Descriptor() protoreflect.EnumDescriptor {
	return file_api_state_status_status_proto_enumTypes[0].Descriptor()
}

func (Common) Type() protoreflect.EnumType {
	return &file_api_state_status_status_proto_enumTypes[0]
}

func (x Common) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Common.Descriptor instead.
func (Common) EnumDescriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{0}
}

type Development int32

const (
	Development_STATUS_DEVELOPMENT_UNKNOWN           Development = 0
	Development_STATUS_DEVELOPMENT_BACKLOG           Development = 1
	Development_STATUS_DEVELOPMENT_IN_PROGRESS       Development = 2
	Development_STATUS_DEVELOPMENT_IN_REVIEW         Development = 3
	Development_STATUS_DEVELOPMENT_READY_FOR_TESTING Development = 4
	Development_STATUS_DEVELOPMENT_TESTING           Development = 5
	Development_STATUS_DEVELOPMENT_TESTING_PASSED    Development = 6
	Development_STATUS_DEVELOPMENT_DONE              Development = 7
	Development_STATUS_DEVELOPMENT_READY_TO_DEPLOY   Development = 8
	Development_STATUS_DEVELOPMENT_CLOSED            Development = 9
	Development_STATUS_DEVELOPMENT_ON_HOLD           Development = 10
)

// Enum value maps for Development.
var (
	Development_name = map[int32]string{
		0:  "STATUS_DEVELOPMENT_UNKNOWN",
		1:  "STATUS_DEVELOPMENT_BACKLOG",
		2:  "STATUS_DEVELOPMENT_IN_PROGRESS",
		3:  "STATUS_DEVELOPMENT_IN_REVIEW",
		4:  "STATUS_DEVELOPMENT_READY_FOR_TESTING",
		5:  "STATUS_DEVELOPMENT_TESTING",
		6:  "STATUS_DEVELOPMENT_TESTING_PASSED",
		7:  "STATUS_DEVELOPMENT_DONE",
		8:  "STATUS_DEVELOPMENT_READY_TO_DEPLOY",
		9:  "STATUS_DEVELOPMENT_CLOSED",
		10: "STATUS_DEVELOPMENT_ON_HOLD",
	}
	Development_value = map[string]int32{
		"STATUS_DEVELOPMENT_UNKNOWN":           0,
		"STATUS_DEVELOPMENT_BACKLOG":           1,
		"STATUS_DEVELOPMENT_IN_PROGRESS":       2,
		"STATUS_DEVELOPMENT_IN_REVIEW":         3,
		"STATUS_DEVELOPMENT_READY_FOR_TESTING": 4,
		"STATUS_DEVELOPMENT_TESTING":           5,
		"STATUS_DEVELOPMENT_TESTING_PASSED":    6,
		"STATUS_DEVELOPMENT_DONE":              7,
		"STATUS_DEVELOPMENT_READY_TO_DEPLOY":   8,
		"STATUS_DEVELOPMENT_CLOSED":            9,
		"STATUS_DEVELOPMENT_ON_HOLD":           10,
	}
)

func (x Development) Enum() *Development {
	p := new(Development)
	*p = x
	return p
}

func (x Development) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Development) Descriptor() protoreflect.EnumDescriptor {
	return file_api_state_status_status_proto_enumTypes[1].Descriptor()
}

func (Development) Type() protoreflect.EnumType {
	return &file_api_state_status_status_proto_enumTypes[1]
}

func (x Development) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Development.Descriptor instead.
func (Development) EnumDescriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{1}
}

type Epic int32

const (
	Epic_STATUS_EPIC_UNKNOWN   Epic = 0
	Epic_STATUS_EPIC_BACKLOG   Epic = 1
	Epic_STATUS_EPIC_PLANNING  Epic = 2
	Epic_STATUS_EPIC_DESIGNING Epic = 3
	Epic_STATUS_EPIC_IT        Epic = 4
	Epic_STATUS_EPIC_DONE      Epic = 5
	Epic_STATUS_EPIC_CLOSED    Epic = 6
	Epic_STATUS_EPIC_ON_HOLD   Epic = 7
)

// Enum value maps for Epic.
var (
	Epic_name = map[int32]string{
		0: "STATUS_EPIC_UNKNOWN",
		1: "STATUS_EPIC_BACKLOG",
		2: "STATUS_EPIC_PLANNING",
		3: "STATUS_EPIC_DESIGNING",
		4: "STATUS_EPIC_IT",
		5: "STATUS_EPIC_DONE",
		6: "STATUS_EPIC_CLOSED",
		7: "STATUS_EPIC_ON_HOLD",
	}
	Epic_value = map[string]int32{
		"STATUS_EPIC_UNKNOWN":   0,
		"STATUS_EPIC_BACKLOG":   1,
		"STATUS_EPIC_PLANNING":  2,
		"STATUS_EPIC_DESIGNING": 3,
		"STATUS_EPIC_IT":        4,
		"STATUS_EPIC_DONE":      5,
		"STATUS_EPIC_CLOSED":    6,
		"STATUS_EPIC_ON_HOLD":   7,
	}
)

func (x Epic) Enum() *Epic {
	p := new(Epic)
	*p = x
	return p
}

func (x Epic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Epic) Descriptor() protoreflect.EnumDescriptor {
	return file_api_state_status_status_proto_enumTypes[2].Descriptor()
}

func (Epic) Type() protoreflect.EnumType {
	return &file_api_state_status_status_proto_enumTypes[2]
}

func (x Epic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Epic.Descriptor instead.
func (Epic) EnumDescriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{2}
}

type Sprint int32

const (
	Sprint_STATUS_SPRINT_UNKNOWN  Sprint = 0
	Sprint_STATUS_SPRINT_ACTIVE   Sprint = 1
	Sprint_STATUS_SPRINT_FINISHED Sprint = 2
)

// Enum value maps for Sprint.
var (
	Sprint_name = map[int32]string{
		0: "STATUS_SPRINT_UNKNOWN",
		1: "STATUS_SPRINT_ACTIVE",
		2: "STATUS_SPRINT_FINISHED",
	}
	Sprint_value = map[string]int32{
		"STATUS_SPRINT_UNKNOWN":  0,
		"STATUS_SPRINT_ACTIVE":   1,
		"STATUS_SPRINT_FINISHED": 2,
	}
)

func (x Sprint) Enum() *Sprint {
	p := new(Sprint)
	*p = x
	return p
}

func (x Sprint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sprint) Descriptor() protoreflect.EnumDescriptor {
	return file_api_state_status_status_proto_enumTypes[3].Descriptor()
}

func (Sprint) Type() protoreflect.EnumType {
	return &file_api_state_status_status_proto_enumTypes[3]
}

func (x Sprint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sprint.Descriptor instead.
func (Sprint) EnumDescriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{3}
}

type IssueStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Status:
	//
	//	*IssueStatus_Common
	//	*IssueStatus_Development
	//	*IssueStatus_Epic
	Status isIssueStatus_Status `protobuf_oneof:"status"`
}

func (x *IssueStatus) Reset() {
	*x = IssueStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_state_status_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStatus) ProtoMessage() {}

func (x *IssueStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_state_status_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStatus.ProtoReflect.Descriptor instead.
func (*IssueStatus) Descriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{0}
}

func (m *IssueStatus) GetStatus() isIssueStatus_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (x *IssueStatus) GetCommon() *IssueStatusCommon {
	if x, ok := x.GetStatus().(*IssueStatus_Common); ok {
		return x.Common
	}
	return nil
}

func (x *IssueStatus) GetDevelopment() *IssueStatusDevelopment {
	if x, ok := x.GetStatus().(*IssueStatus_Development); ok {
		return x.Development
	}
	return nil
}

func (x *IssueStatus) GetEpic() *IssueStatusEpic {
	if x, ok := x.GetStatus().(*IssueStatus_Epic); ok {
		return x.Epic
	}
	return nil
}

type isIssueStatus_Status interface {
	isIssueStatus_Status()
}

type IssueStatus_Common struct {
	Common *IssueStatusCommon `protobuf:"bytes,1,opt,name=common,proto3,oneof"`
}

type IssueStatus_Development struct {
	Development *IssueStatusDevelopment `protobuf:"bytes,2,opt,name=development,proto3,oneof"`
}

type IssueStatus_Epic struct {
	Epic *IssueStatusEpic `protobuf:"bytes,3,opt,name=epic,proto3,oneof"`
}

func (*IssueStatus_Common) isIssueStatus_Status() {}

func (*IssueStatus_Development) isIssueStatus_Status() {}

func (*IssueStatus_Epic) isIssueStatus_Status() {}

type IssueStatusCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Common `protobuf:"varint,1,opt,name=status,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.Common" json:"status,omitempty"`
}

func (x *IssueStatusCommon) Reset() {
	*x = IssueStatusCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_state_status_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStatusCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStatusCommon) ProtoMessage() {}

func (x *IssueStatusCommon) ProtoReflect() protoreflect.Message {
	mi := &file_api_state_status_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStatusCommon.ProtoReflect.Descriptor instead.
func (*IssueStatusCommon) Descriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{1}
}

func (x *IssueStatusCommon) GetStatus() Common {
	if x != nil {
		return x.Status
	}
	return Common_STATUS_COMMON_UNKNOWN
}

type IssueStatusDevelopment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Development `protobuf:"varint,2,opt,name=status,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.Development" json:"status,omitempty"`
}

func (x *IssueStatusDevelopment) Reset() {
	*x = IssueStatusDevelopment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_state_status_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStatusDevelopment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStatusDevelopment) ProtoMessage() {}

func (x *IssueStatusDevelopment) ProtoReflect() protoreflect.Message {
	mi := &file_api_state_status_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStatusDevelopment.ProtoReflect.Descriptor instead.
func (*IssueStatusDevelopment) Descriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{2}
}

func (x *IssueStatusDevelopment) GetStatus() Development {
	if x != nil {
		return x.Status
	}
	return Development_STATUS_DEVELOPMENT_UNKNOWN
}

type IssueStatusEpic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Epic `protobuf:"varint,2,opt,name=status,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.Epic" json:"status,omitempty"`
}

func (x *IssueStatusEpic) Reset() {
	*x = IssueStatusEpic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_state_status_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStatusEpic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStatusEpic) ProtoMessage() {}

func (x *IssueStatusEpic) ProtoReflect() protoreflect.Message {
	mi := &file_api_state_status_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStatusEpic.ProtoReflect.Descriptor instead.
func (*IssueStatusEpic) Descriptor() ([]byte, []int) {
	return file_api_state_status_status_proto_rawDescGZIP(), []int{3}
}

func (x *IssueStatusEpic) GetStatus() Epic {
	if x != nil {
		return x.Status
	}
	return Epic_STATUS_EPIC_UNKNOWN
}

var File_api_state_status_status_proto protoreflect.FileDescriptor

var file_api_state_status_status_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd5, 0x02, 0x0a, 0x0b, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x65, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37,
	0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x12, 0x74, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x04, 0x65, 0x70, 0x69, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x70, 0x69, 0x63, 0x48,
	0x00, 0x52, 0x04, 0x65, 0x70, 0x69, 0x63, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x6d, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63,
	0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x77, 0x0a, 0x16, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32,
	0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x69, 0x0a, 0x0f, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x70, 0x69, 0x63, 0x12, 0x56, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x63, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2a, 0xd2, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x43, 0x4b,
	0x4c, 0x4f, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x04,
	0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f,
	0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x05, 0x12, 0x1f, 0x0a,
	0x1b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x1f,
	0x0a, 0x1b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x07, 0x12,
	0x23, 0x0a, 0x1f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x5f, 0x47, 0x49, 0x56, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x09, 0x12, 0x19,
	0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f,
	0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x0a, 0x2a, 0x88, 0x03, 0x0a, 0x0b, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x42, 0x41, 0x43, 0x4b, 0x4c, 0x4f, 0x47, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x20, 0x0a,
	0x1c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03, 0x12,
	0x28, 0x0a, 0x24, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f,
	0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x06,
	0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c,
	0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x07, 0x12, 0x26, 0x0a,
	0x22, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x44, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44,
	0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x48, 0x4f,
	0x4c, 0x44, 0x10, 0x0a, 0x2a, 0xc8, 0x01, 0x0a, 0x04, 0x45, 0x70, 0x69, 0x63, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x4c, 0x4f, 0x47, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x50,
	0x4c, 0x41, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45,
	0x50, 0x49, 0x43, 0x5f, 0x49, 0x54, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x05, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x43, 0x4c,
	0x4f, 0x53, 0x45, 0x44, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x45, 0x50, 0x49, 0x43, 0x5f, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x07, 0x2a,
	0x59, 0x0a, 0x06, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1a,
	0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x4b, 0x2f, 0x63, 0x72, 0x6e, 0x74, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_state_status_status_proto_rawDescOnce sync.Once
	file_api_state_status_status_proto_rawDescData = file_api_state_status_status_proto_rawDesc
)

func file_api_state_status_status_proto_rawDescGZIP() []byte {
	file_api_state_status_status_proto_rawDescOnce.Do(func() {
		file_api_state_status_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_state_status_status_proto_rawDescData)
	})
	return file_api_state_status_status_proto_rawDescData
}

var file_api_state_status_status_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_state_status_status_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_state_status_status_proto_goTypes = []interface{}{
	(Common)(0),                    // 0: github.constantine27k.crnt_data_manager.api.state.status.Common
	(Development)(0),               // 1: github.constantine27k.crnt_data_manager.api.state.status.Development
	(Epic)(0),                      // 2: github.constantine27k.crnt_data_manager.api.state.status.Epic
	(Sprint)(0),                    // 3: github.constantine27k.crnt_data_manager.api.state.status.Sprint
	(*IssueStatus)(nil),            // 4: github.constantine27k.crnt_data_manager.api.state.status.IssueStatus
	(*IssueStatusCommon)(nil),      // 5: github.constantine27k.crnt_data_manager.api.state.status.IssueStatusCommon
	(*IssueStatusDevelopment)(nil), // 6: github.constantine27k.crnt_data_manager.api.state.status.IssueStatusDevelopment
	(*IssueStatusEpic)(nil),        // 7: github.constantine27k.crnt_data_manager.api.state.status.IssueStatusEpic
}
var file_api_state_status_status_proto_depIdxs = []int32{
	5, // 0: github.constantine27k.crnt_data_manager.api.state.status.IssueStatus.common:type_name -> github.constantine27k.crnt_data_manager.api.state.status.IssueStatusCommon
	6, // 1: github.constantine27k.crnt_data_manager.api.state.status.IssueStatus.development:type_name -> github.constantine27k.crnt_data_manager.api.state.status.IssueStatusDevelopment
	7, // 2: github.constantine27k.crnt_data_manager.api.state.status.IssueStatus.epic:type_name -> github.constantine27k.crnt_data_manager.api.state.status.IssueStatusEpic
	0, // 3: github.constantine27k.crnt_data_manager.api.state.status.IssueStatusCommon.status:type_name -> github.constantine27k.crnt_data_manager.api.state.status.Common
	1, // 4: github.constantine27k.crnt_data_manager.api.state.status.IssueStatusDevelopment.status:type_name -> github.constantine27k.crnt_data_manager.api.state.status.Development
	2, // 5: github.constantine27k.crnt_data_manager.api.state.status.IssueStatusEpic.status:type_name -> github.constantine27k.crnt_data_manager.api.state.status.Epic
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_state_status_status_proto_init() }
func file_api_state_status_status_proto_init() {
	if File_api_state_status_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_state_status_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStatus); i {
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
		file_api_state_status_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStatusCommon); i {
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
		file_api_state_status_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStatusDevelopment); i {
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
		file_api_state_status_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStatusEpic); i {
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
	file_api_state_status_status_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*IssueStatus_Common)(nil),
		(*IssueStatus_Development)(nil),
		(*IssueStatus_Epic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_state_status_status_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_state_status_status_proto_goTypes,
		DependencyIndexes: file_api_state_status_status_proto_depIdxs,
		EnumInfos:         file_api_state_status_status_proto_enumTypes,
		MessageInfos:      file_api_state_status_status_proto_msgTypes,
	}.Build()
	File_api_state_status_status_proto = out.File
	file_api_state_status_status_proto_rawDesc = nil
	file_api_state_status_status_proto_goTypes = nil
	file_api_state_status_status_proto_depIdxs = nil
}
