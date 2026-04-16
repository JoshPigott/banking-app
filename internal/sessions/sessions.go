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

func createSessionStruct(username string, userID string, expiryTime time.Time) *models.Session {
	id := createSessionID()
	session := models.Session{ID: id, LoginStatus: true,
		UserID: userID, Username: username, ExpiryTime: expiryTime.Unix()}
	return &session
}

func CreateSession(username string, userID string, expiryTime time.Time) (string, error) {
	session := createSessionStruct(username, userID, expiryTime)
	err := database.CreateSession(session)
	if err != nil {
		err = fmt.Errorf("There is an error in creating the session id %w", err)
	}
	return session.ID, err
}

// Cleans up expired session every hour so session don't build up
func CleanUpSessions() {
	time.AfterFunc(2*time.Hour, func() {
		currTime := time.Now().Unix()
		fmt.Print("Clean Up ran")
		err := database.CleanUpSessions(currTime)
		if err != nil {
			fmt.Printf("Error doing session cleanup %v", err)
		}
		// Schedule again
		CleanUpSessions()
	})
}
