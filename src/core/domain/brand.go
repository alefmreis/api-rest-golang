package domain

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBrand(id uuid.UUID, name string) Brand {
	return Brand{Id: id, Name: name}
}
