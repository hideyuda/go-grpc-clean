// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: media.proto

// Specify package name to avoid name collision

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Media_MediaType int32

const (
	Media_YOUTUBE   Media_MediaType = 0
	Media_TWITTER   Media_MediaType = 1
	Media_INSTAGRAM Media_MediaType = 2
	Media_FACEBOOK  Media_MediaType = 3
	Media_LINKEDIN  Media_MediaType = 4
	Media_TIKTOK    Media_MediaType = 5
	Media_PINTEREST Media_MediaType = 6
	Media_REDDIT    Media_MediaType = 7
	Media_SNAPCHAT  Media_MediaType = 8
	Media_TUMBLR    Media_MediaType = 9
	Media_TWITCH    Media_MediaType = 10
	Media_WEIBO     Media_MediaType = 11
	Media_WECHAT    Media_MediaType = 12
	Media_WHATSAPP  Media_MediaType = 13
	Media_LINE      Media_MediaType = 14
	Media_TELEGRAM  Media_MediaType = 15
	Media_VK        Media_MediaType = 16
	Media_YAHOO     Media_MediaType = 17
	Media_OTHER     Media_MediaType = 18
)

// Enum value maps for Media_MediaType.
var (
	Media_MediaType_name = map[int32]string{
		0:  "YOUTUBE",
		1:  "TWITTER",
		2:  "INSTAGRAM",
		3:  "FACEBOOK",
		4:  "LINKEDIN",
		5:  "TIKTOK",
		6:  "PINTEREST",
		7:  "REDDIT",
		8:  "SNAPCHAT",
		9:  "TUMBLR",
		10: "TWITCH",
		11: "WEIBO",
		12: "WECHAT",
		13: "WHATSAPP",
		14: "LINE",
		15: "TELEGRAM",
		16: "VK",
		17: "YAHOO",
		18: "OTHER",
	}
	Media_MediaType_value = map[string]int32{
		"YOUTUBE":   0,
		"TWITTER":   1,
		"INSTAGRAM": 2,
		"FACEBOOK":  3,
		"LINKEDIN":  4,
		"TIKTOK":    5,
		"PINTEREST": 6,
		"REDDIT":    7,
		"SNAPCHAT":  8,
		"TUMBLR":    9,
		"TWITCH":    10,
		"WEIBO":     11,
		"WECHAT":    12,
		"WHATSAPP":  13,
		"LINE":      14,
		"TELEGRAM":  15,
		"VK":        16,
		"YAHOO":     17,
		"OTHER":     18,
	}
)

func (x Media_MediaType) Enum() *Media_MediaType {
	p := new(Media_MediaType)
	*p = x
	return p
}

func (x Media_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Media_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_media_proto_enumTypes[0].Descriptor()
}

func (Media_MediaType) Type() protoreflect.EnumType {
	return &file_media_proto_enumTypes[0]
}

func (x Media_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Media_MediaType.Descriptor instead.
func (Media_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0, 0}
}

type MediaTypeRequest_MediaType int32

const (
	MediaTypeRequest_YOUTUBE   MediaTypeRequest_MediaType = 0
	MediaTypeRequest_TWITTER   MediaTypeRequest_MediaType = 1
	MediaTypeRequest_INSTAGRAM MediaTypeRequest_MediaType = 2
	MediaTypeRequest_FACEBOOK  MediaTypeRequest_MediaType = 3
	MediaTypeRequest_LINKEDIN  MediaTypeRequest_MediaType = 4
	MediaTypeRequest_TIKTOK    MediaTypeRequest_MediaType = 5
	MediaTypeRequest_PINTEREST MediaTypeRequest_MediaType = 6
	MediaTypeRequest_REDDIT    MediaTypeRequest_MediaType = 7
	MediaTypeRequest_SNAPCHAT  MediaTypeRequest_MediaType = 8
	MediaTypeRequest_TUMBLR    MediaTypeRequest_MediaType = 9
	MediaTypeRequest_TWITCH    MediaTypeRequest_MediaType = 10
	MediaTypeRequest_WEIBO     MediaTypeRequest_MediaType = 11
	MediaTypeRequest_WECHAT    MediaTypeRequest_MediaType = 12
	MediaTypeRequest_WHATSAPP  MediaTypeRequest_MediaType = 13
	MediaTypeRequest_LINE      MediaTypeRequest_MediaType = 14
	MediaTypeRequest_TELEGRAM  MediaTypeRequest_MediaType = 15
	MediaTypeRequest_VK        MediaTypeRequest_MediaType = 16
	MediaTypeRequest_YAHOO     MediaTypeRequest_MediaType = 17
	MediaTypeRequest_OTHER     MediaTypeRequest_MediaType = 18
)

