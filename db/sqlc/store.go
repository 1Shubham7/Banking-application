package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)

// store provides all funcs to execute db transections and queries
type Store interface {
	Querier
	TransferTransection(ctx context.Context, arg TransferTransectionParams) (TransferTransectionResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct{
	*Queries
	connPool *pgxpool.Pool
}

func NewSQLStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries: New(connPool),
	}
}