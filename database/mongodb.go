package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to mongodb
func Connect() (*mongo.Database, error) {
	user := os.Getenv("USERNAME")
	psw := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	// host:=os.Getenv("")
	uri := fmt.Sprintf("mongodb://%s:%s@mongodb:27017", user, psw)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(database)
	return db, nil
}
