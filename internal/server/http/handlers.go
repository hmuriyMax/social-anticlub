package http

import (
	"github.com/hmuriyMax/social-anticlub/internal/pb/user_service"
	"net/http"
)

type userService struct {
	user_service.UserServiceServer
}

func (userService) get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func (userService) register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func (userService) search(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

type loginService struct{}

func (loginService) login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
