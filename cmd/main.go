package main

import (
	"github.com/Studio56School/university/internal/server"
	"github.com/labstack/echo/v4"
)

func main() {
	//db, err := storage.ConnectDB()
	//
	//if err != nil {
	//	log.Fatalf("cannot initialize db ")
	//}

	//   Перенести логику сервера эхо

	e := echo.New()
	e = server.RoutesGetStudents()
	e.Logger.Fatal(e.Start(":1323"))

	//err = storage.StudentbyID(db, 1)
	//if err != nil {
	//	log.Println(err)
	//}

}
