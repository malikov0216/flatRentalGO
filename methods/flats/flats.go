package flats

import (
	"github.com/malikov0216/flatRentalGO/database"
	"github.com/malikov0216/flatRentalGO/models"
)

func GetList() ([]models.Flat, error) {
	const query = `select * from flats`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]models.Flat, 0)

	for rows.Next() {
		flat := models.Flat{}
		err = rows.Scan(&flat.ID, &flat.Name, &flat.Price, &flat.ResidentID, &flat.IsFree)
		if err != nil {
			return nil, err
		}
		results = append(results, models.Flat{flat.ID, flat.Name, flat.Price, flat.ResidentID, flat.IsFree})
	}

	return results, nil
}

func GetBy(id string) (interface{}, error) {
	const query = `select * from flats f WHERE f.id = ($1)`
	flat := models.Flat{}

	row, err := database.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&flat.ID, &flat.Name, &flat.Price, &flat.ResidentID, &flat.IsFree)
		if err != nil {
			return nil, err
		}
	}

	if flat.ID == 0 {
		return map[string]string{"error": "No flat with that id"}, nil
	} else {
		return flat, nil
	}
}

func Add(flat models.Flat) error {
	const query = `INSERT INTO flats(name, price, resident_id, is_free) VALUES($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, flat.Name, flat.Price, flat.ResidentID, flat.IsFree)
	return err
}

func EditBy(flatEdit models.FlatEdit) error {
	const query = `UPDATE flats SET resident_id = ($1) WHERE id = ($2)`
	const caseQuery = `UPDATE flats SET is_free = CASE WHEN resident_id IS NULL THEN true ELSE false END WHERE id = ($1)`
	_, err := database.DB.Exec(query, flatEdit.ResidentID, flatEdit.ID)
	if err != nil {
		return err
	}

	_, err = database.DB.Query(caseQuery, flatEdit.ID)
	return err
}
