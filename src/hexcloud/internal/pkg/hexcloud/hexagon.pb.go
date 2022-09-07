// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: api/hexagon.proto

package hexcloud

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{0}
}

type HexLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X          int64             `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y          int64             `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z          int64             `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
	HexID      string            `protobuf:"bytes,5,opt,name=HexID,proto3" json:"HexID,omitempty"`
	LocalData  map[string]string `protobuf:"bytes,6,rep,name=LocalData,proto3" json:"LocalData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	GlobalData map[string]string `protobuf:"bytes,7,rep,name=GlobalData,proto3" json:"GlobalData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HexLocation) Reset() {
	*x = HexLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexLocation) ProtoMessage() {}

func (x *HexLocation) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexLocation.ProtoReflect.Descriptor instead.
func (*HexLocation) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{1}
}

func (x *HexLocation) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *HexLocation) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *HexLocation) GetZ() int64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *HexLocation) GetHexID() string {
	if x != nil {
		return x.HexID
	}
	return ""
}

func (x *HexLocation) GetLocalData() map[string]string {
	if x != nil {
		return x.LocalData
	}
	return nil
}

func (x *HexLocation) GetGlobalData() map[string]string {
	if x != nil {
		return x.GlobalData
	}
	return nil
}

type HexInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Data map[string]string `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HexInfo) Reset() {
	*x = HexInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexInfo) ProtoMessage() {}

func (x *HexInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexInfo.ProtoReflect.Descriptor instead.
func (*HexInfo) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{2}
}

func (x *HexInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *HexInfo) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type HexInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexInfo []*HexInfo `protobuf:"bytes,1,rep,name=hexInfo,proto3" json:"hexInfo,omitempty"`
}

func (x *HexInfoList) Reset() {
	*x = HexInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexInfoList) ProtoMessage() {}

func (x *HexInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexInfoList.ProtoReflect.Descriptor instead.
func (*HexInfoList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{3}
}

func (x *HexInfoList) GetHexInfo() []*HexInfo {
	if x != nil {
		return x.HexInfo
	}
	return nil
}

type HexLocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexLoc []*HexLocation `protobuf:"bytes,1,rep,name=hexLoc,proto3" json:"hexLoc,omitempty"`
}

func (x *HexLocationList) Reset() {
	*x = HexLocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexLocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexLocationList) ProtoMessage() {}

func (x *HexLocationList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexLocationList.ProtoReflect.Descriptor instead.
func (*HexLocationList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{4}
}

func (x *HexLocationList) GetHexLoc() []*HexLocation {
	if x != nil {
		return x.HexLoc
	}
	return nil
}

type HexagonGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexLoc *HexLocation `protobuf:"bytes,1,opt,name=hexLoc,proto3" json:"hexLoc,omitempty"`
	Radius int64        `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
	Fill   bool         `protobuf:"varint,3,opt,name=fill,proto3" json:"fill,omitempty"`
}

func (x *HexagonGetRequest) Reset() {
	*x = HexagonGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexagonGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexagonGetRequest) ProtoMessage() {}

func (x *HexagonGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexagonGetRequest.ProtoReflect.Descriptor instead.
func (*HexagonGetRequest) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{5}
}

func (x *HexagonGetRequest) GetHexLoc() *HexLocation {
	if x != nil {
		return x.HexLoc
	}
	return nil
}

func (x *HexagonGetRequest) GetRadius() int64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *HexagonGetRequest) GetFill() bool {
	if x != nil {
		return x.Fill
	}
	return false
}

type HexIDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexID []string `protobuf:"bytes,1,rep,name=HexID,proto3" json:"HexID,omitempty"`
}

func (x *HexIDList) Reset() {
	*x = HexIDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexIDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexIDList) ProtoMessage() {}

func (x *HexIDList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexIDList.ProtoReflect.Descriptor instead.
func (*HexIDList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{6}
}

func (x *HexIDList) GetHexID() []string {
	if x != nil {
		return x.HexID
	}
	return nil
}

type HexIDData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexID   string `protobuf:"bytes,1,opt,name=HexID,proto3" json:"HexID,omitempty"`
	DataKey string `protobuf:"bytes,2,opt,name=dataKey,proto3" json:"dataKey,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HexIDData) Reset() {
	*x = HexIDData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexIDData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexIDData) ProtoMessage() {}

func (x *HexIDData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexIDData.ProtoReflect.Descriptor instead.
func (*HexIDData) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{7}
}

func (x *HexIDData) GetHexID() string {
	if x != nil {
		return x.HexID
	}
	return ""
}

func (x *HexIDData) GetDataKey() string {
	if x != nil {
		return x.DataKey
	}
	return ""
}

func (x *HexIDData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type HexLocData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X       int64  `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y       int64  `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z       int64  `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
	DataKey string `protobuf:"bytes,4,opt,name=dataKey,proto3" json:"dataKey,omitempty"`
	Value   string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HexLocData) Reset() {
	*x = HexLocData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexLocData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexLocData) ProtoMessage() {}

