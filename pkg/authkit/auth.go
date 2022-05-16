package authkit

import (
	"context"
	"log"

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
		log.Printf("{Request: %s}", info.FullMethod)

		if _, ok := publicAPI[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		// check token if is valid and get user_id
		newCtx, err := a.authenticate(ctx)
		if err != nil {
			return nil, err
		}

		return handler(newCtx, req)
	}
}

func (a *AuthService) authenticate(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, MetaDataNotProvided
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, TokenNotProvided
	}

	// get access token
	accessToken := string(values[0][7:])
	// verify token and get userID
	payload, err := a.JWTManager.Verify(accessToken)
	if err != nil {
		return ctx, TokenInvalid
	}

	md.Append("user_id", payload.UserID)
	newCtx := metadata.NewIncomingContext(ctx, md)

	return newCtx, nil
}
