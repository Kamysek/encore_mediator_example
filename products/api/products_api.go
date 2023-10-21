package api

import (
	"context"
	"encore.app/products/features/creating_product/commands"
	creatingProductsDtos "encore.app/products/features/creating_product/dtos"
	"encore.app/products/features/creating_product/events"
	gettingProductByIdDtos "encore.app/products/features/getting_product_by_id/dtos"
	"encore.app/products/features/getting_product_by_id/queries"
	"encore.app/products/repository"
	"encore.app/shared/behaviours"
	"encore.dev/types/uuid"
	"github.com/go-playground/validator"
	"github.com/mehdihadeli/go-mediatr"
	"log"
)

//encore:service
type ProductsController struct {
	validator *validator.Validate
}

func initProductsController() (*ProductsController, error) {

	productRepository := repository.NewInMemoryProductRepository()

	//////////////////////////////////////////////////////////////////////////////////////////////
	// Register request handlers to the mediatr

	createProductCommandHandler := commands.NewCreateProductCommandHandler(productRepository)
	getByIdQueryHandler := queries.NewGetProductByIdHandler(productRepository)

	err := mediatr.RegisterRequestHandler[*commands.CreateProductCommand, *creatingProductsDtos.CreateProductCommandResponse](createProductCommandHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = mediatr.RegisterRequestHandler[*queries.GetProductByIdQuery, *gettingProductByIdDtos.GetProductByIdQueryResponse](getByIdQueryHandler)
	if err != nil {
		log.Fatal(err)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////
	// Register notification handlers to the mediatr
	notificationHandler := events.NewProductCreatedEventHandler()
	err = mediatr.RegisterNotificationHandler[*events.ProductCreatedEvent](notificationHandler)
	if err != nil {
		log.Fatal(err)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////
	// Register request handlers pipeline to the mediatr
	loggerPipeline := &behaviours.RequestLoggerBehaviour{}
	err = mediatr.RegisterRequestPipelineBehaviors(loggerPipeline)
	if err != nil {
		log.Fatal(err)
	}

	return &ProductsController{
		validator: validator.New(),
	}, nil
}

//encore:api public method=POST path=/api/v1/products
func (pc *ProductsController) CreateProduct(ctx context.Context, request creatingProductsDtos.CreateProductRequestDto) (*creatingProductsDtos.CreateProductCommandResponse, error) {

	if err := pc.validator.StructCtx(ctx, request); err != nil {
		return nil, err
	}

	command := commands.NewCreateProductCommand(request.Name, request.Description, request.Price)
	result, err := mediatr.Send[*commands.CreateProductCommand, *creatingProductsDtos.CreateProductCommandResponse](ctx, command)
	if err != nil {
		return nil, err
	}

	return result, err
}

//encore:api public method=GET path=/api/v1/products/:id
func (pc *ProductsController) GetProductByID(ctx context.Context, id string) (*gettingProductByIdDtos.GetProductByIdQueryResponse, error) {

	query := queries.NewGetProductByIdQuery(uuid.FromStringOrNil(id))

	if err := pc.validator.StructCtx(ctx, query); err != nil {
		return nil, err
	}

	queryResult, err := mediatr.Send[*queries.GetProductByIdQuery, *gettingProductByIdDtos.GetProductByIdQueryResponse](ctx, query)
	if err != nil {
		return nil, err
	}

	return queryResult, nil
}
