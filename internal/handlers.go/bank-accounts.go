package handlers

import (
	"banking-app/internal/domain"
	"banking-app/internal/services"
	"fmt"
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

// func getTransferData(r *http.Request) (string, domain.BankAccountType, domain.BankAccountType, ){
// 	cookie, err := r.Cookie("session_id")
// 	if err != nil {
// 		return 0, domain.BankAccountType{}, domain.BankAccountType{}, err
// 	}
// 	if err = r.ParseForm(); err != nil {
// 		return err
// 	}
// 	accountFrom := domain.BankAccountType(r.FormValue("accountFrom"))
// 	accountTo := domain.BankAccountType(r.FormValue("accountTo"))
// 	transferAmountStr := r.FormValue("transferAmount")
// 	return transferAmountStr, accountFrom, accountTo, cookie.Value, err
// }

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Fail to get cookie", http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	accountFrom := domain.BankAccountType(r.FormValue("accountFrom"))
	accountTo := domain.BankAccountType(r.FormValue("accountTo"))

	transferAmountStr := r.FormValue("transferAmount")
	transferAmount, err := strconv.ParseFloat(transferAmountStr, 64)
	transferAmountCents := int(transferAmount * 100)
	fmt.Println("transferAmountCents:", transferAmountCents)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`<div>Invalid transfer amount</div>`))
		return
	}

	if !accountFrom.CanWithdraw() || !accountTo.IsValid() || accountFrom == accountTo {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`<div>Invalid to and from accounts</div>`))
		return
	}
	if !services.IsValidTransferAmount(transferAmountCents, accountFrom, cookie.Value) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`<div>Invalid transfer amount</div>`))
		return
	}
	if err = services.MakeTransfer(transferAmountCents, accountFrom, accountTo, cookie.Value); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`<div>Transfer failed</div>`))
		return
	}
	w.Write([]byte(`<div>Transfer sucessful!</div>`))
}

// IsValidTransferAmount 👍
// If not send back bad request 👍
// If yes make the transfer
// If send back transfer sucess
// Then clean up function I feel like it is dry and go like
