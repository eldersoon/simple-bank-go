package api

import (
	"database/sql"
	"net/http"

	db "github.com/eldersoon/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required" validate:"oneof=USD EUR BRL"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required" validate:"min=1"`
}

func (server *Server) getAccount (ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	Offset int32 `form:"offset"`
	Limit int32 `form:"limit" binding:"required" validate:"min=5, max=20"`
}

func (server *Server) listAccounts (ctx *gin.Context) {
	var req listAccountsRequest
	var totalCount int

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Offset: req.Offset,
		Limit: req.Limit,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)

    if err := server.query.QueryRow("SELECT COUNT(*) FROM accounts").Scan(&totalCount); err != nil {
        ctx.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }

    totalPages := (int32(totalCount) + arg.Limit - 1) / arg.Limit

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"items": accounts,
		"pagination": gin.H{
			"current_page": arg.Offset / arg.Limit + 1,
			"per_page": arg.Limit,
			"count": len(accounts),
			"total_pages": totalPages,
			"total_items": totalCount,
		},
	})
}