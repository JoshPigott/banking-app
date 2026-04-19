package domain

const minPasswordLen = 8

type User struct {
	ID             string
	Name           string
	HashedPassword string
}
