package database

const (
	InsertRescue = `
	INSERT INTO rescues (rescue_count, rescue_type, urgency, description, street, number, neighborhood, zip_code, complement) 
	VALUES (:rescue_count, :rescue_type, :urgency, :description, :street, :number, :neighborhood, :zip_code, :complement)
	RETURNING id, created_at, updated_at;`

	GetRescues = "SELECT * FROM rescues ORDER BY created_at ASC;"

	GetRescue = "SELECT * FROM rescues WHERE id=$1;"
)
