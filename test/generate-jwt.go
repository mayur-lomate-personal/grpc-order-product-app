package main

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("my_secret_key")

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "test-user",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated Token:", tokenString)
}
