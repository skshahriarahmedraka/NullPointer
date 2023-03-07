package handler



import "github.com/gofiber/fiber/v2"

func (H *DatabaseCollections)HotQues(c *fiber.Ctx) error  {
	return c.Next()
}
