// EDIT IT, change to your package, service and message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: extension/aop/watch/api/ioc_golang/aop/watch/watch.proto

package watch

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdid     string     `protobuf:"bytes,1,opt,name=sdid,proto3" json:"sdid,omitempty"`
	Method   string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	MaxDepth int64      `protobuf:"varint,3,opt,name=maxDepth,proto3" json:"maxDepth,omitempty"`
	Matchers []*Matcher `protobuf:"bytes,4,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRequest) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *WatchRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *WatchRequest) GetMaxDepth() int64 {
	if x != nil {
		return x.MaxDepth
	}
	return 0
}

func (x *WatchRequest) GetMatchers() []*Matcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	MatchPath  string `protobuf:"bytes,2,opt,name=matchPath,proto3" json:"matchPath,omitempty"`
	MatchValue string `protobuf:"bytes,3,opt,name=matchValue,proto3" json:"matchValue,omitempty"`
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescGZIP(), []int{1}
}

func (x *Matcher) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Matcher) GetMatchPath() string {
	if x != nil {
		return x.MatchPath
	}
	return ""
}

func (x *Matcher) GetMatchValue() string {
	if x != nil {
		return x.MatchValue
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdid         string   `protobuf:"bytes,1,opt,name=sdid,proto3" json:"sdid,omitempty"`
	MethodName   string   `protobuf:"bytes,2,opt,name=methodName,proto3" json:"methodName,omitempty"`
	Params       []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	ReturnValues []string `protobuf:"bytes,4,rep,name=returnValues,proto3" json:"returnValues,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescGZIP(), []int{2}
}

func (x *WatchResponse) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *WatchResponse) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *WatchResponse) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *WatchResponse) GetReturnValues() []string {
	if x != nil {
		return x.ReturnValues
	}
	return nil
}

var File_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto protoreflect.FileDescriptor

var file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x6f, 0x70, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x61, 0x6f, 0x70, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6f, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x6f, 0x70, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x22, 0x91, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6f,
	0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x6f, 0x70, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x64, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x32, 0x64, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x22, 0x2e,
	0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61, 0x6f, 0x70, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x61,
	0x6f, 0x70, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x69, 0x6f,
	0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x61, 0x6f, 0x70, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescOnce sync.Once
	file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescData = file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDesc
)

func file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescGZIP() []byte {
	file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescOnce.Do(func() {
		file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescData)
	})
	return file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDescData
}

var file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_goTypes = []interface{}{
	(*WatchRequest)(nil),  // 0: ioc_golang.aop.watch.WatchRequest
	(*Matcher)(nil),       // 1: ioc_golang.aop.watch.Matcher
	(*WatchResponse)(nil), // 2: ioc_golang.aop.watch.WatchResponse
}
var file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_depIdxs = []int32{
	1, // 0: ioc_golang.aop.watch.WatchRequest.matchers:type_name -> ioc_golang.aop.watch.Matcher
	0, // 1: ioc_golang.aop.watch.WatchService.Watch:input_type -> ioc_golang.aop.watch.WatchRequest
	2, // 2: ioc_golang.aop.watch.WatchService.Watch:output_type -> ioc_golang.aop.watch.WatchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_init() }
func file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_init() {
	if File_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher); i {
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
		file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
			RawDescriptor: file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_goTypes,
		DependencyIndexes: file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_depIdxs,
		MessageInfos:      file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_msgTypes,
	}.Build()
	File_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto = out.File
	file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_rawDesc = nil
	file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_goTypes = nil
	file_extension_aop_watch_api_ioc_golang_aop_watch_watch_proto_depIdxs = nil
}
