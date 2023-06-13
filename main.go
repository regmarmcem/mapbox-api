package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx"
	"github.com/joho/godotenv"

	"github.com/regmarmcem/mapbox-api/api"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(".env file not found")
	}

	dbHost := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")

	conn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic("failed to connect database")
	}

	r := api.NewRouter(db)
	fmt.Println("web server starting...")

	log.Panic(http.ListenAndServe("localhost:8080", r))
}
