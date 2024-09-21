package postgres

import (
	"context"
	"database/sql"
	"engine/internal/service/repo"
	"engine/pkg/logger"

)



type screenshotRepo struct{
	db *sql.DB
	log logger.Logger
}


func NewScreenshotRepo(db *sql.DB,logger logger.Logger) repo.ScreenshotRepo{
	return &screenshotRepo{
		db:db,
		log:logger,
	}
}

func (repo *screenshotRepo)	CreateScreenshot(ctx context.Context,Id string,path string) (error){
	_,err:=repo.db.ExecContext(ctx,insertScreenshot,Id,path)
	if err!=nil{
		return err
	}
	return nil
}
