package handlers

import (
	"banking-app/internal/domain"
	"banking-app/internal/services"
	"errors"
	"math"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("web/templates/balance.html"))

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	// Gets the account type
	params := r.URL.Query()
	bankAccountType := domain.BankAccountType(params.Get("bankAccountType"))
	if !bankAccountType.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Gets cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Fail to get cookie", http.StatusInternalServerError)
		return
	}
	// Get the account balance from database
	sessionID := cookie.Value
	balanceCents, err := services.GetAccountBalance(sessionID, bankAccountType)
	if err != nil {
		http.Error(w, "Fail to get balance", http.StatusInternalServerError)
		return
	}

	// Create a struc
	accountBalance := domain.AccountBalance{
		BankAccountType: bankAccountType,
		Balance:         (float64(balanceCents) / 100),
	}
	// Uses the struc to fill in a html template and return it as reponse
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = tmpl.Execute(w, accountBalance); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	transferRequest, err := getTransferData(r)
	if err != nil {
		writeError(w, err)
		return
	}

	if err = services.CanTransfer(transferRequest); err != nil {
		writeError(w, err)
		return
	}

	if err = services.MakeTransfer(transferRequest); err != nil {
		writeError(w, err)
		return
	}
	w.Write([]byte(`<div>Transfer sucessful!</div>`))
}

// Get data from reqeust and return a struc of data
func getTransferData(r *http.Request) (domain.TransferRequest, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return domain.TransferRequest{}, errors.New("Fail to get cookie")
	}

	err = r.ParseForm()
	if err != nil {
		return domain.TransferRequest{}, errors.New("Parse error")
	}

	accountFrom := domain.BankAccountType(r.FormValue("accountFrom"))
	accountTo := domain.BankAccountType(r.FormValue("accountTo"))

	transferAmountStr := r.FormValue("transferAmount")

	transferAmount, err := strconv.ParseFloat(transferAmountStr, 64)
	if err != nil {
		return domain.TransferRequest{}, errors.New("Invalid transfer amount")
	}
	AmountCents := int(math.Round(transferAmount * 100))

	return domain.TransferRequest{
		SessionID:   cookie.Value,
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		AmountCents: AmountCents,
	}, nil

}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("<div>" + err.Error() + "</div>"))
}

func writeSuccess(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("<div>" + message + "</div>"))
}
