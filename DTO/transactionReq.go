package dto

import "bankserver/errs"
const WITHDRAWAL = "withdraw"
const DEPOSIT = "deposit"
type TransactionReq struct {
	AccountId  string  `json:"account_id"`
	CustomerId string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
}

func (t TransactionReq) ValidateData() *errs.AppError{ 
	if t.Amount < 0 {
		return errs.NewValidationError("amount cannot be less than 0")
	}
	if t.TransactionType!=WITHDRAWAL && t.TransactionType!=DEPOSIT {
		return errs.NewValidationError("Transaction type only can be deposit and withdrawal")
	}
	return nil
}

func(t TransactionReq)IsTransactionTypeWithdrawal()bool{
	return t.TransactionType==WITHDRAWAL
}