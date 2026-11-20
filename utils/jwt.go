package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	Uid   int64  `json:"uid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(secret string, claims *UserClaims, expire time.Duration) string {
	claims.ExpiresAt = time.Now().Add(expire).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return token
}

func ParseToken(secret, tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}
	return claims, nil
}
