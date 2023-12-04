package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MountFiberMiddlewares(app *fiber.App) {
	app.Use(
		favicon.New(favicon.Config{
			File: "./favicon.ico",
			URL:  "/favicon.ico",
		}),
		cors.New(),
		logger.New(),
	)
}
