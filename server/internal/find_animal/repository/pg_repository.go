package repository

import (
	"github.com/francoggm/save-rs-brazil/internal/models"
	"github.com/jmoiron/sqlx"
)

func CreateFindAnimal(db *sqlx.DB, findAnimal *models.FindAnimal) error {
	result, err := db.NamedQuery(queryInsertFindAnimal, findAnimal)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&findAnimal.ID, &findAnimal.CreatedAt, &findAnimal.UpdatedAt)
	}

	return err
}

func GetFindAnimals(db *sqlx.DB) ([]*models.FindAnimal, error) {
	var findAnimals []*models.FindAnimal

	err := db.Select(&findAnimals, queryGetFindAnimals)
	if err != nil {
		return nil, err
	}

	return findAnimals, nil
}

func GetFindAnimal(db *sqlx.DB, id int) (*models.FindAnimal, error) {
	findAnimal := new(models.FindAnimal)

	err := db.Get(findAnimal, queryGetFindAnimal, id)
	if err != nil {
		return nil, err
	}

	return findAnimal, nil
}
