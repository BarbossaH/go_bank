package domain

import dto "bankserver/DTO"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId string `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithDrawal() bool {
	return t.TransactionType == WITHDRAWAL 
}

func (t Transaction) ToResDto() dto.TransactionRes{
	return dto.TransactionRes{
		AccountId: t.AccountId,
		Amount: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}