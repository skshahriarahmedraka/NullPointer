package middle

import "github.com/gofiber/fiber"


func JwtAuth(  )func(*fiber.Ctx)error {
	  // Set a custom header on all responses:
	//   c.Set("X-Custom-Header", "Hello, World")
	return func (c *fiber.Ctx) error  {
		c.Next()
		return nil
	}
} 
