package controllers

import (
	"eventify/config"
	"eventify/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTicket(c echo.Context) error {
	eventID := c.Param("event_id")
	var event models.Event
	if err := config.DB.First(&event, eventID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}

	userID := c.Get("user_id").(uint)
	if event.UserID != userID {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	ticket := new(models.Ticket)
	if err := c.Bind(ticket); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ticket.EventID = event.ID

	config.DB.Create(&ticket)
	return c.JSON(http.StatusCreated, ticket)
}

func GetTicketsByEvent(c echo.Context) error {
	eventID := c.Param("event_id")
	var tickets []models.Ticket
	config.DB.Where("event_id = ?", eventID).Find(&tickets)
	return c.JSON(http.StatusOK, tickets)
}

func UpdateTicket(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	ticketID := c.Param("ticket_id")

	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Ticket not found")
	}

	var event models.Event
	if err := config.DB.First(&event, ticket.EventID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Related event not found")
	}

	if event.UserID != userID {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	if err := c.Bind(&ticket); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Save(&ticket)
	return c.JSON(http.StatusOK, ticket)
}

func DeleteTicket(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	ticketID := c.Param("ticket_id")

	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Ticket not found")
	}

	var event models.Event
	if err := config.DB.First(&event, ticket.EventID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}

	if event.UserID != userID {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	config.DB.Delete(&ticket)
	return c.JSON(http.StatusOK, "Ticket deleted")
}
