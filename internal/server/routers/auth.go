package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/controller"
	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.uber.org/zap"
)

func RegisterAuthRouter(router fiber.Router) {

	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)

	lib.Log().Info("Routers initializing...", zap.String("type", "authentication"))
}
