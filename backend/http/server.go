package http

import (
	"fmt"
	"github.com/Wnzl/webchat/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/tarent/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Server struct {
	Storage *gorm.DB
	Port    int
}

func (s *Server) Start() error {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	jwtAuth := jwtauth.New(jwa.HS256.String(), []byte(os.Getenv("JWT_SECRET")), nil)
	auth := controllers.NewAuthController(s.Storage, jwtAuth)
	r.Mount("/", auth.Routes())

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtAuth))
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
