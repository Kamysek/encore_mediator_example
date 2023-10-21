package queries

import "encore.dev/types/uuid"

type GetProductByIdQuery struct {
	ProductID uuid.UUID `validate:"required"`
}

func NewGetProductByIdQuery(productID uuid.UUID) *GetProductByIdQuery {
	return &GetProductByIdQuery{ProductID: productID}
}
