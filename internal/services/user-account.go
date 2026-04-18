package services

import (
	"banking-app/internal/database"
	"banking-app/internal/models"
	"fmt"
)

func getUser(username string, hashedPassword string) models.User {
	user := models.User{ID: createID(), Name: username, HashedPassword: hashedPassword}
	return user
}

// Hashes password and stores account and create new bank account
func CreateUserAccount(username string, password string) (string, error) {
	var user models.User
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
