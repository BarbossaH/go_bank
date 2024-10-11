package domain

import (
	"bankserver/errs"
	"bankserver/logger"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	account_db *sqlx.DB
}

func InitAccountRepositoryDB(db *sqlx.DB)AccountRepositoryDB{
	return AccountRepositoryDB{db}
}

func(accRepDB AccountRepositoryDB)InsertAccount(a Account)(*Account, *errs.AppError){
	sqlInsert:="INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values($1,$2,$3,$4,$5)RETURNING account_id"

	err := accRepDB.account_db.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&a.AccountId)
	if err!=nil{
		logger.Error("Error while creating new account: "+ err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &a,nil
}

func (accRepDB AccountRepositoryDB) FindAccountById(id string) (*Account, *errs.AppError) {
	var account Account
	accountSql := `SELECT customer_id, opening_date, account_type, amount, status 
				   FROM accounts 
				   WHERE account_id=$1`
	
	// Pass the 'id' as the argument to the query
	err := accRepDB.account_db.Get(&account, accountSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &account, nil
}


func(accRepDB AccountRepositoryDB)SaveTransaction(t Transaction)(*Transaction,*errs.AppError){
	tx,err:=accRepDB.account_db.Begin()
	if err !=nil {
		logger.Error("Error while starting a new transaction for bank account transaction"+err.Error())
		return nil,errs.NewUnexpectedError("Unexpected database error")
	}

	insertSql := `
		INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
		VALUES ($1, $2, $3, $4) 
		RETURNING transaction_id`
	var transactionId int64
	err = tx.QueryRow(insertSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&transactionId)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while inserting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Update account balance based on transaction type (withdrawal or deposit)
	var updateSql string
	if t.IsWithDrawal() {
		updateSql = `UPDATE accounts SET amount = amount - $1 WHERE account_id = $2`
	} else {
		updateSql = `UPDATE accounts SET amount = amount + $1 WHERE account_id = $2`
	}

	_, err = tx.Exec(updateSql, t.Amount, t.AccountId)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing the transaction: " + err.Error())
		tx.Rollback()
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Fetch updated account balance
	account, appErr := accRepDB.FindAccountById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	// Set transaction ID and updated balance
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}
