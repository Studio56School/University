package main

import (
	"github.com/Studio56School/university/internal/handler"
	"github.com/Studio56School/university/internal/server"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {

	//   Перенести логику сервера эхо
	log := zap.Logger{}
	repos := storage.NewRepository(log)
	handlers := handler.NewHandler(repos)

	e := echo.New()
	e = server.RoutesGetStudents()
	e.Logger.Fatal(e.Start(":1323"))
	handlers.InitRoutes(e)

	//err = storage.StudentbyID(db, 1)
	//if err != nil {
	//	log.Println(err)
	//}

}
