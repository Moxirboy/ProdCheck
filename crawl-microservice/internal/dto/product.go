package dto

type Product struct {
	ManufacturerID        string  `json:"ManufacturerID" `
	ProductName string  `json:"ProductName" `
	SKU       string `json:"SKU" `
	Map       string  `json:"Map" `
}


type GetProductByID struct {
	ID string `query:"id" binding:"required"`
}
type ProductsList struct {
	TotalCount int        `json:"totalCount" :"total_count"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	HasMore    bool       `json:"has_more"`
	Product    []*Product `json:"product"`
}
