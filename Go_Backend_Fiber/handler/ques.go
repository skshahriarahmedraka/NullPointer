package handler



import "github.com/gofiber/fiber/v2"

func (H *DatabaseCollections)Ques(c *fiber.Ctx) error  {
	return c.Next()
}
