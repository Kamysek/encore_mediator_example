package repository

import (
	"context"
	"encore.app/products/models"
	"encore.dev/types/uuid"
)

type InMemoryProductRepository struct {
}

var products []*models.Product

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{}
}

func (p *InMemoryProductRepository) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {

	products = append(products, product)

	return product, nil
}

func (p *InMemoryProductRepository) GetProductById(ctx context.Context, uuid uuid.UUID) (*models.Product, error) {

	var product *models.Product

	for _, p := range products {

		if p.ProductID == uuid {
			product = p
		}
	}

	return product, nil
}
