package repository

import (
	"errors"
	"runtime"
	"sync"

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
func (carRepo *CarRepository) GetCars() []*entity.Car {
	var cars []*entity.Car
	carRepo.db.Find(&cars)
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
		"reg_plate": car.RegPlate,
		"brand":     car.Brand,
		"color":     car.Color,
		"year":      car.Year,
		"is_ready":  car.IsReady,
		"pos_x":     car.PosX,
		"pos_y":     car.PosY,
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

// GetNearestCars return nearest cars by client's coordinates
func (carRepo *CarRepository) GetNearestCars(x, y int) map[uint]*entity.Car {
	// Search distance
	const distance = 5

	// All cars in DB
	allCars := carRepo.GetCars()

	// Result nearest cars
	resultCars := make(map[uint]*entity.Car)

	// Search boundary coordinates
	minX, maxX := x-distance, x+distance
	minY, maxY := y-distance, y+distance

	// Total cars
	n := len(allCars)
	// Cars per goroutine
	step := n / runtime.GOMAXPROCS(0)
	if step == 0 {
		step = n
	}

	var mu sync.Mutex
	var goCount uint = 0
	quit := make(chan bool)

	for i := 0; i < n; i += step {
		max := i + step
		if max > n {
			max = n
		}

		go func(cars []*entity.Car, q chan bool) {
			for _, car := range cars {
				if car.IsReady && (car.PosX >= minX && car.PosX <= maxX) && (car.PosY >= minY && car.PosY <= maxY) {
					mu.Lock()
					resultCars[car.Id] = car
					mu.Unlock()
				}
			}
			q <- true
		}(allCars[i:max], quit)

		goCount++
	}

	for goCount > 0 {
		<-quit
		goCount--
	}

	return resultCars
}

func (carRepo *CarRepository) ReserveCarById(id string) error {
	car, err := carRepo.GetCarById(id)
	if err != nil {
		return errors.New(err.Error())
	}
	carRepo.db.Model(car).Update("IsReady", false)
	return nil
}
