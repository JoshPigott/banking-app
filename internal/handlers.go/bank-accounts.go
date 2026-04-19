package handlers

import (
	"banking-app/internal/domain"
	"banking-app/internal/services"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("web/templates/balance.html"))

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	accountType := domain.AccountType(params.Get("accountType"))
	if !accountType.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Fail to get cookie", http.StatusInternalServerError)
		return
	}

	sessionID := cookie.Value
	balance, err := services.GetAccountBalance(sessionID, accountType)
	if err != nil {
		http.Error(w, "Fail to get balance", http.StatusInternalServerError)
		return
	}

	accountBalance := domain.AccountBalance{
		AccountType: accountType,
		Balance:     balance,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = tmpl.Execute(w, accountBalance); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}
