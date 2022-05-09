package service

import (
	"context"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	pb.UnimplementedPostServer

	postDAO dao.PostDAO
}

func NewService(postDAO dao.PostDAO) *service {
	return &service{
		postDAO: postDAO,
	}
}

func (s *service) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	postID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, ErrInvalidUUID
	}

	post, err := s.postDAO.Get(ctx, postID)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{Post: post.ToProto()}, nil
}

func (s *service) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	posts, err := s.postDAO.List(ctx, int64(req.GetLimit()), int64(req.GetSkip()))
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.PostInfo, 0, len(posts))
	for _, post := range posts {
		pbPosts = append(pbPosts, post.ToProto())
	}

	return &pb.ListPostResponse{Posts: pbPosts}, nil
}

func (s *service) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
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

	if err := s.postDAO.Create(ctx, post); err != nil {
		return nil, err
	}

}
