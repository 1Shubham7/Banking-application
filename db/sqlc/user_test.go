package db

import (
	"context"
	"testing"

	"github.com/1shubham7/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	password, _ := util.HashPassword(util.RandomString(6))
	args := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: password,
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), args)

	assert.NotNil(user)
	assert.NoError(err)

	assert.Equal(user.Username, args.Username)
	assert.Equal(user.HashedPassword, args.HashedPassword)
	assert.Equal(user.FullName, args.FullName)
	assert.Equal(user.Email, args.Email)

	assert.NotZero(user.CreatedAt)
	assert.True(user.PasswordChangedAt.IsZero())
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)
	password, _ := util.HashPassword(util.RandomString(6))
	args := CreateUserParams{
		Username:       util.RandomString(10),
		HashedPassword: password,
		FullName:       "Shubham Singh Mahar",
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), args)

	assert.NotNil(user)
	assert.NoError(err)

	u, err := testStore.GetUser(context.Background(), user.Username)

	assert.NotNil(u)
	assert.NoError(err)

	assert.Equal(u.Username, args.Username)
	assert.Equal(u.HashedPassword, args.HashedPassword)
	assert.Equal(u.FullName, args.FullName)
	assert.Equal(u.Email, args.Email)

	assert.NotZero(u.CreatedAt)
	assert.True(u.PasswordChangedAt.IsZero())
}
