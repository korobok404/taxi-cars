package entity

import (
	"time"
)

type Car struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	RegNum    string `json:"regNum" binding:"required"`
	Brand     string `json:"brand" binding:"required"`
	Color     string `json:"color" binding:"required"`
	Year      uint   `json:"year" binding:"required"`
	IsReady   bool   `json:"isReady"`
	LocX      int    `json:"locX" binding:"required"`
	LocY      int    `json:"locY" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCar() *Car {
	return &Car{}
}
