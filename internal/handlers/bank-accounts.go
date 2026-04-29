package handlers

import (
	"banking-app/internal/domain"
	"banking-app/internal/helpers"
	"errors"
	"fmt"
	"html"
	"math"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("web/templates/account.html"))

func GetAccount(w http.ResponseWriter, r *http.Request) {
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
	balanceCents, err := helpers.GetAccountBalance(sessionID, bankAccountType)
	if err != nil {
		http.Error(w, "Fail to get balance", http.StatusInternalServerError)
		return
	}
	// Create a struc
	account := domain.Account{
		BankAccountType: bankAccountType.GetFormatName(),
		Balance:         fmt.Sprintf("%.2f", float64(balanceCents)/100),
		ImageName:       bankAccountType.GetImageName(),
	}
	// Uses the struc to fill in a html template and return it as reponse
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = tmpl.Execute(w, account); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}

func Payment(w http.ResponseWriter, r *http.Request) {
	paymentRequest, err := getPaymentData(r)
	if err != nil {
		writeError(w, err)
		return
	}
	if err = helpers.IsValidPayment(&paymentRequest); err != nil {
		writeError(w, err)
		return
	}
	if err = helpers.MakePayment(&paymentRequest); err != nil {
		writeError(w, err)
		return
	}
	w.Write([]byte(`<p class="transaction__message--success">Payment sucessful!</p>`))
}

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	transferRequest, err := getTransferData(r)
	if err != nil {
		writeError(w, err)
		return
	}

	if err = helpers.CanTransfer(transferRequest); err != nil {
		writeError(w, err)
		return
	}

	if err = helpers.MakeTransfer(&transferRequest); err != nil {
		writeError(w, err)
		return
	}
	w.Write([]byte(`<p class="transaction__message--success">Transfer sucessful!</p>
	<input id="transfer__amount-input" name="amount" type="number" min="0" max="99999999999999"
	step="0.01" placeholder="Enter Amount done" autocomplete="off" hx-swap-oob="true">`))
}

// Gets username for to create a custom greeting for the user
func GetWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Gets cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte(""))
		return
	}
	username, err := helpers.GetUsername(cookie.Value)
	if err != nil {
		w.Write([]byte(""))
		return
	}
	w.Write(fmt.Appendf(nil, "Hello %s", html.EscapeString(username)))
}

// Gets data from reqeust and return a struc TransferRequest
func getPaymentData(r *http.Request) (domain.PaymentRequest, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return domain.PaymentRequest{}, errors.New("Fail to get cookie")
	}

	err = r.ParseForm()
	if err != nil {
		return domain.PaymentRequest{}, errors.New("Parse error")
	}
	accountFrom := domain.BankAccountType(r.FormValue("account-from"))
	receiverUsername := r.FormValue("receiver-username")

	amountCents, err := getAmount(r)
	if err != nil {
		return domain.PaymentRequest{}, err
	}
	return domain.PaymentRequest{
		SessionID:        cookie.Value,
		AccountFrom:      accountFrom,
		ReceiverUsername: receiverUsername,
		AmountCents:      amountCents,
	}, nil
}

// Gets data from reqeust and return a struc TransferRequest
func getTransferData(r *http.Request) (domain.TransferRequest, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return domain.TransferRequest{}, errors.New("Fail to get cookie")
	}

	err = r.ParseForm()
	if err != nil {
		return domain.TransferRequest{}, errors.New("Parse error")
	}

	accountFrom := domain.BankAccountType(r.FormValue("account-from"))
	accountTo := domain.BankAccountType(r.FormValue("account-to"))

	amountCents, err := getAmount(r)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	return domain.TransferRequest{
		SessionID:   cookie.Value,
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		AmountCents: amountCents,
	}, nil
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("<p class='transaction__message--error'>" + err.Error() + "</p>"))
}

// Gets amount from requst and return in cents
func getAmount(r *http.Request) (int, error) {
	transferAmountStr := r.FormValue("amount")
	transferAmount, err := strconv.ParseFloat(transferAmountStr, 64)
	if err != nil {
		return 0, errors.New("Invalid amount")
	}
	amountCents := int(math.Round(transferAmount * 100))
	return amountCents, nil
}
