package dao

import (
	"context"
	"tx-api/model"

	"gorm.io/gorm"
)

func (d *transactionDAO) GetAccountByID(ctx context.Context, accountID string) (model.Account, error) {
	var account model.Account
	if err := d.pgSvc.FindFirst(ctx, &account, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", accountID)
	}); err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (d *transactionDAO) BatchWriteData(ctx context.Context, savedData []model.TableName, deletedData ...model.TableName) error {
	if err := d.pgSvc.BatchWriteData(ctx, savedData, deletedData...); err != nil {
		return err
	}

	return nil
}
