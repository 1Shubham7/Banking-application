package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// store provides all funcs to execute db transections and queries

type Store struct {
	*Queries
	db *sql.DB
}

type SQLStore struct{
	*Queries
	connPool 
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) executeTransection(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("Rollback error: %v. Transection error: %v", rollbackErr, err)
		}
		//if rollback is successful we will only return the original error when fn(q) failed
		return err
	}

	err = tx.Commit()
	return err
}

type TransferTransectionParams struct {
	FromAccountId int64 `json:"from_acc_id"`
	ToAccountId  int64 `json:"to_acc_id"`
	Amount        int64 `json:"amount"`
}

type TransferTransectionResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_acc"`
	ToAccount Account `json:"to_acc"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

// this will transfer money from one another to another
// this will create a transfer record, add account entries and update account balance within a single database  transection
func (store *Store) TransferTransection(ctx context.Context, arg TransferTransectionParams) (TransferTransectionResult, error) {
	var result TransferTransectionResult

	err := store.executeTransection(ctx, func (q *Queries) error  {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID : arg.ToAccountId,
			Amount: arg.Amount,
		})
		if err != nil{
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount: -arg.Amount,
		})
		if err!= nil{
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount: arg.Amount,
		})
		if err!= nil{
			return err
		}

		// Update acc balance

		acc1, err := q.GetAccount(ctx, arg.FromAccountId)
		if err != nil{
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID: arg.FromAccountId,
			Balance: acc1.Balance - arg.Amount,
		})

		if err != nil{
			return err
		}

		acc2, err := q.GetAccount(ctx, arg.ToAccountId)
		if err != nil{
			return err
		}

		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID: arg.ToAccountId,
			Balance: acc2.Balance + arg.Amount,
		})

		if err != nil{
			return err
		}

		return nil
	})

	return result, nil
}
