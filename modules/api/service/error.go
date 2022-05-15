package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidObjectID      = status.Errorf(codes.InvalidArgument, "invalid ObjectID")
	ErrPostNotFound         = status.Errorf(codes.NotFound, "post not found")
	ErrUserNotFound         = status.Errorf(codes.NotFound, "user not found")
	ErrToHashPWD            = status.Errorf(codes.InvalidArgument, "hash password failed")
	ErrWrongPWD             = status.Errorf(codes.PermissionDenied, "wrong password")
	ErrMetadataNotProivided = status.Errorf(codes.InvalidArgument, "meatadata not provided")
	ErrWrongPWD             = status.Errorf(codes.PermissionDenied, "wrong password")
	ErrUserAlreadyExists    = status.Errorf(codes.AlreadyExists, "user account already exists")
)
