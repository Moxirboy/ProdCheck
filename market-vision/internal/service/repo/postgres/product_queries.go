package postgres


const(

	createProduct=`
	insert into products(product_id,retailer_id,advertised_price,scrape_date,url) values($1,$2,$3,$4,$5) returning id
	`

	getDetails=`
	select product_id,manufacturer_id,product_name,SKU,MAP from products where id=$1
	`
)