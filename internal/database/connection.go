package database

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var currentDb string

func NewConnection(ip, user, password, database string) (*sqlx.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, ip, database)
	conn, err := sqlx.Connect("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %s: %s", database, err.Error())
	}
	currentDb = database
	return conn, nil
}

func Close(conn *sqlx.DB) error {
	err := conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close database %s: %s", currentDb, err.Error())
	}
	return nil
}
