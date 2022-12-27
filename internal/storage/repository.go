package storage

import "github.com/Studio56School/university/internal/model"

type IRepository interface {
	AllStudents() (student []model.Student, err error)
	StudentByID(id int) (student model.Student, err error)
	DeleteStudentById(id int) (err error)
	UpdateStudent(student model.Student, id int) (err error)
	AddNewStudent(student model.Student) (id int, err error)
}
