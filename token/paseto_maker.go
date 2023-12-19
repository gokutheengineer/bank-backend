package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {

	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: required key size must be %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken creates a new paseto token for a specific username and duration
func (m *PasetoMaker) CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(username, role, duration)
	token, err := m.paseto.Encrypt(m.symmetricKey, payload, nil)
	return token, payload, err
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var payload Payload

	err := m.paseto.Decrypt(token, m.symmetricKey, &payload, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	err = payload.Valid()
	if err != nil {
		return nil, fmt.Errorf("payload is invalid: %w", err)
	}

	return &payload, err
}
