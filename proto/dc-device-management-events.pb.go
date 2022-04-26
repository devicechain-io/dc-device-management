// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: proto/dc-device-management-events.proto

package proto

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

//*
// Enumeration of failure reasons.
type FailureReason int32

const (
	FailureReason_Unknown                   FailureReason = 0 // Failed for unknown reason
	FailureReason_Invalid                   FailureReason = 1 // Event was not able to be parsed
	FailureReason_ApiCallFailed             FailureReason = 2 // API call required for resolution failed
	FailureReason_DeviceNotFound            FailureReason = 3 // Device token could not be resolved to a device
	FailureReason_NoActiveDeviceAssignments FailureReason = 4 // Device had no active assignments
)

// Enum value maps for FailureReason.
var (
	FailureReason_name = map[int32]string{
		0: "Unknown",
		1: "Invalid",
		2: "ApiCallFailed",
		3: "DeviceNotFound",
		4: "NoActiveDeviceAssignments",
	}
	FailureReason_value = map[string]int32{
		"Unknown":                   0,
		"Invalid":                   1,
		"ApiCallFailed":             2,
		"DeviceNotFound":            3,
		"NoActiveDeviceAssignments": 4,
	}
)

func (x FailureReason) Enum() *FailureReason {
	p := new(FailureReason)
	*p = x
	return p
}

func (x FailureReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailureReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dc_device_management_events_proto_enumTypes[0].Descriptor()
}

func (FailureReason) Type() protoreflect.EnumType {
	return &file_proto_dc_device_management_events_proto_enumTypes[0]
}

func (x FailureReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailureReason.Descriptor instead.
func (FailureReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_dc_device_management_events_proto_rawDescGZIP(), []int{0}
}

//*
// Event that could not be processed.
type PFailedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason  FailureReason `protobuf:"varint,1,opt,name=reason,proto3,enum=io.devicechain.devicemanagement.FailureReason" json:"reason,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Payload []byte        `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PFailedEvent) Reset() {
	*x = PFailedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dc_device_management_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PFailedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PFailedEvent) ProtoMessage() {}

func (x *PFailedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dc_device_management_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PFailedEvent.ProtoReflect.Descriptor instead.
func (*PFailedEvent) Descriptor() ([]byte, []int) {
	return file_proto_dc_device_management_events_proto_rawDescGZIP(), []int{0}
}

func (x *PFailedEvent) GetReason() FailureReason {
	if x != nil {
		return x.Reason
	}
	return FailureReason_Unknown
}

func (x *PFailedEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PFailedEvent) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_proto_dc_device_management_events_proto protoreflect.FileDescriptor

var file_proto_dc_device_management_events_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x69, 0x6f, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x50,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x69, 0x6f,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x6f, 0x0a, 0x0d, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x70, 0x69, 0x43, 0x61, 0x6c, 0x6c, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x6f, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x10, 0x04, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dc_device_management_events_proto_rawDescOnce sync.Once
	file_proto_dc_device_management_events_proto_rawDescData = file_proto_dc_device_management_events_proto_rawDesc
)

func file_proto_dc_device_management_events_proto_rawDescGZIP() []byte {
	file_proto_dc_device_management_events_proto_rawDescOnce.Do(func() {
		file_proto_dc_device_management_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dc_device_management_events_proto_rawDescData)
	})
	return file_proto_dc_device_management_events_proto_rawDescData
}

var file_proto_dc_device_management_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_dc_device_management_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_dc_device_management_events_proto_goTypes = []interface{}{
	(FailureReason)(0),   // 0: io.devicechain.devicemanagement.FailureReason
	(*PFailedEvent)(nil), // 1: io.devicechain.devicemanagement.PFailedEvent
}
var file_proto_dc_device_management_events_proto_depIdxs = []int32{
	0, // 0: io.devicechain.devicemanagement.PFailedEvent.reason:type_name -> io.devicechain.devicemanagement.FailureReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_dc_device_management_events_proto_init() }
func file_proto_dc_device_management_events_proto_init() {
	if File_proto_dc_device_management_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dc_device_management_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PFailedEvent); i {
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
			RawDescriptor: file_proto_dc_device_management_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_dc_device_management_events_proto_goTypes,
		DependencyIndexes: file_proto_dc_device_management_events_proto_depIdxs,
		EnumInfos:         file_proto_dc_device_management_events_proto_enumTypes,
		MessageInfos:      file_proto_dc_device_management_events_proto_msgTypes,
	}.Build()
	File_proto_dc_device_management_events_proto = out.File
	file_proto_dc_device_management_events_proto_rawDesc = nil
	file_proto_dc_device_management_events_proto_goTypes = nil
	file_proto_dc_device_management_events_proto_depIdxs = nil
}
