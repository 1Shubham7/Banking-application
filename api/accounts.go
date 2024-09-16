package api

import "github.com/gin-gonic/gin"

type createAccountRequest struct{
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency" binding:"required,oneof=USD INR EUR"`
}

func (server *Server) createAccount(ctx *gin.Context){

}