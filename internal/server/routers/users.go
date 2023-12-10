package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/controller"
	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.uber.org/zap"
)

func RegisterUserRouter(router fiber.Router) {

	router.Get("/list", controller.GetUsers)
	router.Get("/user/:username", controller.GetUser)
	router.Post("/create", controller.CreateUser)
	router.Delete("/:userID", controller.DeleteUser)

	lib.Log().Info("Routers initializing...", zap.String("type", "users"))
}
