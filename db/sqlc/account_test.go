package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	// Args for CreateAccount
	arg := CreateAccountParams{
		Owner:    "Lionel Messi",
		Balance:  100,
		Currency: "USD",
	}

	// Create a new account
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// Check if the account is created correctly
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	// Check if the account has an ID and CreatedAt
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
