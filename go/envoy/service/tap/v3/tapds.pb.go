// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/service/tap/v3/tapds.proto

package envoy_service_tap_v3

import (
	context "context"
	v3 "github.com/cilium/proxy/go/envoy/config/tap/v3"
	v31 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// [#not-implemented-hide:] A tap resource is essentially a tap configuration with a name
// The filter TapDS config references this name.
type TapResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the tap configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Tap config to apply
	Config *v3.TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *TapResource) Reset() {
	*x = TapResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_tap_v3_tapds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TapResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TapResource) ProtoMessage() {}

func (x *TapResource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_tap_v3_tapds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TapResource.ProtoReflect.Descriptor instead.
func (*TapResource) Descriptor() ([]byte, []int) {
	return file_envoy_service_tap_v3_tapds_proto_rawDescGZIP(), []int{0}
}

func (x *TapResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TapResource) GetConfig() *v3.TapConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_envoy_service_tap_v3_tapds_proto protoreflect.FileDescriptor

var file_envoy_service_tap_v3_tapds_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x74, 0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x61, 0x70, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x90, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x61, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x32, 0xa9, 0x03, 0x0a, 0x13, 0x54, 0x61, 0x70, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2c,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x7e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x61, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x12, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x9a, 0x01, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x61, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x19, 0x2f, 0x76, 0x33, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x74, 0x61, 0x70, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x3a, 0x01, 0x2a, 0x42, 0x3d,
	0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x61,
	0x70, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x54, 0x61, 0x70, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x88, 0x01, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_tap_v3_tapds_proto_rawDescOnce sync.Once
	file_envoy_service_tap_v3_tapds_proto_rawDescData = file_envoy_service_tap_v3_tapds_proto_rawDesc
)

func file_envoy_service_tap_v3_tapds_proto_rawDescGZIP() []byte {
	file_envoy_service_tap_v3_tapds_proto_rawDescOnce.Do(func() {
		file_envoy_service_tap_v3_tapds_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_tap_v3_tapds_proto_rawDescData)
	})
	return file_envoy_service_tap_v3_tapds_proto_rawDescData
}

var file_envoy_service_tap_v3_tapds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_service_tap_v3_tapds_proto_goTypes = []interface{}{
	(*TapResource)(nil),                // 0: envoy.service.tap.v3.TapResource
	(*v3.TapConfig)(nil),               // 1: envoy.config.tap.v3.TapConfig
	(*v31.DiscoveryRequest)(nil),       // 2: envoy.service.discovery.v3.DiscoveryRequest
	(*v31.DeltaDiscoveryRequest)(nil),  // 3: envoy.service.discovery.v3.DeltaDiscoveryRequest
	(*v31.DiscoveryResponse)(nil),      // 4: envoy.service.discovery.v3.DiscoveryResponse
	(*v31.DeltaDiscoveryResponse)(nil), // 5: envoy.service.discovery.v3.DeltaDiscoveryResponse
}
var file_envoy_service_tap_v3_tapds_proto_depIdxs = []int32{
	1, // 0: envoy.service.tap.v3.TapResource.config:type_name -> envoy.config.tap.v3.TapConfig
	2, // 1: envoy.service.tap.v3.TapDiscoveryService.StreamTapConfigs:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	3, // 2: envoy.service.tap.v3.TapDiscoveryService.DeltaTapConfigs:input_type -> envoy.service.discovery.v3.DeltaDiscoveryRequest
	2, // 3: envoy.service.tap.v3.TapDiscoveryService.FetchTapConfigs:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	4, // 4: envoy.service.tap.v3.TapDiscoveryService.StreamTapConfigs:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	5, // 5: envoy.service.tap.v3.TapDiscoveryService.DeltaTapConfigs:output_type -> envoy.service.discovery.v3.DeltaDiscoveryResponse
	4, // 6: envoy.service.tap.v3.TapDiscoveryService.FetchTapConfigs:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_service_tap_v3_tapds_proto_init() }
func file_envoy_service_tap_v3_tapds_proto_init() {
	if File_envoy_service_tap_v3_tapds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_tap_v3_tapds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TapResource); i {
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
			RawDescriptor: file_envoy_service_tap_v3_tapds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_tap_v3_tapds_proto_goTypes,
		DependencyIndexes: file_envoy_service_tap_v3_tapds_proto_depIdxs,
		MessageInfos:      file_envoy_service_tap_v3_tapds_proto_msgTypes,
	}.Build()
	File_envoy_service_tap_v3_tapds_proto = out.File
	file_envoy_service_tap_v3_tapds_proto_rawDesc = nil
	file_envoy_service_tap_v3_tapds_proto_goTypes = nil
	file_envoy_service_tap_v3_tapds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v31.DiscoveryRequest, opts ...grpc.CallOption) (*v31.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTapDiscoveryServiceClient(cc grpc.ClientConnInterface) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v3.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v31.DiscoveryRequest) error
	Recv() (*v31.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v31.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v31.DiscoveryResponse, error) {
	m := new(v31.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v3.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v31.DeltaDiscoveryRequest) error
	Recv() (*v31.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v31.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v31.DeltaDiscoveryResponse, error) {
	m := new(v31.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v31.DiscoveryRequest, opts ...grpc.CallOption) (*v31.DiscoveryResponse, error) {
	out := new(v31.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v3.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v31.DiscoveryRequest) (*v31.DiscoveryResponse, error)
}

// UnimplementedTapDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapDiscoveryServiceServer struct {
}

func (*UnimplementedTapDiscoveryServiceServer) StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) FetchTapConfigs(context.Context, *v31.DiscoveryRequest) (*v31.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTapConfigs not implemented")
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v31.DiscoveryResponse) error
	Recv() (*v31.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v31.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v31.DiscoveryRequest, error) {
	m := new(v31.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v31.DeltaDiscoveryResponse) error
	Recv() (*v31.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v31.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v31.DeltaDiscoveryRequest, error) {
	m := new(v31.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v31.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v3.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v31.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v3.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v3/tapds.proto",
}
