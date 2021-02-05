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
		logrus.WithError(err).Warning("Error loading .env file")
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
		logrus.WithError(err).Fatal("Can't connect to database")
	}

	a := api.NewAPI(db)
	server, err := routes.Routes(a)
	if err != nil {
		panic(err)
	}

	logrus.Info("Starting rest server")
	logrus.Error(http.ListenAndServe(":8080", server))
}
