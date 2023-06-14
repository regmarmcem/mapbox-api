package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
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

	databaseUrl := fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s&sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)

	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		panic(err)
	}

	r := api.NewRouter(db)
	fmt.Println("web server starting...")

	log.Panic(http.ListenAndServe(":8080", r))
}