func (x *HexLocData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexLocData.ProtoReflect.Descriptor instead.
func (*HexLocData) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{8}
}

func (x *HexLocData) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *HexLocData) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *HexLocData) GetZ() int64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *HexLocData) GetDataKey() string {
	if x != nil {
		return x.DataKey
	}
	return ""
}

func (x *HexLocData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type HexLocDataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexLocData []*HexLocData `protobuf:"bytes,1,rep,name=hexLocData,proto3" json:"hexLocData,omitempty"`
}

func (x *HexLocDataList) Reset() {
	*x = HexLocDataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexLocDataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexLocDataList) ProtoMessage() {}

func (x *HexLocDataList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexLocDataList.ProtoReflect.Descriptor instead.
func (*HexLocDataList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{9}
}

func (x *HexLocDataList) GetHexLocData() []*HexLocData {
	if x != nil {
		return x.HexLocData
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{10}
}

func (x *Status) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{11}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_hexagon_proto protoreflect.FileDescriptor

var file_api_hexagon_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xc3, 0x02, 0x0a,
	0x0b, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01,
	0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x5a, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x09,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0a, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x48, 0x65,
	0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3c, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x7a, 0x0a, 0x07, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x48, 0x65,
	0x78, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x31,
	0x0a, 0x0b, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x07, 0x68, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x68, 0x65, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x37, 0x0a, 0x0f, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x22, 0x65, 0x0a, 0x11, 0x48, 0x65,
	0x78, 0x61, 0x67, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x06, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x68,
	0x65, 0x78, 0x4c, 0x6f, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x6c, 0x22, 0x21, 0x0a, 0x09, 0x48, 0x65, 0x78, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x48,
	0x65, 0x78, 0x49, 0x44, 0x22, 0x51, 0x0a, 0x09, 0x48, 0x65, 0x78, 0x49, 0x44, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x66, 0x0a, 0x0a, 0x48, 0x65, 0x78, 0x4c, 0x6f,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x5a, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x3d, 0x0a, 0x0e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1a,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xfa,
	0x05, 0x0a, 0x0e, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x64, 0x64, 0x48, 0x65, 0x78, 0x61,
	0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e,
	0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0a, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0a, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x0a, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x65,
	0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0c, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x6c, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0a, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x16, 0x52, 0x65,
	0x70, 0x6f, 0x44, 0x65, 0x6c, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0a, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x4d, 0x61, 0x70,
	0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26,
	0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x48,
	0x65, 0x78, 0x4c, 0x6f, 0x63, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x07, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x47, 0x65, 0x74,
	0x12, 0x12, 0x2e, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x0d, 0x4d, 0x61,
	0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x48, 0x65,
	0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12,
	0x10, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x0d, 0x4d, 0x61,
	0x70, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x48, 0x65,
	0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x07, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x28, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x6c, 0x41, 0x6c, 0x6c, 0x48, 0x65,
	0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0c, 0x4d, 0x61,
	0x70, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_hexagon_proto_rawDescOnce sync.Once
	file_api_hexagon_proto_rawDescData = file_api_hexagon_proto_rawDesc
)

func file_api_hexagon_proto_rawDescGZIP() []byte {
	file_api_hexagon_proto_rawDescOnce.Do(func() {
		file_api_hexagon_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hexagon_proto_rawDescData)
	})
	return file_api_hexagon_proto_rawDescData
}

