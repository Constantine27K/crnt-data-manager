// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: state/template/template.proto

package template

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

type Template int32

const (
	Template_TEMPLATE_UNKNOWN     Template = 0
	Template_TEMPLATE_COMMON      Template = 1
	Template_TEMPLATE_DEVELOPMENT Template = 2
)

// Enum value maps for Template.
var (
	Template_name = map[int32]string{
		0: "TEMPLATE_UNKNOWN",
		1: "TEMPLATE_COMMON",
		2: "TEMPLATE_DEVELOPMENT",
	}
	Template_value = map[string]int32{
		"TEMPLATE_UNKNOWN":     0,
		"TEMPLATE_COMMON":      1,
		"TEMPLATE_DEVELOPMENT": 2,
	}
)

func (x Template) Enum() *Template {
	p := new(Template)
	*p = x
	return p
}

func (x Template) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Template) Descriptor() protoreflect.EnumDescriptor {
	return file_state_template_template_proto_enumTypes[0].Descriptor()
}

func (Template) Type() protoreflect.EnumType {
	return &file_state_template_template_proto_enumTypes[0]
}

func (x Template) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Template.Descriptor instead.
func (Template) EnumDescriptor() ([]byte, []int) {
	return file_state_template_template_proto_rawDescGZIP(), []int{0}
}

var File_state_template_template_proto protoreflect.FileDescriptor

var file_state_template_template_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x32, 0x37, 0x6b, 0x2e, 0x63, 0x72, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2a, 0x4f, 0x0a, 0x08, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45, 0x4d, 0x50, 0x4c,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x44,
	0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x49, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x37, 0x4b, 0x2f, 0x63, 0x72, 0x6e, 0x74, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x3b, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_template_template_proto_rawDescOnce sync.Once
	file_state_template_template_proto_rawDescData = file_state_template_template_proto_rawDesc
)

func file_state_template_template_proto_rawDescGZIP() []byte {
	file_state_template_template_proto_rawDescOnce.Do(func() {
		file_state_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_template_template_proto_rawDescData)
	})
	return file_state_template_template_proto_rawDescData
}

var file_state_template_template_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_state_template_template_proto_goTypes = []interface{}{
	(Template)(0), // 0: github.constantine27k.crnt_data_manager.api.state.template.Template
}
var file_state_template_template_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_state_template_template_proto_init() }
func file_state_template_template_proto_init() {
	if File_state_template_template_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_state_template_template_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_template_template_proto_goTypes,
		DependencyIndexes: file_state_template_template_proto_depIdxs,
		EnumInfos:         file_state_template_template_proto_enumTypes,
	}.Build()
	File_state_template_template_proto = out.File
	file_state_template_template_proto_rawDesc = nil
	file_state_template_template_proto_goTypes = nil
	file_state_template_template_proto_depIdxs = nil
}
