package main

import (
	"chat-go/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/chatgo", http.HandlerFunc(handlers.ChatgoEndpoint))

	return mux
}
