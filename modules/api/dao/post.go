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

func (p *Post) ToProto() *pb.PostInfo {
	return &pb.PostInfo{
		Id:        p.ID.String(),
		UserId:    p.UserID.String(),
		Title:     p.Title,
		Content:   p.Content,
		Views:     uint32(p.Views),
		Likes:     uint32(p.Likes),
		Tags:      p.Tags,
		Image:     p.Image,
		CreatedAt: timestamppb.New(p.CreatedAT),
		UpdatedAt: timestamppb.New(p.UpdatedAT),
		DeletedAt: timestamppb.New(p.DeletedAT),
	}
}

type PostDAO interface {
	Get(ctx context.Context, id primitive.ObjectID) (*Post, error)
	List(ctx context.Context, limit, skip int64) ([]*Post, error)
	Create(ctx context.Context, post *Post) (*Post, error)
	Update(ctx context.Context, post *Post) (*Post, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

var (
	ErrPostNotFound = errors.New("post not found")
)
