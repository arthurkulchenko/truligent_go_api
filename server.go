package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/config"
	"github.com/arthurkulchenko/truligent_go_api/app/services"
	"github.com/arthurkulchenko/truligent_go_api/app/controllers"
)

var appConfig *config.AppConfig

func main() {
	appConfig = config.InitializeConfig()
	services.InitializeServices(appConfig)
	e := echo.New()
	// TODO: add middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.GET("/oms_ping", services.OmsPing)
	e.POST("/rotate_token", controllers.RotateToken)

	if err := e.Start(appConfig.PortNumber); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
