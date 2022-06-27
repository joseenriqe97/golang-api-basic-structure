package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joseenriqe97/golang-init/config"
	"github.com/joseenriqe97/golang-init/controller"
	"github.com/joseenriqe97/golang-init/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	config.NewConnectMongo()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	appController := controller.AppController{
		User: controller.UserOrg,
	}

	e = router.NewRouter(e, appController)
	go func() {
		if err := e.Start(config.HTTPListener()); err != nil {
			logrus.WithError(err).Error("shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}
