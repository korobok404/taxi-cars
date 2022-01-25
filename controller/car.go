package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korobok404/taxi-cars/data"
	"github.com/korobok404/taxi-cars/entity"
)

// getCars return all cars
func GetCars(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.GetCars())
}

// addCar add new car
func AddCar(context *gin.Context) {
	car := entity.NewCar()

	if err := json.NewDecoder(context.Request.Body).Decode(car); err != nil {
		log.Fatal(err)
	}

	data.AddCar(car)
	context.IndentedJSON(http.StatusCreated, car)
}

// getCarById return car by unique id
func GetCarById(context *gin.Context) {
	id := context.Param("id")

	car, err := data.GetCarById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, car)
}

// updateCarById change existing car
func UpdateCarById(context *gin.Context) {
	id := context.Param("id")

	car := entity.NewCar()
	json.NewDecoder(context.Request.Body).Decode(car)

	if err := data.UpdateCarById(id, car); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "car was updated"})
}

// deleteCarById delete existing car
func DeleteCarById(context *gin.Context) {
	id := context.Param("id")

	if err := data.DeleteCarById(id); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Car was deleted"})
}
