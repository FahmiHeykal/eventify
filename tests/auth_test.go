package tests

import (
	"eventify/config"
	"eventify/controllers"
	"eventify/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	config.LoadEnv()
	config.InitDB()

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	user := models.User{
		Name:     "Test User",
		Email:    "test@mail.com",
		Password: string(password),
	}
	config.DB.Create(&user)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@mail.com","password":"123456"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := controllers.Login(c); err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "token") {
		t.Fatalf("Expected token in response, but got: %s", rec.Body.String())
	}
}
