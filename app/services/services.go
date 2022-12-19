package services

import(
	// "github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/config"
	// "github.com/golang-jwt/jwt/v4"
	"database/sql"
	// "fmt"
)

var client *Client
var appConfig *config.AppConfig
// const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b"

type Client struct {
	dbConn *sql.DB
}

// type postgresDBRepo struct {
// 	AppConfig *config.AppConfig
// 	DB *sql.DB
// }

func InitializeServices(appConfiguration *config.AppConfig) {
	appConfig = appConfiguration
}
