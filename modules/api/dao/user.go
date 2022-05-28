package dao

import (
	"context"
	"errors"

	"github.com/alice890308/blog-server/modules/api/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Account     string             `bson:"account,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Password    string             `bson:"password,omitempty"`
	Description string             `bson:"description,omitempty"`
	Avator      string             `bson:"avator,omitempty"`
	IG          string             `bson:"ig,omitempty"`
	FB          string             `bson:"fb,omitempty"`
	TW          string             `bson:"tw,omitempty"`
	Email       string             `bson:"email,omitempty"`
}

type UserDAO interface {
	Get(ctx context.Context, id primitive.ObjectID) (*User, error)
	GetByUserAccount(ctx context.Context, account string) (*User, error)
	List(ctx context.Context, limit, skip int64) ([]*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

func (u *User) ToProto() *pb.UserInfo {
	return &pb.UserInfo{
		UserId:      u.ID.Hex(),
		UserName:    u.Name,
		UserAccount: u.Account,
		Description: u.Description,
		Avator:      u.Avator,
		Ig:          u.IG,
		Fb:          u.FB,
		Tw:          u.TW,
		Email:       u.Email,
	}
}

var (
	ErrUserNotFound = errors.New("user not found")
)
