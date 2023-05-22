package handler

import (
	"app/logs"
	"app/model"
	"context"
	"fmt"

	// "encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections)UserDataFlair(c *fiber.Ctx) error  {
	c.Accepts("application/json")

	
	fmt.Println("🚀 ~ file: UserdataFlair.go ~ line 22 ~ ifc.Params ~ c.Params(\"ID\") : ", c.Params("ID"))
	if  c.Params("ID")!= "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbUserData := model.UserData{}
       
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
		err := H.MongoUserCol.FindOne(ctx, bson.M{"_id": id}).Decode(&dbUserData)
		logs.Error("Error while finding user flair data", err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching user flair data",
			})
		}
		dbUserData.Password = ""
		var userFlair model.UserDataFlair
		userFlair.UserID = dbUserData.UserID
		userFlair.UserURL = dbUserData.UserURL
		userFlair.UserName = dbUserData.UserName
		userFlair.UserImage = dbUserData.UserImage
		userFlair.Badges = dbUserData.Badges
		userFlair.Location = dbUserData.Location
		userFlair.Aboutme = dbUserData.Aboutme

		fmt.Println("🚀 ~ file: UserdataFlair.go ~ line 54 ~ ifc.Params ~ userFlair : ", userFlair)
		
        
		return c.Status(fiber.StatusOK).JSON(userFlair)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

}
