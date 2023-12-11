package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/controller"
)

func RegisterUserRouter(router fiber.Router) {

	router.Get("/user/list", controller.GetUsers)
	router.Get("/user/:username", controller.GetUser)
	router.Post("/user/create", controller.CreateUser)
	router.Delete("/user/:userID", controller.DeleteUser)
}
