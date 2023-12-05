package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rakesh-gupta29/sqlite-golang/config"
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

func RateLimiter(app *fiber.App) {
	app.Use(limiter.New())
}

func AllowedOrigins(app *fiber.App) {
	allowedOrigins := config.AppConfig.AllowedDomains
	app.Use(func(c *fiber.Ctx) error {

		origin := c.Get("Origin")
		isAllowed := false
		for _, v := range allowedOrigins {
			if v == origin {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Not allowed from this domain"})
		}

		return c.Next()
	})
}
