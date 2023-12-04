package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RenderHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
