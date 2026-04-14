package database

// import (
// 	"banking-app/internal/models"
// 	"fmt"
// )

// func AddUser(userID string, name string) error {
// 	_, err := DB.Exec(`INSERT INTO users
//    (userID, name) VALUES(?, ?)`, userID, name)
// 	return err
// }

// func GetUser(userID string) (models.User, error) {
// 	var user models.User
// 	querry := "SELECT * FROM USERS WHERE userID = ?"
// 	row := DB.QueryRow(querry, userID)
// 	fmt.Print("this is a row:", *row)
// 	err := row.Scan(&user.UserID, &user.Name)
// 	return user, err
// }
