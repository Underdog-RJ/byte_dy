// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: LikeService.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"interaction/service/pb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	VideoId int64 `protobuf:"varint,2,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
}

func (x *IsLikeRequest) Reset() {
	*x = IsLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LikeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikeRequest) ProtoMessage() {}

func (x *IsLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LikeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikeRequest.ProtoReflect.Descriptor instead.
func (*IsLikeRequest) Descriptor() ([]byte, []int) {
	return file_LikeService_proto_rawDescGZIP(), []int{0}
}

func (x *IsLikeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsLikeRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type IsLikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	IsLike bool   `protobuf:"varint,2,opt,name=IsLike,proto3" json:"IsLike,omitempty"`
}

func (x *IsLikeResponse) Reset() {
	*x = IsLikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LikeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikeResponse) ProtoMessage() {}

func (x *IsLikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_LikeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikeResponse.ProtoReflect.Descriptor instead.
func (*IsLikeResponse) Descriptor() ([]byte, []int) {
	return file_LikeService_proto_rawDescGZIP(), []int{1}
}

func (x *IsLikeResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *IsLikeResponse) GetIsLike() bool {
	if x != nil {
		return x.IsLike
	}
	return false
}

type LikeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @inject_tag: json:"video_id"
	VideoId int64 `protobuf:"varint,2,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	// @inject_tag: json:"action_tpye"
	ActionType uint32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (x *LikeActionRequest) Reset() {
	*x = LikeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LikeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeActionRequest) ProtoMessage() {}

func (x *LikeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LikeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeActionRequest.ProtoReflect.Descriptor instead.
func (*LikeActionRequest) Descriptor() ([]byte, []int) {
	return file_LikeService_proto_rawDescGZIP(), []int{2}
}

func (x *LikeActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *LikeActionRequest) GetActionType() uint32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type LikeActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *LikeActionResponse) Reset() {
	*x = LikeActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LikeService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeActionResponse) ProtoMessage() {}

func (x *LikeActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_LikeService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeActionResponse.ProtoReflect.Descriptor instead.
func (*LikeActionResponse) Descriptor() ([]byte, []int) {
	return file_LikeService_proto_rawDescGZIP(), []int{3}
}

func (x *LikeActionResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type LikeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *LikeListRequest) Reset() {
	*x = LikeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LikeService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeListRequest) ProtoMessage() {}

func (x *LikeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LikeService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeListRequest.ProtoReflect.Descriptor instead.
func (*LikeListRequest) Descriptor() ([]byte, []int) {
	return file_LikeService_proto_rawDescGZIP(), []int{4}
}

func (x *LikeListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LikeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      uint32       `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	VideoList []*VideoInfo `protobuf:"bytes,2,rep,name=VideoList,proto3" json:"VideoList,omitempty"`
}

func (x *LikeListResponse) Reset() {
	*x = LikeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LikeService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeListResponse) ProtoMessage() {}

func (x *LikeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_LikeService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeListResponse.ProtoReflect.Descriptor instead.
func (*LikeListResponse) Descriptor() ([]byte, []int) {
	return file_LikeService_proto_rawDescGZIP(), []int{5}
}

func (x *LikeListResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LikeListResponse) GetVideoList() []*VideoInfo {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_LikeService_proto protoreflect.FileDescriptor

var file_LikeService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x49, 0x73, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0e, 0x49,
	0x73, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x65, 0x0a, 0x11, 0x4c, 0x69, 0x6b,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x28, 0x0a, 0x12, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x4c, 0x69,
	0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x10, 0x4c, 0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a,
	0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xb5, 0x01, 0x0a, 0x0b, 0x4c,
	0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x49, 0x73,
	0x4c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x4c,
	0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x4c,
	0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x6b,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LikeService_proto_rawDescOnce sync.Once
	file_LikeService_proto_rawDescData = file_LikeService_proto_rawDesc
)

func file_LikeService_proto_rawDescGZIP() []byte {
	file_LikeService_proto_rawDescOnce.Do(func() {
		file_LikeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_LikeService_proto_rawDescData)
	})
	return file_LikeService_proto_rawDescData
}

var file_LikeService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_LikeService_proto_goTypes = []interface{}{
	(*IsLikeRequest)(nil),      // 0: pb.IsLikeRequest
	(*IsLikeResponse)(nil),     // 1: pb.IsLikeResponse
	(*LikeActionRequest)(nil),  // 2: pb.LikeActionRequest
	(*LikeActionResponse)(nil), // 3: pb.LikeActionResponse
	(*LikeListRequest)(nil),    // 4: pb.LikeListRequest
	(*LikeListResponse)(nil),   // 5: pb.LikeListResponse
	(*VideoInfo)(nil),          // 6: pb.VideoInfo
}
var file_LikeService_proto_depIdxs = []int32{
	6, // 0: pb.LikeListResponse.VideoList:type_name -> pb.VideoInfo
	0, // 1: pb.LikeService.IsLike:input_type -> pb.IsLikeRequest
	2, // 2: pb.LikeService.LikeAction:input_type -> pb.LikeActionRequest
	4, // 3: pb.LikeService.GetLikeList:input_type -> pb.LikeListRequest
	1, // 4: pb.LikeService.IsLike:output_type -> pb.IsLikeResponse
	3, // 5: pb.LikeService.LikeAction:output_type -> pb.LikeActionResponse
	5, // 6: pb.LikeService.GetLikeList:output_type -> pb.LikeListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LikeService_proto_init() }
func file_LikeService_proto_init() {
	if File_LikeService_proto != nil {
		return
	}
	service.file_videoInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LikeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikeRequest); i {
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
		file_LikeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikeResponse); i {
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
		file_LikeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeActionRequest); i {
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
		file_LikeService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeActionResponse); i {
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
		file_LikeService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeListRequest); i {
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
		file_LikeService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeListResponse); i {
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
			RawDescriptor: file_LikeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_LikeService_proto_goTypes,
		DependencyIndexes: file_LikeService_proto_depIdxs,
		MessageInfos:      file_LikeService_proto_msgTypes,
	}.Build()
	File_LikeService_proto = out.File
	file_LikeService_proto_rawDesc = nil
	file_LikeService_proto_goTypes = nil
	file_LikeService_proto_depIdxs = nil
}