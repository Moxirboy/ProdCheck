package repo

import (
	"context"
	"market-vision/internal/model"
)

type ProductRepo interface{
	GetProductDetails(ctx context.Context,Id string) (model.Products,error)
	CreateProduct(ctx context.Context,price model.Price) (string,error)
}