package service

import (
	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
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
