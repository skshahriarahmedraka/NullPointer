package middle

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)


func JwtAuth(  )func(*fiber.Ctx)error {
	  // Set a custom header on all responses:
	//   c.Set("X-Custom-Header", "Hello, World")
	return func (c *fiber.Ctx) error  {
		token := c.Get(os.Getenv("COOKIE_SECRET_JWT_AUTH1"))
        fmt.Println("🚀 ~ file: main.go ~ line 38 ~ app.Use ~ token : ", token)
		err:=c.Next()
		return err
	}
} 
