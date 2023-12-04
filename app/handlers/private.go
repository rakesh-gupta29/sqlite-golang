package handlers

import "github.com/gofiber/fiber/v2"

func GetAdminProfile(c *fiber.Ctx) error {
	return c.SendString("getting some private info")
}
