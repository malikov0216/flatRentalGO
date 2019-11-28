package residents

import (
	"fmt"

	"github.com/malikov0216/flatRental/database"
	"github.com/malikov0216/flatRental/models"
)

func AddResidentMethod(resident models.Resident) error {
	const query = `INSERT INTO residents(name, contact, checkin_date, checkout_date) VALUES($1, $2, $3, $4)`
	_, err := database.DB.Exec(query, resident.Name, resident.Contact, resident.CheckIn, resident.CheckOut)
	fmt.Println("Test: ", err)
	return err
}
