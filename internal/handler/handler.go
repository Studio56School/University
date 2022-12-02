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

type IHandler interface {
	GetStudents(c echo.Context) error
}

func GetStudents(c echo.Context) error {
	students, _ := s.r.Allstudents()
	return c.JSON(http.StatusOK, students)
}

func (h *Handler) InitRoutes(e) {
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", GetStudents)
}
