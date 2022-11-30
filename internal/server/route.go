package server

import (
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
)

//locahost:8080/StudentById

func RoutesGetStudents() *echo.Echo {
	router := echo.New()
	router.GET("/students", storage.GetStudents)
	return router
}
