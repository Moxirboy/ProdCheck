package dto

type Price struct {
	ID               string `json:"id"`               // Unique identifier for the price
	ProductID        string `json:"product_id"`       // Foreign Key to the product
	RetailerID       string `json:"retailer_id"`      // Foreign Key to the retailer
	AdvertisedPrice   string `json:"advertised_price"` // Price displayed to the customer
	URL              string `json:"url"`              // Link to the product
}
