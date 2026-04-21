package services

import (
	"banking-app/internal/database"
	"banking-app/internal/domain"
	"errors"
)

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

func MakeTransfer(t domain.TransferRequest) error {
	userID, err := database.GetUserID(t.SessionID)
	if err != nil {
		return errors.New("Unable to get user")
	}
	// Get database table names
	accountFromTable := t.AccountFrom.GetTableName()
	accountToTable := t.AccountTo.GetTableName()

	if err = database.MakeTransfer(accountFromTable, accountToTable, t.AmountCents, userID); err != nil {
		return errors.New("Unable to make transfer")
	}
	return err
}

// Check if transfer data is valid
func CanTransfer(t domain.TransferRequest) error {
	if !t.AccountFrom.CanWithdraw() || !t.AccountTo.IsValid() || t.AccountFrom == t.AccountTo {
		return errors.New("Invalid to and from accounts")
	}
	if !isValidTransferAmount(t.AmountCents, t.AccountFrom, t.SessionID) {
		return errors.New("Invalid transfer amount")
	}
	return nil
}

// Checks if the amount to transfer is valid or not
func isValidTransferAmount(transferAmount int, accountFrom domain.BankAccountType, sessionID string) bool {
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
