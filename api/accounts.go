package api

import (
	"database/sql"
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

type getAccountRequest struct{
	ID int64 `uri:"id" binding:"required,min=1`	
}

func (server *Server) getAccount (ctx *gin.Context){
	var req getAccountRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}