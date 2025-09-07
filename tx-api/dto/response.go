package dto

import (
	"github.com/google/uuid"
)

type CreateAccountResponse struct {
	AccountID uuid.UUID `json:"accountID"`
}
