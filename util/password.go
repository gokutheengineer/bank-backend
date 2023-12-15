package util

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

func HashPassword(password string) string {
	hash := blake2b.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(password string, hashedpassword string) bool {
	hash := blake2b.Sum256([]byte(password))
	return string(hash[:]) == hashedpassword
}
