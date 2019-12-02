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

	api := route.Group("/api")
	{
		api.GET("/flats", h.GetFlatsAll)
		api.GET("/flat", h.GetFlatByID)
		api.PUT("/flat", h.EditFlat)
		api.POST("/flat", h.AddFlat)

		api.GET("/residents", h.GetResidents)
		api.GET("/resident", h.GetResidentByID)
		api.POST("/resident", h.AddResident)
		api.PUT("/resident", h.EditResident)

		api.POST("/payement", h.AddPayement)
		api.GET("/payements", h.GetPayements)
		api.GET("/payement", h.GetPayementByResidentID)
	}

	if err := route.Run(":8080"); err != nil {
		panic(err)
	}
}
