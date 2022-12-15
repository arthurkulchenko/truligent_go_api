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
	// TimeNextBlockingSec int
	// TimeNoResponseBlock string
	// TimeBeforeNotificationSec int
	// NewPrivateKey string
}

type ServerAccessOption struct {
	Blocked bool                  `json:"blocked,omitempty"`
	BlockingEnabled bool          `json:"blocking_enabled,omitempty"`
	BlockedMessage string         `json:"blocked_message,omitempty"`
	NotificationText string       `json:"notification_text,omitempty"`
	TimeNextBlockingSec int       `json:"time_next_blocking_sec,omitempty"`
	TimeBeforeNotificationSec int `json:"time_before_notification_sec,omitempty"`
	// NewPrivateKey string          `json:"new_private_key,omitempty"`
	// TimeNoResponseBlock string    `json:"time_no_response_block,omitempty"`
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
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	query := `SELECT
		blocked,
		server_access_options ->> 'blocking_enabled',
		server_access_options ->> 'blocked_message',
		server_access_options ->> 'notification_text',
		server_access_options ->> 'time_next_blocking_sec',
		server_access_options ->> 'time_before_notification_sec'
	FROM companies WHERE server_access_options ->> 'local_company_id' = $1 OR id = ($1)::uuid`
	// row := dbConn.QueryRowContext(ctx, query, requestBody.CompanyId)
	row := dbConn.QueryRowContext(ctx, query, requestBody.CompanyId)
	var sao ServerAccessOption
	// var p []byte
	// err = row.Scan(&p)
	err = row.Scan(
		&sao.Blocked,
		&sao.BlockingEnabled,
		&sao.BlockedMessage,
		&sao.NotificationText,
		&sao.TimeNextBlockingSec,
		&sao.TimeBeforeNotificationSec,
	)
	fmt.Println("WORKS 6")
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("WORKS 7")
	// err = json.Unmarshal(p, &sao)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	fmt.Println("WORKS 8")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"blocked": sao.Blocked,
		"blocking_enabled": sao.BlockingEnabled,
		"blocked_message": sao.BlockedMessage,
		"notification_text": sao.NotificationText,
		"time_next_blocking_sec": sao.TimeNextBlockingSec,
		"time_before_notification_sec": sao.TimeBeforeNotificationSec,
	})
	hmacSampleSecret := "secret"
	fmt.Println("WORKS 1")
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	fmt.Println("WORKS 2")
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("tokenString")
	// UPDATE company
	return c.JSON(http.StatusOK, echo.Map{ "data": tokenString, })
}
