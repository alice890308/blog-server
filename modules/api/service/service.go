package service

import (
	"context"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"github.com/alice890308/blog-server/pkg/authkit"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	pb.UnimplementedPostServer
	pb.UnimplementedUserServer
	pb.UnimplementedSessionServer

	userDAO    dao.UserDAO
	postDAO    dao.PostDAO
	jwtManager authkit.JWT
}

func NewService(postDAO dao.PostDAO, userDAO dao.UserDAO, jwtManager authkit.JWT) *Service {
	return &Service{
		userDAO:    userDAO,
		postDAO:    postDAO,
		jwtManager: jwtManager,
	}
}

func getUserIdFromMetadata(ctx context.Context) (primitive.ObjectID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return primitive.NilObjectID, ErrMetadataNotProivided
	}

	userID, err := primitive.ObjectIDFromHex(md["user_id"][0])
	if err != nil {
		return primitive.NilObjectID, ErrInvalidObjectID
	}

	return userID, nil
}
