package token

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    int64     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(userId int64, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return fmt.Errorf("token expired")
	}
	return nil
}
