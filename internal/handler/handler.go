package handler

import (
	"context"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	repo *storage.Repo
}

func NewHandler(repo *storage.Repo) *Handler {
	return &Handler{repo: repo}
}

type IHandler interface {
	GetStudents(c echo.Context) error
	GetStudentsById(c echo.Context) error
	CreateStudent(c echo.Context) error
	DeleteStudent(c echo.Context) error
}

func (h *Handler) GetStudents(c echo.Context) error {
	students, err := h.repo.AllStudents()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, students)
}

func (h *Handler) GetStudentsById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	student, err := h.repo.StudentByID(context.Background(), id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) CreateStudent(c echo.Context) error {
	var request model.Student
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	student, err := h.repo.AddNewStudent(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) updateStudent(c echo.Context) error {
	var request model.Student
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	student, err := h.repo.UpdateStudent(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	student, err := h.repo.DeleteStudentById(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	e.GET("/students", h.GetStudents)
	e.GET("/students/:id", h.GetStudentsById)
	e.POST("/students/create", h.CreateStudent)
	//e.PUT("/users/update", h.updateStudent)
	e.DELETE("/students/:id", h.DeleteStudent)
}
