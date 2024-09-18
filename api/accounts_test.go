package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
}