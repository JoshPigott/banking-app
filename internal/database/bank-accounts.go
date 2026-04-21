package database

import (
	"database/sql"
	"fmt"
)

func CreateEverdayAccount(userID string) error {
	startBalanceCents := 100000
	query := `INSERT INTO everydayAccount (userID, balanceCents) VALUES(?, ?)`
	_, err := DB.Exec(query, userID, startBalanceCents)
	return err
}

func CreateSaverAccount(userID string) error {
	startBalanceCents := 0
	query := `INSERT INTO saverAccount (userID, balanceCents) VALUES(?, ?)`
	_, err := DB.Exec(query, userID, startBalanceCents)
	return err
}

func CreateKiwiSaverAccount(userID string) error {
	startBalanceCents := 0
	query := `INSERT INTO kiwiSaverAccount (userID, balanceCents) VALUES(?, ?)`
	_, err := DB.Exec(query, userID, startBalanceCents)
	return err
}

func GetAccountBalance(account string, userID string) (int, error) {
	var balanceCents int
	query := fmt.Sprintf("SELECT balanceCents FROM %s WHERE userID = ?", account)
	row := DB.QueryRow(query, userID)
	err := row.Scan(&balanceCents)
	return balanceCents, err
}

type DBTX interface {
	Exec(query string, args ...any) (sql.Result, error)
}

// Note in both withdraw and in deposit account,
// amount, and userID are check input to avoid sql injections

// Takes money out from the account
func WithDraw(db DBTX, account string, amount int, userID string) error {
	var err error
	query := fmt.Sprintf("UPDATE %s SET balanceCents = balanceCents - ? WHERE userID = ?", account)
	res, err := db.Exec(query, amount, userID)
	if err != nil {
		return err
	}
	rowsChanged, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsChanged != 1 {
		return fmt.Errorf("Account not found")
	}
	return err
}

// Adds money to the account
func Deposit(db DBTX, account string, amount int, userID string) error {
	query := fmt.Sprintf("UPDATE %s SET balanceCents = balanceCents + ? WHERE userID = ?", account)
	res, err := db.Exec(query, amount, userID)
	if err != nil {
		return err
	}
	rowsChanged, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsChanged != 1 {
		return fmt.Errorf("Account not found")
	}
	return err
}

func transfer(db DBTX, accountFromTable string, accountToTable string, amount int, userID string) error {
	var err error
	if err = WithDraw(db, accountFromTable, amount, userID); err != nil {
		return err
	}
	if err = Deposit(db, accountToTable, amount, userID); err != nil {
		return err
	}
	return err
}

// Make sure WithDraw can't happen without Deposit happening
func MakeTransfer(accountFromTable string, accountToTable string, amount int, userID string) error {
	return withTx(func(tx *sql.Tx) error {
		return transfer(tx, accountFromTable, accountToTable, amount, userID)
	})
}

// With muilple database changes makes sure they all happen or none
func withTx(fn func(*sql.Tx) error) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit()
}
