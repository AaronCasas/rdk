// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/service/framesystem/v1/frame_system.proto

package v1

import (
	v1 "go.viam.com/rdk/proto/api/common/v1"
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

type FrameConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Pose   *v1.Pose `protobuf:"bytes,2,opt,name=pose,proto3" json:"pose,omitempty"`
}

func (x *FrameConfig) Reset() {
	*x = FrameConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameConfig) ProtoMessage() {}

func (x *FrameConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameConfig.ProtoReflect.Descriptor instead.
func (*FrameConfig) Descriptor() ([]byte, []int) {
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP(), []int{0}
}

func (x *FrameConfig) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *FrameConfig) GetPose() *v1.Pose {
	if x != nil {
		return x.Pose
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FrameConfig *FrameConfig `protobuf:"bytes,2,opt,name=frame_config,json=frameConfig,proto3" json:"frame_config,omitempty"`
	ModelJson   []byte       `protobuf:"bytes,3,opt,name=model_json,json=modelJson,proto3" json:"model_json,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetFrameConfig() *FrameConfig {
	if x != nil {
		return x.FrameConfig
	}
	return nil
}

func (x *Config) GetModelJson() []byte {
	if x != nil {
		return x.ModelJson
	}
	return nil
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP(), []int{2}
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameSystemConfigs []*Config `protobuf:"bytes,1,rep,name=frame_system_configs,json=frameSystemConfigs,proto3" json:"frame_system_configs,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigResponse) GetFrameSystemConfigs() []*Config {
	if x != nil {
		return x.FrameSystemConfigs
	}
	return nil
}

type TransformPoseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      *v1.PoseInFrame `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination string          `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *TransformPoseRequest) Reset() {
	*x = TransformPoseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformPoseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformPoseRequest) ProtoMessage() {}

func (x *TransformPoseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformPoseRequest.ProtoReflect.Descriptor instead.
func (*TransformPoseRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP(), []int{4}
}

func (x *TransformPoseRequest) GetSource() *v1.PoseInFrame {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *TransformPoseRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type TransformPoseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pose *v1.Pose `protobuf:"bytes,1,opt,name=pose,proto3" json:"pose,omitempty"`
}

func (x *TransformPoseResponse) Reset() {
	*x = TransformPoseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformPoseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformPoseResponse) ProtoMessage() {}

func (x *TransformPoseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformPoseResponse.ProtoReflect.Descriptor instead.
func (*TransformPoseResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP(), []int{5}
}

func (x *TransformPoseResponse) GetPose() *v1.Pose {
	if x != nil {
		return x.Pose
	}
	return nil
}

var File_proto_api_service_framesystem_v1_frame_system_proto protoreflect.FileDescriptor

var file_proto_api_service_framesystem_v1_frame_system_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0b, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2d,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x8d, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0c,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x0f, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6c,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5a, 0x0a, 0x14, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x72, 0x0a, 0x14,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x65, 0x49,
	0x6e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x46, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x6f, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x65, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x32, 0xf1, 0x02, 0x0a, 0x12, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x9d, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0xba, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73,
	0x65, 0x12, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x70, 0x6f, 0x73, 0x65, 0x42, 0x61, 0x0a, 0x2d,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x30, 0x67,
	0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_service_framesystem_v1_frame_system_proto_rawDescOnce sync.Once
	file_proto_api_service_framesystem_v1_frame_system_proto_rawDescData = file_proto_api_service_framesystem_v1_frame_system_proto_rawDesc
)

func file_proto_api_service_framesystem_v1_frame_system_proto_rawDescGZIP() []byte {
	file_proto_api_service_framesystem_v1_frame_system_proto_rawDescOnce.Do(func() {
		file_proto_api_service_framesystem_v1_frame_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_service_framesystem_v1_frame_system_proto_rawDescData)
	})
	return file_proto_api_service_framesystem_v1_frame_system_proto_rawDescData
}

var file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_api_service_framesystem_v1_frame_system_proto_goTypes = []interface{}{
	(*FrameConfig)(nil),           // 0: proto.api.service.framesystem.v1.FrameConfig
	(*Config)(nil),                // 1: proto.api.service.framesystem.v1.Config
	(*ConfigRequest)(nil),         // 2: proto.api.service.framesystem.v1.ConfigRequest
	(*ConfigResponse)(nil),        // 3: proto.api.service.framesystem.v1.ConfigResponse
	(*TransformPoseRequest)(nil),  // 4: proto.api.service.framesystem.v1.TransformPoseRequest
	(*TransformPoseResponse)(nil), // 5: proto.api.service.framesystem.v1.TransformPoseResponse
	(*v1.Pose)(nil),               // 6: proto.api.common.v1.Pose
	(*v1.PoseInFrame)(nil),        // 7: proto.api.common.v1.PoseInFrame
}
var file_proto_api_service_framesystem_v1_frame_system_proto_depIdxs = []int32{
	6, // 0: proto.api.service.framesystem.v1.FrameConfig.pose:type_name -> proto.api.common.v1.Pose
	0, // 1: proto.api.service.framesystem.v1.Config.frame_config:type_name -> proto.api.service.framesystem.v1.FrameConfig
	1, // 2: proto.api.service.framesystem.v1.ConfigResponse.frame_system_configs:type_name -> proto.api.service.framesystem.v1.Config
	7, // 3: proto.api.service.framesystem.v1.TransformPoseRequest.source:type_name -> proto.api.common.v1.PoseInFrame
	6, // 4: proto.api.service.framesystem.v1.TransformPoseResponse.pose:type_name -> proto.api.common.v1.Pose
	2, // 5: proto.api.service.framesystem.v1.FrameSystemService.Config:input_type -> proto.api.service.framesystem.v1.ConfigRequest
	4, // 6: proto.api.service.framesystem.v1.FrameSystemService.TransformPose:input_type -> proto.api.service.framesystem.v1.TransformPoseRequest
	3, // 7: proto.api.service.framesystem.v1.FrameSystemService.Config:output_type -> proto.api.service.framesystem.v1.ConfigResponse
	5, // 8: proto.api.service.framesystem.v1.FrameSystemService.TransformPose:output_type -> proto.api.service.framesystem.v1.TransformPoseResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_api_service_framesystem_v1_frame_system_proto_init() }
func file_proto_api_service_framesystem_v1_frame_system_proto_init() {
	if File_proto_api_service_framesystem_v1_frame_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameConfig); i {
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
		file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
		file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformPoseRequest); i {
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
		file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformPoseResponse); i {
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
			RawDescriptor: file_proto_api_service_framesystem_v1_frame_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_service_framesystem_v1_frame_system_proto_goTypes,
		DependencyIndexes: file_proto_api_service_framesystem_v1_frame_system_proto_depIdxs,
		MessageInfos:      file_proto_api_service_framesystem_v1_frame_system_proto_msgTypes,
	}.Build()
	File_proto_api_service_framesystem_v1_frame_system_proto = out.File
	file_proto_api_service_framesystem_v1_frame_system_proto_rawDesc = nil
	file_proto_api_service_framesystem_v1_frame_system_proto_goTypes = nil
	file_proto_api_service_framesystem_v1_frame_system_proto_depIdxs = nil
}
