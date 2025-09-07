package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const TableNameTransaction = "transaction"

type Transaction struct {
	TransactionID   uuid.UUID       `json:"transactionID" gorm:"column:transaction_id;type:uuid;primaryKey"`
	AccountID       uuid.UUID       `json:"accountID" gorm:"column:account_id;type:uuid;not null"`
	OperationTypeID int             ` json:"operationTypeID" gorm:"column:operation_type_id;not null"`
	Amount          decimal.Decimal `json:"amount" gorm:"column:amount; type:numeric"`
	EventDate       time.Time       `json:"eventDate" gorm:"column:event_date"`
}

func (p *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if p.TransactionID == uuid.Nil {
		p.TransactionID = uuid.New()
		p.EventDate = time.Now()
	}
	return
}

func (Transaction) TableName() string {
	return TableNameTransaction
}
