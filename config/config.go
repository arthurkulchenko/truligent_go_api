package config

import(
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
	"database/sql"
	"time"
	"gopkg.in/yaml.v3"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4"
)

const DEV_CLIENT_TOKEN = "8fad429c-f54d-4b8d-87a6-874771c7f68b" // to distinguish server id to choose db
const FILE_PATH = "../companies_access_options.csv"

const DEFAULT_PORT_NUMBER = 4000
const MAX_OPEN_DB_CONN = 10
const MAX_IDLE_DB_CONN = 5
const MAX_DB_LIFETIME = 5 * time.Minute

type AppConfig struct {
	DatabaseConnections map[string]*sql.DB
	FileStorage
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
	dbCredentials := loadDbCreds()
	for _, dbCredential := range dbCredentials {
		url := credentialsToUrl(dbCredential)
		dbConns[DEV_CLIENT_TOKEN], err = connectSQL(url)
		if err != nil { log.Fatal("Can't connect to DB") }
	}
	portNumber := os.Getenv("PORT_NUMBER")
	if portNumber == "" {
		portNumber = fmt.Sprintf("%v", DEFAULT_PORT_NUMBER)
	}
	// return &AppConfig { DatabaseConnections: dbConns, PortNumber: fmt.Sprintf(":%v", portNumber), }
	return &AppConfig { DatabaseConnections: dbConns, PortNumber: fmt.Sprintf(":%v", portNumber), }
}

func connectSQL(databaseUrl string) (*sql.DB, error) {
	db, err := newDatabase(databaseUrl)
	if err != nil { panic(err) }

	db.SetMaxOpenConns(MAX_OPEN_DB_CONN)
	db.SetMaxIdleConns(MAX_IDLE_DB_CONN)
	db.SetConnMaxLifetime(MAX_DB_LIFETIME)
	// dbConn.SQL = db
	// return dbConn, nil
	return db, nil
}

func newDatabase(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil { return nil, err }
	if err = db.Ping(); err != nil { return nil, err }

	return db, nil
}

func credentialsToUrl(c DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", c.Host, c.Port, c.Database, c.Username, c.Password)
}

// TODO: Add multiple Databases creds, so this service could connect to several remote databases not just local main database

func loadDbCreds() []DatabaseConfig {
	// TODO: Load creds from config file
	dbCredentials := make([]DatabaseConfig, 1)
	dbCredentials[0] = DatabaseConfig {
		Adapter: "postgres",
		Host: "localhost",
		Port: "5432",
		Database: "truligent_development",
		Username: "truligent",
		Password: "truligent",
	}
	return dbCredentials
}

func readYaml() {
	var config DatabaseConfig
	// EXAMPLE
	yfile, err := ioutil.ReadFile("config/database.yaml")
	if err != nil { log.Fatal(err) }
	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	config = data[AppConfig.Environment]
	// for k, v := range data {
	// 	fmt.Printf("%s -> %d\n", k, v)
	// }
}
