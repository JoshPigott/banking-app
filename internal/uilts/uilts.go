package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// Makes a random 16 byte string
func CreateID() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	id := hex.EncodeToString(randomBytes)
	return id
}
