package db

import (
	"context"
	"fmt"
)

func (store *SQLStore) executeTransection(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rollBAckErr := tx.Rollback(ctx)
		if rollBAckErr != nil {
			return fmt.Errorf("transection err: %v, Rollback err: %v", err, rollBAckErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
