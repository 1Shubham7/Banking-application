package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// store provides all funcs to execute db transections and queries
type Store interface {
	Querier //this is bring all the funcs inside Queries (querier.go) inside this interface
	TransferTransection(ctx context.Context, arg TransferTransectionParams) (TransferTransectionResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
