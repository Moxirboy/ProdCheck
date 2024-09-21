package usecase

import (
	"database/sql"
	"engine/internal/configs"
	"engine/internal/service/repo/postgres"
	"engine/pkg/logger"
)

type IUseCase interface {
	ScreenshotUsecase() ScreenshotUsecase
}

type UseCase struct {
	connections map[string]interface{}
}

const (
	_ScreenshotUsecase = "ScreenshotUsecase"
)

func New(
	cfg *configs.Config,
	pg *sql.DB,
	logger logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_ScreenshotUsecase] = NewScreenshotRepo(postgres.NewScreenshotRepo(pg,logger))

	return &UseCase{
		connections: connections,
	}
}


func (uc *UseCase) ScreenshotUsecase() ScreenshotUsecase {
	return uc.connections[_ScreenshotUsecase].(ScreenshotUsecase)
}