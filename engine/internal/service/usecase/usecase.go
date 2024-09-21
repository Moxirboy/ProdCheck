package usecase

import "context"

type ScreenshotUsecase interface{
	CreateScreenshot(ctx context.Context,Id string,path string) (error)
}