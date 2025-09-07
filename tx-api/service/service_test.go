package service

import (
	"context"
	"errors"
	"testing"
	"tx-api/dao"
	"tx-api/dto"
	"tx-api/model"

	logger "tx-api/core/logging"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	logger.Init()
	model.DefaultEnum.OperationTypes = []model.OperationType{
		{OperationTypeID: 1, Description: "Payment", IsPositive: false},
		{OperationTypeID: 2, Description: "Deposit", IsPositive: true},
	}
}
func setupServiceWithMock(mockDAO *dao.MockTransactionDAO) *transactionServiceHandler {

	return &transactionServiceHandler{
		dao:    mockDAO,
		logger: logger.GetLogger(),
	}
}

func TestTransactionService_CreateAccount(t *testing.T) {
	var mockDAO dao.MockTransactionDAO
	svc := setupServiceWithMock(&mockDAO)

	req := dto.CreateAccountRequest{DocumentNumber: "123"}
	acc := &model.Account{DocumentNumber: req.DocumentNumber}

	mockDAO.On("SaveAccount", mock.Anything, acc).Return(nil).Run(func(args mock.Arguments) {
		a := args.Get(1).(*model.Account)
		a.AccountID = uuid.New()
	})

	resp, err := svc.CreateAccount(context.Background(), req)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, resp.AccountID)

	mockDAO.AssertExpectations(t)
}

func TestTransactionService_GetAccount(t *testing.T) {
	var mockDAO dao.MockTransactionDAO
	svc := setupServiceWithMock(&mockDAO)

	accountID := uuid.New()
	account := model.Account{AccountID: accountID}
	req := dto.GetAccountRequest{AccountID: accountID}

	mockDAO.On("GetAccountByID", mock.Anything, accountID).Return(account, nil)

	resp, err := svc.GetAccount(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, accountID, resp.Account.AccountID)

	mockDAO.AssertExpectations(t)
}

func TestTransactionService_CreateTransaction(t *testing.T) {
	var mockDAO dao.MockTransactionDAO
	svc := setupServiceWithMock(&mockDAO)

	accountID := uuid.New()
	req := dto.CreateTransactionRequest{
		AccountID:       accountID,
		OperationTypeID: 1,
		Amount:          decimal.NewFromInt(100),
		IsPositive:      true,
	}

	mockDAO.On("SaveTransaction", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		tr := args.Get(1).(*model.Transaction)
		tr.TransactionID = uuid.New()
	})

	resp, err := svc.CreateTransaction(context.Background(), req)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, resp.TransactionID)

	mockDAO.AssertExpectations(t)
}

func TestTransactionService_ErrorPaths(t *testing.T) {
	var mockDAO dao.MockTransactionDAO
	svc := setupServiceWithMock(&mockDAO)

	// SaveAccount fails
	accReq := dto.CreateAccountRequest{DocumentNumber: "123"}
	mockDAO.On("SaveAccount", mock.Anything, mock.Anything).Return(errors.New("db error"))

	_, err := svc.CreateAccount(context.Background(), accReq)
	assert.Error(t, err)
}
