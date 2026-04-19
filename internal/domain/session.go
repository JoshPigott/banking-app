package domain

import (
	utils "banking-app/internal/uilts"
	"time"
)

type Session struct {
	ID         string
	UserID     string
	ExpiryTime int64
}

func NewSession(userID string, expiryTime time.Time) *Session {
	id := utils.CreateID()
	session := Session{
		ID:         id,
		UserID:     userID,
		ExpiryTime: expiryTime.Unix()}
	return &session
}
