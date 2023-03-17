package models

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateTokenAndCookies(email string, password string) (*http.Cookie, error) {
	// membuat token menggunakan data yang diberikan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// menandatangani token menggunakan secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	// Set cookie pada header response
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	return cookie, nil
}
