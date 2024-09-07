package db

import (
	"context"
	"testing"

	"github.com/1shubham7/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	assert := assert.New(t)
	args := CreateAccountParams{
		Owner:    "Shubham",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	assert.Equal(account.Owner, "Shubham")
	assert.Equal(account.Balance, int64(100000000000))
	assert.Equal(account.Currency, "USD")

	assert.NotZero(account.ID)
	assert.NotZero(account.CreatedAt)
}
