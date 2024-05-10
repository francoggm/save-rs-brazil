package repository

const (
	queryInsertFindAnimal = `
	INSERT INTO find_animals (specie, race, gender, description, street, number, neighborhood, zip_code, complement) 
	VALUES (:specie, :race, :gender, :description, :street, :number, :neighborhood, :zip_code, :complement)
	RETURNING id, created_at, updated_at;`

	queryGetFindAnimals = "SELECT * FROM find_animals ORDER BY created_at ASC;"

	queryGetFindAnimal = "SELECT * FROM find_animals WHERE id=$1;"
)
