package repository

import (
	"errors"

	"github.com/korobok404/taxi-cars/entity"
	"gorm.io/gorm"
)

// getCars return all cars
func GetCars(db *gorm.DB) []*entity.Car {
	var cars []*entity.Car
	db.Find(&cars)
	return cars

}

// addCar add new car
func AddCar(car *entity.Car, db *gorm.DB) error {
	err := db.Create(car).Error
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// getCarById return car by unique id
func GetCarById(id string, db *gorm.DB) (*entity.Car, error) {
	car := entity.NewCar()
	err := db.First(car, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return car, nil
}

// updateCarById change existing car
func UpdateCarById(id string, car *entity.Car, db *gorm.DB) error {
	oldCar, err := GetCarById(id, db)
	if err != nil {
		return errors.New(err.Error())
	}

	// When update with struct, GORM will only update non-zero fields
	// https://gorm.io/docs/update.html#Updates-multiple-columns
	err = db.Model(oldCar).Updates(map[string]interface{}{
		"reg_plate":  car.RegPlate,
		"brand":      car.Brand,
		"color":      car.Color,
		"year":       car.Year,
		"is_ready":   car.IsReady,
		"pos_x":      car.PosX,
		"pos_y":      car.PosY,
		"created_at": car.CreatedAt,
		"updated_at": car.UpdatedAt,
	}).Error

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// deleteCarById delete existing car
func DeleteCarById(id string, db *gorm.DB) error {
	car, err := GetCarById(id, db)
	if err != nil {
		return errors.New(err.Error())
	}
	db.Delete(car, car.Id)
	return nil
}
