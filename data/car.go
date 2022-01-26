package data

import (
	"time"

	"errors"

	"github.com/korobok404/taxi-cars/entity"
)

// List of test cars
var cars map[string]*entity.Car = map[string]*entity.Car{
	"123": {
		Id:        "123",
		RegPlate:  "H123",
		Brand:     "Toyota",
		Color:     "Black",
		Year:      2020,
		IsReady:   true,
		PosX:      3,
		PosY:      5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	"234": {
		Id:        "234",
		RegPlate:  "H456",
		Brand:     "BMW",
		Color:     "Grey",
		Year:      2019,
		IsReady:   true,
		PosX:      7,
		PosY:      2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	"567": {
		Id:        "567",
		RegPlate:  "H732",
		Brand:     "Kia",
		Color:     "Blue",
		Year:      2018,
		IsReady:   true,
		PosX:      4,
		PosY:      8,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// getCars return all cars
func GetCars() map[string]*entity.Car {
	return cars
}

// addCar add new car
func AddCar(car *entity.Car) bool {
	cars[car.Id] = car
	return true
}

// getCarById return car by unique id
func GetCarById(id string) (*entity.Car, error) {
	car, ok := cars[id]
	if !ok {
		return nil, errors.New("car not found")
	}
	return car, nil
}

// updateCarById change existing car
func UpdateCarById(id string, car *entity.Car) error {
	_, ok := cars[id]
	if !ok {
		return errors.New("car not found")
	}
	cars[id] = car
	return nil
}

// deleteCarById delete existing car
func DeleteCarById(id string) error {
	if _, ok := cars[id]; !ok {
		return errors.New("car not found")
	}
	delete(cars, id)
	return nil
}
