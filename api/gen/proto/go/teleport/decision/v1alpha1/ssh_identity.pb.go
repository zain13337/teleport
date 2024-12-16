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
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: teleport/decision/v1alpha1/ssh_identity.proto

package decisionpb

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

// SSHIdentity is the identity used for SSH connections.
type SSHIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SSHIdentity) Reset() {
	*x = SSHIdentity{}
	mi := &file_teleport_decision_v1alpha1_ssh_identity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SSHIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHIdentity) ProtoMessage() {}

func (x *SSHIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_ssh_identity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHIdentity.ProtoReflect.Descriptor instead.
func (*SSHIdentity) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescGZIP(), []int{0}
}

var File_teleport_decision_v1alpha1_ssh_identity_proto protoreflect.FileDescriptor

var file_teleport_decision_v1alpha1_ssh_identity_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x73, 0x68,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x0d, 0x0a, 0x0b, 0x53,
	0x53, 0x48, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescOnce sync.Once
	file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescData = file_teleport_decision_v1alpha1_ssh_identity_proto_rawDesc
)

func file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescGZIP() []byte {
	file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescOnce.Do(func() {
		file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescData)
	})
	return file_teleport_decision_v1alpha1_ssh_identity_proto_rawDescData
}

var file_teleport_decision_v1alpha1_ssh_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_decision_v1alpha1_ssh_identity_proto_goTypes = []any{
	(*SSHIdentity)(nil), // 0: teleport.decision.v1alpha1.SSHIdentity
}
var file_teleport_decision_v1alpha1_ssh_identity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_decision_v1alpha1_ssh_identity_proto_init() }
func file_teleport_decision_v1alpha1_ssh_identity_proto_init() {
	if File_teleport_decision_v1alpha1_ssh_identity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_decision_v1alpha1_ssh_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_decision_v1alpha1_ssh_identity_proto_goTypes,
		DependencyIndexes: file_teleport_decision_v1alpha1_ssh_identity_proto_depIdxs,
		MessageInfos:      file_teleport_decision_v1alpha1_ssh_identity_proto_msgTypes,
	}.Build()
	File_teleport_decision_v1alpha1_ssh_identity_proto = out.File
	file_teleport_decision_v1alpha1_ssh_identity_proto_rawDesc = nil
	file_teleport_decision_v1alpha1_ssh_identity_proto_goTypes = nil
	file_teleport_decision_v1alpha1_ssh_identity_proto_depIdxs = nil
}
