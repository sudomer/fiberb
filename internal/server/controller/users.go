package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sudomer/boiler-fiber/internal/server/model"
	"github.com/sudomer/boiler-fiber/pkg/lib"
	"go.uber.org/zap"
)

func GetUsers(c *fiber.Ctx) error {

	users, err := model.GetUsers()
	if err != nil {
		lib.Log().Warn("Server database error", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(lib.JSONError("We can't response with correctly! Please ask to administrator"))
	}
	usr_data := []struct {
		ID       string
		Username string
	}{}
	for _, usr := range users {
		temp := struct {
			ID       string
			Username string
		}{
			ID:       usr.ID.Hex(),
			Username: usr.Username,
		}

		usr_data = append(usr_data, temp)
	}
	return c.Status(fiber.StatusOK).JSON(usr_data)
}
func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	lib.Log().Warn("", zap.String("", username))
	res, err := model.GetUser(username)
	if err != nil {
		lib.Log().Warn("Server database error", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(lib.JSONError("We can't response with correctly! Please ask to administrator"))
	}

	if len(res.Username) < 1 {
		lib.Log().Warn("User not found", zap.Error(err), zap.Any("Result", res))
		return c.Status(fiber.StatusBadRequest).JSON(lib.JSONError("We can't find user"))
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
func DeleteUser(c *fiber.Ctx) error {

	userID := c.Params("usr-del-id")

	err := model.DeleteUser(userID)
	if err != nil {
		lib.Log().Warn("Server database error", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(lib.JSONError("We can't response with correctly! Please ask to administrator"))
	}

	return c.Status(fiber.StatusOK).JSON(lib.JSONSuccess("User deleted"))
}
func CreateUser(c *fiber.Ctx) error {
	valid := validator.New()

	var usr model.User

	err := c.BodyParser(&usr)
	if err != nil {
		lib.Log().Warn("Client requested illegal form", zap.Error(err))
		return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("OOOPS! Please check your form data for user create"))
	}

	if err = valid.Struct(usr); err != nil {
		lib.Log().Warn("Client requested illegal form", zap.Error(err))
		return c.Status(fiber.ErrBadRequest.Code).JSON(lib.JSONError("OOOPS! Please check your form data for user create"))
	}
	res, err := usr.CreateUser()
	if err != nil {
		lib.Log().Warn("Server database error", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(lib.JSONError(res))
	}

	return c.Status(fiber.StatusOK).JSON(lib.JSONSuccess(fmt.Sprintf("UserID: %s", res)))

}
