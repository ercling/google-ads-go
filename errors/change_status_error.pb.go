// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: google/ads/googleads/v11/errors/change_status_error.proto

package errors

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

// Enum describing possible change status errors.
type ChangeStatusErrorEnum_ChangeStatusError int32

const (
	// Enum unspecified.
	ChangeStatusErrorEnum_UNSPECIFIED ChangeStatusErrorEnum_ChangeStatusError = 0
	// The received error code is not known in this version.
	ChangeStatusErrorEnum_UNKNOWN ChangeStatusErrorEnum_ChangeStatusError = 1
	// The requested start date is too old.
	ChangeStatusErrorEnum_START_DATE_TOO_OLD ChangeStatusErrorEnum_ChangeStatusError = 3
	// The change_status search request must specify a finite range filter
	// on last_change_date_time.
	ChangeStatusErrorEnum_CHANGE_DATE_RANGE_INFINITE ChangeStatusErrorEnum_ChangeStatusError = 4
	// The change status search request has specified invalid date time filters
	// that can never logically produce any valid results (for example, start
	// time after end time).
	ChangeStatusErrorEnum_CHANGE_DATE_RANGE_NEGATIVE ChangeStatusErrorEnum_ChangeStatusError = 5
	// The change_status search request must specify a LIMIT.
	ChangeStatusErrorEnum_LIMIT_NOT_SPECIFIED ChangeStatusErrorEnum_ChangeStatusError = 6
	// The LIMIT specified by change_status request should be less than or equal
	// to 10K.
	ChangeStatusErrorEnum_INVALID_LIMIT_CLAUSE ChangeStatusErrorEnum_ChangeStatusError = 7
)

// Enum value maps for ChangeStatusErrorEnum_ChangeStatusError.
var (
	ChangeStatusErrorEnum_ChangeStatusError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		3: "START_DATE_TOO_OLD",
		4: "CHANGE_DATE_RANGE_INFINITE",
		5: "CHANGE_DATE_RANGE_NEGATIVE",
		6: "LIMIT_NOT_SPECIFIED",
		7: "INVALID_LIMIT_CLAUSE",
	}
	ChangeStatusErrorEnum_ChangeStatusError_value = map[string]int32{
		"UNSPECIFIED":                0,
		"UNKNOWN":                    1,
		"START_DATE_TOO_OLD":         3,
		"CHANGE_DATE_RANGE_INFINITE": 4,
		"CHANGE_DATE_RANGE_NEGATIVE": 5,
		"LIMIT_NOT_SPECIFIED":        6,
		"INVALID_LIMIT_CLAUSE":       7,
	}
)

func (x ChangeStatusErrorEnum_ChangeStatusError) Enum() *ChangeStatusErrorEnum_ChangeStatusError {
	p := new(ChangeStatusErrorEnum_ChangeStatusError)
	*p = x
	return p
}

func (x ChangeStatusErrorEnum_ChangeStatusError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeStatusErrorEnum_ChangeStatusError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_change_status_error_proto_enumTypes[0].Descriptor()
}

func (ChangeStatusErrorEnum_ChangeStatusError) Type() protoreflect.EnumType {
	return &file_errors_change_status_error_proto_enumTypes[0]
}

func (x ChangeStatusErrorEnum_ChangeStatusError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeStatusErrorEnum_ChangeStatusError.Descriptor instead.
func (ChangeStatusErrorEnum_ChangeStatusError) EnumDescriptor() ([]byte, []int) {
	return file_errors_change_status_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible change status errors.
type ChangeStatusErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeStatusErrorEnum) Reset() {
	*x = ChangeStatusErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_change_status_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusErrorEnum) ProtoMessage() {}

func (x *ChangeStatusErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_change_status_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusErrorEnum.ProtoReflect.Descriptor instead.
func (*ChangeStatusErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_change_status_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_change_status_error_proto protoreflect.FileDescriptor

var file_errors_change_status_error_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xd6, 0x01, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xbc, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4f, 0x4c, 0x44,
	0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x45,
	0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x43, 0x4c, 0x41,
	0x55, 0x53, 0x45, 0x10, 0x07, 0x42, 0xf6, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x16, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x31, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_change_status_error_proto_rawDescOnce sync.Once
	file_errors_change_status_error_proto_rawDescData = file_errors_change_status_error_proto_rawDesc
)

func file_errors_change_status_error_proto_rawDescGZIP() []byte {
	file_errors_change_status_error_proto_rawDescOnce.Do(func() {
		file_errors_change_status_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_change_status_error_proto_rawDescData)
	})
	return file_errors_change_status_error_proto_rawDescData
}

var file_errors_change_status_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_change_status_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_change_status_error_proto_goTypes = []interface{}{
	(ChangeStatusErrorEnum_ChangeStatusError)(0), // 0: google.ads.googleads.v11.errors.ChangeStatusErrorEnum.ChangeStatusError
	(*ChangeStatusErrorEnum)(nil),                // 1: google.ads.googleads.v11.errors.ChangeStatusErrorEnum
}
var file_errors_change_status_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_change_status_error_proto_init() }
func file_errors_change_status_error_proto_init() {
	if File_errors_change_status_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_change_status_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusErrorEnum); i {
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
			RawDescriptor: file_errors_change_status_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_change_status_error_proto_goTypes,
		DependencyIndexes: file_errors_change_status_error_proto_depIdxs,
		EnumInfos:         file_errors_change_status_error_proto_enumTypes,
		MessageInfos:      file_errors_change_status_error_proto_msgTypes,
	}.Build()
	File_errors_change_status_error_proto = out.File
	file_errors_change_status_error_proto_rawDesc = nil
	file_errors_change_status_error_proto_goTypes = nil
	file_errors_change_status_error_proto_depIdxs = nil
}
