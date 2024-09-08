package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provides all funcs to execute db transections and queries

type Store struct{
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db),
	}
}

func (store *Store) executeTransection (ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	
	if err!= nil{
		return err
	}

	q := New(tx)

	err = fn(q)
	if err!=nil{
		rollbackErr := tx.Rollback()
		if rollbackErr != nil{
			return fmt.Errorf("Rollback error: %v. Transection error: %v", rollbackErr, err)
		}

		//if rollback is successful we will only return the original error when fn(q) failed
		return err
	}

	err = tx.Commit()
	return err
}