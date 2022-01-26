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
)

func main() {
	router := gin.Default()
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
