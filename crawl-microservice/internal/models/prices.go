package models

import "crawl-microservice/internal/dto"

// Price represents the pricing information for a product.
type Price struct {
	id               string
	productId        string
	retailerId       string
	advertaisedPrice string
	url              string
}
type PriceList struct {
	TotalCount int
	TotalPages int
	Page       int
	Size       int
	HasMore    bool
	Price    []*Price
}

func NewPriceFromModel(model dto.Price) *Price {
	return &Price{
		id:   model.ID,
		productId:       model.ProductID,
		retailerId: model.RetailerID,
		advertaisedPrice: model.AdvertisedPrice,
		url:      model.URL,
	}
}

// NewPrice is a constructor for creating a new Price instance.
func NewPrice( productId, retailerId, advertaisedPrice, url string) *Price {
	return &Price{
		productId:        productId,
		retailerId:       retailerId,
		advertaisedPrice: advertaisedPrice,
		url:              url,
	}
}

// GetId returns the ID of the Price.
func (p *Price) GetId() string {
	return p.id
}

// SetId sets the ID of the Price.
func (p *Price) SetId(id string) {
	p.id = id
}

// GetProductId returns the Product ID.
func (p *Price) GetProductId() string {
	return p.productId
}

// SetProductId sets the Product ID.
func (p *Price) SetProductId(productId string) {
	p.productId = productId
}

// GetRetailerId returns the Retailer ID.
func (p *Price) GetRetailerId() string {
	return p.retailerId
}

// SetRetailerId sets the Retailer ID.
func (p *Price) SetRetailerId(retailerId string) {
	p.retailerId = retailerId
}

// GetAdvertaisedPrice returns the advertised price.
func (p *Price) GetAdvertaisedPrice() string {
	return p.advertaisedPrice
}

// SetAdvertaisedPrice sets the advertised price.
func (p *Price) SetAdvertaisedPrice(advertaisedPrice string) {
	p.advertaisedPrice = advertaisedPrice
}

// GetUrl returns the URL.
func (p *Price) GetUrl() string {
	return p.url
}

// SetUrl sets the URL.
func (p *Price) SetUrl(url string) {
	p.url = url
}
