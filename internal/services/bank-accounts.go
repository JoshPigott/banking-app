package services

import (
	"banking-app/internal/database"
	"banking-app/internal/domain"
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

// Checks if the amount to transfer is valid or not
func IsValidTransferAmount(transferAmount int, accountFrom domain.BankAccountType, sessionID string) bool {
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

func MakeTransfer(transferAmount int, accountFrom domain.BankAccountType, accountTo domain.BankAccountType, sessionID string) error {
	userID, err := database.GetUserID(sessionID)
	if err != nil {
		return err
	}
	// Get database table names
	accountFromTable := accountFrom.GetTableName()
	accountToTable := accountTo.GetTableName()

	err = database.MakeTransfer(accountFromTable, accountToTable, transferAmount, userID)
	return err
}
