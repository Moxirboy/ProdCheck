package crawl

import (
	configs "market-vision/internal/config"
	"market-vision/internal/model"
	"market-vision/pkg/client/browser"
	"time"

	helper "market-vision/prod-business"
)

func Crawl(cfg configs.Config, product model.Url) (model.Price, error) {
	browser := browser.InitializeChrome(cfg)
	htmlContent, err := browser.FetchHTML(product.Url, 10*time.Second)
	if err != nil {

	}
	xph, err := helper.NewXPathHelper(htmlContent)
	if err != nil {

	}

	item := getProductDetails(xph, product)

	return item, nil
}
func getProductDetails(xph *helper.XPathHelper, product model.Url) model.Price {
	item := model.NewPrice(
		product.ProdcutId,
		product.RetailerId,
		xph.ExtractByXPath(""),
		xph.ExtractByXPath(""),
	)
	return *item
}
