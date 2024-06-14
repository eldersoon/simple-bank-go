package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	db "github.com/eldersoon/simple-bank/db/sqlc"
	"github.com/eldersoon/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount (t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	} 

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func Test_CreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func Test_GetAccount (t *testing.T) {
	newAccount := CreateRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, account.CreatedAt, time.Second)
}

func Test_GetAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	params := db.ListAccountsParams{
		Offset: 5,
		Limit: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), params)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func Test_DeleteAccount(t *testing.T) {
	newAccount := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.Error(t, err)
	require.Empty(t, account)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func Test_UpdateAccount(t *testing.T) {
	newAccount := CreateRandomAccount(t)

	arg := db.UpdateAccountParams {
		ID: newAccount.ID,
		Balance: utils.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, newAccount.Currency, account.Currency)

	require.NotEqual(t, newAccount.Balance, account.Balance)
}
