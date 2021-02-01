package main

import (
	"fmt"
	"github.com/Wnzl/webchat/http"
	"github.com/Wnzl/webchat/models"
	"github.com/joho/godotenv"
	"github.com/tarent/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	server := http.Server{Storage: db, Port: 8080}

	logrus.Info("Starting rest server")
	logrus.WithError(server.Start()).Fatal("Rest server can't start")
}
