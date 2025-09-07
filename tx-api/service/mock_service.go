package service

import (
	"context"
	"tx-api/dto"

	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of TransactionServiceInterface
type MockService struct {
	mock.Mock
}

func (m *MockService) CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (dto.CreateAccountResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(dto.CreateAccountResponse), args.Error(1)
}

func (m *MockService) GetAccount(ctx context.Context, request dto.GetAccountRequest) (dto.GetAccountResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(dto.GetAccountResponse), args.Error(1)
}

func (m *MockService) CreateTransaction(ctx context.Context, request dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(dto.CreateTransactionResponse), args.Error(1)
}
