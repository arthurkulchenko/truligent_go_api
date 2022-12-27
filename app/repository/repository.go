package repository

import(
	"github.com/arthurkulchenko/truligent_go_api/app/models"
)

// type

type StorageInterface interface {
	GetCompanysServerAccessOptions(clientToken string) (models.ServerAccessOption, error)
	PutCompanysServerAccessOptions(clientToken string, sao models.ServerAccessOption) (string, error)
}

// func () {
// 	config[serverIdToken]
// 	serverIdToken
// }
