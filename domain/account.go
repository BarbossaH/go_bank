package domain

import (
	dto "bankserver/DTO"
	"bankserver/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId   string  `db:"customer_id"`
	OpeningDate  string  `db:"opening_date"`
	AccountType  string  `db:"account_type"`
	Amount       float64 `db:"amount"`
	Status       string  `db:"status"`
}

func(a Account)AccountToDto() dto.CreateAccountRes{
	return dto.CreateAccountRes{
		AccountId: a.AccountId,
	}
}

type IAccountRepository interface {
	InsertAccount(Account) (*Account, *errs.AppError)
	FindAccountById(string)(*Account,*errs.AppError)
	SaveTransaction(t Transaction)(*Transaction,*errs.AppError)
}

func(a Account)CanWithDraw(amount float64)bool{
	return a.Amount>amount
}