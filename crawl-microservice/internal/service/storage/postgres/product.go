package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"crawl-microservice/internal/models"
	"crawl-microservice/internal/service/storage/repo"
	"crawl-microservice/pkg/logger"
	"crawl-microservice/pkg/utils"
	"github.com/pkg/errors"
)

type product struct {
	db  *sql.DB
	log logger.Logger
}

func NewProduct(db *sql.DB, log logger.Logger) repo.Product {
	return &product{
		db:  db,
		log: log,
	}
}

func (r *product) Create(ctx context.Context, product *models.Product) error {
	if err := r.db.QueryRowContext(ctx, CreateProduct,
		product.ManufacturerID, product.ProductName, product.SKU, product.Map,
	).Scan(&product.Id); err != nil {
		r.log.Error("could not create product:" + err.Error())
		return err
	}

	return nil
}

func (r *product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product := models.Product{}

	if err := r.db.QueryRowContext(ctx, GetProduct, id).Scan(
		&product.ManufacturerID, &product.ProductName, &product.SKU, &product.Map); err != nil {
		r.log.Error("could not get product by id: ", err.Error())
		return nil, err
	}

	return &product, nil
}

func (r *product) GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error) {
	var (
		totalCount     int
		GetProductList = `
	select id, manufacturer_id, product_name, SKU, MAP_price,category from products where manufacturer_id=$1;
`
		GetTotalCount = `
SELECT COUNT(id) FROM products WHERE 1=1
`
	)
	if name != "" {
		GetTotalCount = fmt.Sprintf("%s%s", GetTotalCount, " and name LIKE '%"+name+"%'")
		GetProductList = fmt.Sprintf("%s%s", GetProductList, " and name LIKE '%"+name+"%'")
	}
	GetProductList += " ORDER BY created_at OFFSET $1 LIMIT $2"
	if err := r.db.QueryRowContext(ctx, GetTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "product.GetAll.QueryContext")
	}
	if totalCount == 0 {
		return &models.ProductsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Product:    make([]*models.Product, 0),
		}, nil
	}
	rows, err := r.db.QueryContext(ctx, GetProductList, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "product.GetAll.QueryContext")
	}
	defer rows.Close()
	productList := make([]*models.Product, 0, query.GetSize())
	for rows.Next() {
		product := &models.Product{}
		if err = rows.Scan(
			&product.Id, &product.ManufacturerID, &product.ProductName, &product.SKU, &product.Map,
		); err != nil {
			return nil, errors.Wrap(err, "product.GetAll.Scan")
		}
		productList = append(productList, product)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "product.GetAll.Row.Err")
	}
	return &models.ProductsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Product:    productList,
	}, nil
}

func (r *product) GetListViolations(ctx context.Context, name string, query utils.PaginationQuery) (*models.ViolationList, error) {
	var (
		totalCount     int
		GetViolationsList = `
			select id,price_id,violation_date,violation_type,severity,screenshot_path from violations where price_id=$1 and 1=1`

		GetTotalCount = `
SELECT COUNT(id) FROM violations WHERE 1=1
`
	)
	if name != "" {
		GetTotalCount = fmt.Sprintf("%s%s", GetTotalCount, " and name LIKE '%"+name+"%'")
		GetViolationsList = fmt.Sprintf("%s%s", GetViolationsList, " and name LIKE '%"+name+"%'")
	}
	GetViolationsList += " ORDER BY created_at OFFSET $1 LIMIT $2"
	if err := r.db.QueryRowContext(ctx, GetTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "violation.GetAll.QueryContext")
	}
	if totalCount == 0 {
		return &models.ViolationList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Violation:    make([]*models.Violation, 0),
		}, nil
	}
	rows, err := r.db.QueryContext(ctx, GetViolationsList, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "violation.GetAll.QueryContext")
	}
	defer rows.Close()
	ViolationList := make([]*models.Violation, 0, query.GetSize())
	for rows.Next() {
		violation := &models.Violation{}
		if err = rows.Scan(
			&violation.ID, &violation.PriceID, &violation.ViolationDate, &violation.ViolationType, &violation.Severity, &violation.ScreenshotPath,
		); err != nil {
			return nil, errors.Wrap(err, "violation.GetAll.Scan")
		}
		ViolationList = append(ViolationList, violation)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "violation.GetAll.Row.Err")
	}
	return &models.ViolationList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Violation:    ViolationList,
	}, nil
}

func (r *product) GetListCrawledPrices(ctx context.Context, name string, query utils.PaginationQuery) (*models.PriceList, error) {
	var (
		totalCount     int
		GetProductList = `
			select id,product_id,retailer_id,advertised_price,url from crawled_prices where product_id=$1 and 1=1`

		GetTotalCount = `
SELECT COUNT(id) FROM crawled_prices WHERE 1=1
`

	)
	if name != "" {
		GetTotalCount = fmt.Sprintf("%s%s", GetTotalCount, " and name LIKE '%"+name+"%'")
		GetProductList = fmt.Sprintf("%s%s", GetProductList, " and name LIKE '%"+name+"%'")
	}
	GetProductList += " ORDER BY created_at OFFSET $1 LIMIT $2"
	if err := r.db.QueryRowContext(ctx, GetTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "product.GetAll.QueryContext")
	}
	if totalCount == 0 {
		return &models.PriceList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Price:    make([]*models.Price, 0),
		}, nil
	}
	rows, err := r.db.QueryContext(ctx, GetProductList, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "product.GetAll.QueryContext")
	}
	defer rows.Close()
	priceList := make([]*models.Price, 0, query.GetSize())
	for rows.Next() {
		price := &models.Price{}
		var(
			id=""
			productId=""
			retailerId=""
			advertisedPrice=""
			url=""
		)
		
		if err = rows.Scan(
			&id, &productId, &retailerId, &advertisedPrice, &url,
			); err != nil {
			return nil, errors.Wrap(err, "product.GetAll.Scan")
		}
		price.SetId(id)
		price.SetProductId(productId)
		price.SetRetailerId(retailerId)
		price.SetAdvertaisedPrice(advertisedPrice)
		price.SetUrl(url)

		priceList = append(priceList, price)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "product.GetAll.Row.Err")
	}
	return &models.PriceList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),

		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Price:    priceList,
	}, nil
}

