package database

import (
	"banking-app/internal/domain"
)

// I am going to need to remake this with userID, username and password

func CreateUserAccount(user domain.User) error {
	query := `INSERT INTO users (userID, username, hashedPassword) VALUES(?, ?, ?)`
	_, err := DB.Exec(query, user.ID, user.Name, user.HashedPassword)
	return err
}

func GetUserByUsername(username string) (domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE username=?"
	row := DB.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Name, &user.HashedPassword)
	return user, err
}

func GetUsername(userID string) (string, error) {
	var username string
	query := "SELECT username FROM users WHERE userID=?"
	row := DB.QueryRow(query, userID)
	err := row.Scan(&username)
	return username, err
}
