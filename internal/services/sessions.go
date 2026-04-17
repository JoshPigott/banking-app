package services

import (
	"banking-app/internal/database"
	"banking-app/internal/models"
	"fmt"
	"time"
)

func createSessionStruct(userID string, expiryTime time.Time) *models.Session {
	id := createID()
	session := models.Session{ID: id,
		UserID: userID, ExpiryTime: expiryTime.Unix()}
	return &session
}

func CreateSession(userID string) (string, time.Time, error) {
	expiryTime := time.Now().Add(time.Hour)
	session := createSessionStruct(userID, expiryTime)
	err := database.CreateSession(session)
	if err != nil {
		err = fmt.Errorf("There is an error in creating the session %w", err)
	} else {
		fmt.Printf("A session has been made %s\n", session.ID)
	}
	return session.ID, expiryTime, err
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
