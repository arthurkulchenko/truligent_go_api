package dbrepo

import(
	// "database/sql"
	// "database/sql/driver"
	"github.com/arthurkulchenko/truligent_go_api/app/models"
	"github.com/arthurkulchenko/truligent_go_api/config"
	"time"
	"context"
	"log"
)

func GetCompanysServerAccessOptions(companyId string) (models.ServerAccessOption, error) {
	const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b" // to distinguish server id to choose db

	var sao models.ServerAccessOption
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

func CreateOrPutCompanysServerAccessOptions(companyId string, sao models.ServerAccessOption) (string, error) {
	// var sao models.ServerAccessOption
	// TODO:
	// UPDATE companies server_access_options ->> 'time_last_successful_ping_at'
	// UPDATE companies server_access_options ->> 'time_last_ping_at'
}
