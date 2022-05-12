package dao

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoUserDAO struct {
	collection *mongo.Collection
}

var _ UserDAO = (*mongoUserDAO)(nil)

func NewMongoUserDAO(collection *mongo.Collection) *mongoUserDAO {
	return &mongoUserDAO{
		collection: collection,
	}
}

func (dao *mongoUserDAO) Get(ctx context.Context, id primitive.ObjectID) (*User, error) {
	var user User
	if err := dao.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (dao *mongoUserDAO) List(ctx context.Context, limit, skip int64) ([]*User, error) {
	o := options.Find().SetLimit(limit).SetSkip(skip)

	cursor, err := dao.collection.Find(ctx, bson.M{}, o)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	users := make([]*User, 0)
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode((&user)); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (dao *mongoUserDAO) Create(ctx context.Context, user *User) error {
	result, err := dao.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (dao *mongoUserDAO) Update(ctx context.Context, user *User) error {
	if result, err := dao.collection.UpdateByID(
		ctx,
		user.ID,
		bson.M{
			"$set": bson.M{
				"name":        user.Name,
				"description": user.Description,
				"avator":      user.Avator,
			},
		},
	); err != nil {
		return err
	} else if result.ModifiedCount == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (dao *mongoUserDAO) Delete(ctx context.Context, id primitive.ObjectID) error {
	if result, err := dao.collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return err
	} else if result.DeletedCount == 0 {
		return ErrUserNotFound
	}

	return nil
}
