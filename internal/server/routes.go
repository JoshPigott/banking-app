package server

import (
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/login", auth.Login)
	// mux.HandleFunc("/sign-up", auth.Signup)
	return mux
}
