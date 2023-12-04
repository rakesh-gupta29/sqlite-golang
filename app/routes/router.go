package routes

import "github.com/gofiber/fiber/v2"

func MountAllRoutes(app *fiber.App) {
	mountWebRoutes(app)
	MountAPIRoutes(app)
	mountPrivateRoutes(app)
	app.Static("/static", "./public")
	NotFoundRoute(app)
}

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "sorry, endpoint is not found",
			})
		},
	)
}
