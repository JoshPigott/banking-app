package handlers

import "net/http"

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/views/login.html")
}

func GetOnlineBankingPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/views/online-banking.html")
}

func GetTransferPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/views/transfer.html")
}

func GetPaymentPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/views/payment.html")
}
