package service

import (
	"context"
	"errors"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userDAO.GetByUserAccount(ctx, req.GetUserAccount())
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetUserPassword()))
	if err != nil {
		return nil, ErrWrongPWD
	}

	userID := user.ID.Hex()
	token, err := s.jwtManager.Generate(userID)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token, UserId: userID}, nil
}
