package api

import (
	db "github.com/1shubham7/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server  struct{
	store *db.SQLStore
	router *gin.Engine
}

// func NewServer(store *db.Store) *Server{
// 	// server := &Server{store: store}
// 	// router := gin.Default()

// 	// router.POST("/accounts", server.createAccount)
// 	// server.router = router
// 	// return server
// }

