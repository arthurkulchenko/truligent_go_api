package models

type ServerAccessOption struct {
	CompanyId string               `json:"blocked_message,omitempty"`
	Blocked bool                   `json:"blocked,omitempty"`
	BlockingEnabled bool           `json:"blocking_enabled,omitempty"`
	BlockedMessage string          `json:"blocked_message,omitempty"`
	NotificationText string        `json:"notification_text,omitempty"`
	TimeNextBlockingSec int64        `json:"time_next_blocking_sec,omitempty"`
	TimeBeforeNotificationSec int64  `json:"time_before_notification_sec,omitempty"`
	TimeLastSuccessfulPingAt int64   `json:"time_last_successful_ping_at,omitempty"`
	TimeLastPingAt int64             `json:"time_last_ping_at,omitempty"`
	LocalCompanyId string          `json:"local_company_id,omitempty"`
	LocalClientSessionToken string `json:"local_client_session_token,omitempty"`
	CloudClientSessionToken string `json:"cloud_client_session_token,omitempty"`
	TruligentApiClientToken string `json:"truligent_api_client_token,omitempty"`
	CurrentClientPrivateKey string `json:"current_client_private_key,omitempty"`
}
