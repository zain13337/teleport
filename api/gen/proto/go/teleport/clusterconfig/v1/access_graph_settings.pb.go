// Copyright 2024 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: teleport/clusterconfig/v1/access_graph_settings.proto

package clusterconfigv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// AccessGraphSecretsScanConfig is used to configure the parameters for the secrets scanning functionality.
type AccessGraphSecretsScanConfig int32

const (
	// ACCESS_GRAPH_SECRETS_SCAN_CONFIG_UNSPECIFIED is an unknown secrets scan configuration.
	AccessGraphSecretsScanConfig_ACCESS_GRAPH_SECRETS_SCAN_CONFIG_UNSPECIFIED AccessGraphSecretsScanConfig = 0
	// ACCESS_GRAPH_SECRETS_SCAN_CONFIG_DISABLED is a disabled secrets scan configuration.
	AccessGraphSecretsScanConfig_ACCESS_GRAPH_SECRETS_SCAN_CONFIG_DISABLED AccessGraphSecretsScanConfig = 1
	// ACCESS_GRAPH_SECRETS_SCAN_CONFIG_ENABLED is an enabled secrets scan configuration.
	AccessGraphSecretsScanConfig_ACCESS_GRAPH_SECRETS_SCAN_CONFIG_ENABLED AccessGraphSecretsScanConfig = 2
)

// Enum value maps for AccessGraphSecretsScanConfig.
var (
	AccessGraphSecretsScanConfig_name = map[int32]string{
		0: "ACCESS_GRAPH_SECRETS_SCAN_CONFIG_UNSPECIFIED",
		1: "ACCESS_GRAPH_SECRETS_SCAN_CONFIG_DISABLED",
		2: "ACCESS_GRAPH_SECRETS_SCAN_CONFIG_ENABLED",
	}
	AccessGraphSecretsScanConfig_value = map[string]int32{
		"ACCESS_GRAPH_SECRETS_SCAN_CONFIG_UNSPECIFIED": 0,
		"ACCESS_GRAPH_SECRETS_SCAN_CONFIG_DISABLED":    1,
		"ACCESS_GRAPH_SECRETS_SCAN_CONFIG_ENABLED":     2,
	}
)

func (x AccessGraphSecretsScanConfig) Enum() *AccessGraphSecretsScanConfig {
	p := new(AccessGraphSecretsScanConfig)
	*p = x
	return p
}

func (x AccessGraphSecretsScanConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessGraphSecretsScanConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_clusterconfig_v1_access_graph_settings_proto_enumTypes[0].Descriptor()
}

func (AccessGraphSecretsScanConfig) Type() protoreflect.EnumType {
	return &file_teleport_clusterconfig_v1_access_graph_settings_proto_enumTypes[0]
}

func (x AccessGraphSecretsScanConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessGraphSecretsScanConfig.Descriptor instead.
func (AccessGraphSecretsScanConfig) EnumDescriptor() ([]byte, []int) {
	return file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescGZIP(), []int{0}
}

// AccessGraphSettings holds dynamic configuration settings for the Access Graph service.
type AccessGraphSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// kind is the kind of the resource.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// sub_kind is the sub kind of the resource.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// version is the version of the resource.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// metadata is the metadata of the resource.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec is the spec of the resource.
	Spec          *AccessGraphSettingsSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessGraphSettings) Reset() {
	*x = AccessGraphSettings{}
	mi := &file_teleport_clusterconfig_v1_access_graph_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessGraphSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGraphSettings) ProtoMessage() {}

func (x *AccessGraphSettings) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_clusterconfig_v1_access_graph_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGraphSettings.ProtoReflect.Descriptor instead.
func (*AccessGraphSettings) Descriptor() ([]byte, []int) {
	return file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescGZIP(), []int{0}
}

