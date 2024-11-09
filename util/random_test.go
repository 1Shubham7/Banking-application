package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	assert := assert.New(t)

	min, max := int64(1), int64(100)
	randomValue := RandomInt(min, max)
	assert.GreaterOrEqual(randomValue, min)
	assert.LessOrEqual(randomValue, max)
}

func TestRandomString(t *testing.T) {
	assert := assert.New(t)

	length := 10
	randomStr := RandomString(length)
	assert.Equal(len(randomStr), length)

	for _, char := range randomStr {
		assert.Contains(alpha, string(char))
	}
}

func TestRandomOwner(t *testing.T) {
	assert := assert.New(t)

	randomOwner := RandomOwner()
	assert.Equal(len(randomOwner), 10)

	for _, char := range randomOwner {
		assert.Contains(alpha, string(char))
	}
}


func TestRandomCurrency(t *testing.T) {
	assert := assert.New(t)

	currencies := []string{"USD", "INR", "CAD", "EUR", "YUAN"}
	randomCurrency := RandomCurrency()
	assert.Contains(currencies, randomCurrency)
}
