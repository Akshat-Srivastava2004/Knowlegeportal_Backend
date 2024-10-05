package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// ComparePasswords compares the provided password with the hashed password from the database
// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	// Generate a hashed version of the password with a cost factor
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Return an empty string and the error if hashing fails
	}
	return string(hashedPassword), nil // Return the hashed password as a string
}
