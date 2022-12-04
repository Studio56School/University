package main

import (
	"github.com/Studio56School/university/internal/handler"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {

	//   Перенести логику сервера эхо
	log := zap.Logger{}
	repos, err := storage.NewRepository(log)
	if err != nil {
		panic(err)
	}

	handlers := handler.NewHandler(repos)

	e := echo.New()
	handlers.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))

}
