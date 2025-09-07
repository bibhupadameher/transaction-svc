package dao

import (
	"context"
	"errors"
	constants "tx-api/constant"
	logger "tx-api/core/logging"
	apperrors "tx-api/errors"
	"tx-api/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (d *transactionDAO) SaveAccount(ctx context.Context, account *model.Account) error {
	if err := d.pgSvc.BatchWriteData(ctx, []model.TableName{account}); err != nil {
		logger.Error(ctx, "error on batch write data", zap.Error(err))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == string(constants.UNIQUE_VIOLATION) {
				err = apperrors.NewErrUniqueViolation("documentNumber", account.DocumentNumber)
				return err
			}
		}
		err = apperrors.NewErrDatabaseError()
		return err
	}
	return nil
}

func (d *transactionDAO) GetAccountByID(ctx context.Context, accountID uuid.UUID) (model.Account, error) {
	var account model.Account
	if err := d.pgSvc.FindFirst(ctx, &account, func(db *gorm.DB) *gorm.DB {
		return db.Where("account_id = ?", accountID)
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperrors.NewErrRecordNotFound("accountID", accountID.String())
			return account, err
		}
		err = apperrors.NewErrDatabaseError()
		return account, err
	}

	return account, nil
}
func (d *transactionDAO) SaveTransaction(ctx context.Context, trans *model.Transaction) error {
	if err := d.pgSvc.BatchWriteData(ctx, []model.TableName{trans}); err != nil {
		logger.Error(ctx, "error on batch write data", zap.Error(err))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == string(constants.FOREIGN_KEY_VIOLATION) {
				err = apperrors.NewErrForeignKeyViolation("accountID", trans.AccountID.String())
				return err
			}
		}
		err = apperrors.NewErrDatabaseError()
		return err
	}
	return nil
}
