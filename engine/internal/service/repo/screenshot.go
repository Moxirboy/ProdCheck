package repo

import "context"

type ScreenshotRepo interface{
	CreateScreenshot(ctx context.Context,Id string,path string) (error)
}