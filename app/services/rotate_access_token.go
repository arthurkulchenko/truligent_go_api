package services

import(
	"log"
	"time"
	"encoding/base64"
	"github.com/arthurkulchenko/truligent_go_api/app/repository/filerepo"
	// "strings"
	// "github.com/arthurkulchenko/truligent_go_api/app/controllers"
	"github.com/go-jose/go-jose/v3"
	"encoding/json"
	"github.com/arthurkulchenko/truligent_go_api/app/models"
)

func RotateTokenService(companyId string) (string, error) {
	// dataSource := repository.filerepo
	sao, err := filerepo.GetCompanysServerAccessOptions(companyId)
	if err != nil { log.Println(err) }
	var key string
	currentTime := time.Now().Unix()
	if sao.TimeNextBlockingSec > currentTime {
		key = sao.CurrentClientPrivateKey
	} else {
		// Will block company
		key = sao.CloudClientSessionToken
	}
	tokenString, _ := encryptData(sao, key)
	log.Println(sao.TimeLastSuccessfulPingAt)
	sao.TimeLastSuccessfulPingAt = currentTime
	sao.TimeLastPingAt = currentTime
	log.Println(currentTime)

	_, err = filerepo.CreateOrPutCompanysServerAccessOptions(companyId, sao)
	return tokenString, err
}

func encryptData(sao models.ServerAccessOption, key string) (string, error) {
	var err error
	payload, _ := json.Marshal(sao)
	bytePrivatKey := []byte(key)
	byteKey := make([]byte, base64.StdEncoding.EncodedLen(len(bytePrivatKey)))
	base64.StdEncoding.Encode(byteKey, bytePrivatKey)
	if err != nil { panic(err) }
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.PBES2_HS256_A128KW, Key: string(byteKey)}, nil)
	encObj, err := encrypter.Encrypt([]byte(payload))
	tokenString, err := encObj.CompactSerialize()
	if err != nil { panic(err) }
	return string(string(tokenString)), err
}
