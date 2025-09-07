package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameAccount = "account"

type Account struct {
	AccountID      uuid.UUID     `json:"accountID" gorm:"column:account_id;type:uuid;primaryKey"`
	DocumentNumber string        `json:"documentNumber" gorm:"column:document_number;size:255;not null;uniqueIndex"`
	Transactions   []Transaction ` json:"-" gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;"`
}

func (p *Account) BeforeCreate(tx *gorm.DB) (err error) {
	if p.AccountID == uuid.Nil {
		p.AccountID = uuid.New()
	}
	return
}

func (Account) TableName() string {
	return TableNameAccount
}
