package helpers

import (
	"banking-app/internal/database"
	"banking-app/internal/domain"
	utils "banking-app/internal/uilts"
	"fmt"
)

// Hashes password and stores account and create new bank account
func CreateUserAccount(username string, password string) (string, error) {
	var user domain.User
	hashedPassword, err := hashPassword(password)

	if err != nil {
		return user.ID, fmt.Errorf("Fail to hash password: %w\n", err)
	}

	user = getUser(username, hashedPassword)
	err = database.CreateUserAccount(user)
	if err != nil {
		return user.ID, fmt.Errorf("Fail to create users account: %w\n", err)
	}
	// Create bank accounts
	err = createAllAccounts(user.ID)
	if err != nil {
		return user.ID, err
	}
	return user.ID, err
}

// Set up new bank account in the database for the user
func createAllAccounts(userID string) error {
	var err error
	if err = database.CreateEverdayAccount(userID); err != nil {
		return fmt.Errorf("Fail to create everyday bank account: %w\n", err)
	}
	if err = database.CreateSaverAccount(userID); err != nil {
		return fmt.Errorf("Fail to create saver bank account: %w\n", err)
	}
	if err = database.CreateKiwiSaverAccount(userID); err != nil {
		return fmt.Errorf("Fail to create kiwi saver bank account: %w\n", err)
	}
	return err
}

func getUser(username string, hashedPassword string) domain.User {
	user := domain.User{ID: utils.CreateID(), Name: username, HashedPassword: hashedPassword}
	return user
}
