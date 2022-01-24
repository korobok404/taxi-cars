package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.POST("/cars", createCar)
	router.PUT("/cars/:id", updateCar)
	router.DELETE("/cars/:id", deleteCar)

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

}

// updateCar change existing car
func updateCar(context *gin.Context) {

}

// deleteCar delete existing car
func deleteCar(context *gin.Context) {

}
