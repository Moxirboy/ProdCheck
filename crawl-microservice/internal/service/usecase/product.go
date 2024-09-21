package usecase

import (
	"context"
	"crawl-microservice/internal/models"
	"crawl-microservice/internal/service/storage/repo"
	"crawl-microservice/pkg/logger"
	"crawl-microservice/pkg/utils"
)

type ProductUsecase struct {
	log  logger.Logger
	repo repo.Product
}

func NewProductUsecase(repo repo.Product, log logger.Logger) IProductUseCase {
	return &ProductUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ProductUsecase) Create(ctx context.Context, product *models.Product) error {
	if err := uc.repo.Create(ctx, product); err != nil {
		return err
	}

	return nil
}

func (uc *ProductUsecase) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUsecase) GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error) {
	return uc.repo.GetList(
		ctx,
		name,
		query,
	)
}

func (uc *ProductUsecase) GetListViolations(ctx context.Context, name string, query utils.PaginationQuery) (*models.ViolationList, error) {
	return uc.repo.GetListViolations(
		ctx,
		name,
		query,
	)
}
func (uc *ProductUsecase) GetListCrawledPrices(ctx context.Context, name string, query utils.PaginationQuery) (*models.PriceList, error) {
	return uc.repo.GetListCrawledPrices(
		ctx,
		name,
		query,
	)
}