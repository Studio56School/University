package handler

import (
	"context"
	"github.com/Studio56School/university/internal/logger"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/service"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	repo *storage.Repo
	svc  service.IService
	log  *logger.Logger
}

type IHandler interface {
	GetStudents(c echo.Context) error
	GetStudentsById(c echo.Context) error
	CreateStudent(c echo.Context) error
	DeleteStudent(c echo.Context) error
}

func NewHandler(svc service.IService, logger *logger.Logger) *Handler {
	return &Handler{log: logger, svc: svc}
}

func (h *Handler) GetStudents(c echo.Context) error {

	students, err := h.svc.AllStudentsService(c.Request().Context())
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

func (h *Handler) UpdateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	var request model.Student
	err = c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = h.repo.UpdateStudent(context.Background(), request, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, err)
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
