package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"app/logs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnection() *mongo.Database {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=%v&w=majority", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_IP"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_POOL_SIZE"))

	fmt.Println("mongoURI", mongoURI)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	logs.Error("MongoDB New Connection", err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	logs.Error("MongoDB Connection", err)

	err = client.Ping(ctx, nil)

	logs.Error("MongoDB Ping", err)

	DB := client.Database(os.Getenv("MONGO_DB"))

	logs.Error("MongoDB create/connect Database", err)

	if err == nil {
		fmt.Println("😁🤩 MongoDB Connected successfully")
	}

	return DB

}
