package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID   uint    `json:"user_id"`
	User     User    `json:"-" gorm:"foreignKey:UserID"`
	TicketID uint    `json:"ticket_id"`
	Ticket   Ticket  `json:"-" gorm:"foreignKey:TicketID"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
