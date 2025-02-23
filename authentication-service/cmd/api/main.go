package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB    *sql.DB
	Model data.Models
}

func main() {
	log.Println("Starting auth service")

	// connect to db
	conn := connectDB()
	if conn == nil {
		log.Panic("Could not connect to postgres!")
	}
	log.Println("Connected to postgres")
	// create a web server
	app := Config{
		DB:    conn,
		Model: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the web server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB() *sql.DB {
	dsn := os.Getenv("DSN")

	var count int8
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres is not ready yet..")
			count++
		} else {
			return conn
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
