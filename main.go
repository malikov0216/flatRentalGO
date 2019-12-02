package main

import (
	"github.com/malikov0216/flatRental/database"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	h "github.com/malikov0216/flatRental/handlers"
)

func main() {

	database.Open()
	defer database.Close()

	route := gin.Default()

	route.GET("/flats", h.GetFlatsAll)
	route.GET("/flat", h.GetFlatByID)
	route.PUT("/flat", h.EditFlat)
	route.POST("/flat", h.AddFlat)

	route.GET("/residents", h.GetResidents)
	route.GET("/resident", h.GetResidentByID)
	route.POST("/resident", h.AddResident)
	route.PUT("/resident", h.EditResident)

	if err := route.Run(":8080"); err != nil {
		panic(err)
	}
}
