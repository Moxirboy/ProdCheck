package postgres

const(
	createViolation = `INSERT INTO violation (price_id,violation_date,violation_type,severity,screenshot_path) VALUES ($1,$2,$3,$4,$5) RETURNING id`
)