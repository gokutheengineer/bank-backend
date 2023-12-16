package util

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
)

// HashPassword hashes a password using blake2b, encodes it to hex, and returns it as a string
func HashPasswordBlake(password string) string {
	hash := blake2b.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// HashPasswordBcrypt hashes a password using bcrypt, and returns it as a string
func HashPasswordBcrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return string(hashedPassword), nil
}

// VerifyPasswordBlake verifies a password using blake2b. Returns true if the password matches the hash
func VerifyPasswordBlake(password string, hashedpassword string) bool {
	hash := blake2b.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]) == hashedpassword
}

// VerifyPasswordBcrypt verifies a password using bcrypt. Returns true if the password matches the hash
func VerifyPasswordBcrypt(password string, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
