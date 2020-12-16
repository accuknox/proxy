// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/core/v3/address.proto

package envoy_config_core_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

// Enum value maps for SocketAddress_Protocol.
var (
	SocketAddress_Protocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
	}
	SocketAddress_Protocol_value = map[string]int32{
		"TCP": 0,
		"UDP": 1,
	}
)

func (x SocketAddress_Protocol) Enum() *SocketAddress_Protocol {
	p := new(SocketAddress_Protocol)
	*p = x
	return p
}

func (x SocketAddress_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketAddress_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_core_v3_address_proto_enumTypes[0].Descriptor()
}

func (SocketAddress_Protocol) Type() protoreflect.EnumType {
	return &file_envoy_config_core_v3_address_proto_enumTypes[0]
}

func (x SocketAddress_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketAddress_Protocol.Descriptor instead.
func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{1, 0}
}

type Pipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix Domain Socket path. On Linux, paths starting with '@' will use the
	// abstract namespace. The starting '@' is replaced by a null byte by Envoy.
	// Paths starting with '@' will result in an error in environments other than
	// Linux.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The mode for the Pipe. Not applicable for abstract sockets.
	Mode uint32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *Pipe) Reset() {
	*x = Pipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipe) ProtoMessage() {}

func (x *Pipe) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipe.ProtoReflect.Descriptor instead.
func (*Pipe) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{0}
}

func (x *Pipe) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Pipe) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

// [#next-free-field: 7]
type SocketAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.config.core.v3.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The address for this socket. :ref:`Listeners <config_listeners>` will bind
	// to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
	// to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
	// It is possible to distinguish a Listener address via the prefix/suffix matching
	// in :ref:`FilterChainMatch <envoy_api_msg_config.listener.v3.FilterChainMatch>`.] When used
	// within an upstream :ref:`BindConfig <envoy_api_msg_config.core.v3.BindConfig>`, the address
	// controls the source address of outbound connections. For :ref:`clusters
	// <envoy_api_msg_config.cluster.v3.Cluster>`, the cluster type determines whether the
	// address must be an IP (*STATIC* or *EDS* clusters) or a hostname resolved by DNS
	// (*STRICT_DNS* or *LOGICAL_DNS* clusters). Address resolution can be customized
	// via :ref:`resolver_name <envoy_api_field_config.core.v3.SocketAddress.resolver_name>`.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are assignable to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	// The name of the custom resolver. This must have been registered with Envoy. If
	// this is empty, a context dependent default applies. If the address is a concrete
	// IP address, no resolution will occur. If address is a hostname this
	// should be set for resolution other than DNS. Specifying a custom resolver with
	// *STRICT_DNS* or *LOGICAL_DNS* will generate an error at runtime.
	ResolverName string `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	// When binding to an IPv6 address above, this enables `IPv4 compatibility
	// <https://tools.ietf.org/html/rfc3493#page-11>`_. Binding to ``::`` will
	// allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
	// IPv6 space as ``::FFFF:<IPv4-address>``.
	Ipv4Compat bool `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
}

func (x *SocketAddress) Reset() {
	*x = SocketAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketAddress) ProtoMessage() {}

func (x *SocketAddress) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketAddress.ProtoReflect.Descriptor instead.
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{1}
}

func (x *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if x != nil {
		return x.Protocol
	}
	return SocketAddress_TCP
}

func (x *SocketAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (x *SocketAddress) GetPortValue() uint32 {
	if x, ok := x.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (x *SocketAddress) GetNamedPort() string {
	if x, ok := x.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (x *SocketAddress) GetResolverName() string {
	if x != nil {
		return x.ResolverName
	}
	return ""
}

func (x *SocketAddress) GetIpv4Compat() bool {
	if x != nil {
		return x.Ipv4Compat
	}
	return false
}

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	// This is only valid if :ref:`resolver_name
	// <envoy_api_field_config.core.v3.SocketAddress.resolver_name>` is specified below and the
	// named resolver is capable of named port resolution.
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}

func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

type TcpKeepalive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum number of keepalive probes to send without response before deciding
	// the connection is dead. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 9.)
	KeepaliveProbes *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	// The number of seconds a connection needs to be idle before keep-alive probes
	// start being sent. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 7200s (i.e., 2 hours.)
	KeepaliveTime *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	// The number of seconds between keep-alive probes. Default is to use the OS
	// level configuration (unless overridden, Linux defaults to 75s.)
	KeepaliveInterval *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
}

