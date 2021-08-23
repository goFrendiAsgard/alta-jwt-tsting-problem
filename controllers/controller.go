package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int) *jwt.Token {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func CreateSignedString(userId int) (string, error) {
	token := CreateToken(userId)
	return token.SignedString([]byte("wkwkwk"))
}

func ExtractToken(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["userId"].(int))
	return userId
}

func Controller(c echo.Context) error {
	userId := ExtractToken(c)
	return c.JSON(200, map[string]int{
		"userId": userId,
	})
}
