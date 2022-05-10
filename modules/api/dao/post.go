package dao

import (
	"context"
	"errors"
	"time"

	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Views     int                `bson:"views,omitempty"`
	Likes     int                `bson:"likes,omitempty"`
	Tags      []string           `bson:"tags,omitempty"`
	Image     string             `bson:"image,omitempty"`
	CreatedAT time.Time          `bson:"created_at,omitempty"`
	UpdatedAT time.Time          `bson:"updated_at,omitempty"`
	DeletedAT time.Time          `bson:"deleted_at,omitempty"`
}

func (p *Post) ToProto(userName string) *pb.PostInfo {
	return &pb.PostInfo{
		Id:        p.ID.Hex(),
		UserId:    p.UserID.Hex(),
		UserName:  userName,
		Title:     p.Title,
		Content:   p.Content,
		Views:     uint32(p.Views),
		Likes:     uint32(p.Likes),
		Tags:      p.Tags,
		Image:     p.Image,
		CreatedAt: timestamppb.New(p.CreatedAT),
		UpdatedAt: timestamppb.New(p.UpdatedAT),
	}
}

type PostDAO interface {
	Get(ctx context.Context, id primitive.ObjectID) (*Post, error)
	List(ctx context.Context, limit, skip int64) ([]*Post, error)
	ListByUserID(ctx context.Context, user_id primitive.ObjectID, limit, skip int64) ([]*Post, error)
	Create(ctx context.Context, post *Post) (primitive.ObjectID, error)
	UpdateContent(ctx context.Context, post *Post) error
	UpdateLikes(ctx context.Context, id primitive.ObjectID) error
	UpdateViews(ctx context.Context, id primitive.ObjectID) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

var (
	ErrPostNotFound = errors.New("post not found")
)
