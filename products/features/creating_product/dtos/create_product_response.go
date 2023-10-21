package dtos

import "encore.dev/types/uuid"

type CreateProductCommandResponse struct {
	ProductID uuid.UUID `json:"productId"`
}
