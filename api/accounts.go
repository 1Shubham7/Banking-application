package api

import (
	"net/http"

	db "github.com/1shubham7/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct{
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency" binding:"required,oneof=USD INR EUR"`
}

func (server *Server) createAccount(ctx *gin.Context){
	var req createAccountRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Balance: req.Balance,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}