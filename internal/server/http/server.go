package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RouteHandlers(opts ...Option) chi.Router {
	var (
		router = chi.NewRouter()
		userS  = userService{}
		loginS = loginService{}
	)

	router.Use(middleware.CleanPath)
	for _, opt := range opts {
		opt(router)
	}

	router.Route("/user", func(r chi.Router) {
		r.Get("/{nickname}", userS.get)
		r.Post("/register", userS.register)
		r.Get("/search", userS.search)
	})
	router.Post("/login", loginS.login)

	return router
}
