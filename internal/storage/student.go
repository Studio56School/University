package storage

import (
	"context"
	"fmt"
	"github.com/Studio56School/university/internal/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewRepository(log zap.Logger) (*Repo, error) {
	pgDB, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	return &Repo{l: log, DB: pgDB}, nil
}

type Repo struct {
	DB *pgx.Conn
	l  zap.Logger
}

type IRepository interface {
	StudentByID(id int) (student model.Student, err error)
	AllStudents() (students []model.Student, err error)
	DeleteStudentById(ctx context.Context, id int) (err error)
	UpdateStudent(ctx context.Context, student model.Student) (newStudent model.Student, err error)
}

func (r *Repo) StudentByID(ctx context.Context, id int) (student model.Student, err error) {

	query := `select id, name, surname, gender from students where id = $1 `
	err = r.DB.QueryRow(ctx, query, id).Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)

	if err != nil {
		r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return student, err
	}

	return student, err
}

func (r *Repo) AllStudents() (students []model.Student, err error) {

	students = make([]model.Student, 0)
	query := `select id, name, surname, gender from students`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return nil, err
	}
	var student model.Student

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)
		if err != nil {
			r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
			return nil, err
		}

		students = append(students, student)
	}

	defer rows.Close()
	return students, nil
}

func (r *Repo) AddNewStudent(ctx context.Context, student model.Student) (id int, err error) {
	query := `INSERT INTO public.students
	(name, surname, gender)
	VALUES ($1, $2, $3) RETURNING id`

	err = r.DB.QueryRow(ctx, query, student.Name, student.Surname, student.Gender).Scan(&id)
	if err != nil {
		r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return -1, err
	}

	return id, nil
}

func (r *Repo) UpdateStudent(ctx context.Context, student model.Student, id int) (model.Student, error) {
	query := `UPDATE public.students
	SET name=$2, surname = $3, gender = $4 
	WHERE id = $1;`
	err := r.DB.QueryRow(ctx, query, id, student.Name, student.Surname, student.Gender).Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)
	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return student, err
	}

	return student, err
}

func (r *Repo) DeleteStudentById(ctx context.Context, id int) (int int, err error) {
	query := `DELETE FROM students_by_group WHERE student_id = $1`
	query2 := `	DELETE FROM students WHERE id = $1`

	_, err = r.DB.Exec(ctx, query, id)
	_, err = r.DB.Exec(ctx, query2, id)
	if err != nil {
		//r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return -1, err
	}

	return id, err
}
