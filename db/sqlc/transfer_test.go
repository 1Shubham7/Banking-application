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
	userArgs1 := CreateUserParams{
		Username:    util.RandomString(10),
		HashedPassword:  "secret",
		FullName: "Shubham Singh Mahar",
		Email: util.RandomEmail(),
	}

	user1, err := testStore.CreateUser(context.Background(), userArgs1)
	assert.NoError(err)

	account1, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user1.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account1)
	assert.NoError(err)

	userArgs2 := CreateUserParams{
		Username:    util.RandomString(10),
		HashedPassword:  "secret",
		FullName: "Shubham Singh Mahar",
		Email: util.RandomEmail(),
	}

	user2, err := testStore.CreateUser(context.Background(), userArgs2)
	assert.NoError(err)

	account2, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user2.Username,
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

	userArgs1 := CreateUserParams{
		Username:    util.RandomString(10),
		HashedPassword:  "secret",
		FullName: "Shubham Singh Mahar",
		Email: util.RandomEmail(),
	}

	user1, err := testStore.CreateUser(context.Background(), userArgs1)
	assert.NoError(err)

	account1, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user1.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account1)
	assert.NoError(err)

	userArgs2 := CreateUserParams{
		Username:    util.RandomString(10),
		HashedPassword:  "secret",
		FullName: "Shubham Singh Mahar",
		Email: util.RandomEmail(),
	}

	user2, err := testStore.CreateUser(context.Background(), userArgs2)
	assert.NoError(err)

	account2, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user2.Username,
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

	userArgs1 := CreateUserParams{
		Username:    util.RandomString(10),
		HashedPassword:  "secret",
		FullName: "Shubham Singh Mahar",
		Email: util.RandomEmail(),
	}

	user1, err := testStore.CreateUser(context.Background(), userArgs1)
	assert.NoError(err)

	account1, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user1.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	})

	assert.NotNil(account1)
	assert.NoError(err)

	userArgs2 := CreateUserParams{
		Username:    util.RandomString(10),
		HashedPassword:  "secret",
		FullName: "Shubham Singh Mahar",
		Email: util.RandomEmail(),
	}

	user2, err := testStore.CreateUser(context.Background(), userArgs2)
	assert.NoError(err)

	account2, err := testStore.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user2.Username,
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
