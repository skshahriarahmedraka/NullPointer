package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func MongodbConnection() *mongo.Database {

	mongoURI := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=" + os.Getenv("MONGO_MAXPOOLSIZE") + "&w=" + os.Getenv("MONGO_W")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Println("❌ failed to NewCient Client", err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Failed  to connect to mongodb db", err)
	}
	// defer client.Disconnect(ctx)
	Mydb := client.Database(os.Getenv("MONGO_DBNAME"))
	fmt.Println("✨🥰 ~ file: mongodb.go ~ line 32 ~ funcMongodbConnection ~ Mydb : ", Mydb)

	return Mydb

}
