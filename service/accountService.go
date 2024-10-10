package service

import (
	dto "bankserver/DTO"
	"bankserver/domain"
	"bankserver/errs"
	"time"
)

type IAccountService interface {
	CreateAccount(dto.CreateAccountReq)(*dto.CreateAccountRes, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.IAccountRepository
}

func RegisterAccountService(repository domain.IAccountRepository)DefaultAccountService{
	return DefaultAccountService{repository}
}

func(s DefaultAccountService)CreateAccount(req dto.CreateAccountReq)(*dto.CreateAccountRes,*errs.AppError){

	//before creating the account, we need to validate the data first to check data is valid or not
	err:=req.ValidateData()
	if err!=nil {
		return nil, err
	}

	a:=domain.Account{
		AccountId: "",
		CustomerId: req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05Z07:00"),
		// OpeningDate: time.Now().Format(time.RFC3339),
		AccountType: req.AccountType,
		Amount: req.Amount,
		Status: "1",
	}
	newAccount,err:= s.repo.InsertAccount(a)
	if err!=nil {
		return nil, err
	}
	res:=newAccount.AccountToDto()
	return &res, nil
}