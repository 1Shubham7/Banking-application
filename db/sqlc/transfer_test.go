package db

import (
	"context"
	"testing"

	"github.com/1shubham7/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransfer(t *testing.T) {
	assert := assert.New(t)

	// creating two accounts

	account1, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Shubham the legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account1)
	assert.NoError(err)

	account2, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Shubham the real legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account2)
	assert.NoError(err)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10000,
	}

	transfer, err := testStore.CreateTransfer(context.Background(), args)

	assert.NotNil(transfer)
	assert.NoError(err)
}

func TestGetTransfer(t *testing.T) {
	assert := assert.New(t)

	// creating two accounts

	account1, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Shubham the legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account1)
	assert.NoError(err)

	account2, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Shubham the real legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account2)
	assert.NoError(err)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10000,
	}

	transfer, err := testStore.CreateTransfer(context.Background(), args)

	assert.NotNil(transfer)
	assert.NoError(err)

	tf, err := testStore.GetTransfer(context.Background(), transfer.ID)

	assert.NotNil(tf)
	assert.NoError(err)
	assert.Equal(tf.ID, transfer.ID)
}

func TestListTransfer(t *testing.T) {
	assert := assert.New(t)

	// creating two accounts

	account1, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Shubham the legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account1)
	assert.NoError(err)

	account2, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Shubham the real legend",
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account2)
	assert.NoError(err)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10000,
	}

	transfer, err := testStore.CreateTransfer(context.Background(), args)

	assert.NotNil(transfer)
	assert.NoError(err)

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         1,
		Offset:        0,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)

	assert.NoError(err)
	assert.Len(transfers, 1)
}