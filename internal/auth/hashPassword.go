package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	p := []byte(password)
	hashedPass, err := bcrypt.GenerateFromPassword(p, 10)
	if err != nil {
		log.Printf("couldnt generate hashed passowrd: %v", err)
	}

	if password == "" || len(p) < 5 {
		return "", errors.New("Invalid password")
	}

	return string(hashedPass), nil
}
