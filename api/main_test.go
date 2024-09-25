package api

import (
	"os"
	"testing"
	"time"

	db "github.com/1shubham7/bank/db/sqlc"
	"github.com/1shubham7/bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	assert.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
