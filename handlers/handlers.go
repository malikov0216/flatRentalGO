package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malikov0216/flatRental/methods/flats"
	"github.com/malikov0216/flatRental/methods/residents"
	"github.com/malikov0216/flatRental/models"
)

// AddFlat : Adding new flat to Database
func AddFlat(ctx *gin.Context) {
	flat := models.Flat{}
	err := ctx.ShouldBindJSON(&flat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		err = flats.Add(flat)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		}
	}
}

// GetFlatsAll : Getting out all flats
func GetFlatsAll(ctx *gin.Context) {
	result, err := flats.GetList()
	if err != nil {
		ctx.JSON(500, gin.H{"status": err.Error()})
	}

	ctx.JSON(200, result)
}

// EditFlat : Edit Flats resident ID
func EditFlat(ctx *gin.Context) {
	flatEdit := models.FlatEdit{}
	err := ctx.ShouldBindJSON(&flatEdit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		err = flats.EditBy(flatEdit)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

/*
GetFlatByID : Getting out a flat through passing id
id : ID of flat
*/
func GetFlatByID(ctx *gin.Context) {
	id := ctx.Query("id")

	result, err := flats.GetBy(id)
	if err != nil {
		ctx.JSON(500, gin.H{"status": err.Error()})
	}

	ctx.JSON(200, result)
}

// AddResident : Adding new resident to Database
func AddResident(ctx *gin.Context) {
	resident := models.Resident{}
	err := ctx.ShouldBindJSON(&resident)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		err = residents.AddResidentMethod(resident)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		}
	}
}

func GetResidents(ctx *gin.Context) {
	result, err := residents.GetList()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, result)
}

func GetResidentByID(ctx *gin.Context) {
	id := ctx.Query("id")

	result, err := residents.GetBy(id)
	if err != nil {
		ctx.JSON(500, gin.H{"status": err.Error()})
	}

	ctx.JSON(200, result)
}
