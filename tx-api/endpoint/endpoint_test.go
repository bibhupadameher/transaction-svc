package endpoint

import (
	"context"
	"errors"
	"testing"
	"tx-api/dto"
	"tx-api/model"
	"tx-api/service"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMakeCreateAccountEndpoint_Success(t *testing.T) {
	mockSvc := new(service.MockService)
	endpointFn := MakeCreateAccountEndpoint(mockSvc)

	req := CreateAccountRequest{dto.CreateAccountRequest{DocumentNumber: "123"}}
	expectedResp := dto.CreateAccountResponse{AccountID: uuid.New()}

	mockSvc.On("CreateAccount", context.Background(), req.CreateAccountRequest).
		Return(expectedResp, nil)

	resp, err := endpointFn(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp.(CreateAccountResponse).CreateAccountResponse)

	mockSvc.AssertExpectations(t)
}

func TestMakeCreateAccountEndpoint_Error(t *testing.T) {
	mockSvc := new(service.MockService)
	endpointFn := MakeCreateAccountEndpoint(mockSvc)

	req := CreateAccountRequest{dto.CreateAccountRequest{DocumentNumber: "123"}}
	mockSvc.On("CreateAccount", context.Background(), req.CreateAccountRequest).
		Return(dto.CreateAccountResponse{}, errors.New("create error"))

	resp, err := endpointFn(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)

	mockSvc.AssertExpectations(t)
}

func TestMakeGetAccountEndpoint_Success(t *testing.T) {
	mockSvc := new(service.MockService)
	endpointFn := MakeGetAccountEndpoint(mockSvc)

	accountID := uuid.New()
	req := GetAccountRequest{dto.GetAccountRequest{AccountID: accountID}}
	expectedResp := dto.GetAccountResponse{
		Account: model.Account{AccountID: accountID, DocumentNumber: "456"},
	}

	mockSvc.On("GetAccount", context.Background(), req.GetAccountRequest).
		Return(expectedResp, nil)

	resp, err := endpointFn(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp.(GetAccountResponse).GetAccountResponse)

	mockSvc.AssertExpectations(t)
}

func TestMakeGetAccountEndpoint_Error(t *testing.T) {
	mockSvc := new(service.MockService)
	endpointFn := MakeGetAccountEndpoint(mockSvc)

	accountID := uuid.New()
	req := GetAccountRequest{dto.GetAccountRequest{AccountID: accountID}}

	mockSvc.On("GetAccount", context.Background(), req.GetAccountRequest).
		Return(dto.GetAccountResponse{}, errors.New("get error"))

	resp, err := endpointFn(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)

	mockSvc.AssertExpectations(t)
}

func TestMakeCreateTransactionEndpoint_Success(t *testing.T) {
	mockSvc := new(service.MockService)
	endpointFn := MakeCreateTransactionEndpoint(mockSvc)

	req := CreateTransactionRequest{dto.CreateTransactionRequest{
		AccountID:       uuid.New(),
		OperationTypeID: 1,
		Amount:          decimal.NewFromInt(100),
		IsPositive:      true,
	}}
	expectedResp := dto.CreateTransactionResponse{TransactionID: uuid.New()}

	mockSvc.On("CreateTransaction", context.Background(), req.CreateTransactionRequest).
		Return(expectedResp, nil)

	resp, err := endpointFn(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp.(CreateTransactionResponse).CreateTransactionResponse)

	mockSvc.AssertExpectations(t)
}

func TestMakeCreateTransactionEndpoint_Error(t *testing.T) {
	mockSvc := new(service.MockService)
	endpointFn := MakeCreateTransactionEndpoint(mockSvc)

	req := CreateTransactionRequest{dto.CreateTransactionRequest{
		AccountID:       uuid.New(),
		OperationTypeID: 1,
		Amount:          decimal.NewFromInt(100),
		IsPositive:      true,
	}}

	mockSvc.On("CreateTransaction", context.Background(), req.CreateTransactionRequest).
		Return(dto.CreateTransactionResponse{}, errors.New("tx error"))

	resp, err := endpointFn(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)

	mockSvc.AssertExpectations(t)
}
