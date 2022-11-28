package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func ConnectDB() (*pgx.Conn, error) {
	connString := "postgres://dias:1234@127.0.0.1:5432/postgres"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	return conn, nil

}
