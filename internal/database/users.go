package database

import (
	"banking-app/internal/models"
)

// I am going to need to remake this with userID, username and password

func CreateUserAccount(user models.User) error {
	querry := `INSERT INTO users (userID, username, hashedPassword) VALUES(?, ?, ?)`
	_, err := DB.Exec(querry, user.ID, user.Name, user.HashedPassword)
	return err
}

func GetUser(username string) (models.User, error) {
	var user models.User
	querry := "SELECT * FROM USERS WHERE username=?"
	row := DB.QueryRow(querry, username)
	err := row.Scan(&user.ID, &user.Name, &user.HashedPassword)
	return user, err
}
