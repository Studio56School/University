package handler

import (
	_ "github.com/Studio56School/university/docs"
	"github.com/Studio56School/university/internal/service"
	"github.com/Studio56School/university/internal/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	repo *storage.Repo
	svc  service.IService
	log  *zap.Logger
}

type IHandler interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
}

func NewHandler(svc service.IService, logger *zap.Logger) *Handler {
	return &Handler{log: logger, svc: svc}
}
