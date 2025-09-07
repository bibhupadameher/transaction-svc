package dto

import (
	"tx-api/model"

	"github.com/google/uuid"
)

type CreateAccountResponse struct {
	AccountID uuid.UUID `json:"accountID"`
}

type GetAccountResponse struct {
	model.Account
}
type CreateTransactionResponse struct {
	TransactionID uuid.UUID `json:"transactionID"`
}
