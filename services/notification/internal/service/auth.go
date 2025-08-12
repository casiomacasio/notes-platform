package service

import (
	"errors"
	"os"
	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token expired")
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("signingKey")), nil
	})
	if err != nil {
		return 0, err
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return 0, ErrTokenExpired 
		}
		return 0, ErrInvalidToken
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return 0, ErrInvalidToken
	}
	return claims.UserId, nil
}