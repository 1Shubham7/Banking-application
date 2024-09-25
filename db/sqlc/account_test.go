package db

import (
	"context"
	"testing"

	"github.com/1shubham7/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	assert := assert.New(t)

	// now that account onwer must be a user, we will first create a user
	userArgs := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: "secret",
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), userArgs)
	assert.NoError(err)

	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	assert.Equal(account.Owner, user.Username)
	assert.Equal(account.Balance, args.Balance)
	assert.Equal(account.Currency, args.Currency)

	assert.NotZero(account.ID)
	assert.NotZero(account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	assert := assert.New(t)

	userArgs := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: "secret",
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), userArgs)
	assert.NoError(err)

	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	acc, err := testStore.GetAccount(context.Background(), account.ID)

	assert.NotNil(acc)
	assert.NoError(err)

	assert.Equal(acc.Owner, user.Username)
	assert.Equal(acc.Balance, args.Balance)
	assert.Equal(acc.Currency, args.Currency)
}

func TestUpdateAccount(t *testing.T) {
	assert := assert.New(t)

	userArgs := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: "secret",
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), userArgs)
	assert.NoError(err)

	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	updateParams := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomBalance(),
	}

	acc, err := testStore.UpdateAccount(context.Background(), updateParams)

	assert.NotNil(acc)
	assert.NoError(err)

	assert.Equal(acc.Owner, user.Username)
	assert.Equal(acc.Balance, updateParams.Balance)
}

func TestDeleteAccount(t *testing.T) {
	assert := assert.New(t)

	userArgs := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: "secret",
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), userArgs)
	assert.NoError(err)

	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	err = testStore.DeleteAccount(context.Background(), account.ID)

	assert.NoError(err)

	acc, err := testStore.GetAccount(context.Background(), account.ID)

	assert.Error(err)
	assert.EqualError(err, "no rows in result set")

	assert.Equal(acc.ID, int64(0))
	assert.Equal(acc.Balance, int64(0))
}

func TestListAccounts(t *testing.T) {
	assert := assert.New(t)

	userArgs := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: "secret",
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), userArgs)
	assert.NoError(err)

	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), args)

	assert.NotNil(account)
	assert.NoError(err)

	lap := ListAccountsParams{
		Owner:  user.Username,
		Limit:  2,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), lap)

	assert.NoError(err)
	assert.Len(accounts, 1)

	for _, a := range accounts {
		assert.NotEmpty(a)
	}
}
