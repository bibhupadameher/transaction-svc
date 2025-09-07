package service

import (
	"context"
	logger "tx-api/core/logging"
	"tx-api/dao"
	"tx-api/dto"
	"tx-api/model"

	"go.uber.org/zap"
)

type TransactionServiceInterface interface {
	CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (dto.CreateAccountResponse, error)
}

type transactionServiceHandler struct {
	dao    dao.TransactionDAOInterface
	logger *zap.Logger
}

var (
	newDao = dao.New
)

func initHandler(_ context.Context, logger *zap.Logger) (*transactionServiceHandler, error) {
	db, err := newDao()
	if err != nil {
		return nil, err
	}

	return &transactionServiceHandler{
		dao:    db,
		logger: logger,
	}, nil
}

func New(ctx context.Context, logger *zap.Logger) (TransactionServiceInterface, error) {
	serviceLogger := logger.With(zap.String("location", "service"))
	serviceLogger.Info("setting up transaction service handler...")

	var svc TransactionServiceInterface
	var err error
	svc, err = initHandler(ctx, serviceLogger)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *transactionServiceHandler) CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (dto.CreateAccountResponse, error) {
	logger.Info(ctx, "Starting CreateAccount...", zap.Any("request", request))
	var response dto.CreateAccountResponse

	if err := request.Validate(); err != nil {
		logger.Error(ctx, "error on validating request", zap.Any("request", request))
		return response, err
	}

	var newAccount model.Account
	newAccount.DocumentNumber = request.DocumentNumber
	if err := s.dao.BatchWriteData(ctx, []model.TableName{&newAccount}); err != nil {
		logger.Error(ctx, "error on saving account", zap.Any("request", request))
		return response, err
	}

	response.AccountID = newAccount.AccountID
	logger.Info(ctx, "Finishing CreateAccount...", zap.Any("response", response))
	return response, nil
}
