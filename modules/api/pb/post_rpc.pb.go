// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.3
// source: modules/api/proto/post_rpc.proto

package pb

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

var File_modules_api_proto_post_rpc_proto protoreflect.FileDescriptor

var file_modules_api_proto_post_rpc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x04, 0x0a, 0x04, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12,
	0x0b, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0b, 0x12, 0x06, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x62, 0x01, 0x2a, 0x12, 0x51, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x22, 0x06, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x62, 0x01, 0x2a, 0x12,
	0x73, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x1a, 0x13, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x62, 0x01, 0x2a, 0x12, 0x6b, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x11, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x62, 0x01,
	0x2a, 0x12, 0x53, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0b, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x62, 0x01, 0x2a, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_modules_api_proto_post_rpc_proto_goTypes = []interface{}{
	(*GetPostRequest)(nil),            // 0: pb.GetPostRequest
	(*ListPostRequest)(nil),           // 1: pb.ListPostRequest
	(*CreatePostRequest)(nil),         // 2: pb.CreatePostRequest
	(*UpdatePostContentRequest)(nil),  // 3: pb.UpdatePostContentRequest
	(*UpdatePostLikesRequest)(nil),    // 4: pb.UpdatePostLikesRequest
	(*DeletePostRequest)(nil),         // 5: pb.DeletePostRequest
	(*GetPostResponse)(nil),           // 6: pb.GetPostResponse
	(*ListPostResponse)(nil),          // 7: pb.ListPostResponse
	(*CreatePostResponse)(nil),        // 8: pb.CreatePostResponse
	(*UpdatePostContentResponse)(nil), // 9: pb.UpdatePostContentResponse
	(*UpdatePostLikesResponse)(nil),   // 10: pb.UpdatePostLikesResponse
	(*DeletePostResponse)(nil),        // 11: pb.DeletePostResponse
}
var file_modules_api_proto_post_rpc_proto_depIdxs = []int32{
	0,  // 0: pb.Post.GetPost:input_type -> pb.GetPostRequest
	1,  // 1: pb.Post.ListPost:input_type -> pb.ListPostRequest
	2,  // 2: pb.Post.CreatePost:input_type -> pb.CreatePostRequest
	3,  // 3: pb.Post.UpdatePostContent:input_type -> pb.UpdatePostContentRequest
	4,  // 4: pb.Post.UpdatePostLikes:input_type -> pb.UpdatePostLikesRequest
	5,  // 5: pb.Post.DeletePost:input_type -> pb.DeletePostRequest
	6,  // 6: pb.Post.GetPost:output_type -> pb.GetPostResponse
	7,  // 7: pb.Post.ListPost:output_type -> pb.ListPostResponse
	8,  // 8: pb.Post.CreatePost:output_type -> pb.CreatePostResponse
	9,  // 9: pb.Post.UpdatePostContent:output_type -> pb.UpdatePostContentResponse
	10, // 10: pb.Post.UpdatePostLikes:output_type -> pb.UpdatePostLikesResponse
	11, // 11: pb.Post.DeletePost:output_type -> pb.DeletePostResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_modules_api_proto_post_rpc_proto_init() }
func file_modules_api_proto_post_rpc_proto_init() {
	if File_modules_api_proto_post_rpc_proto != nil {
		return
	}
	file_modules_api_proto_post_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_api_proto_post_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_api_proto_post_rpc_proto_goTypes,
		DependencyIndexes: file_modules_api_proto_post_rpc_proto_depIdxs,
	}.Build()
	File_modules_api_proto_post_rpc_proto = out.File
	file_modules_api_proto_post_rpc_proto_rawDesc = nil
	file_modules_api_proto_post_rpc_proto_goTypes = nil
	file_modules_api_proto_post_rpc_proto_depIdxs = nil
}
