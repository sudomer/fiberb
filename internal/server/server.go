package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/routers"
	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.uber.org/zap"
)

func NewServer() {
	app := fiber.New()
	registerRouters(app)
	lib.Log().Info("Routers initialized!", zap.Int("count", len(app.GetRoutes())), zap.String("type", "auth,users"))
	lib.Log().Error("", zap.Error(app.Listen(os.Getenv("EXPOSE_PORT"))))
}

func registerRouters(app *fiber.App) {
	app.Route("/", routers.RegisterBaseRouter)
	app.Route("/clientarea", routers.RegisterUserRouter)
	app.Route("/auth", routers.RegisterAuthRouter)

}
