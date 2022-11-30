package storage

import (
	"context"
	"github.com/Studio56School/university/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var DB *pgx.Conn

func StudentbyID(conn *pgx.Conn, id int) error {
	query := `select id, name, surname from students where id = $1 `
	var name, surname string
	err := conn.QueryRow(context.Background(), query, id).Scan(&id, &name, &surname)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("id : %d, Name: %s, Surname: %s\n", id, name, surname)
	return nil
}

//func GetDBInstance() *pgx.Conn {
//	return DB
//}

func GetStudents(c echo.Context) error {
	students, _ := Allstudents()
	return c.JSON(http.StatusOK, students)
}

func Allstudents() ([]model.Student, error) {
	db, _ := ConnectDB()
	var student model.Student
	students := make([]model.Student, 0)
	query := `select id, name, surname, gender from students`
	rows, err := db.Query(context.Background(), query)
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

func AddNewStudent(conn *pgx.Conn, student model.Student) error {
	query := `INSERT INTO public.students
	(name, surname, gender)
	VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), query, student.Name, student.Surname, student.Gender)
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
