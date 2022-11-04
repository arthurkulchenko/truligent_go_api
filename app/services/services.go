package services

import(
	"net/http"
	"github.com/labstack/echo/v4"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4"
)

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

var client *Client

type Client struct {
	dbConn *sql.DB
}

func OmsPing(c echo.Context) error {
	client := Client { dbConn: AppConfig.DatabaseConnections[c.QueryParam("token")] }
	return c.String(http.StatusOK, client.OmsPingCall()
}

func connectSQL(durl string) (*DB, error) {
	d, err := newDatabase(durl)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL = d
	// err = testDB(d)
	// if err != nil {
	// 	return nil, error
	// }
	return dbConn, nil
}

func newDatabase(durl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", durl)
	if err != nil { return nil, err }
	if err = db.Ping(); err != nil { return nil, err }

	return db, nil
}
