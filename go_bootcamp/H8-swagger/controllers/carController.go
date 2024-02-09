package controllers

import (
	"fmt"
	"go_bootcamp/H8-swagger/database"
	"go_bootcamp/H8-swagger/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

var CarDatas = []models.Car{}

// GetCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars/getCar [get]
func GetAllCar(ctx *gin.Context) {
	var db = database.GetDB()
	// fmt.Println("slebski", *db)
	if db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas : ", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetCar godoc
// @Summary Get details fro a given id
// @Description Get details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{id} [get]
func GetOneCars(ctx *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car

	err := db.First(&carOne, "Id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Data One": carOne})
}

// CreateCars godoc
// @Summary Post details for a given
// @Description Post details of car corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCars(ctx *gin.Context) {
	var db = database.GetDB()
	//Validate Input

	var input models.Car
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carinput := models.Car{Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars}
	db.Create(&carinput)

	ctx.JSON(http.StatusOK, gin.H{"data": carinput})
}

// UpdateCars godoc
// @Summary Update car indetified by the given id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCars(ctx *gin.Context) {
	var db = database.GetDB()

	var car models.Car
	err := db.First(&car, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Car
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

// DeleteCar godoc
// @Summary Delete car identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCar(ctx *gin.Context) {
	var db = database.GetDB()

	var carDelete models.Car

	var car models.Car
	err := db.First(&car, "Id = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
