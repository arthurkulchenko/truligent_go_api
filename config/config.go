package config

import(
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
	"database/sql"
)

const DEFAULT_PORT_NUMBER = 4000

type AppConfig struct {
	DatabaseConnections map[string]*sql.DB
	PortNumber string
	Environment string
}

type DatabaseConfig struct {
	Adapter string
	Host string
	Port string
	Database string
	Username string
	Password string
	Url string
}

func InitializeConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil { log.Fatal("Error loading .env file") }
	dbConns := make(map[string]*sql.DB)
	portNumber := os.Getenv("PORT_NUMBER")
	if portNumber == "" { portNumber = DEFAULT_PORT_NUMBER }
	return &AppConfig { DatabaseConnections: dbConns, PortNumber: fmt.Sprintf(":%v", portNumber), }
}
