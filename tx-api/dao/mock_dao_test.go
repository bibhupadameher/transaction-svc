package dao

import (
	"context"
	"testing"
	"tx-api/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMockTransactionDAO_GetAccountByID(t *testing.T) {
	mockDAO := new(MockTransactionDAO)
	ctx := context.Background()
	id := uuid.New()
	expectedAccount := model.Account{AccountID: uuid.New(), DocumentNumber: "Save Test"}

	mockDAO.On("GetAccountByID", ctx, id).Return(expectedAccount, nil)

	account, err := mockDAO.GetAccountByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)

	mockDAO.AssertCalled(t, "GetAccountByID", ctx, id)
}

func TestMockTransactionDAO_SaveAccount(t *testing.T) {
	mockDAO := new(MockTransactionDAO)
	ctx := context.Background()
	account := &model.Account{AccountID: uuid.New(), DocumentNumber: "Save Test"}

	mockDAO.On("SaveAccount", ctx, account).Return(nil)

	err := mockDAO.SaveAccount(ctx, account)
	assert.NoError(t, err)
	mockDAO.AssertCalled(t, "SaveAccount", ctx, account)
}

func TestMockTransactionDAO_SaveTransaction(t *testing.T) {
	mockDAO := new(MockTransactionDAO)
	ctx := context.Background()
	trans := &model.Transaction{AccountID: uuid.New()}

	mockDAO.On("SaveTransaction", ctx, trans).Return(nil)

	err := mockDAO.SaveTransaction(ctx, trans)
	assert.NoError(t, err)
	mockDAO.AssertCalled(t, "SaveTransaction", ctx, trans)
}
