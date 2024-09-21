package models

import (
	"crawl-microservice/internal/dto"
)
type Product struct {
	Id             string
	ManufacturerID string
	ProductName    string
	SKU            string
	Map            string
}


type ProductsList struct {
	TotalCount int
	TotalPages int
	Page       int
	Size       int
	HasMore    bool
	Product    []*Product
}

func NewProduct(model dto.Product) *Product {
	return &Product{
		ManufacturerID:        model.ManufacturerID,
		ProductName: model.ProductName,
		SKU:       model.SKU,
		Map:       model.Map,
	}
}
