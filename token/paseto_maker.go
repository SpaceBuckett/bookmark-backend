package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(SymmetricKey string) (Maker, error) {
	if len(SymmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("symmetric key must be 32 bytes long.")
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(SymmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(userId int64, duration time.Duration) (string, error) {
	payload, err := NewPayload(userId, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt token: %s", err)
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
