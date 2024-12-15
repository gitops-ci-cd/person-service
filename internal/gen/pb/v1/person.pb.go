// edition = "2023";

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: com/acme/schema/v1/person.proto

package v1

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

type PersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *PersonRequest) Reset() {
	*x = PersonRequest{}
	mi := &file_com_acme_schema_v1_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRequest) ProtoMessage() {}

func (x *PersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_acme_schema_v1_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRequest.ProtoReflect.Descriptor instead.
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return file_com_acme_schema_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *PersonRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type PersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *PersonResponse) Reset() {
	*x = PersonResponse{}
	mi := &file_com_acme_schema_v1_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonResponse) ProtoMessage() {}

func (x *PersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_acme_schema_v1_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonResponse.ProtoReflect.Descriptor instead.
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return file_com_acme_schema_v1_person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PersonResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_com_acme_schema_v1_person_proto protoreflect.FileDescriptor

var file_com_acme_schema_v1_person_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x23, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x0e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x5f, 0x0a, 0x0d, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73,
	0x2d, 0x63, 0x69, 0x2d, 0x63, 0x64, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_acme_schema_v1_person_proto_rawDescOnce sync.Once
	file_com_acme_schema_v1_person_proto_rawDescData = file_com_acme_schema_v1_person_proto_rawDesc
)

func file_com_acme_schema_v1_person_proto_rawDescGZIP() []byte {
	file_com_acme_schema_v1_person_proto_rawDescOnce.Do(func() {
		file_com_acme_schema_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_acme_schema_v1_person_proto_rawDescData)
	})
	return file_com_acme_schema_v1_person_proto_rawDescData
}

var file_com_acme_schema_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_acme_schema_v1_person_proto_goTypes = []any{
	(*PersonRequest)(nil),  // 0: com.acme.schema.v1.PersonRequest
	(*PersonResponse)(nil), // 1: com.acme.schema.v1.PersonResponse
}
var file_com_acme_schema_v1_person_proto_depIdxs = []int32{
	0, // 0: com.acme.schema.v1.PersonService.Fetch:input_type -> com.acme.schema.v1.PersonRequest
	1, // 1: com.acme.schema.v1.PersonService.Fetch:output_type -> com.acme.schema.v1.PersonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_acme_schema_v1_person_proto_init() }
func file_com_acme_schema_v1_person_proto_init() {
	if File_com_acme_schema_v1_person_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_acme_schema_v1_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_acme_schema_v1_person_proto_goTypes,
		DependencyIndexes: file_com_acme_schema_v1_person_proto_depIdxs,
		MessageInfos:      file_com_acme_schema_v1_person_proto_msgTypes,
	}.Build()
	File_com_acme_schema_v1_person_proto = out.File
	file_com_acme_schema_v1_person_proto_rawDesc = nil
	file_com_acme_schema_v1_person_proto_goTypes = nil
	file_com_acme_schema_v1_person_proto_depIdxs = nil
}