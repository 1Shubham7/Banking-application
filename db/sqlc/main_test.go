package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testStore Store

func TestMain(m *testing.M){
	connPool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil{
		log.Fatal("Error occured while connecting to the sql database: ", err)
	}

	// testQueries = New(connPool)
	testStore = NewSQLStore(connPool)

	os.Exit(m.Run())
}