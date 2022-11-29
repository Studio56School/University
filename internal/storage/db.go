package storage

import (
	"context"
	"github.com/jackc/pgx/v5"

	"log"
)
viper
func ConnectDB() (*pgx.Conn, error) {
	connString := "postgres://dias:1234@127.0.0.1:5432/postgres"
	//connString  := "postgres:/" + viper.GetString()

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	return conn, nil

}
