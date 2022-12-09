package services

import(
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/config"
	"database/sql"
	// "time"
	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
)

var client *Client
var appConfig *config.AppConfig
const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b"

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

func OmsPing(c echo.Context) error {
	// token := c.QueryParam("token")
	// client := Client { dbConn: config.AppConfig.DatabaseConnections[serverIdToken] }
	client := Client { dbConn: appConfig.DatabaseConnections[serverIdToken] }
	// client := Client { dbConn: config.AppConfig }
	// fmt.Println(token)
	// return c.String(http.StatusOK, calll())
	return c.String(http.StatusOK, client.OmsPingCall())
}

func RotateToken(c echo.Context) error {
	hmacSampleSecret := "secret"
	companyId := c.FormValue("company_id")

	// client := Client { dbConn: config.AppConfig.DatabaseConnections[serverIdToken] }
	// client := Client { dbConn: appConfig.DatabaseConnections[serverIdToken] }
	dbConn := appConfig.DatabaseConnections[serverIdToken] // postgresql

	// company ||= Company.local_company_id(company_id).or(Company.where(id: company_id)).last

	fmt.Println(companyId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"blocked": "company.blocked",
		"blocking_enabled": "company.blocking_enabled",
		"blocked_message": "company.blocked_message",
		"notification_text": "company.notification_text",
		"time_next_blocking_sec": "company.time_next_blocking_sec",
		"time_no_response_block": "company.time_no_response_block",
		"time_before_notification_sec": "company.time_before_notification_sec",
		"new_private_key": "Base64.encode64(company.new_private_key.to_s).presence",
	})
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil { return err }
	// UPDATE company
	return c.JSON(http.StatusOK, echo.Map{ "data": tokenString, })
}

	// fn wip() {
		// callback := c.QueryParam("callback")
  // var content struct {
	// 	Response  string    `json:"response"`
	// 	Timestamp time.Time `json:"timestamp"`
	// 	Random    int       `json:"random"`
	// }
	// content.Response = "Sent via JSONP"
	// content.Timestamp = time.Now().UTC()
	// content.Random = 1000
	// return c.JSONP(http.StatusOK, callback, &content)
	// }

