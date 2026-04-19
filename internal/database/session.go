package database

import (
	"banking-app/internal/domain"
)

func CreateSession(session *domain.Session) error {
	query := `INSERT INTO sessions (id, userId, expiryTime) VALUES(?, ?, ?)`
	_, err := DB.Exec(query, &session.ID,
		&session.UserID, &session.ExpiryTime)
	return err
}

func DeleteSession(id string) error {
	_, err := DB.Exec(`DELETE FROM sessions where id=?`, id)
	return err
}

func CleanUpSessions(currTime int64) error {
	_, err := DB.Exec(`DELETE FROM sessions where expiryTime <= ?`, currTime)
	return err
}

func GetSession(id string) (domain.Session, error) {
	var session domain.Session
	querry := "SELECT * FROM sessions WHERE id = ?"
	row := DB.QueryRow(querry, id)
	err := row.Scan(&session.ID, &session.UserID, &session.ExpiryTime)
	return session, err
}

func GetUserID(id string) (string, error) {
	var userID string
	querry := "SELECT userID FROM sessions WHERE id = ?"
	row := DB.QueryRow(querry, id)
	err := row.Scan(&userID)
	return userID, err
}
