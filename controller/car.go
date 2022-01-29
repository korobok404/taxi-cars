package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/korobok404/taxi-cars/entity"
	"github.com/korobok404/taxi-cars/repository"
	"github.com/korobok404/taxi-cars/service"
	"gorm.io/gorm"
)

// GetCars return all cars
func GetCars(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, repository.GetCars(context.MustGet("db").(*gorm.DB)))
}

// AddCar add new car
func AddCar(context *gin.Context) {
	car := entity.NewCar()

	if err := json.NewDecoder(context.Request.Body).Decode(car); err != nil {
		log.Fatal(err)
	}

	if err := repository.NewCarRepository(context).AddCar(car); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusCreated, car)
}

// GetCarById return car by unique id
func GetCarById(context *gin.Context) {
	id := context.Param("id")

	car, err := repository.NewCarRepository(context).GetCarById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, car)
}

// UpdateCarById change existing car
func UpdateCarById(context *gin.Context) {
	id := context.Param("id")

	car := entity.NewCar()
	json.NewDecoder(context.Request.Body).Decode(car)

	if err := repository.NewCarRepository(context).UpdateCarById(id, car); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "car was updated"})
}

// DeleteCarById delete existing car
func DeleteCarById(context *gin.Context) {
	id := context.Param("id")

	if err := repository.NewCarRepository(context).DeleteCarById(id); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Car was deleted"})
}

// GetNearestCars return nearest cars that are ready to accept the order
func GetNearestCars(context *gin.Context) {
	clientX, errX := strconv.Atoi(context.Query("x"))

	clientY, errY := strconv.Atoi(context.Query("y"))
	if errX != nil || errY != nil {
		//TODO: Add middleware with validation
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Coordinaate error"})
		return
	}

	cars := service.GetNearestCars(clientX, clientY, context)
	context.IndentedJSON(http.StatusOK, cars)
}
