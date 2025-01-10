package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) error {
	p := []byte(password)
	h := []byte(hash)
	err := bcrypt.CompareHashAndPassword(h, p)
	if err != nil {
		return errors.New("Invalid password, bratan")
	}
	return nil
}
