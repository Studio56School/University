package service

import (
	"context"
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/model"
	"github.com/Studio56School/university/internal/storage"
	"go.uber.org/zap"
)

type IService interface {
	AllStudentsService() (student []model.Student, err error)
	StudentByID(id int) (student model.Student, err error)
	DeleteStudentById(id int) (err error)
	UpdateStudent(student *model.Student) (err error)
	AddNewStudent(student model.Student) (id int, err error)
	CreateNewProfessorService(professor *model.CreateProfessorRequest) (*model.CreateProfessorResponse, error)
}

type Service struct {
	conf   *config.Config
	logger *zap.Logger
	urepo  *storage.Repo
}

func NewService(conf *config.Config, logger *zap.Logger, urepo *storage.Repo) *Service {
	return &Service{conf: conf, logger: logger, urepo: urepo}
}

func (s *Service) AllStudentsService() (student []model.Student, err error) {

	student, err = s.urepo.AllStudents(context.Background())

	return student, err
}

func (s *Service) StudentByID(id int) (student model.Student, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	student, err = s.urepo.StudentByID(ctx, id)

	return student, err
}

func (s *Service) DeleteStudentById(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	err = s.urepo.DeleteStudentById(ctx, id)

	return err
}

func (s *Service) UpdateStudent(student *model.Student) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	err = s.urepo.UpdateStudent(ctx, student)

	return err
}

func (s *Service) AddNewStudent(student model.Student) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()

	id, err = s.urepo.AddNewStudent(ctx, student)

	return id, err
}

func (s *Service) CreateNewProfessorService(professor *model.CreateProfessorRequest) (*model.CreateProfessorResponse, error) {
	res, err := s.urepo.CreateProfessor(professor)
	if err != nil {
		s.logger.Sugar().Error(err)
		return nil, err
	}
	return res, nil
}
