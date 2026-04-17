package services

import (
	"banking-app/internal/database"
	"database/sql"
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Checks if there and number and low and upper case letters
func hasBothCasesAndNumber(password string) bool {
	var hasLowerCase, hasUpperCase, hasNumber bool
	for _, r := range password {
		if unicode.IsUpper(r) {
			hasUpperCase = true
		}
		if unicode.IsLower(r) {
			hasLowerCase = true
		}
		if unicode.IsDigit(r) {
			hasNumber = true
		}
	}
	return hasLowerCase && hasUpperCase && hasNumber
}

// Checks the size and character in the password
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if !hasBothCasesAndNumber(password) {
		return false
	}
	return true
}

func isUsernameUnique(username string) bool {
	_, err := database.GetUser(username)
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}
	return false
}

// Makes sure username is long enough and unique
func IsValidUsername(username string) bool {
	if len(username) < 6 {
		return false
	}
	if !isUsernameUnique(username) {
		return false
	}
	return true
}

func ValidLoginCredentials(username string, password string) (bool, string) {
	user, err := database.GetUser(username)
	if err != nil {
		return false, user.ID
	}
	valid := CheckPasswordHash(password, user.HashedPassword)
	return valid, user.ID
}
