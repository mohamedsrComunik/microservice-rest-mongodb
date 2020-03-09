package repository

import (
	"context"
	"sync"

	"github.com/communik/user-srv/helper"
	"github.com/communik/user-srv/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	db *mongo.Database
	mu sync.RWMutex
}

//RepositoryInterface interface
type RepositoryInterface interface {
	Create(ctx context.Context, user model.User) (string, error)
	GetAll(ctx context.Context) []model.User
}

//NewUserRepo to create new instance of userRepo with *mongo.Database
func NewUserRepo(db *mongo.Database) userRepo {
	return userRepo{db: db}
}

func (u userRepo) Create(ctx context.Context, user model.User) (string, error) {
	u.mu.Lock()
	res, err := u.db.Collection("user").InsertOne(ctx, bson.D{
		{"userName", user.UserName},
		{"createdAt", primitive.DateTime(helper.TimeToMillis(user.CreatedAt))},
		{"modifiedAt", primitive.DateTime(helper.TimeToMillis(user.ModifiedAt))},
	})
	u.mu.Unlock()
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u userRepo) GetAll(ctx context.Context) []model.User {
	c, err := u.db.Collection("user").Find(ctx, bson.D{})
	defer c.Close(ctx)
	var users []model.User
	for c.Next(ctx) {
		elem := &bson.D{}
		if err = c.Decode(elem); err != nil {
			return users
		}
		m := elem.Map()
		t := model.User{
			ID:         m["_id"].(primitive.ObjectID).Hex(),
			CreatedAt:  helper.DataTimeToTime(m["createdAt"].(primitive.DateTime)),
			ModifiedAt: helper.DataTimeToTime(m["modifiedAt"].(primitive.DateTime)),
			UserName:   m["userName"].(string),
		}
		users = append(users, t)
	}
	return users
}
