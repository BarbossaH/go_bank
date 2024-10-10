package domain

import (
	"bankserver/errs"
	"bankserver/logger"

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