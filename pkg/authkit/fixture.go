package authkit

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var publicAPI = map[string]bool{
	"/pb.Post/ListPost":         true,
	"/pb.Post/GetPost":          true,
	"/pb.Post/ListPostByUserID": true,
	"/pb.Post/UpdatePostViews":  true,
	"/pb.User/ListUser":         true,
	"/pb.User/CreateUser":       true,
	"/pb.User/GetUser":          true,
	"/pb.Session/Login":         true,
	"/pb.Session/Health":        true,
}

var (
	MetaDataNotProvided = status.Errorf(codes.Unauthenticated, "metadata is not provided")
	TokenNotProvided    = status.Errorf(codes.Unauthenticated, "token is not provided")
	TokenInvalid        = status.Errorf(codes.Unauthenticated, "token is invalid")
)
