package service

import (
	"github.com/korobok404/taxi-cars/data"
	"github.com/korobok404/taxi-cars/entity"
)

func GetNearestCars(x, y int) map[string]*entity.Car {
	cars := data.GetCars()

	resultCars := make(map[string]*entity.Car)
	for id, car := range cars {
		if car.IsReady {
			resultCars[id] = car
		}
	}
	return resultCars
}
