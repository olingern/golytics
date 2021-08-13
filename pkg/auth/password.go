package auth

import (
	"golang.org/x/crypto/bcrypt"
)

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
