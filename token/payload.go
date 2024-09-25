package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	Expiry   time.Time `json:"expired_at"`
}

func NewPayLoad(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		IssuedAt: time.Now(),
		Expiry:   time.Now().Add(duration),
	}

	return payload, nil
}

var (
	ErrExpiredToken = errors.New("token has already expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expiry) {
		return ErrExpiredToken
	}

	return nil
}