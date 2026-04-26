package server

import (
	"banking-app/internal/handlers"
	"banking-app/internal/middleware"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/account-balance", middleware.RequireAuth(handlers.GetAccountBalance))
	mux.HandleFunc("/sign-up", handlers.SignUp)
	mux.HandleFunc("/login", routeLogin)
	mux.HandleFunc("/payment", middleware.RequireAuth(routePayment))
	mux.HandleFunc("/transfer", middleware.RequireAuth(routeTransfer))
	mux.HandleFunc("/online-banking", middleware.RequireAuth(handlers.GetOnlineBankingPage))
	// Html templates
	mux.HandleFunc("/get-welcome-message", middleware.RequireAuth(handlers.GetWelcomeMessage))
	// Static files
	mux.Handle("/static/", serveStaticFiles())
	return mux
}

func serveStaticFiles() http.Handler {
	dir := http.Dir("./web/assets/")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	return handler
}

// Route request by method

func routeLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.GetLoginPage(w, r)
	}
	if r.Method == "POST" {
		handlers.LoginAuth(w, r)
	}
}

func routeTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.GetTransferPage(w, r)
	}
	if r.Method == "POST" {
		handlers.TransferMoney(w, r)
	}
}

func routePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.GetPaymentPage(w, r)
	}
	if r.Method == "POST" {
		handlers.Payment(w, r)
	}
}
