package main

import (
	"github.com/Wnzl/webchat/http"
	"github.com/tarent/logrus"
)

func main() {
	server := http.Server{Port: 8080}

	logrus.Info("Starting rest server")
	logrus.WithError(server.Start()).Fatal("Rest server can't start")
}
