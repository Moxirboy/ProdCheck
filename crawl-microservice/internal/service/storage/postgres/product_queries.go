package postgres

const (
	CreateProduct = `
	insert into products (manufacturer_id, product_name, SKU, MAP_price) values ($1, $2, $3, $4) returning id;
`
	GetProduct = `
	select id, manufacturer_id, product_name, SKU, MAP_price,category from products where id=$1;
`
	GetViolations=`
	select id,price_id,violation_date,violation_type,severity,screenshot_path from violation where price_id=$1;`
	GetCrawledData=`
	select id,product_id,retailer_id,price,advertised_price,scrape_date,url from prices where retailer_id=$1;
	`
)
