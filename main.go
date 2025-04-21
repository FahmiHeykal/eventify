package main

import (
	"github.com/fahmiheykal/eventify/config"
	"github.com/fahmiheykal/eventify/models"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	config.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Ticket{}, &models.Transaction{})

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
