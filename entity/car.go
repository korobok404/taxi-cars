package entity

import (
	"time"
)

type Car struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	RegPlate  string `json:"regPlate" binding:"required"`
	Brand     string `json:"brand" binding:"required"`
	Color     string `json:"color" binding:"required"`
	Year      uint   `json:"year" binding:"required"`
	IsReady   bool   `json:"isReady"`
	PosX      int    `json:"posX" binding:"required"`
	PosY      int    `json:"posY" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCar() *Car {
	return &Car{}
}
