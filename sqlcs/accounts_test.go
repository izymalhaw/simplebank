package sqlcs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"sqlcs.sqlc.dev/app/util"
)
func createRandomAccount(t *testing.T) Account{
	user:= CreateRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
		CountryCode: util.RandomCC(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account;
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t);
}
func TestGetAccount(t *testing.T){
	account1 := createRandomAccount(t)
	account2,err :=testQueries.GetAccount(context.Background(),account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Owner, account2.Owner)
}