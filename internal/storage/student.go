package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

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

func AllStudents(conn *pgx.Conn) error {
	query := `select id, name, surname from students`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		return err
	}

	for rows.Next() {
		var id int
		var name, surname string
		err := rows.Scan(&id, &name, &surname)
		if err != nil {
			//log.Println("error while scanning ")
			log.Println(err)
		}

		fmt.Printf("id %d, Name: %s, Surname: %s\n", id, name, surname)
	}

	defer rows.Close()
	return nil
}

func AddNewStudent(conn *pgx.Conn, name string, surname string, gender string) error {
	query := `INSERT INTO public.students
	(name, surname, gender)
	VALUES ($1, $2, $3)`

	_, err := conn.Exec(context.Background(), query, name, surname, gender)
	if err != nil {
		log.Println(err)
	}

	log.Printf("student with Name: %s, Surname: %s inserted  \n", name, surname)
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
