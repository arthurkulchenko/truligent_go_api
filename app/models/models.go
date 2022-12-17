package models

type ServerAccessOption struct {
	Blocked bool                   `json:"blocked,omitempty"`
	BlockingEnabled bool           `json:"blocking_enabled,omitempty"`
	BlockedMessage string          `json:"blocked_message,omitempty"`
	NotificationText string        `json:"notification_text,omitempty"`
	TimeNextBlockingSec int        `json:"time_next_blocking_sec,omitempty"`
	TimeBeforeNotificationSec int  `json:"time_before_notification_sec,omitempty"`
	TimeLastSuccessfulPingAt int   `json:"time_last_successful_ping_at,omitempty"`
	TimeLastPingAt int             `json:"time_last_ping_at,omitempty"`

	LocalCompanyId string          `json:"local_company_id,omitempty"`
	LocalClientSessionToken string `json:"local_client_session_token,omitempty"`
	CloudClientSessionToken string `json:"cloud_client_session_toke,omitempty"`
	CurrentClientPrivateKey string `json:"current_client_private_key,omitempty"`
}
