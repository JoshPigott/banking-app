package models

// type User struct {
// 	UserID string
// 	Name   string
// }

type Session struct {
	ID          string
	LoginStatus bool
	UserID      string
	Username    string
	ExpiryTime  int64
}
