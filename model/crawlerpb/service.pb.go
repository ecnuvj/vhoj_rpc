// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: crawlerpb/service.proto

package crawlerpb

import (
	base "github.com/ecnuvj/vhoj_crawler/pkg/sdk/base"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CrawlProblemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteOj        int32  `protobuf:"varint,1,opt,name=remote_oj,json=remoteOj,proto3" json:"remote_oj,omitempty"`
	RemoteProblemId string `protobuf:"bytes,2,opt,name=remote_problem_id,json=remoteProblemId,proto3" json:"remote_problem_id,omitempty"`
}

func (x *CrawlProblemRequest) Reset() {
	*x = CrawlProblemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlerpb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlProblemRequest) ProtoMessage() {}

func (x *CrawlProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crawlerpb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlProblemRequest.ProtoReflect.Descriptor instead.
func (*CrawlProblemRequest) Descriptor() ([]byte, []int) {
	return file_crawlerpb_service_proto_rawDescGZIP(), []int{0}
}

func (x *CrawlProblemRequest) GetRemoteOj() int32 {
	if x != nil {
		return x.RemoteOj
	}
	return 0
}

func (x *CrawlProblemRequest) GetRemoteProblemId() string {
	if x != nil {
		return x.RemoteProblemId
	}
	return ""
}

type CrawlProblemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *base.BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
}

func (x *CrawlProblemResponse) Reset() {
	*x = CrawlProblemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlerpb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlProblemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlProblemResponse) ProtoMessage() {}

func (x *CrawlProblemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crawlerpb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlProblemResponse.ProtoReflect.Descriptor instead.
func (*CrawlProblemResponse) Descriptor() ([]byte, []int) {
	return file_crawlerpb_service_proto_rawDescGZIP(), []int{1}
}

func (x *CrawlProblemResponse) GetBaseResponse() *base.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

var File_crawlerpb_service_proto protoreflect.FileDescriptor

var file_crawlerpb_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x13, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6f, 0x6a, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4f, 0x6a, 0x12, 0x2a,
	0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x14, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4f, 0x0a, 0x0e, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x0c, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x14, 0x2e,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x6e, 0x75, 0x76,
	0x6a, 0x2f, 0x76, 0x68, 0x6f, 0x6a, 0x5f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crawlerpb_service_proto_rawDescOnce sync.Once
	file_crawlerpb_service_proto_rawDescData = file_crawlerpb_service_proto_rawDesc
)

func file_crawlerpb_service_proto_rawDescGZIP() []byte {
	file_crawlerpb_service_proto_rawDescOnce.Do(func() {
		file_crawlerpb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_crawlerpb_service_proto_rawDescData)
	})
	return file_crawlerpb_service_proto_rawDescData
}

var file_crawlerpb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crawlerpb_service_proto_goTypes = []interface{}{
	(*CrawlProblemRequest)(nil),  // 0: CrawlProblemRequest
	(*CrawlProblemResponse)(nil), // 1: CrawlProblemResponse
	(*base.BaseResponse)(nil),    // 2: base.BaseResponse
}
var file_crawlerpb_service_proto_depIdxs = []int32{
	2, // 0: CrawlProblemResponse.base_response:type_name -> base.BaseResponse
	0, // 1: CrawlerService.CrawlProblem:input_type -> CrawlProblemRequest
	1, // 2: CrawlerService.CrawlProblem:output_type -> CrawlProblemResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_crawlerpb_service_proto_init() }
func file_crawlerpb_service_proto_init() {
	if File_crawlerpb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crawlerpb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlProblemRequest); i {
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
		file_crawlerpb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlProblemResponse); i {
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
			RawDescriptor: file_crawlerpb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crawlerpb_service_proto_goTypes,
		DependencyIndexes: file_crawlerpb_service_proto_depIdxs,
		MessageInfos:      file_crawlerpb_service_proto_msgTypes,
	}.Build()
	File_crawlerpb_service_proto = out.File
	file_crawlerpb_service_proto_rawDesc = nil
	file_crawlerpb_service_proto_goTypes = nil
	file_crawlerpb_service_proto_depIdxs = nil
}
