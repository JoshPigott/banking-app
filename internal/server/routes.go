package server

import (
	"banking-app/internal/handlers.go"
	"banking-app/internal/middleware"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/account-balance", middleware.RequireAuth(handlers.GetAccountBalance))
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/sign-up", handlers.SignUp)
	return mux
}
