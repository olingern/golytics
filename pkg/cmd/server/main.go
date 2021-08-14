package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olingern/golytics/pkg/controllers"
	"github.com/olingern/golytics/pkg/db/clients"
)

type Log struct {
	Id        int
	Ip        string
	Ua        string
	Lang      string
	Date      string
	Processed int
}

func checkForNewLogs(db *sqlx.DB) {
	rr := []Log{}

	for {
		err := db.Select(&rr, "select * from logs")

		if err != nil {
			fmt.Println("Error checking for new records")
			fmt.Println(err)
		}

		fmt.Printf("[checkForNewLogs] - Found id: %s\n", rr[0].Ua)
		time.Sleep(10 * time.Second)
	}
}

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	// now := time.Now().Format(time.RFC3339)
	// ip := r.Header.Get("X-FORWARDED-FOR")
	// ua := r.Header.Get("user-agent")
	// lang := r.Header.Get("accept-language")

	// _, err := db.Exec(fmt.Sprintf("insert into logs(ip, ua, lang, date, processed) values('%s', '%s', '%s', '%s', '%d')", ip, ua, lang, now, 0))
	// if err != nil {
	// 	fmt.Println("Didn't insert")
	// 	fmt.Println(err)
	// }

	// c.Status(200)
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbClient, err := sqlx.Open("sqlite3", "./db/database.sqlite3")

	if err != nil {
		log.Fatal("Failed to open sqlite database")
	}

	defer dbClient.Close()

	sqlClient, err := clients.NewSqliteClient(dbClient)

	if err != nil {
		log.Fatal("Failed to create sql client")
	}

	ctrl := controllers.NewController(sqlClient)

	r := mux.NewRouter()

	// go checkForNewLogs(db)

	r.HandleFunc("/api/users/login", ctrl.LoginHandler)

	// r.GET("/api/track/view", func(c *gin.Context) {

	// })

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	// r.Run()
}
