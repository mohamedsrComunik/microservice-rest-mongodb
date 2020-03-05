package repository

import (
	"context"

	"github.com/communik/user-srv/helper"
	"github.com/communik/user-srv/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	db *mongo.Database
}

//RepositoryInterface interface
type RepositoryInterface interface {
	Create(ctx context.Context, user model.User) (string, error)
}

//NewUserRepo to create new instance of userRepo with *mongo.Database
func NewUserRepo(db *mongo.Database) userRepo {
	return userRepo{db: db}
}

func (u userRepo) Create(ctx context.Context, user model.User) (string, error) {
	res, err := u.db.Collection("user").InsertOne(ctx, bson.D{
		{"userName", user.UserName},
		{"createdAt", primitive.DateTime(helper.TimeToMillis(user.CreatedAt))},
		{"modifiedAt", primitive.DateTime(helper.TimeToMillis(user.ModifiedAt))},
	})
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
