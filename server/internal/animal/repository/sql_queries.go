package repository

const (
	queryInsertAnimal = `
	INSERT INTO animals (specie, race, gender, state, description, street, number, neighborhood, zip_code, complement, phone) 
	VALUES (:specie, :race, :gender, :state, :description, :street, :number, :neighborhood, :zip_code, :complement, :phone)
	RETURNING id, created_at, updated_at;`

	queryGetAnimals = "SELECT * FROM animals ORDER BY created_at ASC;"

	queryGetAnimalsByState = "SELECT * FROM animals WHERE state=$1 ORDER BY created_at ASC;"

	queryGetAnimal = "SELECT * FROM animals WHERE id=$1;"

	queryGetAnimalByState = "SELECT * FROM animals WHERE id=$1 AND state=$2"
)
