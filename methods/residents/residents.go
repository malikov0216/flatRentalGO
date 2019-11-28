package residents

import (
	"github.com/malikov0216/flatRental/database"
	"github.com/malikov0216/flatRental/models"
)

func AddResidentMethod(resident models.Resident) error {
	const query = `INSERT INTO residents(name, contact, checkin_date, checkout_date) VALUES($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, resident.Name, resident.Contact, resident.CheckIn, resident.CheckOut)
	return err
}

func GetList() ([]models.Resident, error) {
	const query = `SELECT * from residents`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]models.Resident, 0)

	for rows.Next() {
		resident := models.Resident{}
		err = rows.Scan(&resident.ID, &resident.Name, &resident.Contact, &resident.CheckIn, &resident.CheckOut)
		if err != nil {
			return nil, err
		}
		results = append(results, models.Resident{resident.ID, resident.Name, resident.Contact, resident.CheckIn, resident.CheckOut})
	}

	return results, nil
}

func GetBy(id string) (interface{}, error) {
	const query = `select * from residents r where r.id = ($1)`
	resident := models.Resident{}

	row, err := database.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&resident.ID, &resident.Name, &resident.Contact, &resident.CheckIn, &resident.CheckOut)
		if err != nil {
			return nil, err
		}
	}

	if resident.ID == 0 {
		return map[string]string{"error": "No resident with that id"}, nil
	} else {
		return resident, nil
	}
}
