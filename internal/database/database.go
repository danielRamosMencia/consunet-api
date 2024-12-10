package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	Connx *sql.DB
	once  sync.Once
	err   error
)

func ConnectDatabase() *sql.DB {
	once.Do(func() {
		host := os.Getenv("DB_HOST")
		port := 5432
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		Connx, err = sql.Open("postgres", psqlconn)
		if err != nil {
			log.Panic("Error connecting database", err)
		}

		err = Connx.Ping()
		if err != nil {
			log.Panic("Error reaching database", err)
		}

		log.Println("Database connected successfully")
	})

	return Connx
}
