package main

import (
	"github.com/Wnzl/webchat/http"
	"github.com/Wnzl/webchat/storage"
	"github.com/joho/godotenv"
	"github.com/tarent/logrus"
	"os"
	"strconv"
)

const (
	dbUsernameEnv = "DATABASE_USERNAME"
	dbPasswordEnv = "DATABASE_PASSWORD"
	dbNameEnv     = "DATABASE_NAME"
	dbHostEnv     = "DATABASE_HOST"
	dbPortEnv     = "DATABASE_PORT"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv(dbPortEnv))
	if err != nil {
		logrus.WithError(err).Fatal("Converting port to int")
	}

	postgresql, err := storage.NewPostgreSqlStorage(storage.Config{
		Username:     os.Getenv(dbUsernameEnv),
		Password:     os.Getenv(dbPasswordEnv),
		DatabaseName: os.Getenv(dbNameEnv),
		Host:         os.Getenv(dbHostEnv),
		Port:         port,
	})
	if err != nil {
		logrus.WithError(err).Fatal("Storage driver initializing")
	}

	server := http.Server{Storage: postgresql, Port: 8080}

	logrus.Info("Starting rest server")
	logrus.WithError(server.Start()).Fatal("Rest server can't start")
}
