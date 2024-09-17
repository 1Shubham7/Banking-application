package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/1shubham7/bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M){
	config, err := util.LoadConfig("../..")

	if err != nil{
		log.Fatal("Error occured while loading config: ", err)
	} 

	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil{
		log.Fatal("Error occured while connecting to the sql database: ", err)
	}

	// testQueries = New(connPool)
	testStore = NewStore(connPool)

	os.Exit(m.Run())
}