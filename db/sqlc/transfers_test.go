package db

import (
	"context"
	"github.com/jose-villedaa/simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// createRandomTransfer tests the CreateTransfer function
func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	// Args for CreateTransfer
	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	// Create a new transfer
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	// Check if the transfer is created correctly
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	// Check if the transfer has an ID and CreatedAt
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

// TestCreateTransfer tests the CreateTransfer function
func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

// TestGetTransfer tests the GetTransfer function
func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, 0)
}

// TestListTransfers tests the ListTransfers function
func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer1 := createRandomTransfer(t, account1, account2)
	transfer2 := createRandomTransfer(t, account1, account2)

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Len(t, transfers, 2)

	require.Equal(t, transfer1.ID, transfers[0].ID)
	require.Equal(t, transfer1.FromAccountID, transfers[0].FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfers[0].ToAccountID)
	require.Equal(t, transfer1.Amount, transfers[0].Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfers[0].CreatedAt, 0)

	require.Equal(t, transfer2.ID, transfers[1].ID)
	require.Equal(t, transfer2.FromAccountID, transfers[1].FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfers[1].ToAccountID)
	require.Equal(t, transfer2.Amount, transfers[1].Amount)
	require.WithinDuration(t, transfer2.CreatedAt, transfers[1].CreatedAt, 0)
}
