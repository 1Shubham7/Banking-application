package db

import (
	"context"
)

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
func (store *SQLStore) TransferTransection(ctx context.Context, arg TransferTransectionParams) (TransferTransectionResult, error) {
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

		// trying to prevent deadlock by this
		if arg.FromAccountId < arg.ToAccountId{
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err

	})

	return result, err
}

func addMoney(){
	
}