func (x *AccessGraphSettings) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AccessGraphSettings) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *AccessGraphSettings) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AccessGraphSettings) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AccessGraphSettings) GetSpec() *AccessGraphSettingsSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// AccessGraphSettingsSpec is the spec for the Access Graph service configuration settings.
type AccessGraphSettingsSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// secrets_scan_config is used to configure the parameters for the secrets scanning functionality.
	SecretsScanConfig AccessGraphSecretsScanConfig `protobuf:"varint,1,opt,name=secrets_scan_config,json=secretsScanConfig,proto3,enum=teleport.clusterconfig.v1.AccessGraphSecretsScanConfig" json:"secrets_scan_config,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *AccessGraphSettingsSpec) Reset() {
	*x = AccessGraphSettingsSpec{}
	mi := &file_teleport_clusterconfig_v1_access_graph_settings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessGraphSettingsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGraphSettingsSpec) ProtoMessage() {}

func (x *AccessGraphSettingsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_clusterconfig_v1_access_graph_settings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGraphSettingsSpec.ProtoReflect.Descriptor instead.
func (*AccessGraphSettingsSpec) Descriptor() ([]byte, []int) {
	return file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescGZIP(), []int{1}
}

func (x *AccessGraphSettingsSpec) GetSecretsScanConfig() AccessGraphSecretsScanConfig {
	if x != nil {
		return x.SecretsScanConfig
	}
	return AccessGraphSecretsScanConfig_ACCESS_GRAPH_SECRETS_SCAN_CONFIG_UNSPECIFIED
}

var File_teleport_clusterconfig_v1_access_graph_settings_proto protoreflect.FileDescriptor

var file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x46, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x67, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x37, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0xad, 0x01,
	0x0a, 0x1c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30,
	0x0a, 0x2c, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x53,
	0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x2d, 0x0a, 0x29, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48,
	0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x2c, 0x0a, 0x28, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f,
	0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x47, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x5e, 0x5a,
	0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescOnce sync.Once
	file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescData = file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDesc
)

func file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescGZIP() []byte {
	file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescOnce.Do(func() {
		file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescData)
	})
	return file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDescData
}

var file_teleport_clusterconfig_v1_access_graph_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_clusterconfig_v1_access_graph_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_clusterconfig_v1_access_graph_settings_proto_goTypes = []any{
	(AccessGraphSecretsScanConfig)(0), // 0: teleport.clusterconfig.v1.AccessGraphSecretsScanConfig
	(*AccessGraphSettings)(nil),       // 1: teleport.clusterconfig.v1.AccessGraphSettings
	(*AccessGraphSettingsSpec)(nil),   // 2: teleport.clusterconfig.v1.AccessGraphSettingsSpec
	(*v1.Metadata)(nil),               // 3: teleport.header.v1.Metadata
}
var file_teleport_clusterconfig_v1_access_graph_settings_proto_depIdxs = []int32{
	3, // 0: teleport.clusterconfig.v1.AccessGraphSettings.metadata:type_name -> teleport.header.v1.Metadata
	2, // 1: teleport.clusterconfig.v1.AccessGraphSettings.spec:type_name -> teleport.clusterconfig.v1.AccessGraphSettingsSpec
	0, // 2: teleport.clusterconfig.v1.AccessGraphSettingsSpec.secrets_scan_config:type_name -> teleport.clusterconfig.v1.AccessGraphSecretsScanConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_teleport_clusterconfig_v1_access_graph_settings_proto_init() }
func file_teleport_clusterconfig_v1_access_graph_settings_proto_init() {
	if File_teleport_clusterconfig_v1_access_graph_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_clusterconfig_v1_access_graph_settings_proto_goTypes,
		DependencyIndexes: file_teleport_clusterconfig_v1_access_graph_settings_proto_depIdxs,
		EnumInfos:         file_teleport_clusterconfig_v1_access_graph_settings_proto_enumTypes,
		MessageInfos:      file_teleport_clusterconfig_v1_access_graph_settings_proto_msgTypes,
	}.Build()
	File_teleport_clusterconfig_v1_access_graph_settings_proto = out.File
	file_teleport_clusterconfig_v1_access_graph_settings_proto_rawDesc = nil
	file_teleport_clusterconfig_v1_access_graph_settings_proto_goTypes = nil
	file_teleport_clusterconfig_v1_access_graph_settings_proto_depIdxs = nil
}
