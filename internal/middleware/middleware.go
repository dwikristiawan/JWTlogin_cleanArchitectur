package middleware

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(role string) echo.MiddlewareFunc {
	return func(Next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, error.Error(errors.New("unauthorize")))
			}

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return SecretKey, nil })
			if err != nil || !token.Valid {
				return c.JSON(http.StatusForbidden, errors.New("unauthorize"))
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok || claims["role"] == role {
				return Next(c)
			}
			return c.JSON(http.StatusForbidden, errors.New("unauthorize"))

		}
	}

}
