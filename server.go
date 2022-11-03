package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/config"
	"github.com/arthurkulchenko/truligent_go_api/app/services"
)
var appConfig *config.AppConfig

func main() {
	appConfig = config.InitializeConfig()
	e := echo.New()
	e.GET("/oms_ping", services.OmsPing)

	if err := e.Start(appConfig.PortNumber); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
  }
}
