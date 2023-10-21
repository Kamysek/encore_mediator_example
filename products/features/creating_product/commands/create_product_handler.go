package commands

import (
	"context"
	creatingProductDtos "encore.app/products/features/creating_product/dtos"
	"encore.app/products/features/creating_product/events"
	"encore.app/products/models"
	"encore.app/products/repository"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
)

type CreateProductCommandHandler struct {
	productRepository *repository.InMemoryProductRepository
}

func NewCreateProductCommandHandler(productRepository *repository.InMemoryProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{productRepository: productRepository}
}

func (c *CreateProductCommandHandler) Handle(ctx context.Context, command *CreateProductCommand) (*creatingProductDtos.CreateProductCommandResponse, error) {
	isLoggerPipelineEnabled := ctx.Value("logger_pipeline").(bool)
	if isLoggerPipelineEnabled {
		fmt.Println("[CreateProductCommandHandler]: logging pipeline is enabled")
	}

	product := &models.Product{
		ProductID:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		CreatedAt:   command.CreatedAt,
	}

	createdProduct, err := c.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	response := &creatingProductDtos.CreateProductCommandResponse{ProductID: createdProduct.ProductID}

	// Publish notification event to the mediatr for dispatching to the notification handlers

	productCreatedEvent := events.NewProductCreatedEvent(createdProduct.ProductID, createdProduct.Name, createdProduct.Description, createdProduct.Price, createdProduct.CreatedAt)
	err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	if err != nil {
		return nil, err
	}

	return response, nil
}
