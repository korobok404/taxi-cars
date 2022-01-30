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

	v1 := router.Group("/v1")
	{
		v1.GET("/cars", controller.GetCars)
		v1.POST("/cars", controller.AddCar)
		v1.GET("/cars/:id", controller.GetCarById)
		v1.PUT("/cars/:id", controller.UpdateCarById)
		v1.DELETE("/cars/:id", controller.DeleteCarById)
		v1.GET("/cars/nearest", controller.GetNearestCars)
		v1.PUT("/cars/:id/reserve", controller.ReserveCar)
	}

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
	return db
}
