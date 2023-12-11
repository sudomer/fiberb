package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/model"
	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.uber.org/zap"
)

func Login(c *fiber.Ctx) error {

	var auth model.Auth

	err := c.BodyParser(&auth)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("Please check your credentials!"))
	}

	resp, err := auth.Login()

	if err != nil {
		switch err.Error() {
		case model.ErrUserNotFound:
			return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("OOPS! User not found."))
		case model.ErrWrongPassword:
			return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("OOPS! Wrong password."))
		default:
			lib.Log().Error("Unkown error", zap.String("controller", "login"), zap.String("username", auth.Username), zap.Error(err))
			return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("OOPS! error occurred."))
		}

	}
	// Login process like JWT
	return c.Status(200).JSON(resp)
}

func Register(c *fiber.Ctx) error {

	valid := validator.New()
	var usr model.User

	err := c.BodyParser(&usr)
	if err != nil {
		lib.Log().Error("Requested body is in wrong form",
			zap.String("controller", "register"),
			zap.Any("body", c.Body()),
		)
		return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("Please check your register data!"))
	}

	if err = valid.Struct(usr); err != nil {
		lib.Log().Warn("Client requested illegal form",
			zap.Error(err),
			zap.Any("body", c.Body()),
		)
		return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("OOOPS! Please check your form data for user create"))
	}
	resp, err := usr.CreateUser()
	if err != nil {
		switch err {
		case model.ErrUserAlreadyExist:
			lib.Log().Error("User registration process failed.", zap.String("controller", "register"), zap.Error(err))
			return c.Status(fiber.ErrInternalServerError.Code).JSON(lib.JSONError("User already exist."))
		}
	}
	return c.Status(fiber.StatusOK).JSON(lib.JSONSuccess(resp))
}
