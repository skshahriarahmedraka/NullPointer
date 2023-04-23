package handler

import (
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections)QuesAsk(c *fiber.Ctx) error  {

	// doc := bson.M{
	// 	"timestamp": primitive.NewDateTimeFromTime(t),
	// }
	
	return c.Next()
}
