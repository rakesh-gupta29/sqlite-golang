package handlers

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("app is running fine")
}