package handler

import (
	"context"
	_ "github.com/Studio56School/university/docs"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/service"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	repo *storage.Repo
	svc  service.IService
	log  *zap.Logger
}

type IHandler interface {
	GetStudents(c echo.Context) error
	GetStudentsById(c echo.Context) error
	CreateStudent(c echo.Context) error
	DeleteStudent(c echo.Context) error
}

func NewHandler(svc service.IService, logger *zap.Logger) *Handler {
	return &Handler{log: logger, svc: svc}
}

//	@Summary		GetStudents
//	@Description	Get all students
//	@Tags			students
//	@ID				get-student
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.student
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/v1/students [get]

func (h *Handler) GetStudents(c echo.Context) error {

	students, err := h.svc.AllStudentsService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, students)
}

//	@Summary		GetStudentsById
//	@Description	Get student by id
//	@Tags			students
//	@ID				get-student
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.student
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/v1/students/{id} [get]

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

//	@Summary		CreateStudent
//	@Description	Create Student
//	@Tags			students
//	@ID				create-student
//	@Accept			json
//	@Produce		json
//	@Param input body model.student
//	@Success		200	{object}	model.student
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/v1/students/create [post]

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

//	@Summary		DeleteStudent
//	@Description	Delete Student
//	@Tags			students
//	@ID				delete-student
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.student
//	@Failure		400	{object}	errorResponse
//	@Failure		401	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/v1/students/{id} [delete]

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
