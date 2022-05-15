package service

import (
	"context"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	pb.UnimplementedPostServer
	pb.UnimplementedUserServer

	userDAO dao.UserDAO
	postDAO dao.PostDAO
}

func NewService(postDAO dao.PostDAO, userDAO dao.UserDAO) *Service {
	return &Service{
		userDAO: userDAO,
		postDAO: postDAO,
	}
}

func getUserId(ctx context.Context) (primitive.ObjectID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return primitive.NilObjectID, ErrMetadataNotProivided
	}

	userID, err := primitive.ObjectIDFromHex(md["user_id"][0])
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userID, nil
}
