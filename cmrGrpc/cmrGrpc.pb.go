// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: cmrGrpc.proto

package cmrGrpc

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

type NewCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shortname string `protobuf:"bytes,1,opt,name=shortname,proto3" json:"shortname,omitempty"`
	Version   int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *NewCollection) Reset() {
	*x = NewCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmrGrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCollection) ProtoMessage() {}

func (x *NewCollection) ProtoReflect() protoreflect.Message {
	mi := &file_cmrGrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCollection.ProtoReflect.Descriptor instead.
func (*NewCollection) Descriptor() ([]byte, []int) {
	return file_cmrGrpc_proto_rawDescGZIP(), []int{0}
}

func (x *NewCollection) GetShortname() string {
	if x != nil {
		return x.Shortname
	}
	return ""
}

func (x *NewCollection) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shortname string `protobuf:"bytes,1,opt,name=shortname,proto3" json:"shortname,omitempty"`
	Version   int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	ConceptId string `protobuf:"bytes,3,opt,name=conceptId,proto3" json:"conceptId,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmrGrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_cmrGrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_cmrGrpc_proto_rawDescGZIP(), []int{1}
}

func (x *Collection) GetShortname() string {
	if x != nil {
		return x.Shortname
	}
	return ""
}

func (x *Collection) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Collection) GetConceptId() string {
	if x != nil {
		return x.ConceptId
	}
	return ""
}

var File_cmrGrpc_proto protoreflect.FileDescriptor

var file_cmrGrpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6d, 0x72, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6d, 0x72, 0x47, 0x72, 0x70, 0x63, 0x22, 0x47, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x62, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x74, 0x49, 0x64, 0x32, 0x4f, 0x0a, 0x07, 0x43, 0x6d, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x12, 0x44, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x6d, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x13, 0x2e, 0x63, 0x6d, 0x72, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6d, 0x72, 0x47, 0x72, 0x70, 0x63, 0x3b, 0x63, 0x6d,
	0x72, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmrGrpc_proto_rawDescOnce sync.Once
	file_cmrGrpc_proto_rawDescData = file_cmrGrpc_proto_rawDesc
)

func file_cmrGrpc_proto_rawDescGZIP() []byte {
	file_cmrGrpc_proto_rawDescOnce.Do(func() {
		file_cmrGrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmrGrpc_proto_rawDescData)
	})
	return file_cmrGrpc_proto_rawDescData
}

var file_cmrGrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmrGrpc_proto_goTypes = []interface{}{
	(*NewCollection)(nil), // 0: cmrGrpc.NewCollection
	(*Collection)(nil),    // 1: cmrGrpc.Collection
}
var file_cmrGrpc_proto_depIdxs = []int32{
	0, // 0: cmrGrpc.CmrGrpc.CreateNewCollection:input_type -> cmrGrpc.NewCollection
	1, // 1: cmrGrpc.CmrGrpc.CreateNewCollection:output_type -> cmrGrpc.Collection
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmrGrpc_proto_init() }
func file_cmrGrpc_proto_init() {
	if File_cmrGrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmrGrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCollection); i {
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
		file_cmrGrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
			RawDescriptor: file_cmrGrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmrGrpc_proto_goTypes,
		DependencyIndexes: file_cmrGrpc_proto_depIdxs,
		MessageInfos:      file_cmrGrpc_proto_msgTypes,
	}.Build()
	File_cmrGrpc_proto = out.File
	file_cmrGrpc_proto_rawDesc = nil
	file_cmrGrpc_proto_goTypes = nil
	file_cmrGrpc_proto_depIdxs = nil
}
