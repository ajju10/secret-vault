package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const secret = "Zggo+umm4JPG8PA3qVrpdHsVflvpcjHxOnV7wlwuEhcD465nncHoX6sNuelZOQPVlROYAV6blVOJhNqxOY0ct1meJFa7DHm/" +
	"dhKgGxAy9cfM6kgSeSra7hDZqTwGJ+IxeV5mE4h30/eU/PcRjDd4DrDgilXSCsEDnB4FDJJ3nBxeUoDwOoG+LY7vIzG8hztwJFvab/irP0IXFlVNvFc" +
	"S9R2ecmtKg3/cnJIVVQ=="

var jwtKey = []byte(secret)

type authClaims struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

func generateToken(user DbUser) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expiresAt,
		},
		UserID: user.UID,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(tokenString string) (string, string, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", "", err
	}
	if !token.Valid {
		return "", "", errors.New("invalid token")
	}
	id := claims.UserID
	username := claims.Subject
	return id, username, nil
}
