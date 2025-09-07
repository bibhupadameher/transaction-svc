package service

import (
	"context"
	"errors"
	"testing"
	"tx-api/dto"
	"tx-api/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMockService_CreateAccount(t *testing.T) {
	mockSvc := new(MockService)
	req := dto.CreateAccountRequest{DocumentNumber: "123"}
	expectedResp := dto.CreateAccountResponse{AccountID: uuid.New()}

	mockSvc.On("CreateAccount", context.Background(), req).Return(expectedResp, nil)

	resp, err := mockSvc.CreateAccount(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)

	mockSvc.AssertExpectations(t)
}

func TestMockService_GetAccount(t *testing.T) {
	mockSvc := new(MockService)
	accountID := uuid.New()
	req := dto.GetAccountRequest{AccountID: accountID}
	expectedResp := dto.GetAccountResponse{Account: model.Account{AccountID: uuid.New(), DocumentNumber: "test"}}

	mockSvc.On("GetAccount", context.Background(), req).Return(expectedResp, nil)

	resp, err := mockSvc.GetAccount(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)

	mockSvc.AssertExpectations(t)
}

func TestMockService_CreateTransaction(t *testing.T) {
	mockSvc := new(MockService)
	req := dto.CreateTransactionRequest{
		AccountID:       uuid.New(),
		OperationTypeID: 1,
		Amount:          decimal.NewFromInt(100),
		IsPositive:      true,
	}
	expectedResp := dto.CreateTransactionResponse{TransactionID: uuid.New()}

	mockSvc.On("CreateTransaction", context.Background(), req).Return(expectedResp, nil)

	resp, err := mockSvc.CreateTransaction(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)

	mockSvc.AssertExpectations(t)
}

func TestMockService_ErrorPath(t *testing.T) {
	mockSvc := new(MockService)
	req := dto.CreateAccountRequest{DocumentNumber: "123"}

	mockSvc.On("CreateAccount", context.Background(), req).Return(dto.CreateAccountResponse{}, errors.New("service error"))

	_, err := mockSvc.CreateAccount(context.Background(), req)
	assert.Error(t, err)

	mockSvc.AssertExpectations(t)
}
