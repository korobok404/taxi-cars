package data

import (
	"time"

	"errors"

	"github.com/korobok404/taxi-cars/entity"
)

// List of test cars
var cars map[string]*entity.Car = map[string]*entity.Car{
	"123": {"123", "H123", "Toyota", "Black", 2020, true, time.Now(), time.Now()},
	"234": {"234", "H456", "BMW", "Grey", 2019, false, time.Now(), time.Now()},
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
