package controllers

import (
	"eventify/config"
	"eventify/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	input := new(models.User)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var user models.User
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenStr,
	})
}
