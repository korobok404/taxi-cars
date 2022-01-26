package entity

import "time"

type Car struct {
	Id        string    `json:"id"`
	RegPlate  string    `json:"regPlate"`
	Brand     string    `json:"brand"`
	Color     string    `json:"color"`
	Year      uint      `json:"year"`
	IsReady   bool      `json:"isReady"`
	PosX      int       `json:"posX"`
	PosY      int       `json:"posY"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCar() *Car {
	return &Car{}
}
