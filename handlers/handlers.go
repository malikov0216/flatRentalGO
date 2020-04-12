package handlers

import (
	"fmt"
	"reflect"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	// "github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/malikov0216/flatRentalGO/methods/flats"
	"github.com/malikov0216/flatRentalGO/methods/payements"
	"github.com/malikov0216/flatRentalGO/methods/residents"
	"github.com/malikov0216/flatRentalGO/models"
)

// AddFlat : Adding new flat to Database
func AddFlat(ctx *gin.Context) {
	flat := models.Flat{}
	err := ctx.ShouldBindJSON(&flat)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		err = flats.Add(flat)
		if err != nil {
			ctx.JSON(500, gin.H{"status": err.Error()})
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

// EditFlat : Edit Flats resident ID
func EditFlat(ctx *gin.Context) {
	flatEdit := models.FlatEdit{}
	err := ctx.ShouldBindJSON(&flatEdit)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		err = flats.EditBy(flatEdit)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
	}
}

// AddResident : Adding new resident to Database
func AddResident(ctx *gin.Context) {
	resident := models.Resident{}
	err := ctx.ShouldBindJSON(&resident)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		err = residents.Add(resident)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
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

func EditResident(ctx *gin.Context) {
	resident := models.Resident{}
	err := ctx.ShouldBindJSON(&resident)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		err = residents.Edit(resident)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
	}
}

func AddPayement(ctx *gin.Context) {
	payement := models.Payement{}
	err := ctx.ShouldBindJSON(&payement)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		err = payements.Add(payement)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
	}
}

func GetPayements(ctx *gin.Context) {
	result, err := payements.GetList()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, result)
}

func GetPayementByResidentID(ctx *gin.Context) {
	residentID := ctx.Query("residentID")
	fmt.Println(reflect.TypeOf(residentID))
	result, err := payements.GetByResidentId(residentID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, result)
}

var (
	mySigningKey = "mySecretKey"
)

func GenerateToken(ctx *gin.Context) {

	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"email":   "mastaok02@gmail.com",
		"telefon": "+77777739673",
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not generate token"})
	}
	ctx.JSON(200, gin.H{"token": tokenString})
}
