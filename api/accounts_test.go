package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	mockdb "github.com/1shubham7/bank/db/mock"
	db "github.com/1shubham7/bank/db/sqlc"
	"github.com/1shubham7/bank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T){
	account := db.Account{
		ID: util.RandomInt(1,1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).Return(account, nil)

	url := fmt.Sprintf("/accounts/%d", account.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	assert.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	requireBodyMatchAccount(t, recorder.Body, account)
}

func TestCreateAccount(t *testing.T){
	assert := assert.New(t)
	account := db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	arg := db.CreateAccountParams{
		Owner:    account.Owner,
		Currency: account.Currency,
		Balance:  account.Balance,
	}

	store.EXPECT().
		CreateAccount(gomock.Any(), gomock.Eq(arg)).
		Times(1).
		Return(account, nil)

	reqBody := createAccountRequest{
			Owner:    account.Owner,
			Balance:  account.Balance,
			Currency: account.Currency,
	}
	body, err := json.Marshal(reqBody)
	assert.NoError(err)
	
	req, err := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader((body)))

	assert.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	server.router.ServeHTTP(recorder, req)

	assert.Equal(http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	assert := assert.New(t)
	data, err := io.ReadAll(body)
	assert.NoError(err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	assert.NoError(err)
	assert.Equal(account, gotAccount)
}