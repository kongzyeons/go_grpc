package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return ""
	}
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}
