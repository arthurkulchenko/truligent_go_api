package services

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

func (c *Client) OmsPingCall() string {
	connection := fmt.Sprintf("%v", c.dbConn)
	fmt.Println(connection)
	return "Pong"
}

func OmsPing(c echo.Context) error {
	// token := c.QueryParam("token")
	// client := Client { dbConn: config.AppConfig.DatabaseConnections[serverIdToken] }
	// client := Client { dbConn: config.AppConfig }
	// fmt.Println(token)
	// return c.String(http.StatusOK, calll())

	// client := Client { dbConn: appConfig.DatabaseConnections[serverIdToken] }
	return c.String(http.StatusOK, "client.OmsPingCall()")
}
