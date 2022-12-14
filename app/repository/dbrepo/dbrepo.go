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

func GetCompanysServerAccessOptions(clientToken string) (models.ServerAccessOption, error) {
	const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b" // to distinguish server id to choose db

	var sao models.ServerAccessOption
	dbConn := appConfig.DatabaseConnections[serverIdToken] // postgresql
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	// query := `SELECT
	// 	COALESCE(blocked, false),
	// 	COALESCE(CAST(nullif(server_access_options ->> 'blocking_enabled', 'NUL') AS bool), false),
	// 	COALESCE(CAST(nullif(server_access_options ->> 'blocked_message', 'NULL') AS text), ''),
	// 	COALESCE(CAST(nullif(server_access_options ->> 'notification_text', 'NULL') AS text), ''),
	// 	COALESCE(CAST(server_access_options ->> 'time_next_blocking_sec' AS int), 0),
	// 	COALESCE(CAST(server_access_options ->> 'time_before_notification_sec' AS int), 0)
	// FROM companies WHERE server_access_options ->> 'local_company_id' = $1 OR id = ($1)::uuid`
	query := `SELECT
		id,
		COALESCE(blocked, false),
		COALESCE(CAST(nullif(server_access_options ->> 'blocking_enabled', 'NUL') AS bool), false),
		COALESCE(CAST(nullif(server_access_options ->> 'blocked_message', 'NULL') AS text), ''),
		COALESCE(CAST(nullif(server_access_options ->> 'notification_text', 'NULL') AS text), ''),
		COALESCE(CAST(server_access_options ->> 'time_next_blocking_sec' AS int), 0),
		COALESCE(CAST(server_access_options ->> 'time_before_notification_sec' AS int), 0),
		COALESCE(CAST(server_access_options ->> 'time_last_successful_ping_at' AS int), 0),
		COALESCE(CAST(server_access_options ->> 'time_last_ping_at' AS int), 0),
		COALESCE(CAST(nullif(server_access_options ->> 'local_company_id', 'NULL') AS text), ''),
		COALESCE(CAST(nullif(server_access_options ->> 'local_client_session_token', 'NULL') AS text), ''),
		COALESCE(CAST(nullif(server_access_options ->> 'cloud_client_session_token', 'NULL') AS text), ''),
		COALESCE(CAST(nullif(server_access_options ->> 'truligent_api_client_token', 'NULL') AS text), ''),
		COALESCE(CAST(nullif(server_access_options ->> 'current_client_private_key', 'NULL') AS text), ''),
	FROM companies WHERE server_access_options ->> 'truligent_api_client_token' = $1 OR id = ($1)::uuid`
	err := dbConn.QueryRowContext(ctx, query, clientToken).Scan(
		&sao.CompanyId,
		&sao.Blocked,
		&sao.BlockingEnabled,
		&sao.BlockedMessage,
		&sao.NotificationText,
		&sao.TimeNextBlockingSec,
		&sao.TimeBeforeNotificationSec,
		&sao.TimeLastSuccessfulPingAt,
		&sao.TimeLastPingAt,
		&sao.LocalCompanyId,
		&sao.LocalClientSessionToken,
		&sao.CloudClientSessionToken,
		&sao.TruligentApiClientToken,
		&sao.CurrentClientPrivateKey,
	)
	if err != nil {
		log.Println(err)
		return sao, err
	}
	return sao, err
}

func PutCompanysServerAccessOptions(clientToken string, sao models.ServerAccessOption) (string, error) {
	// var sao models.ServerAccessOption
	// TODO:
	// UPDATE companies server_access_options ->> 'time_last_successful_ping_at'
	// UPDATE companies server_access_options ->> 'time_last_ping_at'
	panic("NOT IMPLEMENTED")
}

// func FindCompanyIdByAccessToken(token string) string {
// 	// TODO
// 	panic("NOT IMPLEMENTED")
// }
