package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", createCar)
	router.GET("/cars/:id", getCarById)
	router.PUT("/cars/:id", updateCarById)
	router.DELETE("/cars/:id", deleteCarById)

	router.Run(":8080")
}

type Car struct {
	Id        string    `json:"id"`
	RegPlate  string    `json:"regPlate"`
	Brand     string    `json:"brand"`
	Color     string    `json:"color"`
	Year      uint      `json:"year"`
	IsReady   bool      `json:"isReady"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// List of test cars
var cars map[string]*Car = map[string]*Car{
	"123": {"123", "H123", "Toyota", "Black", 2020, true, time.Now(), time.Now()},
	"234": {"234", "H456", "BMW", "Grey", 2019, false, time.Now(), time.Now()},
}

// getCars return all cars
func getCars(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, cars)
}

// createCar add new car
func createCar(context *gin.Context) {
	car := new(Car)
	if err := json.NewDecoder(context.Request.Body).Decode(car); err != nil {
		log.Fatal(err)
	}

	cars[car.Id] = car
	context.IndentedJSON(http.StatusCreated, car)
}

// getCarById return car by unique id
func getCarById(context *gin.Context) {
	id := context.Param("id")
	car, ok := cars[id]
	if !ok {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, car)
}

// updateCarById change existing car
func updateCarById(context *gin.Context) {
	id := context.Param("id")
	car, ok := cars[id]
	if !ok {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}
	json.NewDecoder(context.Request.Body).Decode(car)
	context.IndentedJSON(http.StatusOK, car)
}

// deleteCarById delete existing car
func deleteCarById(context *gin.Context) {
	id := context.Param("id")
	if _, ok := cars[id]; !ok {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
		return
	}
	delete(cars, id)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Car was deleted"})
}
