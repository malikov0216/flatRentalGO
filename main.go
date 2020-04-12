package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/malikov0216/flatRentalGO/database"
	h "github.com/malikov0216/flatRentalGO/handlers"
)

var (
	mySigningKey = "mySecretKey"
)

func main() {

	database.Open()
	defer database.Close()

	router := gin.Default()

	api := router.Group("/api")
	{
		api.Use(cors.Default())
		api.Use(jwt.Auth(mySigningKey))

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

		api.POST("/generateToken", h.GenerateToken)
	}

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
