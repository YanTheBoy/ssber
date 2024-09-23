// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.17.3
// source: api/proto/fm_flats.proto

package flats_manager_api

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

type FlatType int32

const (
	FlatType_FLAT_TYPE_UNSPECIFIED FlatType = 0
	FlatType_FLAT_TYPE_APARTMENT   FlatType = 1
	FlatType_FLAT_TYPE_ROOM        FlatType = 2
)

// Enum value maps for FlatType.
var (
	FlatType_name = map[int32]string{
		0: "FLAT_TYPE_UNSPECIFIED",
		1: "FLAT_TYPE_APARTMENT",
		2: "FLAT_TYPE_ROOM",
	}
	FlatType_value = map[string]int32{
		"FLAT_TYPE_UNSPECIFIED": 0,
		"FLAT_TYPE_APARTMENT":   1,
		"FLAT_TYPE_ROOM":        2,
	}
)

func (x FlatType) Enum() *FlatType {
	p := new(FlatType)
	*p = x
	return p
}

func (x FlatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlatType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_fm_flats_proto_enumTypes[0].Descriptor()
}

func (FlatType) Type() protoreflect.EnumType {
	return &file_api_proto_fm_flats_proto_enumTypes[0]
}

func (x FlatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlatType.Descriptor instead.
func (FlatType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_fm_flats_proto_rawDescGZIP(), []int{0}
}

type CreateFlatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      FlatType `protobuf:"varint,3,opt,name=type,proto3,enum=FlatType" json:"type,omitempty"`
	Price     string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	CreatedBy string   `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *CreateFlatRequest) Reset() {
	*x = CreateFlatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_fm_flats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlatRequest) ProtoMessage() {}

func (x *CreateFlatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_fm_flats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlatRequest.ProtoReflect.Descriptor instead.
func (*CreateFlatRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_fm_flats_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateFlatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFlatRequest) GetType() FlatType {
	if x != nil {
		return x.Type
	}
	return FlatType_FLAT_TYPE_UNSPECIFIED
}

func (x *CreateFlatRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateFlatRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type Flat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Flat) Reset() {
	*x = Flat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_fm_flats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flat) ProtoMessage() {}

func (x *Flat) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_fm_flats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flat.ProtoReflect.Descriptor instead.
func (*Flat) Descriptor() ([]byte, []int) {
	return file_api_proto_fm_flats_proto_rawDescGZIP(), []int{1}
}

func (x *Flat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_proto_fm_flats_proto protoreflect.FileDescriptor

var file_api_proto_fm_flats_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6d, 0x5f, 0x66,
	0x6c, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x09, 0x2e, 0x46, 0x6c, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x2a, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x52, 0x0a, 0x08, 0x46, 0x6c, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x46, 0x4c, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46,
	0x4c, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x41, 0x52, 0x54, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4c, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x32, 0x39, 0x0a, 0x0c, 0x46, 0x6c, 0x61, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x46, 0x6c, 0x61,
	0x74, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x66, 0x6c, 0x61, 0x74, 0x73, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_fm_flats_proto_rawDescOnce sync.Once
	file_api_proto_fm_flats_proto_rawDescData = file_api_proto_fm_flats_proto_rawDesc
)

func file_api_proto_fm_flats_proto_rawDescGZIP() []byte {
	file_api_proto_fm_flats_proto_rawDescOnce.Do(func() {
		file_api_proto_fm_flats_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_fm_flats_proto_rawDescData)
	})
	return file_api_proto_fm_flats_proto_rawDescData
}

var file_api_proto_fm_flats_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_fm_flats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_fm_flats_proto_goTypes = []any{
	(FlatType)(0),             // 0: FlatType
	(*CreateFlatRequest)(nil), // 1: CreateFlatRequest
	(*Flat)(nil),              // 2: Flat
}
var file_api_proto_fm_flats_proto_depIdxs = []int32{
	0, // 0: CreateFlatRequest.type:type_name -> FlatType
	1, // 1: FlatsService.CreateFlat:input_type -> CreateFlatRequest
	2, // 2: FlatsService.CreateFlat:output_type -> Flat
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_fm_flats_proto_init() }
func file_api_proto_fm_flats_proto_init() {
	if File_api_proto_fm_flats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_fm_flats_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFlatRequest); i {
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
		file_api_proto_fm_flats_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Flat); i {
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
			RawDescriptor: file_api_proto_fm_flats_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_fm_flats_proto_goTypes,
		DependencyIndexes: file_api_proto_fm_flats_proto_depIdxs,
		EnumInfos:         file_api_proto_fm_flats_proto_enumTypes,
		MessageInfos:      file_api_proto_fm_flats_proto_msgTypes,
	}.Build()
	File_api_proto_fm_flats_proto = out.File
	file_api_proto_fm_flats_proto_rawDesc = nil
	file_api_proto_fm_flats_proto_goTypes = nil
	file_api_proto_fm_flats_proto_depIdxs = nil
}