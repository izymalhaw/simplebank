package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashedpassword(Password string) (string, error) {
	hashedpassword,err := bcrypt.GenerateFromPassword([]byte (Password) ,bcrypt.DefaultCost)

	if err != nil {
		return "",fmt.Errorf("faile to hash password: %w", err)
	}

	return string(hashedpassword),nil
}

func CheckPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}