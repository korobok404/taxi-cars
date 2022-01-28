package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/korobok404/taxi-cars/controller"
	"github.com/korobok404/taxi-cars/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/cars", controller.GetCars)
	router.POST("/cars", controller.AddCar)
	router.GET("/cars/:id", controller.GetCarById)
	router.PUT("/cars/:id", controller.UpdateCarById)
	router.DELETE("/cars/:id", controller.DeleteCarById)
	router.GET("/cars/nearest", controller.GetNearestCars)

	// Server settings
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("Error starting server:", err)
			os.Exit(1)
		}
	}()

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Gracefully shutting down.")

	// Gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down:", err)
	}

}

func initDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("taxi.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate models
	db.AutoMigrate(&entity.Car{})

	// Create test car
	// db.Create(&entity.Car{
	// 	Id:        "555",
	// 	RegPlate:  "H321",
	// 	Brand:     "Toyota",
	// 	Color:     "Black",
	// 	Year:      2020,
	// 	IsReady:   true,
	// 	PosX:      3,
	// 	PosY:      5,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })
	return db
}
