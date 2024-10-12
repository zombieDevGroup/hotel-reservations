package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel-reservations/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "John",
		LastName:  "Smith",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "John"})
}
