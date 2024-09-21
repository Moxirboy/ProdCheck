package usecase

import "engine/internal/service/repo"
import "context"


type screenshotUsecase struct{
	repo repo.ScreenshotRepo
}


func NewScreenshotRepo(repo repo.ScreenshotRepo) repo.ScreenshotRepo{
	return &screenshotUsecase{
		repo:repo,
	}
}

func (uc *screenshotUsecase)CreateScreenshot(ctx context.Context,Id string,path string) (error){
	return uc.repo.CreateScreenshot(ctx,Id,path)
}