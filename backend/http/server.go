package http

import (
	"fmt"
	"github.com/Wnzl/webchat/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/tarent/logrus"
	"net/http"
)

type Server struct {
	Storage *storage.PostgreSqlStorage
	Port    int
}

func (s *Server) Start() error {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	logrus.Info(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), router))
	return nil
}
