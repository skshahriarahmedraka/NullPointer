package route

import (
	//
	"app/database"
	// "app/handler"
	// "fmt"
	// "net/http"

	// "net/http"
	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// DATABASE CONNECTION
	// DB1 := database.PostgresInit()
	DB2 := database.MongodbConnection()
	H := database.DatabaseInitialization( DB2)
	// fmt.Println("🚀 ~ file: route.go ~ line 19 ~ funcRouter ~ H : ", H)
	// r.LoadHTMLGlob("dist/*")
	// r.GET("/",handler.Home)
	// r.Static("/","public/")

	// r.StaticFS("/",http.Dir("dist/"))

	// r.Use(static.Serve("/", static.LocalFile("dist/", true)))
	r.POST("/api/login", H.Login)

	r.POST("/api/register", H.Register)
	r.GET("/:UID", H.UserProfile)

	// userGroup:= r.Group("/")
	// userGroup.Use()
	// {
	// 	r.GET("/",H.Home)

	// 	r.GET("/:UID",H.UserProfile)
	// 	r.POST("/login",H.Login)
	// 	r.POST("/register",H.Register)
	// 	r.GET("/:UID/s",H.UserSpaces)
	// 	r.GET("/:UID/fh",H.FavoriteHash)
	// 	r.GET("/:UID/settings",handler.UserSettings)
	// 	r.POST("/:UID/postq",handler.PostQuestion)
	// 	r.POST("/:UID/posta",handler.PostAnswer)
	// }

	// quesGroup:=r.Group("/q/")
	// quesGroup.Use()
	// {
	// 	r.GET("/q",H.FiltedQuestion)
	// 	r.GET("/q/:QID",H.Question)
	// 	r.GET("/q/hot",H.HotQuestion)
	// 	r.GET("/q/blogs",H.Blogs)
	// 	r.GET("/rq/:QID",H.RelatedQuestionListHandler)
	// }

	// spaceGroup:= r.Group("/s/")
	// spaceGroup.Use()
	// {
	// 	r.GET("/s/:SID",handler.Space)
	// }

	// r.Static("")
	// r.StaticFile("/favicon.ico", "public/favicon.ico")
	// r.StaticFS("/",http.Dir("public/*"))

	// r.GET("/",handler.Home)
	// r.GET("/articler/view/:id",handler.ArticleID)
	// routegroup := r.Group("/u")
	// {

	// }
}
