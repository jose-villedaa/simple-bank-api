package db

import (
	"context"
	"github.com/jose-villedaa/simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// createRandomAccount tests the CreateAccount function
func createRandomAccount(t *testing.T) Account {
	// Args for CreateAccount
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	// Create a new account
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// Check if the account is created correctly
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	// Check if the account has an ID and CreatedAt
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// TestCreateAccount tests the CreateAccount function
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// TestGetAccount tests the GetAccount function
func TestGetAccount(t *testing.T) {
	// Create a new account
	account1 := createRandomAccount(t)

	// Get the account
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	// Check if the account is the same
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}

// TestListAccounts tests the ListAccounts function
func TestListAccounts(t *testing.T) {
	// Create 10 random accounts
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	// Limit and Offset for ListAccounts
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	// List the accounts
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
}

// TestUpdateAccount tests the UpdateAccount function
func TestUpdateAccount(t *testing.T) {
	// Create a new account
	account1 := createRandomAccount(t)

	// Args for UpdateAccount
	args := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	// Update the account
	account2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	// Check if the account is updated correctly
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, 0)
}

// TestDeleteAccount tests the DeleteAccount function
func TestDeleteAccount(t *testing.T) {
	// Create a new account
	account1 := createRandomAccount(t)

	// Delete the account
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	// Get the account
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}
