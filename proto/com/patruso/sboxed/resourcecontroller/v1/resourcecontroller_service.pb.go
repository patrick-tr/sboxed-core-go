// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: proto/com/patruso/sboxed/resourcecontroller/v1/resourcecontroller_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto protoreflect.FileDescriptor

var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x74, 0x72,
	0x75, 0x73, 0x6f, 0x2f, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74,
	0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x4d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x74,
	0x72, 0x75, 0x73, 0x6f, 0x2f, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xef, 0x04, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xc4, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x65, 0x73, 0x12, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73,
	0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0xb9, 0x01, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f,
	0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e,
	0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0xce, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62,
	0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61,
	0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x7e, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61,
	0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x62, 0x6f, 0x78, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x28, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x61, 0x74, 0x72, 0x75, 0x73, 0x6f, 0x2e, 0x73, 0x62, 0x6f, 0x78, 0x65, 0x64,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_goTypes = []any{
	(*ListSandboxesRequest)(nil),    // 0: proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesRequest
	(*ListRoutesRequest)(nil),       // 1: proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesRequest
	(*ListRouteGroupsRequest)(nil),  // 2: proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsRequest
	(*ListSandboxesResponse)(nil),   // 3: proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesResponse
	(*ListRoutesResponse)(nil),      // 4: proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesResponse
	(*ListRouteGroupsResponse)(nil), // 5: proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsResponse
}
var file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_depIdxs = []int32{
	0, // 0: proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService.ListSandboxes:input_type -> proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesRequest
	1, // 1: proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService.ListRoutes:input_type -> proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesRequest
	2, // 2: proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService.ListRouteGroups:input_type -> proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsRequest
	3, // 3: proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService.ListSandboxes:output_type -> proto.com.patruso.sboxed.resourcecontroller.v1.ListSandboxesResponse
	4, // 4: proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService.ListRoutes:output_type -> proto.com.patruso.sboxed.resourcecontroller.v1.ListRoutesResponse
	5, // 5: proto.com.patruso.sboxed.resourcecontroller.v1.ResourceControllerService.ListRouteGroups:output_type -> proto.com.patruso.sboxed.resourcecontroller.v1.ListRouteGroupsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_init()
}
func file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_init() {
	if File_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto != nil {
		return
	}
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_goTypes,
		DependencyIndexes: file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_depIdxs,
	}.Build()
	File_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto = out.File
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_rawDesc = nil
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_goTypes = nil
	file_proto_com_patruso_sboxed_resourcecontroller_v1_resourcecontroller_service_proto_depIdxs = nil
}