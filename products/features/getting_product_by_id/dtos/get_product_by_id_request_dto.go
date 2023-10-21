package dtos

import "encore.dev/types/uuid"

type GetProductByIdRequestDto struct {
	ProductId uuid.UUID `param:"id" json:"-"`
}
