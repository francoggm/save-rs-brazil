package repository

import (
	"github.com/francoggm/save-rs-brazil/internal/models"
	"github.com/jmoiron/sqlx"
)

func CreateRescue(db *sqlx.DB, rescue *models.Rescue) error {
	result, err := db.NamedQuery(queryInsertRescue, rescue)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&rescue.ID, &rescue.CreatedAt, &rescue.UpdatedAt)
	}

	return err
}

func GetRescues(db *sqlx.DB) ([]*models.Rescue, error) {
	var rescues []*models.Rescue

	err := db.Select(&rescues, queryGetRescues)
	if err != nil {
		return nil, err
	}

	return rescues, nil
}

func GetRescue(db *sqlx.DB, id int) (*models.Rescue, error) {
	rescue := new(models.Rescue)

	err := db.Get(rescue, queryGetRescue, id)
	if err != nil {
		return nil, err
	}

	return rescue, nil
}
