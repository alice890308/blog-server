package service

import (
	"context"
	"errors"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user_service struct {
	pb.UnimplementedPostServer

	userDAO dao.UserDAO
	postDAO dao.PostDAO
}

func NewService(postDAO dao.PostDAO) *user_service {
	return &user_service{
		postDAO: postDAO,
	}
}

func (s *user_service) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, ErrInvalidUUID
	}

	post, err := s.postDAO.Get(ctx, postID)
	if err != nil {
		return nil, err
	}

	user, err := s.userDAO.Get(ctx, post.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{Post: post.ToProto(user.Name)}, nil
}

func (s *user_service) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	posts, err := s.postDAO.List(ctx, int64(req.GetLimit()), int64(req.GetSkip()))
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.PostInfo, 0, len(posts))
	for _, post := range posts {
		user, err := s.userDAO.Get(ctx, post.UserID)
		if err != nil {
			return nil, err
		}

		pbPosts = append(pbPosts, post.ToProto(user.Name))
	}

	return &pb.ListPostResponse{Posts: pbPosts}, nil
}

func (s *user_service) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	post := &dao.Post{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		Views:   0,
		Likes:   0,
		Tags:    req.Tags,
	}

	result, err := s.postDAO.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{Id: result.Hex()}, nil
}

func (s *user_service) UpdateContent(ctx context.Context, req *pb.UpdatePostContentRequest) (*pb.UpdatePostContentResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	post := &dao.Post{
		ID:      userID,
		Title:   req.Title,
		Content: req.Content,
		Tags:    req.Tags,
	}

	if err := s.postDAO.UpdateContent(ctx, post); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostContentResponse{}, nil
}

func (s *user_service) UpdateLikes(ctx context.Context, req *pb.UpdatePostLikesRequest) (*pb.UpdatePostLikesResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	if err := s.postDAO.UpdateLikes(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostLikesResponse{}, nil
}

func (s *user_service) UpdateViews(ctx context.Context, req *pb.UpdatePostViewsRequest) (*pb.UpdatePostViewsResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	if err := s.postDAO.UpdateViews(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostViewsResponse{}, nil
}

func (s *user_service) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	if err := s.postDAO.Delete(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.DeletePostResponse{}, nil
}
