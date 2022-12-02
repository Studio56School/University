package storage

import (
	"context"
	"fmt"
	"github.com/Studio56School/university/internal/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"log"
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
	StudentbyID(conn *pgx.Conn, id int) error
	Allstudents() ([]model.Student, error)
	AddNewStudent(student model.Student) error
}

func (r *Repo) StudentbyID(id int) error {
	query := `select id, name, surname from students where id = $1 `
	var name, surname string
	err := r.DB.QueryRow(context.Background(), query, id).Scan(&id, &name, &surname)
	if err != nil {
		r.l.Sugar().Error(fmt.Sprintf("Не отработался запрос студентам по id: %s", err))
		return err
	}

	r.l.Sugar().Error(fmt.Sprintf("id : %d, Name: %s, Surname: %s\n", id, name, surname))
	return nil
}

func (r *Repo) Allstudents() ([]model.Student, error) {

	var student model.Student
	students := make([]model.Student, 0)
	query := `select id, name, surname, gender from students`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Gender)
		if err != nil {
			log.Println(err)
		}
		students = append(students, student)
	}

	defer rows.Close()
	return students, nil
}

func (r *Repo) AddNewStudent(student model.Student) error {
	query := `INSERT INTO public.students
	(name, surname, gender)
	VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(context.Background(), query, student.Name, student.Surname, student.Gender)
	if err != nil {
		log.Println(err)
	}

	log.Printf("student with Name: %s, Surname: %s inserted  \n", student.Name, student.Surname)
	return nil
}

func UpdateStudentSurname(conn *pgx.Conn, surname string, newSurname string) error {
	query := `UPDATE public.students
	SET surname=$2
	WHERE surname = $1;`
	_, err := conn.Exec(context.Background(), query, surname, newSurname)
	if err != nil {
		log.Println(err)
	}

	log.Printf("student with Surname: %s Updated to %s\n", surname, newSurname)

	return nil
}

func DeleteStudentbyId(conn *pgx.Conn, id int) error {
	query := `DELETE FROM public.students
	WHERE id = $1
	RETURNING name, surname`
	var name, surname string
	err := conn.QueryRow(context.Background(), query, id).Scan(&name, &surname)
	if err != nil {
		log.Println(err)
		return err
	}

	if err != nil {
		log.Println(err)
	}

	log.Printf("Student %s %s has been deleted", name, surname)

	return nil
}
