package service

import (
	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"github.com/alice890308/blog-server/pkg/authkit"
)

type Service struct {
	pb.UnimplementedPostServer
	pb.UnimplementedUserServer

	userDAO    dao.UserDAO
	postDAO    dao.PostDAO
	jwtManager authkit.JWTManager
}

func NewService(postDAO dao.PostDAO, userDAO dao.UserDAO, jwtManager authkit.JWTManager) *Service {
	return &Service{
		userDAO:    userDAO,
		postDAO:    postDAO,
		jwtManager: jwtManager,
	}
}
