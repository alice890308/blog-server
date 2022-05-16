package authkit

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthService struct {
	JWTManager *JWTManager
}

func NewAuthService(jwtManager *JWTManager) *AuthService {
	return &AuthService{
		JWTManager: jwtManager,
	}
}

func (a *AuthService) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// check method if needing authentication
		fmt.Println("request path: " + info.FullMethod)

		if _, ok := publicAPI[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		// check token if is valid and get user_id
		err := a.authenticate(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (a *AuthService) authenticate(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return MetaDataNotProvided
	}

	values := md["authorization"]
	if len(values) == 0 {
		return TokenNotProvided
	}

	accessToken := values[0]
	fmt.Println(accessToken)
	payload, err := a.JWTManager.Verify(accessToken)
	if err != nil {
		return TokenInvalid
	}

	// append user id to context
	md.Append("user_id", payload.Id)

	return nil
}
