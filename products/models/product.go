package models

import (
	"encore.dev/types/uuid"
	"time"
)

// Product model
type Product struct {
	ProductID   uuid.UUID `json:"productId" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
