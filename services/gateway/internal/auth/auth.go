package auth

import (
	"errors"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("signingKey")), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return "", ErrTokenExpired
		}
		return "", ErrInvalidToken
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return "", ErrInvalidToken
	}

	return strconv.Itoa(claims.UserId), nil
}
