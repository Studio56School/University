package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"

	"log"
)

func ConnectDB() (*pgx.Conn, error) {

	viper.AddConfigPath("./heml/")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	connString := "postgres://" +
		viper.GetString("db.username") + ":" + viper.GetString("db.password") +
		"@" + viper.GetString("db.host") + ":" + viper.GetString("db.port") +
		"/" + viper.GetString("db.name_db")

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	return conn, nil

}
