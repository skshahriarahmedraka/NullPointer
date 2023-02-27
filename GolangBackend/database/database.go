package database

import (
	"app/handler"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	// "gorm.io/gorm"
)


func DatabaseInitialization(db1 *mongo.Database) handler.DatabaseCollections {
    fmt.Println("🚀 ~ file: database.go ~ line 15 ~ funcDatabaseInitialization ~ mongodb collection : ", db1)
	return handler.DatabaseCollections{Mongo:db1}
}

