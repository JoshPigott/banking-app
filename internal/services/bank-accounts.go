package services

import (
	"banking-app/internal/database"
	"banking-app/internal/domain"
)

// Get userID then the account balance
func GetAccountBalance(sessionID string, accountType domain.AccountType) (float64, error) {
	var balance float64
	userID, err := database.GetUserID(sessionID)
	if err != nil {
		return balance, err
	}
	tableName := accountType.GetTableName()
	balance, err = database.GetAccountBalance(tableName, userID)
	return balance, err
}
