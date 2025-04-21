package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	EventID uint    `json:"event_id"`
	Event   Event   `json:"-" gorm:"foreignKey:EventID"`
}
