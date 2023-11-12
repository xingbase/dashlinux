package http

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type Service struct {
	Jwt *jwtauth.JWTAuth
}

func (s *Service) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
