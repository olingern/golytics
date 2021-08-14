package clients

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/olingern/golytics/pkg/auth"
)

type DBClient interface {
	ClientType() string
	LogUserIn() (bool, error)
}

type SqliteClient struct {
	clientType string
	dbClient   *sqlx.DB
}

func NewSqliteClient(dbClient *sqlx.DB) (*SqliteClient, error) {
	return &SqliteClient{
		clientType: "SQLite",
		dbClient:   dbClient,
	}, nil
}

func (c *SqliteClient) ClientType() string {
	return c.clientType
}

type user struct {
	id              int64
	username        string
	hashed_password string
}

func (c *SqliteClient) LogUserIn(username string, password string) (bool, error) {

	var hashedPassword string
	err := c.dbClient.QueryRow("SELECT hashed_password FROM users WHERE username = ?", username).Scan(&hashedPassword)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	fmt.Println(hashedPassword)

	result, err := auth.VerifyPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false, err
	}

	return result, nil
}
