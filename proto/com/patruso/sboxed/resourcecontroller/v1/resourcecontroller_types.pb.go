// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: proto/com/patruso/sboxed/resourcecontroller/v1/resourcecontroller_types.proto

package v1

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

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Route) Reset() {
	*x = Route{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{0}
}

type RouteGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RouteGroup) Reset() {
	*x = RouteGroup{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteGroup) ProtoMessage() {}

func (x *RouteGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteGroup.ProtoReflect.Descriptor instead.
func (*RouteGroup) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{1}
}

type Sandbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sandbox) Reset() {
	*x = Sandbox{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sandbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sandbox) ProtoMessage() {}

func (x *Sandbox) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sandbox.ProtoReflect.Descriptor instead.
func (*Sandbox) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{2}
}

type ListRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoutesRequest) Reset() {
	*x = ListRoutesRequest{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesRequest) ProtoMessage() {}

func (x *ListRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{3}
}

type ListRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *ListRoutesResponse) Reset() {
	*x = ListRoutesResponse{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesResponse) ProtoMessage() {}

func (x *ListRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{4}
}

func (x *ListRoutesResponse) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type ListRouteGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRouteGroupsRequest) Reset() {
	*x = ListRouteGroupsRequest{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRouteGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRouteGroupsRequest) ProtoMessage() {}

func (x *ListRouteGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRouteGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListRouteGroupsRequest) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{5}
}

type ListRouteGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteGroups []*RouteGroup `protobuf:"bytes,1,rep,name=route_groups,json=routeGroups,proto3" json:"route_groups,omitempty"`
}

func (x *ListRouteGroupsResponse) Reset() {
	*x = ListRouteGroupsResponse{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRouteGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRouteGroupsResponse) ProtoMessage() {}

func (x *ListRouteGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRouteGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListRouteGroupsResponse) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{6}
}

func (x *ListRouteGroupsResponse) GetRouteGroups() []*RouteGroup {
	if x != nil {
		return x.RouteGroups
	}
	return nil
}

type ListSandboxesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSandboxesRequest) Reset() {
	*x = ListSandboxesRequest{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSandboxesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSandboxesRequest) ProtoMessage() {}

func (x *ListSandboxesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSandboxesRequest.ProtoReflect.Descriptor instead.
func (*ListSandboxesRequest) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{7}
}

type ListSandboxesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sandboxes []*Sandbox `protobuf:"bytes,1,rep,name=sandboxes,proto3" json:"sandboxes,omitempty"`
}

func (x *ListSandboxesResponse) Reset() {
	*x = ListSandboxesResponse{}
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSandboxesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSandboxesResponse) ProtoMessage() {}

func (x *ListSandboxesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSandboxesResponse.ProtoReflect.Descriptor instead.
func (*ListSandboxesResponse) Descriptor() ([]byte, []int) {
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP(), []int{8}
}

func (x *ListSandboxesResponse) GetSandboxes() []*Sandbox {
	if x != nil {
		return x.Sandboxes
	}
	return nil
}

var File_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto protoreflect.FileDescriptor

var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x74, 0x72,
	0x75, 0x73, 0x6f, 0x2f, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75,
	0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22,
	0x07, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x09, 0x0a, 0x07, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f,
	0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5d, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22,
	0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x09, 0x73, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0x7e, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x28, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescOnce sync.Once
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescData = file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDesc
)

func file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescGZIP() []byte {
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescOnce.Do(func() {
		file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescData)
	})
	return file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDescData
}

var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_goTypes = []any{
	(*Route)(nil),                   // 0: proto.com.patruso.sboxed.resourcecontroller.v1.Route
	(*RouteGroup)(nil),              // 1: proto.com.patruso.sboxed.resourcecontroller.v1.RouteGroup
	(*Sandbox)(nil),                 // 2: proto.com.patruso.sboxed.resourcecontroller.v1.Sandbox
	(*ListRoutesRequest)(nil),       // 3: proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesRequest
	(*ListRoutesResponse)(nil),      // 4: proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesResponse
	(*ListRouteGroupsRequest)(nil),  // 5: proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsRequest
	(*ListRouteGroupsResponse)(nil), // 6: proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsResponse
	(*ListSandboxesRequest)(nil),    // 7: proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesRequest
	(*ListSandboxesResponse)(nil),   // 8: proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesResponse
}
var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_depIdxs = []int32{
	0, // 0: proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesResponse.routes:type_name -> proto.com.patruso.sboxed.resourcecontroller.v1.Route
	1, // 1: proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsResponse.route_groups:type_name -> proto.com.patruso.sboxed.resourcecontroller.v1.RouteGroup
	2, // 2: proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesResponse.sandboxes:type_name -> proto.com.patruso.sboxed.resourcecontroller.v1.Sandbox
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_init()
}
func file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_init() {
	if File_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_goTypes,
		DependencyIndexes: file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_depIdxs,
		MessageInfos:      file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_msgTypes,
	}.Build()
	File_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto = out.File
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_rawDesc = nil
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_goTypes = nil
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_depIdxs = nil
}