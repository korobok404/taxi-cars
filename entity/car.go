package entity

import (
	"time"
)

type Car struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	RegPlate  string `json:"regPlate"`
	Brand     string `json:"brand"`
	Color     string `json:"color"`
	Year      uint   `json:"year"`
	IsReady   bool   `json:"isReady"`
	PosX      int    `json:"posX"`
	PosY      int    `json:"posY"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCar() *Car {
	return &Car{}
}
