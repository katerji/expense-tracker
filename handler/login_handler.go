package handler

import (
	"net/http"
)

const LoginRoute = "/auth/login"

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	return
}
