package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/katerji/expense-tracker/handler"
	"net/http"
)

func InitWebServer() {
	r := chi.NewRouter()

	r.Use(AuthMiddleware)
	r.Post(handler.LoginRoute, handler.LoginHandler)
	r.Post(handler.RecordTransactionRoute, handler.RecordTransactionHandler)

	http.ListenAndServe(":3000", r)
}
