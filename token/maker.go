package token

import "time"

// Maker is an interface that creates and verifies tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration.
	// Returns token, payload and error
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}