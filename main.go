package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	// "github.com/rakesh-gupta29/sqlite-golang/app/database"
	// "github.com/rakesh-gupta29/sqlite-golang/app/handlers"
	"github.com/rakesh-gupta29/sqlite-golang/app/routes"
	"github.com/rakesh-gupta29/sqlite-golang/config"
	"github.com/rakesh-gupta29/sqlite-golang/pkg/logger"
	"github.com/rakesh-gupta29/sqlite-golang/pkg/middlewares"
	"github.com/rakesh-gupta29/sqlite-golang/pkg/validator"
)

func main() {

	// if err := database.Connect("data/forms.db"); err != nil {
	// 	log.Fatal("Failed to connect to database.")
	// }

	// if err := handlers.SeedData(); err != nil {
	// 	log.Fatal(err)
	// }

	err := config.LoadAllConfigs(".env")

	if err != nil {
		fmt.Print("error finding the .env file")
	}

	logger.SetUpLogger()
	logr := logger.GetLogger()

	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	middlewares.MountFiberMiddlewares(app)
	routes.MountAllRoutes(app)
	validator.MountValidators(app)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		logr.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	serverAddr := fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port)

	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}
