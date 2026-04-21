package helpers

import (
	"banking-app/internal/database"
	"banking-app/internal/domain"
	"errors"
)

const defaultReceiverAccount = "everydayAccount"

// Get userID then the account balance
func GetAccountBalance(sessionID string, bankAccountType domain.BankAccountType) (int, error) {
	var balance int
	userID, err := database.GetUserID(sessionID)
	if err != nil {
		return balance, err
	}
	tableName := bankAccountType.GetTableName()
	balance, err = database.GetAccountBalance(tableName, userID)
	return balance, err
}

func MakePayment(p *domain.PaymentRequest) error {
	var err error
	// Get database table names
	p.AccountFromTable = p.AccountFrom.GetTableName()
	p.AccountToTable = defaultReceiverAccount

	if err := database.MakePayment(*p); err != nil {
		return errors.New("Unable to make transfer")
	}
	return err
}

func MakeTransfer(t *domain.TransferRequest) error {
	var err error
	t.UserID, err = database.GetUserID(t.SessionID)
	if err != nil {
		return errors.New("Unable to get user")
	}
	// Get database table names
	t.AccountFromTable = t.AccountFrom.GetTableName()
	t.AccountToTable = t.AccountTo.GetTableName()

	// accountFromTable, accountToTable, t.AmountCents, userID
	if err = database.MakeTransfer(*t); err != nil {
		return errors.New("Unable to make transfer")
	}
	return err
}

func IsValidPayment(p *domain.PaymentRequest) error {
	if !p.AccountFrom.CanWithdraw() {
		return errors.New("Invalid account")
	}
	if !isValidAmount(p.AmountCents, p.AccountFrom, p.SessionID) {
		return errors.New("Invalid payment amount")
	}
	if !receiver(p) {
		return errors.New("Invalid receiver")
	}
	return nil
}

// Check if transfer data is valid
func CanTransfer(t domain.TransferRequest) error {
	if !t.AccountFrom.CanWithdraw() || !t.AccountTo.IsValid() || t.AccountFrom == t.AccountTo {
		return errors.New("Invalid to and from accounts")
	}
	if !isValidAmount(t.AmountCents, t.AccountFrom, t.SessionID) {
		return errors.New("Invalid transfer amount")
	}
	return nil
}

// This is where an error is happening
// Checks if the amount to transfer is valid or not
func isValidAmount(transferAmount int, accountFrom domain.BankAccountType, sessionID string) bool {
	if transferAmount <= 0 {
		return false
	}

	balance, err := GetAccountBalance(sessionID, accountFrom)
	if err != nil {
		return false
	}
	if balance < transferAmount {
		return false
	}
	return true
}

// Get receive user id and return if receiver is valid
func receiver(p *domain.PaymentRequest) bool {
	user, err := database.GetUser(p.ReceiverUsername)
	if err != nil {
		return false
	}
	// Updates payment request
	p.ReceiveUserID = user.ID
	p.UserID, err = database.GetUserID(p.SessionID)
	if err != nil {
		return false
	}

	// You can't make payment to yourself
	if p.ReceiveUserID == p.UserID {
		return false
	}
	return true
}