// Enum value maps for MediaTypeRequest_MediaType.
var (
	MediaTypeRequest_MediaType_name = map[int32]string{
		0:  "YOUTUBE",
		1:  "TWITTER",
		2:  "INSTAGRAM",
		3:  "FACEBOOK",
		4:  "LINKEDIN",
		5:  "TIKTOK",
		6:  "PINTEREST",
		7:  "REDDIT",
		8:  "SNAPCHAT",
		9:  "TUMBLR",
		10: "TWITCH",
		11: "WEIBO",
		12: "WECHAT",
		13: "WHATSAPP",
		14: "LINE",
		15: "TELEGRAM",
		16: "VK",
		17: "YAHOO",
		18: "OTHER",
	}
	MediaTypeRequest_MediaType_value = map[string]int32{
		"YOUTUBE":   0,
		"TWITTER":   1,
		"INSTAGRAM": 2,
		"FACEBOOK":  3,
		"LINKEDIN":  4,
		"TIKTOK":    5,
		"PINTEREST": 6,
		"REDDIT":    7,
		"SNAPCHAT":  8,
		"TUMBLR":    9,
		"TWITCH":    10,
		"WEIBO":     11,
		"WECHAT":    12,
		"WHATSAPP":  13,
		"LINE":      14,
		"TELEGRAM":  15,
		"VK":        16,
		"YAHOO":     17,
		"OTHER":     18,
	}
)

func (x MediaTypeRequest_MediaType) Enum() *MediaTypeRequest_MediaType {
	p := new(MediaTypeRequest_MediaType)
	*p = x
	return p
}

func (x MediaTypeRequest_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaTypeRequest_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_media_proto_enumTypes[1].Descriptor()
}

func (MediaTypeRequest_MediaType) Type() protoreflect.EnumType {
	return &file_media_proto_enumTypes[1]
}

