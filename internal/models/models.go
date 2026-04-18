package models

// I think in the further this should be split and each should have methods
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
