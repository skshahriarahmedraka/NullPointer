package handler

import (
	"app/logs"
	"app/model"
	"context"
	"strconv"
	// "strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func (H *DatabaseCollections)BlogsList(c *fiber.Ctx) error  {
	ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if c.Query("num") != "" {
		var Blogs []model.BlogData
		findOpts := 	options.Find().SetLimit(20).SetSort(bson.D{ {"WrittenTime",1 }})
		cur,err := H.MongoBlogCol.Find(ctx, bson.D{},findOpts)
		logs.Error("Error while finding blogs", err)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		for cur.Next(ctx) {
			var blog model.BlogData
			err := cur.Decode(&blog)
			logs.Error("Error while decoding blogs", err)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
			Blogs = append(Blogs, blog)
		}
		return c.Status(fiber.StatusOK).JSON(Blogs)
	}else{
		numOfBlogs,err := strconv.ParseInt(c.Query("num"),10,64)
		var Blogs []model.BlogData
		findOpts := 	options.Find().SetLimit(numOfBlogs).SetSort(bson.D{ {"WrittenTime",1 }})
		cur,err := H.MongoBlogCol.Find(ctx, bson.D{},findOpts)
		logs.Error("Error while finding blogs", err)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		for cur.Next(ctx) {
			var blog model.BlogData
			err := cur.Decode(&blog)
			logs.Error("Error while decoding blogs", err)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
			Blogs = append(Blogs, blog)
		}
		return c.Status(fiber.StatusOK).JSON(Blogs)
	}

}
