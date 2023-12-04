package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/sqlite-golang/app/routes"
	"github.com/rakesh-gupta29/sqlite-golang/config"
	"github.com/rakesh-gupta29/sqlite-golang/pkg/logger"
	"github.com/rakesh-gupta29/sqlite-golang/pkg/middlewares"
)

func main() {
	err := config.LoadAllConfigs(".env")
	if err != nil {

		fmt.Println(config.AppConfig)

		fmt.Print("error finding the .env file")
	}

	logger.SetUpLogger()
	logr := logger.GetLogger()

	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)
	middlewares.MountFiberMiddlewares(app)
	routes.MountAllRoutes(app)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// get the meaning of the code
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
