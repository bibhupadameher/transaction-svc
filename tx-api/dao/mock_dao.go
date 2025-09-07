package dao

import (
	"context"
	"tx-api/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockTransactionDAO is a mock implementation of TransactionDAOInterface
type MockTransactionDAO struct {
	mock.Mock
}

func (m *MockTransactionDAO) GetAccountByID(ctx context.Context, id uuid.UUID) (model.Account, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockTransactionDAO) SaveAccount(ctx context.Context, account *model.Account) error {
	args := m.Called(ctx, account)
	return args.Error(0)
}

func (m *MockTransactionDAO) SaveTransaction(ctx context.Context, trans *model.Transaction) error {
	args := m.Called(ctx, trans)
	return args.Error(0)
}
