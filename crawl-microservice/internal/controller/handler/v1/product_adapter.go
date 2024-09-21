package v1

import (
	"crawl-microservice/internal/dto"
	"crawl-microservice/internal/models"
)

func ToProductResponse(
	product models.Product,
) *dto.Product {
	return &dto.Product{
		ManufacturerID: product.ManufacturerID,
		ProductName:    product.ProductName,
		SKU:            product.SKU,
		Map:            product.Map,
	}
}
func toProductListResponse(
	products models.ProductsList,
) *dto.ProductsList {
	list := make([]*dto.Product, len(products.Product))
	for i, product := range products.Product {
		list[i] = ToProductResponse(*product)
	}

	return &dto.ProductsList{
		TotalCount: products.TotalCount,
		TotalPages: products.TotalPages,
		Page:       products.Page,
		Size:       products.Size,
		HasMore:    products.HasMore,
		Product:    list,
	}
}
