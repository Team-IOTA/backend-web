package functions

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func DecodeJWT(e echo.Context) (jwt.MapClaims, error) {
	tokenString := e.Request().Header.Get("AuthToken")
	if tokenString == "" {
		return nil, errors.New("Invalid Token")
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("evoTopSecret%12345&"), nil
	})
	if fmt.Sprintf("%v", err) != "signature is invalid" {
		return nil, err
	}
	return claims, nil
}