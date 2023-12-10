package routers

import "github.com/gofiber/fiber/v2"

func RegisterBaseRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hi")
	})
}
