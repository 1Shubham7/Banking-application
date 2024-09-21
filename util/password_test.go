package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	assert := assert.New(t)
	password := RandomString(6)

	hashPassword, err := HashPassword(password)
	assert.NoError(err)
	assert.NotEmpty(hashPassword)

	err = ValidatePassword(password, hashPassword)
	assert.NoError(err)
}
