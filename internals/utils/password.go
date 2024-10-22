package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return string(bytes), err
}

func VerifyPasswrod(password, hashed []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, password)
	return err == nil
}
