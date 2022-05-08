package dao

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Account  string             `bson:"acount,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type UserDAO interface {
	Get(ctx context.Context, id primitive.ObjectID) (*User, error)
	List(ctx context.Context, limit, skip int64) ([]*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

var (
	ErrUserNotFound = errors.New("user not found")
)
