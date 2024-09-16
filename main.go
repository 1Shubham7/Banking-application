package main

import (
	"context"
	"log"

	"github.com/1shubham7/bank/api"
	db "github.com/1shubham7/bank/db/sqlc"
	"github.com/1shubham7/bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)


const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main(){
	config, err := util.LoadConfig(".")

	if err != nil{
		log.Fatal("Can't load config: ", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil{
		log.Fatal("Error occured while connecting to the sql database: ", err)
	}

	// testQueries = New(connPool)
	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("can't start server: ", err)
	}	
}