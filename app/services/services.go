package services

import(
	"net/http"
	"github.com/labstack/echo/v4"
)

// type jsonResponse struct {
// 	OK bool `json:"ok"`
// 	Message string `json:"message"`
// }

	// response.Header().Set("Content-Type", "application/json")
	// e.GET("/json", func(c echo.Context) error {
	// 	resp := jsonResponse { OK: true, Message: "Available!" }
	// 	return c.JSON(http.StatusOK, &resp)
	// })
	// 

type User struct {
	DatabaseConnection *sql.DB
}

func getUserToken(c echo.Context) error {
	return "token"
}

func OmsPing(c echo.Context) error {
	userToken := getUserToken(c)
	AppConfig.DatabaseConnections[userToken]
	return c.String(http.StatusOK, OmsPingCall(client))
}
