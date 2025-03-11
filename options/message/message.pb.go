// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: options/message.proto

package message

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_options_message_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         3001,
		Name:          "jsonschema.message.ignore",
		Tag:           "varint,3001,opt,name=ignore",
		Filename:      "options/message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         3002,
		Name:          "jsonschema.message.all_fields_required",
		Tag:           "varint,3002,opt,name=all_fields_required",
		Filename:      "options/message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         3003,
		Name:          "jsonschema.message.allow_null_values",
		Tag:           "varint,3003,opt,name=allow_null_values",
		Filename:      "options/message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         3004,
		Name:          "jsonschema.message.disallow_additional_properties",
		Tag:           "varint,3004,opt,name=disallow_additional_properties",
		Filename:      "options/message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         3005,
		Name:          "jsonschema.message.enums_as_constants",
		Tag:           "varint,3005,opt,name=enums_as_constants",
		Filename:      "options/message.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Messages tagged with this will not be processed
	//
	// optional bool ignore = 3001;
	E_Ignore = &file_options_message_proto_extTypes[0]
	// Messages tagged with this will have all fields marked as "required":
	//
	// optional bool all_fields_required = 3002;
	E_AllFieldsRequired = &file_options_message_proto_extTypes[1]
	// Messages tagged with this will additionally accept null values for all properties:
	//
	// optional bool allow_null_values = 3003;
	E_AllowNullValues = &file_options_message_proto_extTypes[2]
	// Messages tagged with this will have all fields marked as not allowing additional properties:
	//
	// optional bool disallow_additional_properties = 3004;
	E_DisallowAdditionalProperties = &file_options_message_proto_extTypes[3]
	// Messages tagged with this will have all nested enums encoded to use constants instead of simple types (supports value annotations):
	//
	// optional bool enums_as_constants = 3005;
	E_EnumsAsConstants = &file_options_message_proto_extTypes[4]
)

var File_options_message_proto protoreflect.FileDescriptor

var file_options_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x38, 0x0a,
	0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x50, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xba, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x4c, 0x0a, 0x11, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xbb, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x75, 0x6c,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x3a, 0x66, 0x0a, 0x1e, 0x64, 0x69, 0x73, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbc, 0x17, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x3a,
	0x4e, 0x0a, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x41, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x42,
	0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6a,
	0x61, 0x73, 0x6f, 0x6e, 0x6c, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_options_message_proto_goTypes = []any{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_options_message_proto_depIdxs = []int32{
	0, // 0: jsonschema.message.ignore:extendee -> google.protobuf.MessageOptions
	0, // 1: jsonschema.message.all_fields_required:extendee -> google.protobuf.MessageOptions
	0, // 2: jsonschema.message.allow_null_values:extendee -> google.protobuf.MessageOptions
	0, // 3: jsonschema.message.disallow_additional_properties:extendee -> google.protobuf.MessageOptions
	0, // 4: jsonschema.message.enums_as_constants:extendee -> google.protobuf.MessageOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_message_proto_init() }
func file_options_message_proto_init() {
	if File_options_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_options_message_proto_goTypes,
		DependencyIndexes: file_options_message_proto_depIdxs,
		ExtensionInfos:    file_options_message_proto_extTypes,
	}.Build()
	File_options_message_proto = out.File
	file_options_message_proto_rawDesc = nil
	file_options_message_proto_goTypes = nil
	file_options_message_proto_depIdxs = nil
}
