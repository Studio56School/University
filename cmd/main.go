package main

import (
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/server"
	"go.uber.org/zap"
	"log"
	"time"
)

func main() {
	var err error
	time.Local, err = time.LoadLocation("Asia/Almaty")
	if err != nil {
		log.Printf("error loading '%s': %v\n", time.Local, err)
	}

	logger, _ := zap.NewProduction()

	//logger = logger.NewConsoleLogger(logger.INFO, logger.JSON)

	conf, err := config.NewAppConfig()

	if err != nil {
		log.Fatal("[app] Ошибка при инициализации конфигурации приложения: ", err)
	}

	httpServer, err := server.NewServer(conf, logger)

	if err != nil {
		log.Fatal("Ошибка при инициализации http сервера: ", err)
	}
	err = httpServer.RunBlocking()
	if err != nil {
		log.Fatal("Ошибка при запуске http сервера: ", err)
	}
}
