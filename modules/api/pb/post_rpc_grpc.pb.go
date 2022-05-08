// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: modules/api/pb/post_rpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostClient interface {
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostResponse, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	UpdatePostContent(ctx context.Context, in *UpdatePostContentRequest, opts ...grpc.CallOption) (*UpdatePostContentResponse, error)
	UpdatePostLikes(ctx context.Context, in *UpdatePostLikesRequest, opts ...grpc.CallOption) (*UpdatePostLikesResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, "/pb.Post/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostResponse, error) {
	out := new(ListPostResponse)
	err := c.cc.Invoke(ctx, "/pb.Post/ListPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/pb.Post/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) UpdatePostContent(ctx context.Context, in *UpdatePostContentRequest, opts ...grpc.CallOption) (*UpdatePostContentResponse, error) {
	out := new(UpdatePostContentResponse)
	err := c.cc.Invoke(ctx, "/pb.Post/UpdatePostContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) UpdatePostLikes(ctx context.Context, in *UpdatePostLikesRequest, opts ...grpc.CallOption) (*UpdatePostLikesResponse, error) {
	out := new(UpdatePostLikesResponse)
	err := c.cc.Invoke(ctx, "/pb.Post/UpdatePostLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/pb.Post/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
// All implementations must embed UnimplementedPostServer
// for forward compatibility
type PostServer interface {
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	ListPost(context.Context, *ListPostRequest) (*ListPostResponse, error)
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	UpdatePostContent(context.Context, *UpdatePostContentRequest) (*UpdatePostContentResponse, error)
	UpdatePostLikes(context.Context, *UpdatePostLikesRequest) (*UpdatePostLikesResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	mustEmbedUnimplementedPostServer()
}

// UnimplementedPostServer must be embedded to have forward compatible implementations.
type UnimplementedPostServer struct {
}

func (UnimplementedPostServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServer) ListPost(context.Context, *ListPostRequest) (*ListPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (UnimplementedPostServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServer) UpdatePostContent(context.Context, *UpdatePostContentRequest) (*UpdatePostContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostContent not implemented")
}
func (UnimplementedPostServer) UpdatePostLikes(context.Context, *UpdatePostLikesRequest) (*UpdatePostLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostLikes not implemented")
}
func (UnimplementedPostServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServer) mustEmbedUnimplementedPostServer() {}

// UnsafePostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServer will
// result in compilation errors.
type UnsafePostServer interface {
	mustEmbedUnimplementedPostServer()
}

func RegisterPostServer(s grpc.ServiceRegistrar, srv PostServer) {
	s.RegisterService(&Post_ServiceDesc, srv)
}

func _Post_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Post/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Post/ListPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).ListPost(ctx, req.(*ListPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Post/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_UpdatePostContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).UpdatePostContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Post/UpdatePostContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).UpdatePostContent(ctx, req.(*UpdatePostContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_UpdatePostLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).UpdatePostLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Post/UpdatePostLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).UpdatePostLikes(ctx, req.(*UpdatePostLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Post/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Post_ServiceDesc is the grpc.ServiceDesc for Post service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Post_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPost",
			Handler:    _Post_GetPost_Handler,
		},
		{
			MethodName: "ListPost",
			Handler:    _Post_ListPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Post_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePostContent",
			Handler:    _Post_UpdatePostContent_Handler,
		},
		{
			MethodName: "UpdatePostLikes",
			Handler:    _Post_UpdatePostLikes_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Post_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/api/pb/post_rpc.proto",
}