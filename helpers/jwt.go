package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func GenToken(userId string) (token string, err error) {
	claims := &Claims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * (30 * 24))),
		},
	}
	singedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := singedToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}
	if !tkn.Valid {
		return claims, errors.New("invalid token")
	}
	return claims, nil
}
