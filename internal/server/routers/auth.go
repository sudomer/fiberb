package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/controller"
)

func RegisterAuthRouter(router fiber.Router) {

	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)
}
