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
		Owner:    "Shubham the legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	assert.Equal(account.Owner, "Shubham the legend")
	assert.Equal(account.Balance, args.Balance)
	assert.Equal(account.Currency, args.Currency)

	assert.NotZero(account.ID)
	assert.NotZero(account.CreatedAt)
}

func TestGetAccount(t *testing.T){
	assert := assert.New(t)

	args := CreateAccountParams{
		Owner:    "Shubham the OG",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	acc, err := testQueries.GetAccount(context.Background(), account.ID)

	assert.NotNil(acc)
	assert.NoError(err)

	assert.Equal(acc.Owner, "Shubham the OG")
	assert.Equal(acc.Balance, args.Balance)
	assert.Equal(acc.Currency, args.Currency)
}

func TestUpdateAccount(t *testing.T){
	assert := assert.New(t)

	args := CreateAccountParams{
		Owner:    "Shubham the OG",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	updateParam := UpdateAccountParams{
		ID: account.ID,
		Balance: util.RandomBalance(),
	}

	acc, err := testQueries.UpdateAccount(context.Background(), updateParam)

	assert.NotNil(acc)
	assert.NoError(err)

	assert.Equal(acc.Owner, "Shubham the OG")
	assert.Equal(acc.Balance, updateParam.Balance)
}

func TestDeleteAccount(t *testing.T){
	assert := assert.New(t)

	args := CreateAccountParams{
		Owner:    "Shubham the legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	err = testQueries.DeleteAccount(context.Background(), account.ID)

	assert.NoError(err)
	assert.Nil(account)
}