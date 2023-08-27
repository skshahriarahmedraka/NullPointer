package handler

import (

	//
	"app/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "gorm.io/gorm"
)
type DatabaseCollections struct {
	Mongo *mongo.Database
	// Postgres *gorm.DB 
}




type Sample struct {
	Name string `json:"Name"`
	Roll int `json:"Roll"`
	Address string `json:"Address"`
}

func (H *DatabaseCollections)UserProfile( c *gin.Context){
	//SEE THE COOKIES 
	cookie_value, err := c.Cookie("Auth1")
	if err!=nil {

		fmt.Println("🚀🔥❌ ~ file: handler.go ~ line 31 ~ func ~ err : ", err)
	}
    fmt.Println("🚀🚀✨ ~ file: handler.go ~ line 31 ~ func ~ cookie_value : ", cookie_value)
	// allow Cors
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	//

	fmt.Println("🚀 ~ file: handler.go ~ line 17 ~ func ~ H : ", H)
	s:= Sample{Name:"sk shahriar ahmed raka",Roll:180107,Address: "Dhaka 1206,Bangladesh" }
	_=s
	z:=model.UserData{
		UserID:"skraka",
		UserName:"Sk Shahriar Ahmed Raka",
		// UserTitle:"Working at Google Inc.",
		UserImage:"https://res.cloudinary.com/dqo0ssnti/image/upload/v1653060640/samples/jpeg_1_qlbtcn.jpg",
		Badges: struct {
			Reputation int `json:Reputation bson:"Reputation"`
			Gold       int `json:"Gold" bson:"Gold"`
			Silver     int `json:"Silver" bson:"Silver"`
			Bronze     int `json:"Bronze" 		bson:"Bronze"`
		} {
				Reputation: 234267,
				Gold:999,
				Silver:888,
				Bronze:777},
		// Badges : map[string]int{
		// 	"Reputation":354,
		// 	"Gold":444,
		// 	"Silver":508,
		// 	"Bronze":676},
		Follower: []string{},
		Location:"Dhaka, Bangladesh",
		MembershipTime:primitive.Timestamp{T:uint32(time.Now().Unix())},
		LastSeen:primitive.Timestamp{T:uint32(time.Now().Unix())},
		Aboutme:`###  Hi ( Peace be upon you ) , I'm "Sk Shahriar Ahmed Raka" 
# A Curious Learner, Full-Stack Web Developer, Security Researcher <br>		
### Here are my skills and strengths:<br>
✓  Expert in **Golang** <br/>
✓  Expert in **Fiber** framework (using Golang) <br/>
✓  Expert in **WebAssembly** (using Golang) <br/>
✓  Expert in database design, development, optimization, and migration <br/>
(PostgreSQL, MySQL, MongoDB , Redis)<br/>
✓  Expert in ( Grpc, protocol buffer )<br/>
✓  Experienced in ( WebSocket, Socket.io, WebRTC ) for real-time client and server applications <br/>
✓  Experienced in ( Vue.js ) and some knowledge in ( TypeScript )<br/>
✓  Good understanding of ( Docker, Bash, PowerShell, Git,<br/>
Nginx, Kubernetes )<br/>
`,
		Mysite:"https://www.skshahriarahmedraka.com/",
		Github:"https://github.com/skshahriarahmedraka",
		Twitter:"https://twitter.com/shahriarraka7",
		Linkedin:"http://linkedin.com/raka",
		TopTags:[]string{"go","rust","python","svelte","postgresql"},
		TopTagsPercent:[]int{50,30,20,5,2},
		SelectedPanel :"Profile"}
	x,err := json.Marshal(z)
	if err !=nil {
		log.Fatalln("🔥❌",err)
	}
	var y model.UserData
	json.Unmarshal(x,&y)
    // fmt.Println("🚀 ~ file: handler.go ~ line 64 ~ func ~ y : ", y)
	

 _=x
 _=y
	c.JSON(http.StatusOK,z)
}
// func Home(c *gin.Context){
// 	fmt.Println("🚀 ~ file: handler.go ~ line 17 ~ func ~ H : ", H)
// 	s:= Sample{Name:"sk shahriar ahmed raka",Roll:180107,Address: "Dhaka 1206,Bangladesh" }

// 	c.JSON(http.StatusOK,s)
	
// }


func ArticleID(c *gin.Context){

}

// USER

func UserProfile(c *gin.Context){

}
func UserSettings(c *gin.Context){

}

func PostQuestion(c *gin.Context){

}

func PostAnswer(c *gin.Context){

}

// QUESTION GROUP

func Question(c *gin.Context){

}
func Answer(c *gin.Context){

}

// SPACE 

func Space(c *gin.Context){

}