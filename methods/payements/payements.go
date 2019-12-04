package payements

import (
	"github.com/malikov0216/flatRentalGO/database"
	"github.com/malikov0216/flatRentalGO/models"
)

// Add : Adding new payement to DB
func Add(payement models.Payement) error {
	const query = `INSERT INTO payements(reason, accepted_person, amount, electric, resident_id, date) VALUES($1, $2, $3, $4, $5, $6)`
	_, err := database.DB.Exec(query, payement.Reason, payement.AcceptedPerson, payement.Amount, payement.Electric, payement.ResidentID, payement.Date)
	return err
}

func GetList() ([]models.Payement, error) {
	const query = `SELECT * from payements`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]models.Payement, 0)

	for rows.Next() {
		payement := models.Payement{}
		err = rows.Scan(&payement.ID, &payement.Reason, &payement.AcceptedPerson, &payement.Amount, &payement.Electric, &payement.ResidentID, &payement.Date)
		if err != nil {
			return nil, err
		}
		results = append(results, models.Payement{payement.ID, payement.Reason, payement.AcceptedPerson, payement.Amount, payement.Electric, payement.ResidentID, payement.Date})
	}

	return results, nil
}

func GetByResidentId(residentID string) ([]models.Payement, error) {
	const query = `SELECT * from payements where resident_id = ($1)`

	rows, err := database.DB.Query(query, residentID)
	if err != nil {
		return nil, err
	}

	results := make([]models.Payement, 0)

	for rows.Next() {
		payement := models.Payement{}
		err = rows.Scan(&payement.ID, &payement.Reason, &payement.AcceptedPerson, &payement.Amount, &payement.Electric, &payement.ResidentID, &payement.Date)
		if err != nil {
			return nil, err
		}
		results = append(results, models.Payement{payement.ID, payement.Reason, payement.AcceptedPerson, payement.Amount, payement.Electric, payement.ResidentID, payement.Date})
	}

	return results, nil
}
