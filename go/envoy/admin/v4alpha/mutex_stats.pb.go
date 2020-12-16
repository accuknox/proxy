// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/admin/v4alpha/mutex_stats.proto

package envoy_admin_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Proto representation of the statistics collected upon absl::Mutex contention, if Envoy is run
// under :option:`--enable-mutex-tracing`. For more information, see the `absl::Mutex`
// [docs](https://abseil.io/about/design/mutex#extra-features).
//
// *NB*: The wait cycles below are measured by `absl::base_internal::CycleClock`, and may not
// correspond to core clock frequency. For more information, see the `CycleClock`
// [docs](https://github.com/abseil/abseil-cpp/blob/master/absl/base/internal/cycleclock.h).
type MutexStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of individual mutex contentions which have occurred since startup.
	NumContentions uint64 `protobuf:"varint,1,opt,name=num_contentions,json=numContentions,proto3" json:"num_contentions,omitempty"`
	// The length of the current contention wait cycle.
	CurrentWaitCycles uint64 `protobuf:"varint,2,opt,name=current_wait_cycles,json=currentWaitCycles,proto3" json:"current_wait_cycles,omitempty"`
	// The lifetime total of all contention wait cycles.
	LifetimeWaitCycles uint64 `protobuf:"varint,3,opt,name=lifetime_wait_cycles,json=lifetimeWaitCycles,proto3" json:"lifetime_wait_cycles,omitempty"`
}

func (x *MutexStats) Reset() {
	*x = MutexStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v4alpha_mutex_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutexStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutexStats) ProtoMessage() {}

func (x *MutexStats) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v4alpha_mutex_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutexStats.ProtoReflect.Descriptor instead.
func (*MutexStats) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v4alpha_mutex_stats_proto_rawDescGZIP(), []int{0}
}

func (x *MutexStats) GetNumContentions() uint64 {
	if x != nil {
		return x.NumContentions
	}
	return 0
}

func (x *MutexStats) GetCurrentWaitCycles() uint64 {
	if x != nil {
		return x.CurrentWaitCycles
	}
	return 0
}

func (x *MutexStats) GetLifetimeWaitCycles() uint64 {
	if x != nil {
		return x.LifetimeWaitCycles
	}
	return 0
}

var File_envoy_admin_v4alpha_mutex_stats_proto protoreflect.FileDescriptor

var file_envoy_admin_v4alpha_mutex_stats_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9,
	0x01, 0x0a, 0x0a, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x69, 0x74,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x61,
	0x69, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x3a, 0x20, 0x9a, 0xc5, 0x88, 0x1e, 0x1b, 0x0a,
	0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x75, 0x74, 0x65, 0x78, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x3e, 0x0a, 0x21, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x0f, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_admin_v4alpha_mutex_stats_proto_rawDescOnce sync.Once
	file_envoy_admin_v4alpha_mutex_stats_proto_rawDescData = file_envoy_admin_v4alpha_mutex_stats_proto_rawDesc
)

func file_envoy_admin_v4alpha_mutex_stats_proto_rawDescGZIP() []byte {
	file_envoy_admin_v4alpha_mutex_stats_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v4alpha_mutex_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v4alpha_mutex_stats_proto_rawDescData)
	})
	return file_envoy_admin_v4alpha_mutex_stats_proto_rawDescData
}

var file_envoy_admin_v4alpha_mutex_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_admin_v4alpha_mutex_stats_proto_goTypes = []interface{}{
	(*MutexStats)(nil), // 0: envoy.admin.v4alpha.MutexStats
}
var file_envoy_admin_v4alpha_mutex_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_admin_v4alpha_mutex_stats_proto_init() }
func file_envoy_admin_v4alpha_mutex_stats_proto_init() {
	if File_envoy_admin_v4alpha_mutex_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v4alpha_mutex_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutexStats); i {
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
			RawDescriptor: file_envoy_admin_v4alpha_mutex_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v4alpha_mutex_stats_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v4alpha_mutex_stats_proto_depIdxs,
		MessageInfos:      file_envoy_admin_v4alpha_mutex_stats_proto_msgTypes,
	}.Build()
	File_envoy_admin_v4alpha_mutex_stats_proto = out.File
	file_envoy_admin_v4alpha_mutex_stats_proto_rawDesc = nil
	file_envoy_admin_v4alpha_mutex_stats_proto_goTypes = nil
	file_envoy_admin_v4alpha_mutex_stats_proto_depIdxs = nil
}