func (x *TcpKeepalive) Reset() {
	*x = TcpKeepalive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpKeepalive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpKeepalive) ProtoMessage() {}

func (x *TcpKeepalive) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpKeepalive.ProtoReflect.Descriptor instead.
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{2}
}

func (x *TcpKeepalive) GetKeepaliveProbes() *wrappers.UInt32Value {
	if x != nil {
		return x.KeepaliveProbes
	}
	return nil
}

func (x *TcpKeepalive) GetKeepaliveTime() *wrappers.UInt32Value {
	if x != nil {
		return x.KeepaliveTime
	}
	return nil
}

func (x *TcpKeepalive) GetKeepaliveInterval() *wrappers.UInt32Value {
	if x != nil {
		return x.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address to bind to when creating a socket.
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Whether to set the *IP_FREEBIND* option when creating the socket. When this
	// flag is set to true, allows the :ref:`source_address
	// <envoy_api_field_config.cluster.v3.UpstreamBindConfig.source_address>` to be an IP address
	// that is not configured on the system running Envoy. When this flag is set
	// to false, the option *IP_FREEBIND* is disabled on the socket. When this
	// flag is not set (default), the socket is not modified, i.e. the option is
	// neither enabled nor disabled.
	Freebind *wrappers.BoolValue `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions []*SocketOption `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
}

func (x *BindConfig) Reset() {
	*x = BindConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindConfig) ProtoMessage() {}

func (x *BindConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindConfig.ProtoReflect.Descriptor instead.
func (*BindConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{3}
}

func (x *BindConfig) GetSourceAddress() *SocketAddress {
	if x != nil {
		return x.SourceAddress
	}
	return nil
}

func (x *BindConfig) GetFreebind() *wrappers.BoolValue {
	if x != nil {
		return x.Freebind
	}
	return nil
}

func (x *BindConfig) GetSocketOptions() []*SocketOption {
	if x != nil {
		return x.SocketOptions
	}
	return nil
}

// Addresses specify either a logical or physical address and port, which are
// used to tell Envoy where to bind/listen, connect to upstream and find
// management servers.
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address isAddress_Address `protobuf_oneof:"address"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_address_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_address_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{4}
}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *Address) GetSocketAddress() *SocketAddress {
	if x, ok := x.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (x *Address) GetPipe() *Pipe {
	if x, ok := x.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}

func (*Address_Pipe) isAddress_Address() {}

// CidrRange specifies an IP Address and a prefix length to construct
// the subnet mask for a `CIDR <https://tools.ietf.org/html/rfc4632>`_ range.
type CidrRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IPv4 or IPv6 address, e.g. ``192.0.0.0`` or ``2001:db8::``.
	AddressPrefix string `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	// Length of prefix, e.g. 0, 32.
	PrefixLen *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
}

func (x *CidrRange) Reset() {
	*x = CidrRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_address_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CidrRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CidrRange) ProtoMessage() {}

func (x *CidrRange) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_address_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CidrRange.ProtoReflect.Descriptor instead.
func (*CidrRange) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_address_proto_rawDescGZIP(), []int{5}
}

func (x *CidrRange) GetAddressPrefix() string {
	if x != nil {
		return x.AddressPrefix
	}
	return ""
}

func (x *CidrRange) GetPrefixLen() *wrappers.UInt32Value {
	if x != nil {
		return x.PrefixLen
	}
	return nil
}

var File_envoy_config_core_v3_address_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_address_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x60, 0x0a, 0x04, 0x50, 0x69, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x2a, 0x03, 0x18, 0xff, 0x03, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x3a, 0x1d, 0x9a, 0xc5, 0x88, 0x1e, 0x18, 0x0a, 0x16, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x22, 0xf6, 0x02, 0x0a, 0x0d, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x52, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20,
	0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0xff, 0xff, 0x03, 0x48, 0x00, 0x52, 0x09, 0x70, 0x6f, 0x72,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x70, 0x76, 0x34, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x22, 0x1c, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x3a, 0x26, 0x9a, 0xc5, 0x88,
	0x1e, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x15, 0x0a, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x90, 0x02, 0x0a, 0x0c, 0x54,
	0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x6b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0e, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x6b, 0x65, 0x65, 0x70,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x3a, 0x25, 0x9a, 0xc5, 0x88, 0x1e, 0x20, 0x0a, 0x1e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x8a, 0x02,
	0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x0e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x66, 0x72, 0x65, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x49, 0x0a, 0x0e, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x23, 0x9a, 0xc5, 0x88, 0x1e, 0x1e, 0x0a, 0x1c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x0e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x69, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x69, 0x70, 0x65, 0x3a, 0x20, 0x9a, 0xc5, 0x88, 0x1e, 0x1b, 0x0a, 0x19, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xa6, 0x01, 0x0a, 0x09, 0x43, 0x69, 0x64,
	0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x2a, 0x03, 0x18,
	0x80, 0x01, 0x52, 0x09, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x3a, 0x22, 0x9a,
	0xc5, 0x88, 0x1e, 0x1d, 0x0a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x42, 0x3c, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_address_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_address_proto_rawDescData = file_envoy_config_core_v3_address_proto_rawDesc
)

func file_envoy_config_core_v3_address_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_address_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_address_proto_rawDescData)
	})
	return file_envoy_config_core_v3_address_proto_rawDescData
}

var file_envoy_config_core_v3_address_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_core_v3_address_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_config_core_v3_address_proto_goTypes = []interface{}{
	(SocketAddress_Protocol)(0),  // 0: envoy.config.core.v3.SocketAddress.Protocol
	(*Pipe)(nil),                 // 1: envoy.config.core.v3.Pipe
	(*SocketAddress)(nil),        // 2: envoy.config.core.v3.SocketAddress
	(*TcpKeepalive)(nil),         // 3: envoy.config.core.v3.TcpKeepalive
	(*BindConfig)(nil),           // 4: envoy.config.core.v3.BindConfig
	(*Address)(nil),              // 5: envoy.config.core.v3.Address
	(*CidrRange)(nil),            // 6: envoy.config.core.v3.CidrRange
	(*wrappers.UInt32Value)(nil), // 7: google.protobuf.UInt32Value
	(*wrappers.BoolValue)(nil),   // 8: google.protobuf.BoolValue
	(*SocketOption)(nil),         // 9: envoy.config.core.v3.SocketOption
}
var file_envoy_config_core_v3_address_proto_depIdxs = []int32{
	0,  // 0: envoy.config.core.v3.SocketAddress.protocol:type_name -> envoy.config.core.v3.SocketAddress.Protocol
	7,  // 1: envoy.config.core.v3.TcpKeepalive.keepalive_probes:type_name -> google.protobuf.UInt32Value
	7,  // 2: envoy.config.core.v3.TcpKeepalive.keepalive_time:type_name -> google.protobuf.UInt32Value
	7,  // 3: envoy.config.core.v3.TcpKeepalive.keepalive_interval:type_name -> google.protobuf.UInt32Value
	2,  // 4: envoy.config.core.v3.BindConfig.source_address:type_name -> envoy.config.core.v3.SocketAddress
	8,  // 5: envoy.config.core.v3.BindConfig.freebind:type_name -> google.protobuf.BoolValue
	9,  // 6: envoy.config.core.v3.BindConfig.socket_options:type_name -> envoy.config.core.v3.SocketOption
	2,  // 7: envoy.config.core.v3.Address.socket_address:type_name -> envoy.config.core.v3.SocketAddress
	1,  // 8: envoy.config.core.v3.Address.pipe:type_name -> envoy.config.core.v3.Pipe
	7,  // 9: envoy.config.core.v3.CidrRange.prefix_len:type_name -> google.protobuf.UInt32Value
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_address_proto_init() }
func file_envoy_config_core_v3_address_proto_init() {
	if File_envoy_config_core_v3_address_proto != nil {
		return
	}
	file_envoy_config_core_v3_socket_option_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipe); i {
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
		file_envoy_config_core_v3_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketAddress); i {
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
		file_envoy_config_core_v3_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpKeepalive); i {
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
		file_envoy_config_core_v3_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindConfig); i {
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
		file_envoy_config_core_v3_address_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_envoy_config_core_v3_address_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CidrRange); i {
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
	file_envoy_config_core_v3_address_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
	file_envoy_config_core_v3_address_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_address_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_address_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_address_proto_depIdxs,
		EnumInfos:         file_envoy_config_core_v3_address_proto_enumTypes,
		MessageInfos:      file_envoy_config_core_v3_address_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_address_proto = out.File
	file_envoy_config_core_v3_address_proto_rawDesc = nil
	file_envoy_config_core_v3_address_proto_goTypes = nil
	file_envoy_config_core_v3_address_proto_depIdxs = nil
}
