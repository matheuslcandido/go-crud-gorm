package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        string
	Name      string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(name string, amount int) *Product {
	return &Product{
		ID:     uuid.New().String(),
		Name:   name,
		Amount: amount,
	}
}
