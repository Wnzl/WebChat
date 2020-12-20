package http

import (
	"fmt"
	"github.com/Wnzl/webchat/storage"
	"github.com/go-chi/chi"
	"github.com/tarent/logrus"
	"net/http"
)

type Server struct {
	Storage *storage.PostgreSqlStorage
	Port    int
}

func (s *Server) Start() error {
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	logrus.Info(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), router))
	return nil
}
