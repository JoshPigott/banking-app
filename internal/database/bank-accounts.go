package database

import (
	"banking-app/internal/domain"
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

// Makes sure WithDraw can't happen without Deposit happening
func MakePayment(p domain.PaymentRequest) error {
	return withTx(func(tx *sql.Tx) error {
		return payment(tx, p)
	})
}

// Makes sure WithDraw can't happen without Deposit happening
func MakeTransfer(t domain.TransferRequest) error {
	return withTx(func(tx *sql.Tx) error {
		return transfer(tx, t)
	})
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

// accountFromTable string, accountToTable string, amount
func payment(db DBTX, p domain.PaymentRequest) error {
	var err error
	if err = WithDraw(db, p.AccountFromTable, p.AmountCents, p.UserID); err != nil {
		return err
	}
	if err = Deposit(db, p.AccountToTable, p.AmountCents, p.ReceiveUserID); err != nil {
		return err
	}
	return err
}

func transfer(db DBTX, t domain.TransferRequest) error {
	var err error
	if err = WithDraw(db, t.AccountFromTable, t.AmountCents, t.UserID); err != nil {
		return err
	}
	if err = Deposit(db, t.AccountToTable, t.AmountCents, t.UserID); err != nil {
		return err
	}
	return err
}

type DBTX interface {
	Exec(query string, args ...any) (sql.Result, error)
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
