package dao

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Account  string             `bson:"acount,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password,omitempty"`
	Alice    string
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

func getUserKey(id primitive.ObjectID) string {
	return "getUser:" + id.Hex()
}

func listUserKey(limit, skip int64) string {
	return fmt.Sprintf("listUser:%d:%d", limit, skip)
}

func NewFakeUser() *User {
	id := primitive.NewObjectID()

	return &User{
		ID:       id,
		Account:  "fake account",
		Name:     "fake user name",
		Password: "fake password",
	}
}
