package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/sqlite-golang/app/handlers"
)

func mountWebRoutes(app *fiber.App) {

	routes := app.Group("/")

	routes.Get("/", handlers.RenderHome)

}
