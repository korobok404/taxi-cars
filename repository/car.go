package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/korobok404/taxi-cars/entity"
	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(context *gin.Context) *CarRepository {
	return &CarRepository{context.MustGet("db").(*gorm.DB)}
}

// GetCars return all cars
func GetCars(db *gorm.DB) []*entity.Car {
	var cars []*entity.Car
	db.Find(&cars)
	return cars

}

// AddCar add new car
func (carRepo *CarRepository) AddCar(car *entity.Car) error {
	err := carRepo.db.Create(car).Error
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetCarById return car by unique id
func (carRepo *CarRepository) GetCarById(id string) (*entity.Car, error) {
	car := entity.NewCar()
	err := carRepo.db.First(car, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return car, nil
}

// UpdateCarById change existing car
func (carRepo *CarRepository) UpdateCarById(id string, car *entity.Car) error {
	oldCar, err := carRepo.GetCarById(id)
	if err != nil {
		return errors.New(err.Error())
	}

	// When update with struct, GORM will only update non-zero fields
	// https://gorm.io/docs/update.html#Updates-multiple-columns
	err = carRepo.db.Model(oldCar).Updates(map[string]interface{}{
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

// DeleteCarById delete existing car
func (carRepo *CarRepository) DeleteCarById(id string) error {
	car, err := carRepo.GetCarById(id)
	if err != nil {
		return errors.New(err.Error())
	}
	carRepo.db.Delete(car, car.Id)
	return nil
}
