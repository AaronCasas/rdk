// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/v1/gantry.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GantryServiceCurrentPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GantryServiceCurrentPositionRequest) Reset() {
	*x = GantryServiceCurrentPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gantry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GantryServiceCurrentPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GantryServiceCurrentPositionRequest) ProtoMessage() {}

func (x *GantryServiceCurrentPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gantry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GantryServiceCurrentPositionRequest.ProtoReflect.Descriptor instead.
func (*GantryServiceCurrentPositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gantry_proto_rawDescGZIP(), []int{0}
}

func (x *GantryServiceCurrentPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GantryServiceCurrentPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []float64 `protobuf:"fixed64,1,rep,packed,name=positions,proto3" json:"positions,omitempty"`
}

func (x *GantryServiceCurrentPositionResponse) Reset() {
	*x = GantryServiceCurrentPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gantry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GantryServiceCurrentPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GantryServiceCurrentPositionResponse) ProtoMessage() {}

func (x *GantryServiceCurrentPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gantry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GantryServiceCurrentPositionResponse.ProtoReflect.Descriptor instead.
func (*GantryServiceCurrentPositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gantry_proto_rawDescGZIP(), []int{1}
}

func (x *GantryServiceCurrentPositionResponse) GetPositions() []float64 {
	if x != nil {
		return x.Positions
	}
	return nil
}

type GantryServiceMoveToPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Positions []float64 `protobuf:"fixed64,2,rep,packed,name=positions,proto3" json:"positions,omitempty"`
}

func (x *GantryServiceMoveToPositionRequest) Reset() {
	*x = GantryServiceMoveToPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gantry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GantryServiceMoveToPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GantryServiceMoveToPositionRequest) ProtoMessage() {}

func (x *GantryServiceMoveToPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gantry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GantryServiceMoveToPositionRequest.ProtoReflect.Descriptor instead.
func (*GantryServiceMoveToPositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gantry_proto_rawDescGZIP(), []int{2}
}

func (x *GantryServiceMoveToPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GantryServiceMoveToPositionRequest) GetPositions() []float64 {
	if x != nil {
		return x.Positions
	}
	return nil
}

type GantryServiceMoveToPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GantryServiceMoveToPositionResponse) Reset() {
	*x = GantryServiceMoveToPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gantry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GantryServiceMoveToPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GantryServiceMoveToPositionResponse) ProtoMessage() {}

func (x *GantryServiceMoveToPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gantry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GantryServiceMoveToPositionResponse.ProtoReflect.Descriptor instead.
func (*GantryServiceMoveToPositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gantry_proto_rawDescGZIP(), []int{3}
}

type GantryServiceLengthsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GantryServiceLengthsRequest) Reset() {
	*x = GantryServiceLengthsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gantry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GantryServiceLengthsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GantryServiceLengthsRequest) ProtoMessage() {}

func (x *GantryServiceLengthsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gantry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GantryServiceLengthsRequest.ProtoReflect.Descriptor instead.
func (*GantryServiceLengthsRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gantry_proto_rawDescGZIP(), []int{4}
}

func (x *GantryServiceLengthsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GantryServiceLengthsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lengths []float64 `protobuf:"fixed64,1,rep,packed,name=lengths,proto3" json:"lengths,omitempty"`
}

func (x *GantryServiceLengthsResponse) Reset() {
	*x = GantryServiceLengthsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_gantry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GantryServiceLengthsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GantryServiceLengthsResponse) ProtoMessage() {}

func (x *GantryServiceLengthsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_gantry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GantryServiceLengthsResponse.ProtoReflect.Descriptor instead.
func (*GantryServiceLengthsResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_gantry_proto_rawDescGZIP(), []int{5}
}

func (x *GantryServiceLengthsResponse) GetLengths() []float64 {
	if x != nil {
		return x.Lengths
	}
	return nil
}

var File_proto_api_component_v1_gantry_proto protoreflect.FileDescriptor

var file_proto_api_component_v1_gantry_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x23, 0x47,
	0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x24, 0x47, 0x61, 0x6e, 0x74, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x56, 0x0a, 0x22,
	0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x76,
	0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x23, 0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x1b, 0x47,
	0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38,
	0x0a, 0x1c, 0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x07, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x32, 0xc6, 0x04, 0x0a, 0x0d, 0x47, 0x61, 0x6e,
	0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc6, 0x01, 0x0a, 0x0f, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x32, 0x12, 0x30, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0xc3, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x76,
	0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x1a, 0x30, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74, 0x72,
	0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x74, 0x6f,
	0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa5, 0x01, 0x0a, 0x07, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x73, 0x12, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x6e, 0x74,
	0x72, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x73, 0x42, 0x4d, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_component_v1_gantry_proto_rawDescOnce sync.Once
	file_proto_api_component_v1_gantry_proto_rawDescData = file_proto_api_component_v1_gantry_proto_rawDesc
)

func file_proto_api_component_v1_gantry_proto_rawDescGZIP() []byte {
	file_proto_api_component_v1_gantry_proto_rawDescOnce.Do(func() {
		file_proto_api_component_v1_gantry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_v1_gantry_proto_rawDescData)
	})
	return file_proto_api_component_v1_gantry_proto_rawDescData
}

var file_proto_api_component_v1_gantry_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_api_component_v1_gantry_proto_goTypes = []interface{}{
	(*GantryServiceCurrentPositionRequest)(nil),  // 0: proto.api.component.v1.GantryServiceCurrentPositionRequest
	(*GantryServiceCurrentPositionResponse)(nil), // 1: proto.api.component.v1.GantryServiceCurrentPositionResponse
	(*GantryServiceMoveToPositionRequest)(nil),   // 2: proto.api.component.v1.GantryServiceMoveToPositionRequest
	(*GantryServiceMoveToPositionResponse)(nil),  // 3: proto.api.component.v1.GantryServiceMoveToPositionResponse
	(*GantryServiceLengthsRequest)(nil),          // 4: proto.api.component.v1.GantryServiceLengthsRequest
	(*GantryServiceLengthsResponse)(nil),         // 5: proto.api.component.v1.GantryServiceLengthsResponse
}
var file_proto_api_component_v1_gantry_proto_depIdxs = []int32{
	0, // 0: proto.api.component.v1.GantryService.CurrentPosition:input_type -> proto.api.component.v1.GantryServiceCurrentPositionRequest
	2, // 1: proto.api.component.v1.GantryService.MoveToPosition:input_type -> proto.api.component.v1.GantryServiceMoveToPositionRequest
	4, // 2: proto.api.component.v1.GantryService.Lengths:input_type -> proto.api.component.v1.GantryServiceLengthsRequest
	1, // 3: proto.api.component.v1.GantryService.CurrentPosition:output_type -> proto.api.component.v1.GantryServiceCurrentPositionResponse
	3, // 4: proto.api.component.v1.GantryService.MoveToPosition:output_type -> proto.api.component.v1.GantryServiceMoveToPositionResponse
	5, // 5: proto.api.component.v1.GantryService.Lengths:output_type -> proto.api.component.v1.GantryServiceLengthsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_api_component_v1_gantry_proto_init() }
func file_proto_api_component_v1_gantry_proto_init() {
	if File_proto_api_component_v1_gantry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_v1_gantry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GantryServiceCurrentPositionRequest); i {
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
		file_proto_api_component_v1_gantry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GantryServiceCurrentPositionResponse); i {
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
		file_proto_api_component_v1_gantry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GantryServiceMoveToPositionRequest); i {
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
		file_proto_api_component_v1_gantry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GantryServiceMoveToPositionResponse); i {
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
		file_proto_api_component_v1_gantry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GantryServiceLengthsRequest); i {
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
		file_proto_api_component_v1_gantry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GantryServiceLengthsResponse); i {
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
			RawDescriptor: file_proto_api_component_v1_gantry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_v1_gantry_proto_goTypes,
		DependencyIndexes: file_proto_api_component_v1_gantry_proto_depIdxs,
		MessageInfos:      file_proto_api_component_v1_gantry_proto_msgTypes,
	}.Build()
	File_proto_api_component_v1_gantry_proto = out.File
	file_proto_api_component_v1_gantry_proto_rawDesc = nil
	file_proto_api_component_v1_gantry_proto_goTypes = nil
	file_proto_api_component_v1_gantry_proto_depIdxs = nil
}
