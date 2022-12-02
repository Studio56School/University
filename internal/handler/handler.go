package handler

import (
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	repo *storage.Repo
}

func NewHandler(repo *storage.Repo) *Handler {
	return &Handler{repo: repo}
}

func GetStudents(c echo.Context) error {
	students, _ := Allstudents()
	return c.JSON(http.StatusOK, students)
}
