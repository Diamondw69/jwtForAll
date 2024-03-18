package data

import (
	"github.com/golang-jwt/jwt/v5"
	"jwtTesting/Golang/internal/connections"
	"os"
	"time"
)

func GenerateAccessToken(userid int64) (string, error) {
	connections.EnvLoader(connections.Path)
	mySigningKey := os.Getenv("SIGNINGKEY")

	type CustomClaims struct {
		UserId int64 `json:"userId"`
		jwt.RegisteredClaims
	}

	claims := CustomClaims{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "almaz",
			Subject:   "client",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}
