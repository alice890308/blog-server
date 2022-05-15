package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var publicAPI = map[string]bool{
	"aaa": true,
}

var (
	MetaDataNotProvided = status.Errorf(codes.Unauthenticated, "metadata is not provided")
	TokenNotProvided    = status.Errorf(codes.Unauthenticated, "token is not provided")
	TokenInvalid        = status.Errorf(codes.Unauthenticated, "token is invalid")
)
