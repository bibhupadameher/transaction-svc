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
	GetAccount(ctx context.Context, request dto.GetAccountRequest) (dto.GetAccountResponse, error)
	CreateTransaction(ctx context.Context, request dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error)
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
	defer logger.Info(ctx, "Finishing CreateAccount...", zap.Any("response", response))

	if err := request.Validate(); err != nil {
		logger.Error(ctx, "error on validating request", zap.Any("request", request))
		return response, err
	}

	var newAccount model.Account
	newAccount.DocumentNumber = request.DocumentNumber
	if err := s.dao.SaveAccount(ctx, &newAccount); err != nil {
		logger.Error(ctx, "error on saving account", zap.Error(err))
		return response, err
	}
	response.AccountID = newAccount.AccountID

	return response, nil
}

func (s *transactionServiceHandler) GetAccount(ctx context.Context, request dto.GetAccountRequest) (dto.GetAccountResponse, error) {
	logger.Info(ctx, "Starting GetAccount...", zap.Any("request", request))
	var response dto.GetAccountResponse
	defer logger.Info(ctx, "Finishing GetAccount...", zap.Any("response", response))

	if err := request.Validate(); err != nil {
		logger.Error(ctx, "error on validating request", zap.Any("request", request))
		return response, err
	}

	account, err := s.dao.GetAccountByID(ctx, request.AccountID)
	if err != nil {
		logger.Error(ctx, "error on get account id", zap.Error(err), zap.Any("account_id", request.AccountID))
		return response, err
	}
	response.Account = account

	return response, nil
}

func (s *transactionServiceHandler) CreateTransaction(ctx context.Context, request dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error) {
	logger.Info(ctx, "Starting CreateTransaction...", zap.Any("request", request))
	var response dto.CreateTransactionResponse
	defer logger.Info(ctx, "Finishing CreateTransaction...", zap.Any("response", response))

	if err := request.Validate(); err != nil {
		logger.Error(ctx, "error on validating request", zap.Any("request", request))
		return response, err
	}

	var tx model.Transaction
	tx.AccountID = request.AccountID
	tx.OperationTypeID = request.OperationTypeID
	if request.IsPositive {
		tx.Amount = request.Amount
	} else {
		tx.Amount = request.Amount.Neg()
	}

	if err := s.dao.SaveTransaction(ctx, &tx); err != nil {
		logger.Error(ctx, "error on saving transaction", zap.Error(err))
		return response, err
	}
	response.TransactionID = tx.TransactionID

	return response, nil
}
