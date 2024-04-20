package sqlcs

import (
	"context"
	"database/sql"
	"fmt"
)

//this struct provides all functions to executed db queries
type Store struct {
	*Queries
	db *sql.DB
}

//creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

//exec
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("txerr:%v, rberr: %v", err, rberr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Ammount       int64 `json:"ammount"`
}

type TransferTxResult struct {
	Transfer      Transfer `json:"transfer"`
	FromAccountID Account  `json:"from_account"`
	ToAccountID   Account  `json:"to_account"`
	FromEntry     Entry    `json:"from_entry"`
	ToEntry       Entry    `json:"to_entry"`
}

//this function creates Transfer record, add account entries and update accounts balance 
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Ammount:       arg.Ammount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Ammount:   -arg.Ammount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Ammount:   arg.Ammount,
		})

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
