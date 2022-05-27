package service

import (
	"context"
	"errors"
	"time"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.GetPostId())
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
	total, err := s.postDAO.TotalCount(ctx)
	if err != nil {
		return nil, err
	}

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

	return &pb.ListPostResponse{Posts: pbPosts, Total: total}, nil
}

func (s *Service) ListPostByUserID(ctx context.Context, req *pb.ListPostByUserIDRequest) (*pb.ListPostByUserIDResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	posts, err := s.postDAO.ListByUserID(ctx, userID, req.GetLimit(), req.GetSkip())
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
	userID, err := getUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	post := &dao.Post{
		UserID:    userID,
		Title:     req.GetTitle(),
		Content:   req.GetContent(),
		Image:     req.GetImage(),
		Views:     0,
		Likes:     0,
		Tags:      req.GetTags(),
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	postID, err := s.postDAO.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{PostId: postID.Hex()}, nil
}

func (s *Service) UpdatePostContent(ctx context.Context, req *pb.UpdatePostContentRequest) (*pb.UpdatePostContentResponse, error) {
	userID, err := getUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	postID, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	post := &dao.Post{
		ID:      postID,
		UserID:  userID,
		Title:   req.GetTitle(),
		Content: req.GetContent(),
		Image:   req.GetImage(),
		Tags:    req.GetTags(),
	}

	if err := s.postDAO.UpdateContent(ctx, post); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	return &pb.UpdatePostContentResponse{}, nil
}

func (s *Service) UpdatePostLikes(ctx context.Context, req *pb.UpdatePostLikesRequest) (*pb.UpdatePostLikesResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.postDAO.UpdateLikes(ctx, postID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostLikesResponse{}, nil
}

func (s *Service) UpdatePostViews(ctx context.Context, req *pb.UpdatePostViewsRequest) (*pb.UpdatePostViewsResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.postDAO.UpdateViews(ctx, postID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.UpdatePostViewsResponse{}, nil
}

func (s *Service) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	userID, err := getUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	postID, err := primitive.ObjectIDFromHex(req.GetPostId())
	if err != nil {
		return nil, ErrInvalidObjectID
	}

	if err := s.postDAO.Delete(ctx, postID, userID); err != nil {
		if errors.Is(err, dao.ErrPostNotFound) {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	return &pb.DeletePostResponse{}, nil
}
