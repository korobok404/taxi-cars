package service

import (
	"github.com/korobok404/taxi-cars/data"
	"github.com/korobok404/taxi-cars/entity"
)

func GetNearestCars(x, y int) map[string]*entity.Car {
	cars := data.GetCars()

	resultCars := make(map[string]*entity.Car)
	for _, car := range cars {
		if car.IsReady {
			resultCars[car.Id] = car
		}
	}
	return resultCars
}
