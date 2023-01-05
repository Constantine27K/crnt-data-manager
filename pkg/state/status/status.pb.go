// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: state/status/status.proto

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

type TaskCommon int32

const (
	TaskCommon_STATUS_COMMON_UNKNOWN           TaskCommon = 0
	TaskCommon_STATUS_COMMON_BACKLOG           TaskCommon = 1
	TaskCommon_STATUS_COMMON_IN_PROGRESS       TaskCommon = 2
	TaskCommon_STATUS_COMMON_DONE              TaskCommon = 3
	TaskCommon_STATUS_COMMON_READY_FOR_REVIEW  TaskCommon = 4
	TaskCommon_STATUS_COMMON_IN_REVIEW         TaskCommon = 5
	TaskCommon_STATUS_COMMON_REVIEW_FAILED     TaskCommon = 6
	TaskCommon_STATUS_COMMON_REVIEW_PASSED     TaskCommon = 7
	TaskCommon_STATUS_COMMON_GIVEN_TO_CUSTOMER TaskCommon = 8
	TaskCommon_STATUS_COMMON_CLOSED            TaskCommon = 9
	TaskCommon_STATUS_COMMON_ON_HOLD           TaskCommon = 10
)

// Enum value maps for TaskCommon.
var (
	TaskCommon_name = map[int32]string{
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
	TaskCommon_value = map[string]int32{
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

func (x TaskCommon) Enum() *TaskCommon {
	p := new(TaskCommon)
	*p = x
	return p
}

func (x TaskCommon) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskCommon) Descriptor() protoreflect.EnumDescriptor {
	return file_state_status_status_proto_enumTypes[0].Descriptor()
}

func (TaskCommon) Type() protoreflect.EnumType {
	return &file_state_status_status_proto_enumTypes[0]
}

func (x TaskCommon) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskCommon.Descriptor instead.
func (TaskCommon) EnumDescriptor() ([]byte, []int) {
	return file_state_status_status_proto_rawDescGZIP(), []int{0}
}

type TaskDevelopment int32

const (
	TaskDevelopment_STATUS_DEVELOPMENT_UNKNOWN           TaskDevelopment = 0
	TaskDevelopment_STATUS_DEVELOPMENT_BACKLOG           TaskDevelopment = 1
	TaskDevelopment_STATUS_DEVELOPMENT_IN_PROGRESS       TaskDevelopment = 2
	TaskDevelopment_STATUS_DEVELOPMENT_IN_REVIEW         TaskDevelopment = 3
	TaskDevelopment_STATUS_DEVELOPMENT_READY_FOR_TESTING TaskDevelopment = 4
	TaskDevelopment_STATUS_DEVELOPMENT_TESTING           TaskDevelopment = 5
	TaskDevelopment_STATUS_DEVELOPMENT_TESTING_PASSED    TaskDevelopment = 6
	TaskDevelopment_STATUS_DEVELOPMENT_DONE              TaskDevelopment = 7
	TaskDevelopment_STATUS_DEVELOPMENT_READY_TO_DEPLOY   TaskDevelopment = 8
	TaskDevelopment_STATUS_DEVELOPMENT_CLOSED            TaskDevelopment = 9
	TaskDevelopment_STATUS_DEVELOPMENT_ON_HOLD           TaskDevelopment = 10
)

// Enum value maps for TaskDevelopment.
var (
	TaskDevelopment_name = map[int32]string{
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
	TaskDevelopment_value = map[string]int32{
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

func (x TaskDevelopment) Enum() *TaskDevelopment {
	p := new(TaskDevelopment)
	*p = x
	return p
}

func (x TaskDevelopment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskDevelopment) Descriptor() protoreflect.EnumDescriptor {
	return file_state_status_status_proto_enumTypes[1].Descriptor()
}

func (TaskDevelopment) Type() protoreflect.EnumType {
	return &file_state_status_status_proto_enumTypes[1]
}

func (x TaskDevelopment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskDevelopment.Descriptor instead.
func (TaskDevelopment) EnumDescriptor() ([]byte, []int) {
	return file_state_status_status_proto_rawDescGZIP(), []int{1}
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
	return file_state_status_status_proto_enumTypes[2].Descriptor()
}

func (Sprint) Type() protoreflect.EnumType {
	return &file_state_status_status_proto_enumTypes[2]
}

func (x Sprint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sprint.Descriptor instead.
func (Sprint) EnumDescriptor() ([]byte, []int) {
	return file_state_status_status_proto_rawDescGZIP(), []int{2}
}

type TaskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Status:
	//
	//	*TaskStatus_Common
	//	*TaskStatus_Development
	Status isTaskStatus_Status `protobuf_oneof:"status"`
}

func (x *TaskStatus) Reset() {
	*x = TaskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_status_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStatus) ProtoMessage() {}

func (x *TaskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_state_status_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStatus.ProtoReflect.Descriptor instead.
func (*TaskStatus) Descriptor() ([]byte, []int) {
	return file_state_status_status_proto_rawDescGZIP(), []int{0}
}

func (m *TaskStatus) GetStatus() isTaskStatus_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (x *TaskStatus) GetCommon() TaskCommon {
	if x, ok := x.GetStatus().(*TaskStatus_Common); ok {
		return x.Common
	}
	return TaskCommon_STATUS_COMMON_UNKNOWN
}

func (x *TaskStatus) GetDevelopment() TaskDevelopment {
	if x, ok := x.GetStatus().(*TaskStatus_Development); ok {
		return x.Development
	}
	return TaskDevelopment_STATUS_DEVELOPMENT_UNKNOWN
}

type isTaskStatus_Status interface {
	isTaskStatus_Status()
}

type TaskStatus_Common struct {
	Common TaskCommon `protobuf:"varint,1,opt,name=common,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.TaskCommon,oneof"`
}

type TaskStatus_Development struct {
	Development TaskDevelopment `protobuf:"varint,2,opt,name=development,proto3,enum=github.constantine27k.crnt_data_manager.api.state.status.TaskDevelopment,oneof"`
}

func (*TaskStatus_Common) isTaskStatus_Status() {}

func (*TaskStatus_Development) isTaskStatus_Status() {}

var File_state_status_status_proto protoreflect.FileDescriptor

var file_state_status_status_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32,
	0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x5e, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37,
	0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xd6, 0x02,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x4c, 0x4f, 0x47,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x04, 0x12, 0x1b, 0x0a,
	0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x49,
	0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x07, 0x12, 0x23, 0x0a, 0x1f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x47, 0x49,
	0x56, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10,
	0x08, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x4f, 0x4e, 0x5f,
	0x48, 0x4f, 0x4c, 0x44, 0x10, 0x0a, 0x2a, 0x8c, 0x03, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x42, 0x41, 0x43, 0x4b, 0x4c, 0x4f, 0x47, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03,
	0x12, 0x28, 0x0a, 0x24, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c,
	0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45,
	0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x07, 0x12, 0x26,
	0x0a, 0x22, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x48,
	0x4f, 0x4c, 0x44, 0x10, 0x0a, 0x2a, 0x59, 0x0a, 0x06, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x54,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x50, 0x52, 0x49, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x4b, 0x2f, 0x63, 0x72,
	0x6e, 0x74, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_status_status_proto_rawDescOnce sync.Once
	file_state_status_status_proto_rawDescData = file_state_status_status_proto_rawDesc
)

func file_state_status_status_proto_rawDescGZIP() []byte {
	file_state_status_status_proto_rawDescOnce.Do(func() {
		file_state_status_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_status_status_proto_rawDescData)
	})
	return file_state_status_status_proto_rawDescData
}

var file_state_status_status_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_state_status_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_state_status_status_proto_goTypes = []interface{}{
	(TaskCommon)(0),      // 0: github.constantine27k.crnt_data_manager.api.state.status.TaskCommon
	(TaskDevelopment)(0), // 1: github.constantine27k.crnt_data_manager.api.state.status.TaskDevelopment
	(Sprint)(0),          // 2: github.constantine27k.crnt_data_manager.api.state.status.Sprint
	(*TaskStatus)(nil),   // 3: github.constantine27k.crnt_data_manager.api.state.status.TaskStatus
}
var file_state_status_status_proto_depIdxs = []int32{
	0, // 0: github.constantine27k.crnt_data_manager.api.state.status.TaskStatus.common:type_name -> github.constantine27k.crnt_data_manager.api.state.status.TaskCommon
	1, // 1: github.constantine27k.crnt_data_manager.api.state.status.TaskStatus.development:type_name -> github.constantine27k.crnt_data_manager.api.state.status.TaskDevelopment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_state_status_status_proto_init() }
func file_state_status_status_proto_init() {
	if File_state_status_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_status_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStatus); i {
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
	file_state_status_status_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TaskStatus_Common)(nil),
		(*TaskStatus_Development)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_state_status_status_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_status_status_proto_goTypes,
		DependencyIndexes: file_state_status_status_proto_depIdxs,
		EnumInfos:         file_state_status_status_proto_enumTypes,
		MessageInfos:      file_state_status_status_proto_msgTypes,
	}.Build()
	File_state_status_status_proto = out.File
	file_state_status_status_proto_rawDesc = nil
	file_state_status_status_proto_goTypes = nil
	file_state_status_status_proto_depIdxs = nil
}
