package services

import(
	"log"
	"time"
	"github.com/arthurkulchenko/truligent_go_api/app/repository"
	// "github.com/arthurkulchenko/truligent_go_api/app/controllers"
	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/arthurkulchenko/truligent_go_api/app/models"
)

func RotateTokenService(companyId string) (string, error) {
	dataSource := repository.filerepo
	sao, err := dataSource.GetCompanysServerAccessOptions(companyId)
	if err != nil { log.Println(err) }
	// TODO:
	// Token Rotation
	var key string
	if sao.TimeNextBlockingSec > time.Now().Unix() {
		key = sao.LocalClientSessionToken
	} else {
		key = sao.CloudClientSessionToken
	}
	tokenString, _ := encryptData(sao, key)
	// TODO:
	// TimeLastSuccessfulPingAt
	// TimeLastPingAt
	_, err = CreateOrPutCompanysServerAccessOptions(companyId, sao)
	return tokenString, err
}

func encryptData(sao models.ServerAccessOption, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"blocked": sao.Blocked,
		"blocking_enabled": sao.BlockingEnabled,
		"blocked_message": sao.BlockedMessage,
		"notification_text": sao.NotificationText,
		"time_next_blocking_sec": sao.TimeNextBlockingSec,
		"time_before_notification_sec": sao.TimeBeforeNotificationSec,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, err
}
