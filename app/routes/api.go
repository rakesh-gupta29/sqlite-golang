package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/sqlite-golang/app/handlers"
)

func MountAPIRoutes(app *fiber.App) {
	route := app.Group("/api/v1")
	route.Get("/healthcheck", handlers.HealthCheck)

	route.Get("/admin/all", handlers.GetAllSubmissions)
	route.Post("/submit", handlers.Submit)

}
