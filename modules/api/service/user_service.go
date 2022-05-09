package service

import (
	"context"
	"errors"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	pb.UnimplementedUserServer

	UserDAO dao.UserDAO
}

func NewService(userDAO dao.UserDAO) *service {
	return &service{
		UserDAO: userDAO,
	}
}

func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPWD, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrToHashPWD
	}
	user := &dao.User{
		Name:     req.GetUserName(),
		Account:  req.GetUserAccount(),
		Password: string(hashedPWD),
	}

	err = s.UserDAO.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{}, nil
}

func (s *service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	user, err := s.UserDAO.Get(ctx, id)
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &pb.GetUserResponse{User: user.ToProto()}, nil
}

func (s *service) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := s.UserDAO.List(ctx, req.GetLimit(), req.GetSkip())
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*pb.UserInfo, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, user.ToProto())
	}

	return &pb.ListUserResponse{Users: pbUsers}, nil
}

func (s *service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return &pb.UpdateUserResponse{}, ErrInvalidObjectID
	}

	user := &dao.User{
		ID:          id,
		Name:        req.GetUserName(),
		Description: req.GetDescription(),
		Avator:      req.GetAvator(),
	}

	err = s.UserDAO.Update(ctx, user)
	if err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return &pb.UpdateUserResponse{}, ErrUserNotFound
		}

		return &pb.UpdateUserResponse{}, err
	}

	return &pb.UpdateUserResponse{}, nil
}

func (s *service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.UserDAO.Delete(ctx, id); err != nil {
		if errors.Is(err, dao.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &pb.DeleteUserResponse{}, nil
}
