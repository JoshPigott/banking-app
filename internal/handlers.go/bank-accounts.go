package handlers

import (
	"banking-app/internal/database"
	"banking-app/internal/models"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("web/templates/balance.html"))

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	accountType := models.AccountType(params.Get("accountType"))
	if !accountType.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	sessionID := cookie.Value

	userID, err := database.GetUserID(sessionID)
	if err != nil {
		http.Error(w, "Fail to get userID", http.StatusInternalServerError)
		return
	}

	tableName := accountType.GetTableName()
	balance, err := database.GetAccountBalance(tableName, userID)
	if err != nil {
		http.Error(w, "Fail to get balance", http.StatusInternalServerError)
		return
	}

	accountBalance := models.AccountBalance{
		AccountType: accountType,
		Balance:     balance,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = tmpl.Execute(w, accountBalance); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}
