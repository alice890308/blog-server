package dao

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoPostDAO struct {
	collection *mongo.Collection
}

var _ PostDAO = (*mongoPostDAO)(nil)

func NewMongoPostDAO(collection *mongo.Collection) *mongoPostDAO {
	return &mongoPostDAO{
		collection: collection,
	}
}

func (dao *mongoPostDAO) Get(ctx context.Context, id primitive.ObjectID) (*Post, error) {
	var post Post

	if err := dao.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post); err != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	return &post, nil
}

func (dao *mongoPostDAO) List(ctx context.Context, limit, skip int64) ([]*Post, error) {
	o := options.Find().SetLimit(limit).SetSkip(skip)

	cursor, err := dao.collection.Find(ctx, bson.M{}, o)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	posts := make([]*Post, 0)
	for cursor.Next(ctx) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func (dao *mongoPostDAO) ListByUserID(ctx context.Context, user_id primitive.ObjectID, limit, skip int64) ([]*Post, error) {
	o := options.Find().SetLimit(limit).SetSkip(skip)

	cursor, err := dao.collection.Find(ctx, bson.M{"user_id": user_id}, o)
	if err != nil {
		return nil, err
	}

	posts := make([]*Post, 0)
	for cursor.Next(ctx) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func (dao *mongoPostDAO) Create(ctx context.Context, post *Post) (primitive.ObjectID, error) {
	result, err := dao.collection.InsertOne(ctx, post)
	if err != nil {
		return primitive.NilObjectID, err
	}

	post.ID = result.InsertedID.(primitive.ObjectID)

	return post.ID, nil
}

func (dao *mongoPostDAO) UpdateContent(ctx context.Context, post *Post) error {
	if result, err := dao.collection.UpdateOne(
		ctx,
		bson.M{
			"_id":     post.ID,
			"user_id": post.UserID,
		},
		bson.M{
			"$set": bson.M{
				"title":   post.Title,
				"content": post.Content,
				"tags":    post.Tags,
			},
		},
	); err != nil {
		return err
	} else if result.ModifiedCount == 0 {
		return ErrPostNotFound
	}

	return nil
}

func (dao *mongoPostDAO) UpdateLikes(ctx context.Context, id primitive.ObjectID) error {
	if result, err := dao.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$inc": bson.M{
				"likes": 1,
			},
		},
	); err != nil {
		return err
	} else if result.ModifiedCount == 0 {
		return ErrPostNotFound
	}

	return nil
}

func (dao *mongoPostDAO) UpdateViews(ctx context.Context, id primitive.ObjectID) error {
	if result, err := dao.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$inc": bson.M{
				"views": 1,
			},
		},
	); err != nil {
		return err
	} else if result.ModifiedCount == 0 {
		return ErrPostNotFound
	}

	return nil
}

func (dao *mongoPostDAO) Delete(ctx context.Context, id primitive.ObjectID, user_id primitive.ObjectID) error {
	if result, err := dao.collection.DeleteOne(ctx, bson.M{
		"_id":     id,
		"user_id": user_id,
	}); err != nil {
		return err
	} else if result.DeletedCount == 0 {
		return ErrPostNotFound
	}

	return nil
}
