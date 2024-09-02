package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed password using bcrypt
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost menentukan kompleksitas hashing, semakin tinggi semakin aman tapi lebih lambat
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plain text password
func ComparePassword(hashedPassword, password string) error {
	// Fungsi CompareHashAndPassword akan return nil jika password cocok, atau error jika tidak cocok
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}