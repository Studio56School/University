package service

import (
	"context"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/storage"
)

type IService interface {
	AllStudentsService(ctx context.Context) (student model.Student, err error)
}

type Service struct {
	conf   *config.Config
	logger *logger.Logger
	urepo  *storage.Repo
}

func NewService(conf *config.Config, logger *logger.Logger, urepo *storage.Repo) IService {
	return &Service{
		conf:   conf,
		logger: logger,
		urepo:  urepo,
	}
}

func (s *Service) AllStudentsService(ctx context.Context) (student []model.Student, err error) {
	student, err = s.urepo.AllStudents(ctx)

	return student, err
}
