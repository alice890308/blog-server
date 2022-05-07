package dao

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type PostDAO interface {
	Get(ctx context.Context, id primitive.ObjectID) (*Post, error)
	List(ctx context.Context, limit, skip int64) ([]*Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

var (
	ErrPostNotFound = errors.New("post not found")
)
