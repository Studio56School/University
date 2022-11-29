package main

import (
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	//   Перенести логику сервера эхо
	e := echo.New()

	e.Use(echo.MiddlewareFunc())

	db, err := storage.ConnectDB()

	if err != nil {
		log.Fatalf("cannot initialize db ")
	}

	err = storage.StudentbyID(db, 1)
	if err != nil {
		log.Println(err)
	}

}
