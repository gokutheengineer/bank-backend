package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Payload is the payload of the token
type Payload struct {
	// ID is the id of the token
	ID uuid.UUID `json:"id"`
	// Username is the username of the token user
	Username string `json:"username"`
	// Role is the role of the token user
	Role string `json:"role"`
	// IssuedAt is the time the token was issued
	IssuedAt time.Time `json:"issued_at"`
	// ExpiredAt is the time the token will expire
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new payload for the token
func NewPayload(username string, role string, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

// Valid checks if the token is valid or not. If the token is invalid, it returns an error
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}
