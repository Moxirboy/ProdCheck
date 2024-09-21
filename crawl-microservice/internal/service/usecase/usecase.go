package usecase

import (
	"context"
	"crawl-microservice/internal/models"
	"crawl-microservice/pkg/utils"
)

type IAuthUseCase interface {
	Check(ctx context.Context, accessToken string) (id, role string, err error)
	ReNew(ctx context.Context, refreshToken string) (tokens *models.Tokens, err error)
	New(ctx context.Context, id, role string) (tokens *models.Tokens, err error)
}

type IAccountUseCase interface {
	CreateUser(
		ctx context.Context,
		User *models.User,
	) error
	LoginUser(
		ctx context.Context,
		login, password string,
	) (*models.User, error)
}

type IProductUseCase interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error)
	GetListViolations(ctx context.Context, name string, query utils.PaginationQuery) (*models.ViolationList, error)
	GetListCrawledPrices(ctx context.Context, name string, query utils.PaginationQuery) (*models.PriceList, error)
}