var file_api_hexagon_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_api_hexagon_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: Empty
	(*HexLocation)(nil),       // 1: HexLocation
	(*HexInfo)(nil),           // 2: HexInfo
	(*HexInfoList)(nil),       // 3: HexInfoList
	(*HexLocationList)(nil),   // 4: HexLocationList
	(*HexagonGetRequest)(nil), // 5: HexagonGetRequest
	(*HexIDList)(nil),         // 6: HexIDList
	(*HexIDData)(nil),         // 7: HexIDData
	(*HexLocData)(nil),        // 8: HexLocData
	(*HexLocDataList)(nil),    // 9: HexLocDataList
	(*Status)(nil),            // 10: Status
	(*Result)(nil),            // 11: Result
	nil,                       // 12: HexLocation.LocalDataEntry
	nil,                       // 13: HexLocation.GlobalDataEntry
	nil,                       // 14: HexInfo.DataEntry
}
var file_api_hexagon_proto_depIdxs = []int32{
	12, // 0: HexLocation.LocalData:type_name -> HexLocation.LocalDataEntry
	13, // 1: HexLocation.GlobalData:type_name -> HexLocation.GlobalDataEntry
	14, // 2: HexInfo.Data:type_name -> HexInfo.DataEntry
	2,  // 3: HexInfoList.hexInfo:type_name -> HexInfo
	1,  // 4: HexLocationList.hexLoc:type_name -> HexLocation
	1,  // 5: HexagonGetRequest.hexLoc:type_name -> HexLocation
	8,  // 6: HexLocDataList.hexLocData:type_name -> HexLocData
	3,  // 7: HexagonService.RepoAddHexagonInfo:input_type -> HexInfoList
	6,  // 8: HexagonService.RepoGetHexagonInfo:input_type -> HexIDList
	7,  // 9: HexagonService.RepoGetHexagonInfoData:input_type -> HexIDData
	0,  // 10: HexagonService.RepoGetAllHexagonInfo:input_type -> Empty
	6,  // 11: HexagonService.RepoDelHexagonInfo:input_type -> HexIDList
	7,  // 12: HexagonService.RepoDelHexagonInfoData:input_type -> HexIDData
	4,  // 13: HexagonService.MapAdd:input_type -> HexLocationList
	9,  // 14: HexagonService.MapAddData:input_type -> HexLocDataList
	5,  // 15: HexagonService.MapGet:input_type -> HexagonGetRequest
	1,  // 16: HexagonService.MapUpdate:input_type -> HexLocation
	1,  // 17: HexagonService.MapUpdateData:input_type -> HexLocation
	4,  // 18: HexagonService.MapRemove:input_type -> HexLocationList
	1,  // 19: HexagonService.MapRemoveData:input_type -> HexLocation
	0,  // 20: HexagonService.GetStatusServer:input_type -> Empty
	0,  // 21: HexagonService.GetStatusStorage:input_type -> Empty
	0,  // 22: HexagonService.GetStatusClients:input_type -> Empty
	0,  // 23: HexagonService.RepoDelAllHexagonInfo:input_type -> Empty
	0,  // 24: HexagonService.MapRemoveAll:input_type -> Empty
	11, // 25: HexagonService.RepoAddHexagonInfo:output_type -> Result
	3,  // 26: HexagonService.RepoGetHexagonInfo:output_type -> HexInfoList
	7,  // 27: HexagonService.RepoGetHexagonInfoData:output_type -> HexIDData
	3,  // 28: HexagonService.RepoGetAllHexagonInfo:output_type -> HexInfoList
	11, // 29: HexagonService.RepoDelHexagonInfo:output_type -> Result
	11, // 30: HexagonService.RepoDelHexagonInfoData:output_type -> Result
	11, // 31: HexagonService.MapAdd:output_type -> Result
	11, // 32: HexagonService.MapAddData:output_type -> Result
	4,  // 33: HexagonService.MapGet:output_type -> HexLocationList
	11, // 34: HexagonService.MapUpdate:output_type -> Result
	11, // 35: HexagonService.MapUpdateData:output_type -> Result
	11, // 36: HexagonService.MapRemove:output_type -> Result
	11, // 37: HexagonService.MapRemoveData:output_type -> Result
	10, // 38: HexagonService.GetStatusServer:output_type -> Status
	10, // 39: HexagonService.GetStatusStorage:output_type -> Status
	10, // 40: HexagonService.GetStatusClients:output_type -> Status
	11, // 41: HexagonService.RepoDelAllHexagonInfo:output_type -> Result
	11, // 42: HexagonService.MapRemoveAll:output_type -> Result
	25, // [25:43] is the sub-list for method output_type
	7,  // [7:25] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_hexagon_proto_init() }
func file_api_hexagon_proto_init() {
	if File_api_hexagon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hexagon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_hexagon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexLocation); i {
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
		file_api_hexagon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexInfo); i {
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
		file_api_hexagon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexInfoList); i {
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
		file_api_hexagon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexLocationList); i {
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
		file_api_hexagon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexagonGetRequest); i {
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
		file_api_hexagon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexIDList); i {
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
		file_api_hexagon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexIDData); i {
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
		file_api_hexagon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexLocData); i {
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
		file_api_hexagon_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexLocDataList); i {
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
		file_api_hexagon_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_api_hexagon_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_api_hexagon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hexagon_proto_goTypes,
		DependencyIndexes: file_api_hexagon_proto_depIdxs,
		MessageInfos:      file_api_hexagon_proto_msgTypes,
	}.Build()
	File_api_hexagon_proto = out.File
	file_api_hexagon_proto_rawDesc = nil
	file_api_hexagon_proto_goTypes = nil
	file_api_hexagon_proto_depIdxs = nil
}
