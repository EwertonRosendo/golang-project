package validations

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashPassword))
}