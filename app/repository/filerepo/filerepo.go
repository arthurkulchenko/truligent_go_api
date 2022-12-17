package filerepo

import(
	"encoding/csv"
	"os"
	"log"
	"fmt"
	"github.com/arthurkulchenko/truligent_go_api/app/models"
	"github.com/arthurkulchenko/truligent_go_api/config"
)

func GetCompanysServerAccessOptions(companyId string) (models.ServerAccessOption, error) {
	var sao models.ServerAccessOption
	// func takeDataFromCsv(sao models.ServerAccessOption, serverIdToken string) models.ServerAccessOption {
	file, err := os.Open(config.FILE_PATH)
	if err != nil { log.Println(err) }
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	log.Println(records)
	return sao, err
}

func CreateOrPutCompanysServerAccessOptions(companyId string, sao models.ServerAccessOption) (string, error) {
	// var sao models.ServerAccessOption
	// func putDataIntoCsv(sao models.ServerAccessOption) {
	file, err := os.Open(config.FILE_PATH)
	if err != nil {
		file, err = os.Create(config.FILE_PATH)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}
	return "sao.id", err
}
