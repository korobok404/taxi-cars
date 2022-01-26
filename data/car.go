package data

import (
	"errors"
	"time"

	"github.com/korobok404/taxi-cars/entity"
)

// getCars return all cars
func GetCars() []*entity.Car {
	return cars
}

// addCar add new car
func AddCar(car *entity.Car) bool {
	cars = append(cars, car)
	return true
}

// getCarById return car by unique id
func GetCarById(id string) (*entity.Car, error) {
	for _, car := range cars {
		if car.Id == id {
			return car, nil
		}
	}
	return nil, errors.New("car not found")
}

// updateCarById change existing car
func UpdateCarById(id string, car *entity.Car) error {
	oldCar, err := GetCarById(id)
	if err != nil {
		return err
	}
	*oldCar = *car
	// TODO do not overwrite Id
	oldCar.Id = id
	return nil
}

// deleteCarById delete existing car
func DeleteCarById(id string) error {
	for i, car := range cars {
		if car.Id == id {
			cars = append(cars[:i], cars[i+1:]...)
			return nil
		}
	}
	return errors.New("car not found")
}

// List of test cars
var cars []*entity.Car = []*entity.Car{
	{
		Id:        "123",
		RegPlate:  "H321",
		Brand:     "Toyota",
		Color:     "Black",
		Year:      2020,
		IsReady:   true,
		PosX:      3,
		PosY:      5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "234",
		RegPlate:  "H432",
		Brand:     "BMW",
		Color:     "Grey",
		Year:      2019,
		IsReady:   true,
		PosX:      7,
		PosY:      2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "567",
		RegPlate:  "H765",
		Brand:     "Kia",
		Color:     "Blue",
		Year:      2018,
		IsReady:   true,
		PosX:      4,
		PosY:      8,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "891",
		RegPlate:  "H198",
		Brand:     "Mazda",
		Color:     "Red",
		Year:      2021,
		IsReady:   true,
		PosX:      8,
		PosY:      2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "912",
		RegPlate:  "H219",
		Brand:     "Lexus",
		Color:     "Black",
		Year:      2020,
		IsReady:   true,
		PosX:      4,
		PosY:      8,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "900",
		RegPlate:  "H009",
		Brand:     "Kia",
		Color:     "Blue",
		Year:      2018,
		IsReady:   true,
		PosX:      3,
		PosY:      3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "901",
		RegPlate:  "H109",
		Brand:     "Kia",
		Color:     "White",
		Year:      2022,
		IsReady:   true,
		PosX:      8,
		PosY:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "502",
		RegPlate:  "H205",
		Brand:     "Honda",
		Color:     "Red",
		Year:      2022,
		IsReady:   true,
		PosX:      7,
		PosY:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "303",
		RegPlate:  "H303",
		Brand:     "Mercedes",
		Color:     "Blue",
		Year:      2022,
		IsReady:   true,
		PosX:      1,
		PosY:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
