package postgres

import (
	"context"
	"database/sql"
	"market-vision/internal/model"
	"market-vision/internal/service/repo"
	"time"
)



type productRepo struct{
	db *sql.DB
}


func NewProductRepo(db *sql.DB) repo.ProductRepo{
	return &productRepo{
		db:db,
	}
}

func (repo *productRepo)GetProductDetails(ctx context.Context,Id string) (model.Products,error){
	product:=model.Products{}
	err:=repo.db.QueryRowContext(ctx,getDetails, Id).Scan(&
		product.Id,
		product.ManufacturerID,
		product.ProductName,
		product.SKU,
		product.Map,
	)
	if err!=nil{
		return model.Products{},err
	}

	return product,nil
}

func (repo *productRepo)CreateProduct(ctx context.Context,price model.Price) (string,error){
	id:=""
	err:=repo.db.QueryRowContext(ctx,createProduct,
		price.GetId(),
		price.GetRetailerId(),
		price.GetAdvertaisedPrice(),
		time.Now().Format("2006-01-02"),
		price.GetUrl(),
	).Scan(&id)
	if err!=nil{
		return "",err
	}
	return id, nil
}