package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sqlcs.sqlc.dev/app/sqlcs"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required, min = 1 "`
	ToAccountID   int64  `json:"to_account_id" binding:"required, min = 1"`
	Ammount       int64  `json:"ammount" binding:"required, gt=0 "`
	Currency      string `json:"currency" binding:"required, oneof=USD EUR"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlcs.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Ammount:       req.Ammount,
	}


	if !server.validity(ctx, req.ToAccountID,req.Currency){
		return
	}

	if !server.validity(ctx, req.FromAccountID,req.Currency){
		return
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func(server *Server) validity(ctx *gin.Context,accountId int64,currency string) bool {
	account, err := server.store.GetAccount(ctx,accountId)
	if (err != nil){
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account currency mismatch")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
