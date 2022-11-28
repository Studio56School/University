package main

import (
	"github.com/Studio56School/University/internal/storage"
	"log"
)

func main() {
	db, err := storage.ConnectDB()

	if err != nil {
		log.Fatalf("cannot initialize db ")
	}

	err = storage.StudentbyID(db, 1)
	if err != nil {
		log.Println(err)
	}

}
