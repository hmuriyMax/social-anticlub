package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

type (
	Option func(p *chi.Mux)
)

func WithHandlersTimeout(timeout time.Duration) Option {
	return func(h *chi.Mux) {
		h.Use(middleware.Timeout(timeout))
	}
}
