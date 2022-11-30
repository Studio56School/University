package main

import (
	"github.com/Studio56School/university/internal/storage"
	"log"
)

func main() {
	db, err := storage.ConnectDB()

	if err != nil {
		log.Fatalf("cannot initialize db ")
	}

	//   Перенести логику сервера эхо

	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, storage.AllStudents(c, db))
	//})
	//e.Logger.Fatal(e.Start(":1323"))

	//
	err = storage.AllStudents(db)
	if err != nil {
		log.Println(err)
	}

}
