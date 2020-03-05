package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to mongodb
func Connect() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:user@mongodb:27017/comunik"))
	if err != nil {
		return nil, err
	}
	db := client.Database("comunik")
	return db, nil
}
