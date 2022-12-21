package repository

import(
	"github.com/arthurkulchenko/truligent_go_api/app/models"
)

// type

type StorageInterface interface {
	GetCompanysServerAccessOptions(companyId string) (models.ServerAccessOption, error)
	CreateOrPutCompanysServerAccessOptions(companyId string, sao models.ServerAccessOption) (string, error)
}

// func () {
// 	config[serverIdToken]
// 	serverIdToken
// }
