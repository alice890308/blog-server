package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidObjectID = status.Errorf(codes.InvalidArgument, "invalid ObjectID")
	ErrPostNotFound    = status.Errorf(codes.NotFound, "post not found")
)
