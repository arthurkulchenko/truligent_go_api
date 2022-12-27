package filerepo

import(
	"encoding/csv"
	"errors"
	"os"
	"log"
	"fmt"
	"time"
	"strconv"
	"github.com/arthurkulchenko/truligent_go_api/app/models"
	// "github.com/arthurkulchenko/truligent_go_api/config"
)

func GetCompanysServerAccessOptions(clientToken string) (models.ServerAccessOption, error) {
	// const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b" // to distinguish server id to choose db
	var sao models.ServerAccessOption
	var err error
	// path := config.FILE_PATH
	path := "./datastore/data/client_1/companies_access_options.csv"
	// func takeDataFromCsv(sao models.ServerAccessOption, serverIdToken string) models.ServerAccessOption {
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	// log.Println(file)
	if err != nil { log.Println(err) }
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil { log.Println(err) }
	for _, el := range records {
		if el[12] == clientToken {
			sao.CompanyId = el[0]
			sao.Blocked, err = strconv.ParseBool(el[1])
			sao.BlockingEnabled, err = strconv.ParseBool(el[2])
			sao.BlockedMessage = el[3]
			sao.NotificationText = el[4]
			sao.TimeNextBlockingSec, err = strconv.ParseInt(el[5], 10, 64)
			sao.TimeBeforeNotificationSec, err = strconv.ParseInt(el[6], 10, 64)
			sao.TimeLastSuccessfulPingAt, err = strconv.ParseInt(el[7], 10, 64)
			sao.TimeLastPingAt, err = strconv.ParseInt(el[8], 10, 64)
			sao.LocalCompanyId = el[9]
			sao.LocalClientSessionToken = el[10]
			sao.CloudClientSessionToken = el[11]
			sao.TruligentApiClientToken = el[12]
			sao.CurrentClientPrivateKey = el[13]
		}
	}
	if sao.CompanyId == "" { return sao, errors.New(fmt.Sprintf("Company with token %s Not found", clientToken)) }
	return sao, err
}

func PutCompanysServerAccessOptions(token string, sao models.ServerAccessOption) (string, error) {
	var saoFields = []string {
		fmt.Sprintf("%v", sao.CompanyId),
		fmt.Sprintf("%v", sao.Blocked),
		fmt.Sprintf("%v", sao.BlockingEnabled),
		fmt.Sprintf("%v", sao.BlockedMessage),
		fmt.Sprintf("%v", sao.NotificationText),
		fmt.Sprintf("%v", sao.TimeNextBlockingSec),
		fmt.Sprintf("%v", sao.TimeBeforeNotificationSec),
		fmt.Sprintf("%v", sao.TimeLastSuccessfulPingAt),
		fmt.Sprintf("%v", sao.TimeLastPingAt),
		fmt.Sprintf("%v", sao.LocalCompanyId),
		fmt.Sprintf("%v", sao.LocalClientSessionToken),
		fmt.Sprintf("%v", sao.CloudClientSessionToken),
		fmt.Sprintf("%v", sao.TruligentApiClientToken),
		fmt.Sprintf("%v", sao.CurrentClientPrivateKey),
	}
	clientId := "client_1"
	path := fmt.Sprintf("./datastore/data/%v/companies_access_options.csv", clientId)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		file, err = os.Create(path)
		if err != nil { log.Println(err) }
	}
	reader := csv.NewReader(file)
	writer := csv.NewWriter(file)
	records, _ := reader.ReadAll()

	var matchRecordIndex int
	var nilRecord int
	// log.Println("TOKEN:", token)
	for lineNum, el := range records { if token == el[12] { matchRecordIndex = lineNum } }
	if matchRecordIndex == nilRecord { return "", errors.New(fmt.Sprintf("Company with token %s Not found", token)) }

	log.Println(time.Now())
	log.Println("UPDATE (*) companies_access_options fields where company.id =", sao.CompanyId)
	records[matchRecordIndex] = saoFields
	if err != nil { log.Println(err) }

	oldPath := fmt.Sprintf("./datastore/data/%v/companies_access_options_old.csv", clientId)
	os.Rename(path, oldPath)
	file, err = os.Create(path)
	writer = csv.NewWriter(file)
	os.Remove(oldPath)
	err = writer.WriteAll(records)
	if err != nil { log.Println(err) }
	return sao.CompanyId, err
}

// func FindCompanyIdByAccessToken(token string) string {
// 	clientId := "client_1"
// 	path := fmt.Sprintf("./datastore/data/%v/access_tokens.csv", clientId)
// 	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
// 	if err != nil { log.Println(err) }
// 	reader := csv.NewReader(file)
// 	records, _ := reader.ReadAll()
// 	var companyId string
// 	for _, el := range records {
// 		if token == el[2] { companyId = el[1] }
// 	}
// 	return companyId
// }

