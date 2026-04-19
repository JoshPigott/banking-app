package services

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
	// Create an erveryday bank account
	err = database.CreateEverdayAccount(user.ID)
	if err != nil {
		return user.ID, fmt.Errorf("Fail to create everyday bank account: %w\n", err)
	}
	return user.ID, err
}

func getUser(username string, hashedPassword string) domain.User {
	user := domain.User{ID: utils.CreateID(), Name: username, HashedPassword: hashedPassword}
	return user
}
