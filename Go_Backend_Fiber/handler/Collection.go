package handler



import "github.com/gofiber/fiber/v2"

func (H *DatabaseCollections)Collection(c *fiber.Ctx) error  {
	return c.Next()
}
