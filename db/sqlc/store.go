package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store instance
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Create a new transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new Queries instance with the transaction
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	// Commit the transaction
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// Execute the transfer transaction
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Params for the transfer
		transfer := CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		}

		// Create the transfer
		result.Transfer, err = q.CreateTransfer(ctx, transfer)
		if err != nil {
			return err
		}

		// Params for FromEntry
		fromEntryParams := CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		}

		// Create the FromEntry
		result.FromEntry, err = q.CreateEntry(ctx, fromEntryParams)
		if err != nil {
			return err
		}

		// Params for ToEntry
		toEntryParams := CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		}

		// Create the ToEntry
		result.ToEntry, err = q.CreateEntry(ctx, toEntryParams)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
