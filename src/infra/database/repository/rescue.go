package repository

import (
	"github.com/francoggm/save-rs-brazil/infra/database"
	"github.com/francoggm/save-rs-brazil/models"
)

func CreateRescue(db *database.DB, rescue *models.Rescue) error {
	result, err := db.Conn.NamedQuery(database.InsertRescue, rescue)
	if err != nil {
		return err
	}

	if result.Next() {
		err = result.Scan(&rescue.ID, &rescue.CreatedAt, &rescue.UpdatedAt)
	}

	return err
}

func GetRescues(db *database.DB) ([]*models.Rescue, error) {
	var rescues []*models.Rescue

	err := db.Conn.Select(&rescues, database.GetRescues)
	if err != nil {
		return nil, err
	}

	return rescues, nil
}

func GetRescue(db *database.DB, id int) (*models.Rescue, error) {
	rescue := new(models.Rescue)

	err := db.Conn.Get(rescue, database.GetRescue, id)
	if err != nil {
		return nil, err
	}

	return rescue, nil
}
