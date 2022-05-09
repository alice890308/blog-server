package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidUUID  = status.Errorf(codes.InvalidArgument, "invalid UUID")
	ErrPostNotFound = status.Errorf(codes.NotFound, "post not found")
)
