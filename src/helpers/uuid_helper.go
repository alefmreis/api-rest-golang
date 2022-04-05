package helpers

import "github.com/google/uuid"

type UUID struct{}

func NewUUID() UUID {
	return UUID{}
}

func (u UUID) New() uuid.UUID {
	return uuid.New()
}
