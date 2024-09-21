package postgres

import (
	"context"
	"database/sql"
	"market-vision/internal/service/repo"
	"time"
)



type violationRepo struct{
	db *sql.DB
}


func NewViolationRepo(db *sql.DB) repo.IViolationRepo{
	return &violationRepo{
		db:db,
	}
}

func (repo *violationRepo)CreateViolation(ctx context.Context,Id string) (string,error){
	id:=""
	err:=repo.db.QueryRowContext(ctx,
		createViolation, 
		Id,
		time.Now().Format("2006-01-02"),
		"MAP",
		"",
		"",
	).Scan(&id)
	if err!=nil{
		return id,err
	}

	return id ,nil
}
