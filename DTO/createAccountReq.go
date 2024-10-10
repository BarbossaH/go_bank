package dto

import (
	"bankserver/errs"
	"strings"
)

type CreateAccountReq struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (createReq CreateAccountReq) ValidateData() *errs.AppError{
	if(createReq.Amount<100){
		return errs.NewValidationError("To open a new account you need to deposit at least 100")
	}
	aType:=strings.ToLower(createReq.AccountType)
	if aType!="saving" && aType!="checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}

	return nil
}