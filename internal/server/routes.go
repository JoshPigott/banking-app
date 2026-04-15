package server

import (
	"banking-app/internal/handlers.go"
	"banking-app/internal/middleware"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create-session", handlers.CreateSession)
	mux.HandleFunc("/test", middleware.RequireAuth(handlers.Test))
	mux.HandleFunc("/login", handlers.Login)
	return mux
}
