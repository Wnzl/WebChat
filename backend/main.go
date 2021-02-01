package main

import (
	"fmt"
	"github.com/Wnzl/webchat/api"
	"github.com/Wnzl/webchat/models"
	"github.com/Wnzl/webchat/routes"
	"github.com/joho/godotenv"
	"github.com/tarent/logrus"
	"net/http"
	"os"
)

// @title WebChat API
// @version 1.0
// @description WebChat golang backend server.
// @host localhost:8080
func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	db, err := models.NewPostgresDB(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		panic(err)
	}

	a := api.NewAPI(db)
	server, err := routes.Routes(a)

	logrus.Info("Starting rest server")
	logrus.Error(http.ListenAndServe(":8080", server))
}
