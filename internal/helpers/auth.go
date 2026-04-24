package helpers

import (
	"banking-app/internal/database"
	"banking-app/internal/domain"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ValidLoginCredentials(username string, password string) (bool, string) {
	user, err := database.GetUserByUsername(username)
	if err != nil {
		return false, user.ID
	}
	valid := checkPasswordHash(password, user.HashedPassword)
	return valid, user.ID
}

func IsValidCredentials(u string, p string) bool {
	password := domain.NewPassword(p)
	if !password.IsStrong() {
		return false
	}
	if len(u) < 6 {
		return false
	}
	if !isUsernameUnique(u) {
		return false
	}
	return true
}

func isUsernameUnique(username string) bool {
	_, err := database.GetUserByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}
	return false
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
