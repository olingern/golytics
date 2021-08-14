package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const dummyPass = "foobar"

func TestHashAndSalt(t *testing.T) {
	actual, err := HashAndSalt([]byte(dummyPass))
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, actual, "password should be defined")
}

func TestPasswordFlow(t *testing.T) {

	storedPwd, err := HashAndSalt([]byte(dummyPass))

	assert.Nil(t, err, "Error should be nil")

	result, err := VerifyPassword(storedPwd, []byte(dummyPass))
	assert.Nil(t, err, "Error should be nil")
	assert.True(t, result, "dummyPassword comparison to itself should be true")
}

func TestPasswordFailFlow(t *testing.T) {

	storedPwd, err := HashAndSalt([]byte(dummyPass))

	assert.Nil(t, err, "Error should be nil")

	result, err := VerifyPassword(storedPwd, []byte("BAD_PASSWORD"))
	assert.Equal(t, err.Error(), "crypto/bcrypt: hashedPassword is not the hash of the given password")
	assert.False(t, result, "dummyPassword comparison to BAD_PASSWORD should be false")
}

func TestEnvPasswordIsValid(t *testing.T) {
	prev := os.Getenv("ADMIN_PASSWORD")
	os.Setenv("ADMIN_PASSWORD", "fail")
	assert.False(t, EnvPasswordIsValid(), "password 'fail' does not meet minimum requirements and should faile")

	// reset
	os.Setenv("ADMIN_PASSWORD", prev)
}
