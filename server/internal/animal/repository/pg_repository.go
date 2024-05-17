package repository

import (
	"github.com/francoggm/save-rs-brazil/internal/models"
	"github.com/jmoiron/sqlx"
)

func CreateAnimal(db *sqlx.DB, animal *models.Animal) error {
	result, err := db.NamedQuery(queryInsertAnimal, animal)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&animal.ID, &animal.CreatedAt, &animal.UpdatedAt)
	}

	return err
}

func GetAnimals(state string, db *sqlx.DB) ([]*models.Animal, error) {
	var animals []*models.Animal
	var err error

	switch state {
	case "lost":
		err = db.Select(&animals, queryGetAnimalsByState, 1)
	case "found":
		err = db.Select(&animals, queryGetAnimalsByState, 2)
	case "adoption":
		err = db.Select(&animals, queryGetAnimalsByState, 3)
	default:
		err = db.Select(&animals, queryGetAnimals)
	}

	if err != nil {
		return nil, err
	}

	return animals, nil
}

func GetAnimal(id int, state string, db *sqlx.DB) (*models.Animal, error) {
	animal := new(models.Animal)

	var err error

	switch state {
	case "lost":
		err = db.Get(animal, queryGetAnimalByState, 1)
	case "found":
		err = db.Get(animal, queryGetAnimalByState, 2)
	case "adoption":
		err = db.Get(animal, queryGetAnimalByState, 3)
	default:
		err = db.Get(animal, queryGetAnimal, id)
	}
 
	if err != nil {
		return nil, err
	}

	return animal, nil
}
