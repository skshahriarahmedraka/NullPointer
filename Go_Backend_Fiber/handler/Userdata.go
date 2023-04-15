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

func (H *DatabaseCollections)UserData(c *fiber.Ctx) error  {
	c.Accepts("application/json")

	// s:=`
	// {
	// 	UserID: 'skraka',
	// 	UserName: 'Sk Shahriar Ahmed Raka',
	// 	Email: 'skshahra@gmail.com',
	// 	Password: '123456',
	// 	UserTitle : 'Full-Stack Web Developer',
	// 	UserImage:
	// 		'https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg',
	// 	Badges: {
	// 		Reputation: 13452543,
	// 		Gold: 999,
	// 		Silver: 888,
	// 		Bronze: 777
	// 	},
	// 	Follower: ['RKA', 'SHAHRIAR	'],
	// 	Following: ['RKA', 'SHAHRIAR'],
	// 	// Badges map[string]int
	// 	Location: 'Dhaka, Bangladesh',
	// 	MembershipTime: '3 year 5 Month',
	// 	LastSeen: 'This Week',
	// 	Aboutme:
	// 		'A Curious Learner, Full-Stack Web Developer, Security Researcher\nHere are my skills and strengths:\n✓ Expert in Golang\n ✓ Expert in Fiber framework (using Golang) \n ✓ Expert in WebAssembly (using Golang)  Expert in Golang     ✓ Expert in Fiber framework (using Golang)    ✓ Expert in WebAssembly (using Golang) ✓ Expert in database design, development, optimization, and migration    (PostgreSQL, MySQL, MongoDB , Redis) ✓ Expert in ( Grpc, protocol buffer ) ✓ Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications ✓ Experienced in ( Svelte.js ) and some knowledge in ( TypeScript ) ✓ Good understanding of ( Docker, Bash, PowerShell, Git,    Nginx, Kubernetes )        Github : github.com/skshahriarahmedraka    Upwork : upwork.com/o/profiles/users/~0107ef3405bffbe57e    Linkedin : linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/  Telegram : t.me/shahriarraka ',
	// 	Mysite: 'www.shahriarraka.me',
	// 	Github: 'www.github.com/skshahriarahmedraka',
	// 	Twitter: 'www.twitter.com/shahriarraka7',
	// 	Linkedin: 'www.linkedin.com/in/sk-shahriar-ahmed-raka-862a31193/',
	// 	// "TopTags"    : ["Go","Rust","Python","Svelte","PostgreSQL"],
	// 	TopTags: [
	// 		{ Name: 'Go', Score: 12, NumberOfPost: 4 },
	// 		{ Name: 'Rust', Score: 10, NumberOfPost: 6 },
	// 		{ Name: 'Python', Score: 7, NumberOfPost: 4 },
	// 		{ Name: 'Svelte', Score: 12, NumberOfPost: 4 },
	// 		{ Name: 'PostgreSQL', Score: 12, NumberOfPost: 4 }
	// 	],
	// 	SelectedPanel: 'Profile',
	// 	AccountType: 'regular'
	// }
	// `
	// var UserData model.UserData
	//  json.Unmarshal([]byte(s), &UserData)
	// return c.Status(fiber.StatusOK).JSON(UserData)
	if  c.Params("ID")!= "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbUserData := model.UserData{}
        fmt.Println("🚀 ~ file: Userdata.go ~ line 67 ~ ifc.Params ~ c.Query(\"ID\") : ", c.Params("ID"))
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
        fmt.Println("🚀 ~ file: Userdata.go ~ line 68 ~ ifc.Params ~ id : ", id)
		err := H.MongoUserCol.FindOne(ctx, bson.M{"_id": id}).Decode(&dbUserData)
		logs.Error("Error while finding user data", err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching user data",
			})
		}
		dbUserData.Password = ""
		
		return c.Status(fiber.StatusOK).JSON(dbUserData)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

}
