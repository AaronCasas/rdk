// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/v1/imu.proto

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

// AngularVelocity contains angular velocity in deg/s across x/y/z axes.
type AngularVelocity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Velocity in deg/s across the x-axis
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	// Velocity in deg/s across the y-axis
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	// Velocity in deg/s across the z-axis
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *AngularVelocity) Reset() {
	*x = AngularVelocity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AngularVelocity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AngularVelocity) ProtoMessage() {}

func (x *AngularVelocity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AngularVelocity.ProtoReflect.Descriptor instead.
func (*AngularVelocity) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{0}
}

func (x *AngularVelocity) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *AngularVelocity) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *AngularVelocity) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

// EulerAngles are three angles used to represent the rotation of an object in 3D Euclidean space
// The Tait–Bryan angle formalism is used, with rotations around three distinct axes in the z-y′-x″ sequence.
type EulerAngles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rotation in deg around the x-axis
	Roll float64 `protobuf:"fixed64,1,opt,name=roll,proto3" json:"roll,omitempty"`
	// Rotation in deg around the y-axis
	Pitch float64 `protobuf:"fixed64,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// Rotation in deg around the z-axis
	Yaw float64 `protobuf:"fixed64,3,opt,name=yaw,proto3" json:"yaw,omitempty"`
}

func (x *EulerAngles) Reset() {
	*x = EulerAngles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EulerAngles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EulerAngles) ProtoMessage() {}

func (x *EulerAngles) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EulerAngles.ProtoReflect.Descriptor instead.
func (*EulerAngles) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{1}
}

func (x *EulerAngles) GetRoll() float64 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *EulerAngles) GetPitch() float64 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *EulerAngles) GetYaw() float64 {
	if x != nil {
		return x.Yaw
	}
	return 0
}

type IMUServiceAngularVelocityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an IMU
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IMUServiceAngularVelocityRequest) Reset() {
	*x = IMUServiceAngularVelocityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceAngularVelocityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceAngularVelocityRequest) ProtoMessage() {}

func (x *IMUServiceAngularVelocityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceAngularVelocityRequest.ProtoReflect.Descriptor instead.
func (*IMUServiceAngularVelocityRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{2}
}

func (x *IMUServiceAngularVelocityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IMUServiceAngularVelocityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AngularVelocity contains angular velocity in deg/s across x/y/z axes.
	AngularVelocity *AngularVelocity `protobuf:"bytes,1,opt,name=angular_velocity,json=angularVelocity,proto3" json:"angular_velocity,omitempty"`
}

func (x *IMUServiceAngularVelocityResponse) Reset() {
	*x = IMUServiceAngularVelocityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceAngularVelocityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceAngularVelocityResponse) ProtoMessage() {}

func (x *IMUServiceAngularVelocityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceAngularVelocityResponse.ProtoReflect.Descriptor instead.
func (*IMUServiceAngularVelocityResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{3}
}

func (x *IMUServiceAngularVelocityResponse) GetAngularVelocity() *AngularVelocity {
	if x != nil {
		return x.AngularVelocity
	}
	return nil
}

type IMUServiceOrientationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an IMU
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IMUServiceOrientationRequest) Reset() {
	*x = IMUServiceOrientationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceOrientationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceOrientationRequest) ProtoMessage() {}

func (x *IMUServiceOrientationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceOrientationRequest.ProtoReflect.Descriptor instead.
func (*IMUServiceOrientationRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{4}
}

func (x *IMUServiceOrientationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IMUServiceOrientationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EulerAngles are three angles used to represent the rotation of an object in 3D Euclidean space
	// The Tait–Bryan angle formalism is used, with rotations around three distinct axes in the z-y′-x″ sequence.
	Orientation *EulerAngles `protobuf:"bytes,1,opt,name=orientation,proto3" json:"orientation,omitempty"`
}

func (x *IMUServiceOrientationResponse) Reset() {
	*x = IMUServiceOrientationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceOrientationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceOrientationResponse) ProtoMessage() {}

func (x *IMUServiceOrientationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceOrientationResponse.ProtoReflect.Descriptor instead.
func (*IMUServiceOrientationResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{5}
}

func (x *IMUServiceOrientationResponse) GetOrientation() *EulerAngles {
	if x != nil {
		return x.Orientation
	}
	return nil
}

var File_proto_api_component_v1_imu_proto protoreflect.FileDescriptor

var file_proto_api_component_v1_imu_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x41, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x7a, 0x22, 0x49, 0x0a, 0x0b, 0x45, 0x75, 0x6c, 0x65, 0x72, 0x41, 0x6e,
	0x67, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x79, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x79, 0x61, 0x77,
	0x22, 0x36, 0x0a, 0x20, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6e,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x21, 0x49, 0x4d, 0x55, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c,
	0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x10, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x0f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74,
	0x79, 0x22, 0x32, 0x0a, 0x1c, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x1d, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x75, 0x6c, 0x65, 0x72, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73,
	0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xfb, 0x02,
	0x0a, 0x0a, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbd, 0x01, 0x0a,
	0x0f, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x38, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6e,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x69, 0x6d, 0x75, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0xac, 0x01, 0x0a,
	0x0b, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x12, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x75, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4d, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_api_component_v1_imu_proto_rawDescOnce sync.Once
	file_proto_api_component_v1_imu_proto_rawDescData = file_proto_api_component_v1_imu_proto_rawDesc
)

func file_proto_api_component_v1_imu_proto_rawDescGZIP() []byte {
	file_proto_api_component_v1_imu_proto_rawDescOnce.Do(func() {
		file_proto_api_component_v1_imu_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_v1_imu_proto_rawDescData)
	})
	return file_proto_api_component_v1_imu_proto_rawDescData
}

var file_proto_api_component_v1_imu_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_api_component_v1_imu_proto_goTypes = []interface{}{
	(*AngularVelocity)(nil),                   // 0: proto.api.component.v1.AngularVelocity
	(*EulerAngles)(nil),                       // 1: proto.api.component.v1.EulerAngles
	(*IMUServiceAngularVelocityRequest)(nil),  // 2: proto.api.component.v1.IMUServiceAngularVelocityRequest
	(*IMUServiceAngularVelocityResponse)(nil), // 3: proto.api.component.v1.IMUServiceAngularVelocityResponse
	(*IMUServiceOrientationRequest)(nil),      // 4: proto.api.component.v1.IMUServiceOrientationRequest
	(*IMUServiceOrientationResponse)(nil),     // 5: proto.api.component.v1.IMUServiceOrientationResponse
}
var file_proto_api_component_v1_imu_proto_depIdxs = []int32{
	0, // 0: proto.api.component.v1.IMUServiceAngularVelocityResponse.angular_velocity:type_name -> proto.api.component.v1.AngularVelocity
	1, // 1: proto.api.component.v1.IMUServiceOrientationResponse.orientation:type_name -> proto.api.component.v1.EulerAngles
	2, // 2: proto.api.component.v1.IMUService.AngularVelocity:input_type -> proto.api.component.v1.IMUServiceAngularVelocityRequest
	4, // 3: proto.api.component.v1.IMUService.Orientation:input_type -> proto.api.component.v1.IMUServiceOrientationRequest
	3, // 4: proto.api.component.v1.IMUService.AngularVelocity:output_type -> proto.api.component.v1.IMUServiceAngularVelocityResponse
	5, // 5: proto.api.component.v1.IMUService.Orientation:output_type -> proto.api.component.v1.IMUServiceOrientationResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_api_component_v1_imu_proto_init() }
func file_proto_api_component_v1_imu_proto_init() {
	if File_proto_api_component_v1_imu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_v1_imu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AngularVelocity); i {
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
		file_proto_api_component_v1_imu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EulerAngles); i {
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
		file_proto_api_component_v1_imu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceAngularVelocityRequest); i {
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
		file_proto_api_component_v1_imu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceAngularVelocityResponse); i {
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
		file_proto_api_component_v1_imu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceOrientationRequest); i {
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
		file_proto_api_component_v1_imu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceOrientationResponse); i {
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
			RawDescriptor: file_proto_api_component_v1_imu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_v1_imu_proto_goTypes,
		DependencyIndexes: file_proto_api_component_v1_imu_proto_depIdxs,
		MessageInfos:      file_proto_api_component_v1_imu_proto_msgTypes,
	}.Build()
	File_proto_api_component_v1_imu_proto = out.File
	file_proto_api_component_v1_imu_proto_rawDesc = nil
	file_proto_api_component_v1_imu_proto_goTypes = nil
	file_proto_api_component_v1_imu_proto_depIdxs = nil
}
