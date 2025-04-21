package controllers

import (
	"eventify/config"
	"eventify/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateEvent(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	event := new(models.Event)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	event.UserID = userID
	config.DB.Create(&event)
	return c.JSON(http.StatusCreated, event)
}

func GetAllEvents(c echo.Context) error {
	var events []models.Event
	config.DB.Find(&events)
	return c.JSON(http.StatusOK, events)
}

func GetEventByID(c echo.Context) error {
	id := c.Param("id")
	var event models.Event
	if result := config.DB.First(&event, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}
	return c.JSON(http.StatusOK, event)
}

func UpdateEvent(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	id := c.Param("id")
	var event models.Event
	if err := config.DB.First(&event, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}
	if event.UserID != userID {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Save(&event)
	return c.JSON(http.StatusOK, event)
}

func DeleteEvent(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	id := c.Param("id")
	var event models.Event
	if err := config.DB.First(&event, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Event not found")
	}
	if event.UserID != userID {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	config.DB.Delete(&event)
	return c.JSON(http.StatusOK, "Event deleted")
}
