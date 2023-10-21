package dtos

import (
	"encore.app/products/dtos"
)

type GetProductByIdQueryResponse struct {
	Product *dtos.ProductDto `json:"product"`
}
