package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	UserID      uint   `json:"user_id"` // foreign key ke User
	User        User   `json:"-" gorm:"foreignKey:UserID"`
}
