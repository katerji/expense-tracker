package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/katerji/expense-tracker/handler"
)

func InitWebServer() {
	r := chi.NewRouter()
	r.Use(AuthMiddleware)
	r.Post(handler.LoginRoute, handler.LoginHandler)
}
