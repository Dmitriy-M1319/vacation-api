package database

import (
	"fmt"

	"github.com/jackc/pgx"
)

var currentDb string

func NewConnection(ip, user, password, database string) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, password, ip, database)
	config, err := pgx.ParseConnectionString(connString)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials for database %s", database)
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %s: %s", database, err.Error())
	}
	currentDb = database
	return conn, nil
}

func Close(conn *pgx.Conn) error {
	err := conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close database %s: %s", currentDb, err.Error())
	}
	return nil
}
