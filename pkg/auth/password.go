package auth

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func EnvPasswordIsValid() bool {
	envPassword := os.Getenv("ADMIN_PASSWORD")

	return len(envPassword) >= 6
}

func GetPasswordFromEnv() string {
	return os.Getenv("ADMIN_PASSWORD")
}

func VerifyPassword(storedPasswordHash []byte, givenPassword []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(storedPasswordHash, givenPassword)

	if err != nil {
		return false, err
	}

	return true, nil

}

func HashAndSalt(pwd []byte) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}
