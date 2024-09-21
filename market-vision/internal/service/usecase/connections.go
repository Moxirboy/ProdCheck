package usecase

import (
	"database/sql"
	configs "market-vision/internal/config"
	"market-vision/internal/service/repo/postgres"
	"market-vision/pkg/logger"
)

type IUseCase interface {
	IProductUseCase() IProductUseCase
	IViolationUseCase() IViolationUseCase	
}

type UseCase struct {
	connections map[string]interface{}
}

const (
	_ProductUseCase = "ProductUseCase"
	_ViolationUseCase = "ViolationUseCase"
)

func New(
	cfg *configs.Config,
	pg *sql.DB,
	logger logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_ProductUseCase] = NewProductUseCase(postgres.NewProductRepo(pg),logger)

	connections[_ViolationUseCase] = NewViolationUseCase(postgres.NewViolationRepo(pg),logger)

	return &UseCase{
		connections: connections,
	}
}


func (uc *UseCase) IProductUseCase() IProductUseCase {
	return uc.connections[_ProductUseCase].(IProductUseCase)
}

func (uc *UseCase) IViolationUseCase() IViolationUseCase {
	return uc.connections[_ViolationUseCase].(IViolationUseCase)
}
