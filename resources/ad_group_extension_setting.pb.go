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
// source: google/ads/googleads/v11/resources/ad_group_extension_setting.proto

package resources

import (
	enums "github.com/ercling/google-ads-go/enums"
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

// An ad group extension setting.
type AdGroupExtensionSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad group extension setting.
	// AdGroupExtensionSetting resource names have the form:
	//
	// `customers/{customer_id}/adGroupExtensionSettings/{ad_group_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The extension type of the ad group extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v11.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Immutable. The resource name of the ad group. The linked extension feed items will
	// serve under this ad group.
	// AdGroup resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	AdGroup *string `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3,oneof" json:"ad_group,omitempty"`
	// The resource names of the extension feed items to serve under the ad group.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []string `protobuf:"bytes,7,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v11.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
}

func (x *AdGroupExtensionSetting) Reset() {
	*x = AdGroupExtensionSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_ad_group_extension_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupExtensionSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupExtensionSetting) ProtoMessage() {}

func (x *AdGroupExtensionSetting) ProtoReflect() protoreflect.Message {
	mi := &file_resources_ad_group_extension_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupExtensionSetting.ProtoReflect.Descriptor instead.
func (*AdGroupExtensionSetting) Descriptor() ([]byte, []int) {
	return file_resources_ad_group_extension_setting_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupExtensionSetting) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if x != nil {
		return x.ExtensionType
	}
	return enums.ExtensionTypeEnum_ExtensionType(0)
}

func (x *AdGroupExtensionSetting) GetAdGroup() string {
	if x != nil && x.AdGroup != nil {
		return *x.AdGroup
	}
	return ""
}

func (x *AdGroupExtensionSetting) GetExtensionFeedItems() []string {
	if x != nil {
		return x.ExtensionFeedItems
	}
	return nil
}

func (x *AdGroupExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if x != nil {
		return x.Device
	}
	return enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice(0)
}

var File_resources_ad_group_extension_setting_proto protoreflect.FileDescriptor

var file_resources_ad_group_extension_setting_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x05, 0x0a, 0x17, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x5d, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xe0, 0x41,
	0x05, 0xfa, 0x41, 0x32, 0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6b, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x48, 0x0a, 0x08, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52,
	0x07, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x61, 0x0a, 0x14, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2f, 0xfa, 0x41, 0x2c, 0x0a, 0x2a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x69,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x87, 0x01, 0xea, 0x41, 0x83, 0x01,
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x4f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x7d, 0x7e, 0x7b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x8e, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1c, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_ad_group_extension_setting_proto_rawDescOnce sync.Once
	file_resources_ad_group_extension_setting_proto_rawDescData = file_resources_ad_group_extension_setting_proto_rawDesc
)

func file_resources_ad_group_extension_setting_proto_rawDescGZIP() []byte {
	file_resources_ad_group_extension_setting_proto_rawDescOnce.Do(func() {
		file_resources_ad_group_extension_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_ad_group_extension_setting_proto_rawDescData)
	})
	return file_resources_ad_group_extension_setting_proto_rawDescData
}

var file_resources_ad_group_extension_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_ad_group_extension_setting_proto_goTypes = []interface{}{
	(*AdGroupExtensionSetting)(nil),                              // 0: google.ads.googleads.v11.resources.AdGroupExtensionSetting
	(enums.ExtensionTypeEnum_ExtensionType)(0),                   // 1: google.ads.googleads.v11.enums.ExtensionTypeEnum.ExtensionType
	(enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice)(0), // 2: google.ads.googleads.v11.enums.ExtensionSettingDeviceEnum.ExtensionSettingDevice
}
var file_resources_ad_group_extension_setting_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.resources.AdGroupExtensionSetting.extension_type:type_name -> google.ads.googleads.v11.enums.ExtensionTypeEnum.ExtensionType
	2, // 1: google.ads.googleads.v11.resources.AdGroupExtensionSetting.device:type_name -> google.ads.googleads.v11.enums.ExtensionSettingDeviceEnum.ExtensionSettingDevice
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resources_ad_group_extension_setting_proto_init() }
func file_resources_ad_group_extension_setting_proto_init() {
	if File_resources_ad_group_extension_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_ad_group_extension_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupExtensionSetting); i {
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
	file_resources_ad_group_extension_setting_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_ad_group_extension_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_ad_group_extension_setting_proto_goTypes,
		DependencyIndexes: file_resources_ad_group_extension_setting_proto_depIdxs,
		MessageInfos:      file_resources_ad_group_extension_setting_proto_msgTypes,
	}.Build()
	File_resources_ad_group_extension_setting_proto = out.File
	file_resources_ad_group_extension_setting_proto_rawDesc = nil
	file_resources_ad_group_extension_setting_proto_goTypes = nil
	file_resources_ad_group_extension_setting_proto_depIdxs = nil
}
