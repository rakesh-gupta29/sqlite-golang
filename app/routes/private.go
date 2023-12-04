package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/sqlite-golang/app/handlers"
)

func mountPrivateRoutes(app *fiber.App) {
	route := app.Group("/private")

	route.Get("/admin", handlers.GetAdminProfile)
}
