package routes

import (
	"github.com/Wnzl/webchat/api"
	"github.com/Wnzl/webchat/auth"
	_ "github.com/Wnzl/webchat/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func Routes(api *api.API) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/login", api.Users.UserLogin)
	r.Post("/signup", api.Users.UserSignup)

	// private routes that requires to be authorized
	r.Group(func(r chi.Router) {
		r.Use(auth.JwtMiddleware.Handler)
		r.Get("/info", api.Users.UserInfo)
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r, nil
}
