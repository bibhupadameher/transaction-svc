package dao

import (
	"context"
	"tx-api/core/postgres"
	"tx-api/model"
)

type TransactionDAOInterface interface {
	GetAccountByID(ctx context.Context, id string) (model.Account, error)
	BatchWriteData(ctx context.Context, savedData []model.TableName, deletedData ...model.TableName) error
}

// Implementation
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
