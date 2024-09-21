package usecase

import (
	"context"
	"market-vision/internal/model"
	"market-vision/internal/service/repo"
	"market-vision/pkg/logger"
)





type productUseCase struct {
	repo repo.ProductRepo
	logger logger.Logger
}

func NewProductUseCase(repo repo.ProductRepo,logger logger.Logger) IProductUseCase {
	return &productUseCase{
		repo: repo,
		logger:logger,
	}
}


func (uc *productUseCase) GetProductDetails(ctx context.Context,Id string) (model.Products, error) {
	product,err:=uc.repo.GetProductDetails(ctx,Id)
	if err!=nil{
		uc.logger.Error(err)
		return model.Products{},err
	}
	return product,nil
}

func (uc *productUseCase) CreateProduct(ctx context.Context,details model.Price) (string,error) {
	id,err:=uc.repo.CreateProduct(ctx,details)
	if err!=nil{
		uc.logger.Error(err)
		return id,err
	}
	return id,nil
}