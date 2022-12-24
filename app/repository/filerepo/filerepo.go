package filerepo

import(
	"encoding/csv"
	"os"
	"log"
	"fmt"
	"strconv"
	"github.com/arthurkulchenko/truligent_go_api/app/models"
	// "github.com/arthurkulchenko/truligent_go_api/config"
)

func GetCompanysServerAccessOptions(clientToken string) (models.ServerAccessOption, error) {
	const serverIdToken = "8fad429c-f54d-4b8d-87a6-874771c7f68b" // to distinguish server id to choose db
	var sao models.ServerAccessOption
	var err error
	// path := config.FILE_PATH
	path := "./datastore/data/client_1/companies_access_options.csv"
	// func takeDataFromCsv(sao models.ServerAccessOption, serverIdToken string) models.ServerAccessOption {
	file, err := os.Open(path)
	if err != nil { log.Println(err) }
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
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
	return sao, err
}

func CreateOrPutCompanysServerAccessOptions(companyId string, sao models.ServerAccessOption) (string, error) {
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
	// var sao models.ServerAccessOption
	// func putDataIntoCsv(sao models.ServerAccessOption) {
	// path := config.FILE_PATH
	path := "./datastore/data/client_1/companies_access_options.csv"
	// path := "../companies_access_options.csv"
	// file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	file, err := os.Open(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil { log.Println(err) }
	}
	reader := csv.NewReader(file)
	writer := csv.NewWriter(file)
	records, _ := reader.ReadAll()
	var matchRecord int
	var nilRecord int
	var csvHeaderExists bool
	if len(records) > 0 && records[0] != nil { csvHeaderExists = true }
	for lineNum, el := range records {
		if companyId == el[0] {
			matchRecord = lineNum
		}
	}
	// TODO DEBUG
	if matchRecord != nilRecord {
		log.Println("Substitude")
		records[matchRecord] = saoFields
		err = writer.WriteAll(records)
		if err != nil { log.Println(err) }
		return sao.CompanyId, err
	}
	if csvHeaderExists {
		log.Println("Append")
		records = append(records, saoFields)
		err = writer.WriteAll(records)
	} else {
		log.Println("Fulfill")
		var data = [][]string{
			{ "company_id", "blocked", "blocking_enabled", "blocked_message", "notification_text", "time_next_blocking_sec", "time_before_notification_sec", "time_last_successful_ping_at", "time_last_ping_at", "local_company_id", "local_client_session_token", "cloud_client_session_toke", "current_client_private_key", },
			saoFields,
		}
		err = writer.WriteAll(data)
	}
	if err != nil { log.Println(err) }
	return sao.CompanyId, err
}
