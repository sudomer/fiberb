package lib

import "github.com/gofiber/fiber/v2"

func JSONError(message string) fiber.Map {
	return fiber.Map{
		"status":  "Error!",
		"message": message,
	}
}

func JSONSuccess(message string) fiber.Map {
	return fiber.Map{
		"status":  "Success!",
		"message": message,
	}
}
