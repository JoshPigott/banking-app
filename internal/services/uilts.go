package services

import (
	"crypto/rand"
	"encoding/hex"
)

// Makes a random 16 byte string
func createID() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	id := hex.EncodeToString(randomBytes)
	return id
}
