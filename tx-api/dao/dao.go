package dao

import (
	"context"
	"tx-api/core/postgres"
	"tx-api/model"

	"github.com/google/uuid"
)

type TransactionDAOInterface interface {
	GetAccountByID(ctx context.Context, id uuid.UUID) (model.Account, error)
	SaveAccount(ctx context.Context, account *model.Account) error
	SaveTransactions(ctx context.Context, trans []model.Transaction) error
	GetActiveTransactionsByAccountID(ctx context.Context, id uuid.UUID) ([]model.Transaction, error) //Active tx == transactions with balance<0
}

type transactionDAO struct {
	pgSvc postgres.DBService
}

func New() (TransactionDAOInterface, error) {
	pgSvc, err := postgres.NewPostgresService()
	if err != nil {
		return nil, err
	}
	return &transactionDAO{pgSvc: pgSvc}, nil
}
