package service

import (
	"context"
	"errors"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, ErrInvalidObjectID
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

func (s *Service) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostResponse, error) {
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

func (s *Service) ListPostByUserID(ctx context.Context, req *pb.ListPostByUserIDRequest) (*pb.ListPostByUserIDResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	posts, err := s.postDAO.ListByUserID(ctx, userID, req.Limit, req.Skip)
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.PostInfo, 0, len(posts))

	var userName string
	if len(posts) > 0 {
		user, err := s.userDAO.Get(ctx, posts[0].UserID)
		if err != nil {
			return nil, err
		}

		userName = user.Name
	}

	for _, post := range posts {
		pbPosts = append(pbPosts, post.ToProto(userName))
	}

	return &pb.ListPostByUserIDResponse{Posts: pbPosts}, nil
}

func (s *Service) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, ErrInvalidObjectID
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

func (s *Service) UpdateContent(ctx context.Context, req *pb.UpdatePostContentRequest) (*pb.UpdatePostContentResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidObjectID
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

func (s *Service) UpdateLikes(ctx context.Context, req *pb.UpdatePostLikesRequest) (*pb.UpdatePostLikesResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.postDAO.UpdateLikes(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostLikesResponse{}, nil
}

func (s *Service) UpdateViews(ctx context.Context, req *pb.UpdatePostViewsRequest) (*pb.UpdatePostViewsResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.postDAO.UpdateViews(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostViewsResponse{}, nil
}

func (s *Service) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.postDAO.Delete(ctx, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.DeletePostResponse{}, nil
}
