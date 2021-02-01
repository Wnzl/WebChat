package http

import (
	"fmt"
	"github.com/Wnzl/webchat/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/tarent/logrus"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	Storage *gorm.DB
	Port    int
}

func (s *Server) Start() error {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	uc := controllers.NewUsersController(s.Storage)
	r.Post("/login", uc.Login)
	r.Post("/signup", uc.Signup)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(uc.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			_, claims, _ := jwtauth.FromContext(r.Context())
			_, _ = w.Write([]byte(fmt.Sprintf("pong, user: %v", claims["user_id"])))
		})
	})

	logrus.Info(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), r))
	return nil
}
