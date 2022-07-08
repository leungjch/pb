// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: style_transfer.proto

package __

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StyleTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentImage []byte `protobuf:"bytes,1,opt,name=content_image,json=contentImage,proto3" json:"content_image,omitempty"`
	StyleImage   []byte `protobuf:"bytes,2,opt,name=style_image,json=styleImage,proto3" json:"style_image,omitempty"`
	// Must be between 0 to 1
	Alpha float32 `protobuf:"fixed32,3,opt,name=alpha,proto3" json:"alpha,omitempty"`
}

func (x *StyleTransferRequest) Reset() {
	*x = StyleTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_style_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StyleTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StyleTransferRequest) ProtoMessage() {}

func (x *StyleTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_style_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StyleTransferRequest.ProtoReflect.Descriptor instead.
func (*StyleTransferRequest) Descriptor() ([]byte, []int) {
	return file_style_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *StyleTransferRequest) GetContentImage() []byte {
	if x != nil {
		return x.ContentImage
	}
	return nil
}

func (x *StyleTransferRequest) GetStyleImage() []byte {
	if x != nil {
		return x.StyleImage
	}
	return nil
}

func (x *StyleTransferRequest) GetAlpha() float32 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

type StyleTransferURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentImageUrl string  `protobuf:"bytes,1,opt,name=content_image_url,json=contentImageUrl,proto3" json:"content_image_url,omitempty"`
	StyleImageUrl   string  `protobuf:"bytes,2,opt,name=style_image_url,json=styleImageUrl,proto3" json:"style_image_url,omitempty"`
	Alpha           float32 `protobuf:"fixed32,3,opt,name=alpha,proto3" json:"alpha,omitempty"`
}

func (x *StyleTransferURLRequest) Reset() {
	*x = StyleTransferURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_style_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StyleTransferURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StyleTransferURLRequest) ProtoMessage() {}

func (x *StyleTransferURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_style_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StyleTransferURLRequest.ProtoReflect.Descriptor instead.
func (*StyleTransferURLRequest) Descriptor() ([]byte, []int) {
	return file_style_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *StyleTransferURLRequest) GetContentImageUrl() string {
	if x != nil {
		return x.ContentImageUrl
	}
	return ""
}

func (x *StyleTransferURLRequest) GetStyleImageUrl() string {
	if x != nil {
		return x.StyleImageUrl
	}
	return ""
}

func (x *StyleTransferURLRequest) GetAlpha() float32 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

type StyleTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StylizedImage []byte `protobuf:"bytes,1,opt,name=stylized_image,json=stylizedImage,proto3" json:"stylized_image,omitempty"`
}

func (x *StyleTransferResponse) Reset() {
	*x = StyleTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_style_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StyleTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StyleTransferResponse) ProtoMessage() {}

func (x *StyleTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_style_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StyleTransferResponse.ProtoReflect.Descriptor instead.
func (*StyleTransferResponse) Descriptor() ([]byte, []int) {
	return file_style_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *StyleTransferResponse) GetStylizedImage() []byte {
	if x != nil {
		return x.StylizedImage
	}
	return nil
}

var File_style_transfer_proto protoreflect.FileDescriptor

var file_style_transfer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x79, 0x6c, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x3e,
	0x0a, 0x15, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x79, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x73, 0x74, 0x79, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x32, 0xe4,
	0x01, 0x0a, 0x14, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0d, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22,
	0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x10,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x12, 0x18, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x75,
	0x72, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_style_transfer_proto_rawDescOnce sync.Once
	file_style_transfer_proto_rawDescData = file_style_transfer_proto_rawDesc
)

func file_style_transfer_proto_rawDescGZIP() []byte {
	file_style_transfer_proto_rawDescOnce.Do(func() {
		file_style_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_style_transfer_proto_rawDescData)
	})
	return file_style_transfer_proto_rawDescData
}

var file_style_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_style_transfer_proto_goTypes = []interface{}{
	(*StyleTransferRequest)(nil),    // 0: StyleTransferRequest
	(*StyleTransferURLRequest)(nil), // 1: StyleTransferURLRequest
	(*StyleTransferResponse)(nil),   // 2: StyleTransferResponse
}
var file_style_transfer_proto_depIdxs = []int32{
	0, // 0: StyleTransferService.StyleTransfer:input_type -> StyleTransferRequest
	1, // 1: StyleTransferService.StyleTransferURL:input_type -> StyleTransferURLRequest
	2, // 2: StyleTransferService.StyleTransfer:output_type -> StyleTransferResponse
	2, // 3: StyleTransferService.StyleTransferURL:output_type -> StyleTransferResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_style_transfer_proto_init() }
func file_style_transfer_proto_init() {
	if File_style_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_style_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StyleTransferRequest); i {
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
		file_style_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StyleTransferURLRequest); i {
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
		file_style_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StyleTransferResponse); i {
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
			RawDescriptor: file_style_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_style_transfer_proto_goTypes,
		DependencyIndexes: file_style_transfer_proto_depIdxs,
		MessageInfos:      file_style_transfer_proto_msgTypes,
	}.Build()
	File_style_transfer_proto = out.File
	file_style_transfer_proto_rawDesc = nil
	file_style_transfer_proto_goTypes = nil
	file_style_transfer_proto_depIdxs = nil
}
