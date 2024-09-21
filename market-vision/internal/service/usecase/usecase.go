package usecase

import (
	"context"
	"market-vision/internal/model"
)

type IProductUseCase interface {
	GetProductDetails(ctx context.Context,Id string) (model.Products, error) 
	CreateProduct(ctx context.Context,details model.Price) (string,error) 
}

type IViolationUseCase interface {
	CreateViolation(ctx context.Context,Id string) (string,error)
}