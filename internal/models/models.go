package models

type User struct {
	ID             string
	Name           string
	HashedPassword string
}

type Session struct {
	ID         string
	UserID     string
	ExpiryTime int64
}
