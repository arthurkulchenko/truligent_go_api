package services

import(
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/arthurkulchenko/truligent_go_api/config"
	"database/sql"
	"database/sql/driver"
	"context"
	"time"
	"io/ioutil"
	"log"
	"encoding/json"
	"errors"
	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
)

type Company struct {
	Blocked bool
	ServerAccessOptions ServerAccessOption
	// BlockingEnabled bool
	// BlockedMessage string
	// NotificationText string
	// TimeNextBlocking_sec int
	// TimeNoResponseBlock string
	// TimeBeforeNotificationSec int
	// NewPrivateKey string
}

type ServerAccessOption struct {
	BlockingEnabled bool `json:"blocking_enabled,omitempty"`
	BlockedMessage string `json:"blocked_message,omitempty"`
	NotificationText string `json:"notification_text,omitempty"`
	TimeNextBlocking_sec int `json:"time_next_blocking_sec,omitempty"`
	TimeNoResponseBlock string `json:"time_no_response_block,omitempty"`
	TimeBeforeNotificationSec int `json:"time_before_notification_sec,omitempty"`
	NewPrivateKey string `json:"new_private_key,omitempty"`
	// Name        string   `json:"name,omitempty"`
  //   Ingredients []string `json:"ingredients,omitempty"`
  //   Organic     bool     `json:"organic,omitempty"`
  //   Dimensions  struct {
  //       Weight float64 `json:"weight,omitempty"`
  //   } `json:"dimensions,omitempty"`
}

func (a *ServerAccessOption) Scan(value interface{}) error {
	b, ok := value.([]byte)
	fmt.Println(value)
	if !ok { return errors.New("type assertion to []byte failed") }
	return json.Unmarshal(b, &a)
}

func (a ServerAccessOption) Value() (driver.Value, error) {
	return json.Marshal(a)
}

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

type RequestStruct struct {
	CompanyId string `json:"company_id"`
}

func RotateToken(c echo.Context) error {
	var requestBody = RequestStruct{}
	defer c.Request().Body.Close()
	var err error
	var body []byte
	body, err = ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("failed to read rotate_token request body: %s", err)
		return c.String(http.StatusInternalServerError, "failed")
	}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		log.Println("failed to unmatshal rotate_token request body: %s", err)
		return c.String(http.StatusInternalServerError, "failed")
	}
	dbConn := appConfig.DatabaseConnections[serverIdToken] // postgresql
	// fmt.Println("\ndbConn", dbConn, "|")
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	query := `SELECT (blocked, server_access_options) FROM companies WHERE server_access_options ->> 'local_company_id' = $1 OR id = ($1)::uuid`
	row := dbConn.QueryRowContext(ctx, query, requestBody.CompanyId)
	// company := Company{ Blocked: false, ServerAccessOptions: ServerAccessOption{}, }
	var company Company
	// var count int
	// fmt.Println("count is :", count)
	// err = row.Scan(&count)
	err = row.Scan(&company.Blocked, &company.ServerAccessOptions,)
	// err = row.Scan(&company)
	// fmt.Println("ServerAccessOptions is :", company.ServerAccessOptions.BlockingEnabled)
	if err != nil {
		log.Println(err)
		return err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "hello": "hi",
		"blocked": company.Blocked,
		"blocking_enabled": company.ServerAccessOptions.BlockingEnabled,
		"blocked_message": company.ServerAccessOptions.BlockedMessage,
		"notification_text": company.ServerAccessOptions.NotificationText,
		"time_next_blocking_sec": company.ServerAccessOptions.TimeNextBlocking_sec,
		"time_no_response_block": company.ServerAccessOptions.TimeNoResponseBlock,
		"time_before_notification_sec": company.ServerAccessOptions.TimeBeforeNotificationSec,
		"new_private_key": company.ServerAccessOptions.NewPrivateKey, // "Base64.encode64(company.new_private_key.to_s).presence",
	})
	hmacSampleSecret := "secret"
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("tokenString")
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

