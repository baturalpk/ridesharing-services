//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mapping.proto

package genproto

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

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start       *Location `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Destination *Location `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *Route) GetStart() *Location {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Route) GetDestination() *Location {
	if x != nil {
		return x.Destination
	}
	return nil
}

type CalculateDistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route *Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *CalculateDistanceRequest) Reset() {
	*x = CalculateDistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateDistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateDistanceRequest) ProtoMessage() {}

func (x *CalculateDistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateDistanceRequest.ProtoReflect.Descriptor instead.
func (*CalculateDistanceRequest) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateDistanceRequest) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

type CalculateDistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance float32 `protobuf:"fixed32,1,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *CalculateDistanceResponse) Reset() {
	*x = CalculateDistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateDistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateDistanceResponse) ProtoMessage() {}

func (x *CalculateDistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateDistanceResponse.ProtoReflect.Descriptor instead.
func (*CalculateDistanceResponse) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateDistanceResponse) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type MakeTimeEstimationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route *Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *MakeTimeEstimationRequest) Reset() {
	*x = MakeTimeEstimationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeTimeEstimationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeTimeEstimationRequest) ProtoMessage() {}

func (x *MakeTimeEstimationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeTimeEstimationRequest.ProtoReflect.Descriptor instead.
func (*MakeTimeEstimationRequest) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{3}
}

func (x *MakeTimeEstimationRequest) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

type MakeTimeEstimationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eta float32 `protobuf:"fixed32,1,opt,name=eta,proto3" json:"eta,omitempty"`
}

func (x *MakeTimeEstimationResponse) Reset() {
	*x = MakeTimeEstimationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeTimeEstimationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeTimeEstimationResponse) ProtoMessage() {}

func (x *MakeTimeEstimationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeTimeEstimationResponse.ProtoReflect.Descriptor instead.
func (*MakeTimeEstimationResponse) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{4}
}

func (x *MakeTimeEstimationResponse) GetEta() float32 {
	if x != nil {
		return x.Eta
	}
	return 0
}

type DoGeocodingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DoGeocodingRequest) Reset() {
	*x = DoGeocodingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoGeocodingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoGeocodingRequest) ProtoMessage() {}

func (x *DoGeocodingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoGeocodingRequest.ProtoReflect.Descriptor instead.
func (*DoGeocodingRequest) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{5}
}

func (x *DoGeocodingRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DoGeocodingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=Location,proto3" json:"Location,omitempty"`
}

func (x *DoGeocodingResponse) Reset() {
	*x = DoGeocodingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoGeocodingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoGeocodingResponse) ProtoMessage() {}

func (x *DoGeocodingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoGeocodingResponse.ProtoReflect.Descriptor instead.
func (*DoGeocodingResponse) Descriptor() ([]byte, []int) {
	return file_mapping_proto_rawDescGZIP(), []int{6}
}

func (x *DoGeocodingResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_mapping_proto protoreflect.FileDescriptor

var file_mapping_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x18, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x37, 0x0a,
	0x19, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x19, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x2e, 0x0a, 0x1a, 0x4d, 0x61, 0x6b,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x65, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x6f, 0x47,
	0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x6f, 0x47,
	0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x95,
	0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a,
	0x12, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61,
	0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b,
	0x44, 0x6f, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x6f, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x6f, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x0a, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x0e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mapping_proto_rawDescOnce sync.Once
	file_mapping_proto_rawDescData = file_mapping_proto_rawDesc
)

func file_mapping_proto_rawDescGZIP() []byte {
	file_mapping_proto_rawDescOnce.Do(func() {
		file_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_mapping_proto_rawDescData)
	})
	return file_mapping_proto_rawDescData
}

var file_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mapping_proto_goTypes = []interface{}{
	(*Route)(nil),                      // 0: mapping.Route
	(*CalculateDistanceRequest)(nil),   // 1: mapping.CalculateDistanceRequest
	(*CalculateDistanceResponse)(nil),  // 2: mapping.CalculateDistanceResponse
	(*MakeTimeEstimationRequest)(nil),  // 3: mapping.MakeTimeEstimationRequest
	(*MakeTimeEstimationResponse)(nil), // 4: mapping.MakeTimeEstimationResponse
	(*DoGeocodingRequest)(nil),         // 5: mapping.DoGeocodingRequest
	(*DoGeocodingResponse)(nil),        // 6: mapping.DoGeocodingResponse
	(*Location)(nil),                   // 7: common.Location
}
var file_mapping_proto_depIdxs = []int32{
	7, // 0: mapping.Route.start:type_name -> common.Location
	7, // 1: mapping.Route.destination:type_name -> common.Location
	0, // 2: mapping.CalculateDistanceRequest.route:type_name -> mapping.Route
	0, // 3: mapping.MakeTimeEstimationRequest.route:type_name -> mapping.Route
	7, // 4: mapping.DoGeocodingResponse.Location:type_name -> common.Location
	1, // 5: mapping.MappingService.CalculateDistance:input_type -> mapping.CalculateDistanceRequest
	3, // 6: mapping.MappingService.MakeTimeEstimation:input_type -> mapping.MakeTimeEstimationRequest
	5, // 7: mapping.MappingService.DoGeocoding:input_type -> mapping.DoGeocodingRequest
	2, // 8: mapping.MappingService.CalculateDistance:output_type -> mapping.CalculateDistanceResponse
	4, // 9: mapping.MappingService.MakeTimeEstimation:output_type -> mapping.MakeTimeEstimationResponse
	6, // 10: mapping.MappingService.DoGeocoding:output_type -> mapping.DoGeocodingResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mapping_proto_init() }
func file_mapping_proto_init() {
	if File_mapping_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_mapping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateDistanceRequest); i {
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
		file_mapping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateDistanceResponse); i {
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
		file_mapping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeTimeEstimationRequest); i {
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
		file_mapping_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeTimeEstimationResponse); i {
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
		file_mapping_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoGeocodingRequest); i {
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
		file_mapping_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoGeocodingResponse); i {
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
			RawDescriptor: file_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mapping_proto_goTypes,
		DependencyIndexes: file_mapping_proto_depIdxs,
		MessageInfos:      file_mapping_proto_msgTypes,
	}.Build()
	File_mapping_proto = out.File
	file_mapping_proto_rawDesc = nil
	file_mapping_proto_goTypes = nil
	file_mapping_proto_depIdxs = nil
}
