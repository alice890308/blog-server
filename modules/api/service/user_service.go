package service

import (
	"context"
	"errors"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// check if the user account already exists
	_, err := s.userDAO.GetByUserAccount(ctx, req.GetUserAccount())
	if err != ErrUserNotFound {
		return nil, ErrUserAlreadyExists
	}

	hashedPWD, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrToHashPWD
	}
	user := &dao.User{
		Name:     req.GetUserName(),
		Account:  req.GetUserAccount(),
		Password: string(hashedPWD),
	}

	err = s.userDAO.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{}, nil
}

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	user, err := s.userDAO.Get(ctx, id)
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &pb.GetUserResponse{User: user.ToProto()}, nil
}

func (s *Service) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := s.userDAO.List(ctx, req.GetLimit(), req.GetSkip())
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*pb.UserInfo, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, user.ToProto())
	}

	return &pb.ListUserResponse{Users: pbUsers}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	userID, err := getUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	user := &dao.User{
		ID:          userID,
		Name:        req.GetUserName(),
		Description: req.GetDescription(),
		Avator:      req.GetAvator(),
	}

	err = s.userDAO.Update(ctx, user)
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &pb.UpdateUserResponse{}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	userID, err := getUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.userDAO.Delete(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &pb.DeleteUserResponse{}, nil
}
