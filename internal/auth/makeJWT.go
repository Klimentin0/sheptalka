package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	Issuer    string
	IssuedAt  time.Time
	ExpiresAt time.Time
	Subject   string
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	method := jwt.SigningMethodHS256
	expiration := time.Now().Add(expiresIn)
	userIdStr := userID.String()
	claims := &jwt.RegisteredClaims{
		Issuer:    "sheptalka",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiration),
		Subject:   userIdStr,
	}
	token := jwt.NewWithClaims(method, claims)
	signedToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
