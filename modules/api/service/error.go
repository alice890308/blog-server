package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidObjectID = status.Errorf(codes.InvalidArgument, "invalid ObjectID")
	ErrPostNotFound    = status.Errorf(codes.NotFound, "post not found")
	ErrUserNotFound    = status.Errorf(codes.NotFound, "user not found")
	ErrToHashPWD       = status.Errorf(codes.InvalidArgument, "hash password failed")
)
