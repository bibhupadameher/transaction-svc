package dao

import (
	"context"
	"errors"
	"testing"
	constants "tx-api/constant"
	logger "tx-api/core/logging"
	"tx-api/core/postgres"
	apperrors "tx-api/errors"
	"tx-api/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupDAO() (*transactionDAO, *postgres.MockDBService) {
	mockDB := new(postgres.MockDBService)
	dao := &transactionDAO{pgSvc: mockDB}
	return dao, mockDB
}

func TestSaveAccount_Success(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	account := &model.Account{DocumentNumber: "123", AccountID: uuid.New()}
	logger.Init()
	defer logger.Sync()

	mockDB.On("BatchWriteData", ctx, []model.TableName{account}, mock.Anything).Return(nil)

	err := dao.SaveAccount(ctx, account)
	assert.NoError(t, err)
	mockDB.AssertCalled(t, "BatchWriteData", ctx, []model.TableName{account}, mock.Anything)
}

func TestSaveAccount_UniqueViolation(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	account := &model.Account{DocumentNumber: "123", AccountID: uuid.New()}
	logger.Init()
	defer logger.Sync()

	pgErr := &pgconn.PgError{Code: string(constants.UNIQUE_VIOLATION)}
	mockDB.On("BatchWriteData", ctx, []model.TableName{account}, mock.Anything).Return(pgErr)

	err := dao.SaveAccount(ctx, account)
	assert.ErrorIs(t, err, apperrors.NewErrUniqueViolation("documentNumber", "123"))
	mockDB.AssertCalled(t, "BatchWriteData", ctx, []model.TableName{account}, mock.Anything)
}

func TestSaveAccount_GenericDBError(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	account := &model.Account{DocumentNumber: "123"}
	logger.Init()
	defer logger.Sync()

	mockDB.On("BatchWriteData", ctx, []model.TableName{account}, mock.Anything).Return(errors.New("db error"))

	err := dao.SaveAccount(ctx, account)
	assert.ErrorIs(t, err, apperrors.NewErrDatabaseError())
	mockDB.AssertCalled(t, "BatchWriteData", ctx, []model.TableName{account}, mock.Anything)
}

func TestGetAccountByID_Success(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	id := uuid.New()
	var account model.Account

	logger.Init()
	defer logger.Sync()

	mockDB.On("FindFirst", ctx, &account, mock.Anything).Return(nil)

	account, err := dao.GetAccountByID(ctx, id)
	assert.NoError(t, err)

	mockDB.AssertCalled(t, "FindFirst", ctx, &account, mock.Anything)
}

func TestGetAccountByID_NotFound(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	id := uuid.New()
	var account model.Account
	logger.Init()
	defer logger.Sync()

	mockDB.On("FindFirst", ctx, &account, mock.Anything).Return(gorm.ErrRecordNotFound)

	got, err := dao.GetAccountByID(ctx, id)
	assert.ErrorIs(t, err, apperrors.NewErrRecordNotFound("accountID", id.String()))
	assert.Equal(t, model.Account{}, got)
}

func TestSaveTransaction_Success(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	trans := &model.Transaction{AccountID: uuid.New()}
	logger.Init()
	defer logger.Sync()

	mockDB.On("BatchWriteData", ctx, []model.TableName{trans}, mock.Anything).Return(nil)

	err := dao.SaveTransaction(ctx, trans)
	assert.NoError(t, err)
	mockDB.AssertCalled(t, "BatchWriteData", ctx, []model.TableName{trans}, mock.Anything)
}

func TestSaveTransaction_ForeignKeyViolation(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	trans := &model.Transaction{AccountID: uuid.New()}
	logger.Init()
	defer logger.Sync()

	pgErr := &pgconn.PgError{Code: string(constants.FOREIGN_KEY_VIOLATION)}
	mockDB.On("BatchWriteData", ctx, []model.TableName{trans}, mock.Anything).Return(pgErr)

	err := dao.SaveTransaction(ctx, trans)
	assert.ErrorIs(t, err, apperrors.NewErrForeignKeyViolation("accountID", trans.AccountID.String()))
	mockDB.AssertCalled(t, "BatchWriteData", ctx, []model.TableName{trans}, mock.Anything)
}

func TestSaveTransaction_GenericDBError(t *testing.T) {
	dao, mockDB := setupDAO()
	ctx := context.Background()
	trans := &model.Transaction{AccountID: uuid.New()}
	logger.Init()
	defer logger.Sync()

	mockDB.On("BatchWriteData", ctx, []model.TableName{trans}, mock.Anything).Return(errors.New("db error"))

	err := dao.SaveTransaction(ctx, trans)
	assert.ErrorIs(t, err, apperrors.NewErrDatabaseError())
	mockDB.AssertCalled(t, "BatchWriteData", ctx, []model.TableName{trans}, mock.Anything)
}
