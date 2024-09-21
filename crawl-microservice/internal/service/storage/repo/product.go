package repo

import (
	"context"
	"crawl-microservice/internal/models"
	"crawl-microservice/pkg/utils"
)

type Product interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error)
	GetListViolations(ctx context.Context, name string, query utils.PaginationQuery) (*models.ViolationList, error) 
	GetListCrawledPrices(ctx context.Context, name string, query utils.PaginationQuery) (*models.PriceList, error)
}
