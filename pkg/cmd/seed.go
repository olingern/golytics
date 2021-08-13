package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olingern/golytics/pkg/auth"
)

const insertAdmin = "INSERT INTO users (username, hashed_password) VALUES (?, ?)"

var err error

func main() {

	fmt.Println("Starting seed ...")

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqlx.Open("sqlite3", "./db/database.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	prepared, err := db.Prepare(insertAdmin)

	if err != nil {
		log.Fatal(err)
	}

	hashedPwd, err := auth.HashAndSalt([]byte("kagome"))

	if err != nil {
		log.Fatal(err)
	}

	_, err = prepared.Exec("admin", hashedPwd)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
}
