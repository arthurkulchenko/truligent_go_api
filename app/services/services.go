package services

import(
	"net/http"
	"github.com/labstack/echo/v4"
)

var client *Client

type Client struct {
	dbConn *sql.DB
}

func OmsPing(c echo.Context) error {
	// token := c.QueryParam("token")
	token := "8fad429c-f54d-4b8d-87a6-874771c7f68b"
	client := Client { dbConn: AppConfig.DatabaseConnections[token] }
	return c.String(http.StatusOK, client.OmsPingCall()
}
