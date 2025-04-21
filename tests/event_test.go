package tests

import (
	"eventify/config"
	"eventify/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateEvent_Unauthorized(t *testing.T) {
	config.LoadEnv()
	config.InitDB()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/events", strings.NewReader(`{"title":"Concert","description":"Live","location":"Jakarta","date":"2025-12-01"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := controllers.CreateEvent(c); err == nil {
		t.Fatalf("Expected error when unauthorized")
	}
}
