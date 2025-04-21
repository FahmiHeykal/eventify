package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		tokenStr := strings.Replace(auth, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			
			return []byte("SECRET_KEY"), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token: "+err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(float64)
			if !ok {
				return c.JSON(http.StatusUnauthorized, "Invalid user_id in token")
			}

			c.Set("user_id", uint(userID)) 
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}
	}
}

type JwtCustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
