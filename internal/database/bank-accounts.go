package database

import "fmt"

func CreateEverdayAccount(userID string) error {
	startBalance := 1000
	query := `INSERT INTO everydayAccount (userID, balance) VALUES(?, ?)`
	_, err := DB.Exec(query, userID, startBalance)
	return err
}

func CreateSaverAccount(userID string) error {
	startBalance := 0
	query := `INSERT INTO saverAccount (userID, balance) VALUES(?, ?)`
	_, err := DB.Exec(query, userID, startBalance)
	return err
}

func CreateKiwiSaverAccount(userID string) error {
	startBalance := 0
	query := `INSERT INTO kiwiSaverAccount (userID, balance) VALUES(?, ?)`
	_, err := DB.Exec(query, userID, startBalance)
	return err
}

func GetAccountBalance(account string, userID string) (float64, error) {
	var balance float64
	querry := fmt.Sprintf("SELECT balance FROM %s WHERE userID = ?", account)
	row := DB.QueryRow(querry, userID)
	err := row.Scan(&balance)
	return balance, err
}
