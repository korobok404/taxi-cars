package service

import (
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/korobok404/taxi-cars/entity"
	"github.com/korobok404/taxi-cars/repository"
	"gorm.io/gorm"
)

// Search distance
const distance = 5

func GetNearestCars(x, y int, context *gin.Context) map[string]*entity.Car {
	allCars := repository.GetCars(context.MustGet("db").(*gorm.DB))

	// Result nearest cars
	resultCars := make(map[string]*entity.Car)

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
