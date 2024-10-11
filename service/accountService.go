package service

import (
	dto "bankserver/DTO"
	"bankserver/domain"
	"bankserver/errs"
	"time"
)
const dbTSLayout="2006-01-02 15:04:05Z07:00"
type IAccountService interface {
	CreateAccount(dto.CreateAccountReq)(*dto.CreateAccountRes, *errs.AppError)
	MakeTransaction( dto.TransactionReq)(*dto.TransactionRes,*errs.AppError)
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
		OpeningDate: time.Now().Format(dbTSLayout),
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

func(s DefaultAccountService)MakeTransaction(reqDto dto.TransactionReq)(*dto.TransactionRes,*errs.AppError){
	//check the body data
	err:=reqDto.ValidateData()
	if err!=nil {
		return nil,err
	}
	//withdraw or deposit
	if reqDto.IsTransactionTypeWithdrawal(){
		//withdraw money from the given account
		
		//find the account
		account,err:=s.repo.FindAccountById(reqDto.AccountId)
		if err!=nil {
			return nil, err
		}
		//check the balance is sufficient
		if !account.CanWithDraw(reqDto.Amount){
			return nil, errs.NewValidationError("Balance is insufficient")
		}
	}
		//if all is well, then make the transaction

		//transform the dto to original data format
		t:=domain.Transaction{
			AccountId: reqDto.AccountId,
			Amount: reqDto.Amount,
			TransactionType: reqDto.TransactionType,
			TransactionDate: time.Now().Format(dbTSLayout),
		}
		transaction, appError:=s.repo.SaveTransaction(t)
		if appError!=nil {
			return nil, appError
		}
		resDto:=transaction.ToResDto()
		return &resDto, nil

}

/*func(s DefaultAccountService)MakeTransaction(req dto.TransactionReq)(*dto.TransactionRes,*errs.AppError){
	err:=req.ValidateData()
	if err!=nil {
		return nil,err
	}
	if req.IsTransactionTypeWithdrawal() {
		account, err:=s.repo.FindAccountById(req.AccountId)
		if err!=nil {
			return nil,err
		}
		if !account.CanWithDraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	t:=domain.Transaction{
		AccountId: req.AccountId,
		Amount: req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError:=s.repo.SaveTransaction(t)
	if appError!=nil {
		return nil, appError
	}

	response:=transaction.
}
	*/