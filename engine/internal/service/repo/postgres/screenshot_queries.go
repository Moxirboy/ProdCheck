package postgres


const(
	insertScreenshot = `UPDATE violation SET screenshot_path = $1 WHERE id = $2;
`
)