package sessions

import (
	"banking-app/internal/database"
	"banking-app/internal/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// Makes a random 16 byte string
func createSessionID() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	id := hex.EncodeToString(randomBytes)
	return id
}

func createSessionStruct(username string, userID string) *models.Session {
	expiryTime := time.Now().Add(5 * time.Minute).Unix()
	id := createSessionID()
	session := models.Session{ID: id, LoginStatus: true,
		UserID: userID, Username: username, ExpiryTime: expiryTime}
	return &session
}

func CreateSession(username string, userID string) (string, error) {
	session := createSessionStruct(username, userID)

	err := database.CreateSession(session)
	if err != nil {
		err = fmt.Errorf("There is an error in creating the session id %w", err)
	}
	return session.ID, err
}

// Now I need to call this every hour
func CleanUpSessions() error {
	currTime := time.Now().Unix()
	err := database.CleanUpSessions(currTime)
	return err
}
