package services

import(
	"context"
	// "database/sql"
	// "database/sql/driver"
	"encoding/csv"
	"encoding/json"
	// "errors"
	"github.com/arthurkulchenko/truligent_go_api/config"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
)

type ServerAccessOption struct {
	Blocked bool                  `json:"blocked,omitempty"`
	BlockingEnabled bool          `json:"blocking_enabled,omitempty"`
	BlockedMessage string         `json:"blocked_message,omitempty"`
	NotificationText string       `json:"notification_text,omitempty"`
	TimeNextBlockingSec int       `json:"time_next_blocking_sec,omitempty"`
	TimeBeforeNotificationSec int `json:"time_before_notification_sec,omitempty"`
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
	var sao ServerAccessOption
	// takeDataFromDb(sao, requestBody.CompanyId)
	sao = takeDataFromCsv(sao, requestBody.CompanyId)
	tokenString, _ := encryptData(sao)
	// updateDataInDb(sao, id)
	putDataIntoCsv(sao)
	// TODO: UPDATE companies server_access_options ->> 'time_last_successful_ping_at'
	return c.JSON(http.StatusOK, echo.Map{ "data": tokenString, })
}

func encryptData(sao ServerAccessOption) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"blocked": sao.Blocked,
		"blocking_enabled": sao.BlockingEnabled,
		"blocked_message": sao.BlockedMessage,
		"notification_text": sao.NotificationText,
		"time_next_blocking_sec": sao.TimeNextBlockingSec,
		"time_before_notification_sec": sao.TimeBeforeNotificationSec,
	})
	hmacSampleSecret := "secret"
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, err
}

func takeDataFromCsv(sao ServerAccessOption, serverIdToken string) ServerAccessOption {
	file, err := os.Open(config.FILE_PATH)
	if err != nil { fmt.Println(err) }
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Println(records)
	return sao
}

func putDataIntoCsv(sao ServerAccessOption) {
	file, err := os.Open(config.FILE_PATH)
	if err != nil {
		file, err = os.Create(config.FILE_PATH)
		if err != nil {
			fmt.Println(err)
		}
	}
	writer := csv.NewWriter(file)
	var data = [][]string{
		{ "blocked", "blocking_enabled", "blocked_message", "notification_text", "time_next_blocking_sec", "time_before_notification_sec", },
		{
			fmt.Sprintf("%v", sao.Blocked),
			fmt.Sprintf("%v", sao.BlockingEnabled),
			fmt.Sprintf("%v", sao.BlockedMessage),
			fmt.Sprintf("%v", sao.NotificationText),
			fmt.Sprintf("%v", sao.TimeNextBlockingSec),
			fmt.Sprintf("%v", sao.TimeBeforeNotificationSec),
		},
	}
	err = writer.WriteAll(data)
	if err != nil {
	fmt.Println(err)
	}
}

const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b" // to distinguish server id to choose db

func takeDataFromDb(companyId string, sao ServerAccessOption) (ServerAccessOption, error) {
	dbConn := appConfig.DatabaseConnections[serverIdToken] // postgresql
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	query := `SELECT
		COALESCE(blocked, false),
		COALESCE(CAST(nullif(server_access_options ->> 'blocking_enabled', 'NUL') AS bool), false),
		COALESCE(CAST(nullif(server_access_options ->> 'blocked_message', 'NULL') AS text), ''),
		COALESCE(CAST(nullif(server_access_options ->> 'notification_text', 'NULL') AS text), ''),
		COALESCE(CAST(server_access_options ->> 'time_next_blocking_sec' AS int), 0),
		COALESCE(CAST(server_access_options ->> 'time_before_notification_sec' AS int), 0)
	FROM companies WHERE server_access_options ->> 'local_company_id' = $1 OR id = ($1)::uuid`
	err := dbConn.QueryRowContext(ctx, query, companyId).Scan(
		&sao.Blocked,
		&sao.BlockingEnabled,
		&sao.BlockedMessage,
		&sao.NotificationText,
		&sao.TimeNextBlockingSec,
		&sao.TimeBeforeNotificationSec,
	)
	if err != nil {
		log.Println(err)
		return sao, err
	}
	return sao, err
}

func updateDataInDb(sao ServerAccessOption, id string) {
	// TODO:
}
