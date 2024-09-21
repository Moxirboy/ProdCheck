package usecase

import (
	"context"
	"market-vision/internal/service/repo"
	"market-vision/pkg/logger"
)


type violationUseCase struct {
	violationRepo repo.IViolationRepo
	logger logger.Logger
}

func NewViolationUseCase(repo repo.IViolationRepo,logger logger.Logger) IViolationUseCase {
	return &violationUseCase{
		violationRepo: repo,
		logger:logger,
	}
}

func (uc *violationUseCase) CreateViolation(ctx context.Context,Id string) (string,error) {
	id,err:=uc.violationRepo.CreateViolation(ctx,Id)
	if err!=nil{
		uc.logger.Error(err)
		return id,err
	}
	return id,nil
}
