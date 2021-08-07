package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
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

func main() {

	db, err := sqlx.Open("sqlite3", "./db/database.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.StaticFile("/admin", "./admin/public/index.html")
	r.StaticFile("/global.css", "./admin/public/global.css")
	r.Static("/build", "./admin/public/build")

	go checkForNewLogs(db)

	r.GET("/api/dashboard", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/api/logs", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/api/track/view", func(c *gin.Context) {
		lang := c.Request.Header.Get("Accept-Language")
		now := time.Now().Format(time.RFC3339)

		_, err := db.Exec(fmt.Sprintf("insert into logs(ip, ua, lang, date, processed) values('%s', '%s', '%s', '%s', '%d')", c.ClientIP(), c.Request.UserAgent(), lang, now, 0))
		if err != nil {
			fmt.Println("Didn't insert")
			fmt.Println(err)
		}

		c.Status(200)
	})

	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
