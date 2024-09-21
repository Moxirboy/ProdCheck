package dto

import "market-vision/internal/model"



type ProductIncome struct {
	ProductID string `json:"product_id"`
	Url string `json:"url"`
	RetailerID string `json:"retailer_id"`
}



func FromProductIncome(productIncome ProductIncome) *model.Url {
	return &model.Url{
		ProdcutId: productIncome.ProductID,
		Url: productIncome.Url,
		RetailerId: productIncome.RetailerID,
	}
}