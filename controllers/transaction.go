package controllers

import (
	"eventify/config"
	"eventify/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BuyTicketInput struct {
	TicketID uint `json:"ticket_id"`
	Quantity int  `json:"quantity"`
}

func BuyTicket(c echo.Context) error {
	var input BuyTicketInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var ticket models.Ticket
	if err := config.DB.First(&ticket, input.TicketID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Ticket not found")
	}

	if input.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, "Quantity must be greater than 0")
	}

	if ticket.Stock < input.Quantity {
		return c.JSON(http.StatusBadRequest, "Not enough stock")
	}

	userID := c.Get("user_id").(uint)
	total := float64(input.Quantity) * ticket.Price

	transaction := models.Transaction{
		UserID:   userID,
		TicketID: ticket.ID,
		Quantity: input.Quantity,
		Total:    total,
	}

	ticket.Stock -= input.Quantity

	config.DB.Save(&ticket)
	config.DB.Create(&transaction)

	return c.JSON(http.StatusCreated, transaction)
}

func GetMyTransactions(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	var transactions []models.Transaction
	config.DB.Where("user_id = ?", userID).Find(&transactions)
	return c.JSON(http.StatusOK, transactions)
}
