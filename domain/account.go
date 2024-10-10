package domain

import (
	dto "bankserver/DTO"
	"bankserver/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string 
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func(a Account)AccountToDto() dto.CreateAccountRes{
	return dto.CreateAccountRes{
		AccountId: a.AccountId,
	}
}

type IAccountRepository interface {
	InsertAccount(Account) (*Account, *errs.AppError)
}