func (x MediaTypeRequest_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaTypeRequest_MediaType.Descriptor instead.
func (MediaTypeRequest_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{5, 0}
}

// ユーザー情報を表すメッセージ型
type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid        string          `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserId      int64           `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	MediaType   Media_MediaType `protobuf:"varint,6,opt,name=media_type,json=mediaType,proto3,enum=media.Media_MediaType" json:"media_type,omitempty"`
	CreatedAt   string          `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string          `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDeleted   bool            `protobuf:"varint,9,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	ScriptList  []*Script       `protobuf:"bytes,10,rep,name=script_list,json=scriptList,proto3" json:"script_list,omitempty"`
	KeywordList []*Keyword      `protobuf:"bytes,12,rep,name=keyword_list,json=keywordList,proto3" json:"keyword_list,omitempty"` // repeated media.Media media_list = 11;
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Media) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Media) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Media) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Media) GetMediaType() Media_MediaType {
	if x != nil {
		return x.MediaType
	}
	return Media_YOUTUBE
}

func (x *Media) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Media) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Media) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *Media) GetScriptList() []*Script {
	if x != nil {
		return x.ScriptList
	}
	return nil
}

func (x *Media) GetKeywordList() []*Keyword {
	if x != nil {
		return x.KeywordList
	}
	return nil
}

type MediaIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MediaIdRequest) Reset() {
	*x = MediaIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaIdRequest) ProtoMessage() {}

func (x *MediaIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaIdRequest.ProtoReflect.Descriptor instead.
func (*MediaIdRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{1}
}

func (x *MediaIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MediaIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdList []int64 `protobuf:"varint,1,rep,packed,name=id_list,json=idList,proto3" json:"id_list,omitempty"`
}

func (x *MediaIdListRequest) Reset() {
	*x = MediaIdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaIdListRequest) ProtoMessage() {}

func (x *MediaIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaIdListRequest.ProtoReflect.Descriptor instead.
func (*MediaIdListRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{2}
}

func (x *MediaIdListRequest) GetIdList() []int64 {
	if x != nil {
		return x.IdList
	}
	return nil
}

type MediaBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MediaBoolResponse) Reset() {
	*x = MediaBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaBoolResponse) ProtoMessage() {}

func (x *MediaBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaBoolResponse.ProtoReflect.Descriptor instead.
func (*MediaBoolResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{3}
}

func (x *MediaBoolResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type MediaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaList []*Media `protobuf:"bytes,1,rep,name=media_list,json=mediaList,proto3" json:"media_list,omitempty"`
}

func (x *MediaList) Reset() {
	*x = MediaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaList) ProtoMessage() {}

func (x *MediaList) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaList.ProtoReflect.Descriptor instead.
func (*MediaList) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{4}
}

func (x *MediaList) GetMediaList() []*Media {
	if x != nil {
		return x.MediaList
	}
	return nil
}

// ユーザーの情報リクエストに使用するメッセージ型
type MediaTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType MediaTypeRequest_MediaType `protobuf:"varint,1,opt,name=media_type,json=mediaType,proto3,enum=media.MediaTypeRequest_MediaType" json:"media_type,omitempty"`
}

func (x *MediaTypeRequest) Reset() {
	*x = MediaTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaTypeRequest) ProtoMessage() {}

func (x *MediaTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaTypeRequest.ProtoReflect.Descriptor instead.
func (*MediaTypeRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{5}
}

func (x *MediaTypeRequest) GetMediaType() MediaTypeRequest_MediaType {
	if x != nil {
		return x.MediaType
	}
	return MediaTypeRequest_YOUTUBE
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd,
	0x04, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x2f, 0x0a,
	0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0xf8, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x4e, 0x53, 0x54, 0x41, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41,
	0x43, 0x45, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x4b,
	0x45, 0x44, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x49, 0x4b, 0x54, 0x4f, 0x4b,
	0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10,
	0x06, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x44, 0x44, 0x49, 0x54, 0x10, 0x07, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x4e, 0x41, 0x50, 0x43, 0x48, 0x41, 0x54, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x55, 0x4d, 0x42, 0x4c, 0x52, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x57, 0x49, 0x54, 0x43,
	0x48, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x45, 0x49, 0x42, 0x4f, 0x10, 0x0b, 0x12, 0x0a,
	0x0a, 0x06, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x48,
	0x41, 0x54, 0x53, 0x41, 0x50, 0x50, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x4c, 0x45, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x0f,
	0x12, 0x06, 0x0a, 0x02, 0x56, 0x4b, 0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x59, 0x41, 0x48, 0x4f,
	0x4f, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x12, 0x22, 0x20,
	0x0a, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2d, 0x0a, 0x12, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x29, 0x0a, 0x11, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x09, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0xcf, 0x02, 0x0a, 0x10, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x09,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x59, 0x4f, 0x55,
	0x54, 0x55, 0x42, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x47, 0x52, 0x41, 0x4d,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x43, 0x45, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x49, 0x4b, 0x54, 0x4f, 0x4b, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x44,
	0x44, 0x49, 0x54, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4e, 0x41, 0x50, 0x43, 0x48, 0x41,
	0x54, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x55, 0x4d, 0x42, 0x4c, 0x52, 0x10, 0x09, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x57, 0x49, 0x54, 0x43, 0x48, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x57,
	0x45, 0x49, 0x42, 0x4f, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54,
	0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x48, 0x41, 0x54, 0x53, 0x41, 0x50, 0x50, 0x10, 0x0d,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45,
	0x4c, 0x45, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x0f, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x4b, 0x10, 0x10,
	0x12, 0x09, 0x0a, 0x05, 0x59, 0x41, 0x48, 0x4f, 0x4f, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x4f,
	0x54, 0x48, 0x45, 0x52, 0x10, 0x12, 0x32, 0xd4, 0x03, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x1a,
	0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_media_proto_rawDescOnce sync.Once
	file_media_proto_rawDescData = file_media_proto_rawDesc
)

func file_media_proto_rawDescGZIP() []byte {
	file_media_proto_rawDescOnce.Do(func() {
		file_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_proto_rawDescData)
	})
	return file_media_proto_rawDescData
}

var file_media_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_media_proto_goTypes = []interface{}{
	(Media_MediaType)(0),            // 0: media.Media.MediaType
	(MediaTypeRequest_MediaType)(0), // 1: media.MediaTypeRequest.MediaType
	(*Media)(nil),                   // 2: media.Media
	(*MediaIdRequest)(nil),          // 3: media.MediaIdRequest
	(*MediaIdListRequest)(nil),      // 4: media.MediaIdListRequest
	(*MediaBoolResponse)(nil),       // 5: media.MediaBoolResponse
	(*MediaList)(nil),               // 6: media.MediaList
	(*MediaTypeRequest)(nil),        // 7: media.MediaTypeRequest
	(*Script)(nil),                  // 8: script.Script
	(*Keyword)(nil),                 // 9: keyword.Keyword
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_media_proto_depIdxs = []int32{
	0,  // 0: media.Media.media_type:type_name -> media.Media.MediaType
	8,  // 1: media.Media.script_list:type_name -> script.Script
	9,  // 2: media.Media.keyword_list:type_name -> keyword.Keyword
	2,  // 3: media.MediaList.media_list:type_name -> media.Media
	1,  // 4: media.MediaTypeRequest.media_type:type_name -> media.MediaTypeRequest.MediaType
	2,  // 5: media.MediaService.Create:input_type -> media.Media
	2,  // 6: media.MediaService.Update:input_type -> media.Media
	3,  // 7: media.MediaService.Delete:input_type -> media.MediaIdRequest
	3,  // 8: media.MediaService.GetById:input_type -> media.MediaIdRequest
	7,  // 9: media.MediaService.GetListByType:input_type -> media.MediaTypeRequest
	3,  // 10: media.MediaService.GetListByUserId:input_type -> media.MediaIdRequest
	4,  // 11: media.MediaService.GetMediaListByIdList:input_type -> media.MediaIdListRequest
	10, // 12: media.MediaService.GetAll:input_type -> google.protobuf.Empty
	2,  // 13: media.MediaService.Create:output_type -> media.Media
	5,  // 14: media.MediaService.Update:output_type -> media.MediaBoolResponse
	5,  // 15: media.MediaService.Delete:output_type -> media.MediaBoolResponse
	2,  // 16: media.MediaService.GetById:output_type -> media.Media
	6,  // 17: media.MediaService.GetListByType:output_type -> media.MediaList
	6,  // 18: media.MediaService.GetListByUserId:output_type -> media.MediaList
	6,  // 19: media.MediaService.GetMediaListByIdList:output_type -> media.MediaList
	6,  // 20: media.MediaService.GetAll:output_type -> media.MediaList
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_media_proto_init() }
func file_media_proto_init() {
	if File_media_proto != nil {
		return
	}
	file_script_proto_init()
	file_keyword_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaIdRequest); i {
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
		file_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaIdListRequest); i {
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
		file_media_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaBoolResponse); i {
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
		file_media_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaList); i {
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
		file_media_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaTypeRequest); i {
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
			RawDescriptor: file_media_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_media_proto_goTypes,
		DependencyIndexes: file_media_proto_depIdxs,
		EnumInfos:         file_media_proto_enumTypes,
		MessageInfos:      file_media_proto_msgTypes,
	}.Build()
	File_media_proto = out.File
	file_media_proto_rawDesc = nil
	file_media_proto_goTypes = nil
	file_media_proto_depIdxs = nil
}