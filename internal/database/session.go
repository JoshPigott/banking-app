package database

import (
	"banking-app/internal/models"
)

func CreateSession(session *models.Session) error {
	query := `INSERT INTO sessions 
	(id, loginStatus, userId, username, expiryTime)
	VALUES(?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, &session.ID, &session.LoginStatus,
		&session.UserID, &session.Username, &session.ExpiryTime)
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

// Sets login to true
func GiveLoginStatus(id string) error {
	_, err := DB.Exec(`UPDATE sessions 
  SET loginStatus = true WHERE id=?`, id)
	return err
}

// Sets login status to false
func RevokeLoginStatus(id string) error {
	_, err := DB.Exec(`UPDATE sessions 
  SET loginStatus = false WHERE id=?`, id)
	return err
}

func GetLoginStatus(id string) (bool, error) {
	var loginStatus bool
	querry := `SELECT loginStatus FROM sessions WHERE id=?`
	row := DB.QueryRow(querry, id)
	err := row.Scan(&loginStatus)
	return loginStatus, err
}

func GetSession(id string) (models.Session, error) {
	var session models.Session
	querry := "SELECT * FROM sessions WHERE id = ?"
	row := DB.QueryRow(querry, id)
	err := row.Scan(&session.ID, &session.LoginStatus,
		&session.UserID, &session.Username, &session.ExpiryTime)
	return session, err
}
