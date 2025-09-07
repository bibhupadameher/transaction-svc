package dto

import (
	"tx-api/core/apperror"
	"tx-api/errors"
	"tx-api/model"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"documentNumber" valid:"required"`
}

func (req *CreateAccountRequest) Validate() apperror.Errors {
	var valErrors apperror.Errors
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		valErrors = req.mapValidatorErr2WpErr(err)
	}

	if len(valErrors) > 0 {
		return valErrors
	}
	return nil
}

func (req *CreateAccountRequest) mapValidatorErr2WpErr(err error) apperror.Errors {
	var wpErr apperror.Errors
	errs := err.(govalidator.Errors).Errors()
	for _, e := range errs {
		if gErr, ok := e.(govalidator.Errors); ok {
			wpErr = append(wpErr, req.mapValidatorErr2WpErr(gErr)...)
			continue
		}
		gErr := e.(govalidator.Error)
		wpErr = append(wpErr, errors.NewErrFieldMissingorInvalid(gErr.Name).AppError)
	}
	return wpErr
}

type GetAccountRequest struct {
	AccountID uuid.UUID `json:"accountID" valid:"required,uuid"`
}

func (req *GetAccountRequest) Validate() apperror.Errors {
	var valErrors apperror.Errors
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		valErrors = req.mapValidatorErr2WpErr(err)
	}

	if len(valErrors) > 0 {
		return valErrors
	}
	return nil
}

func (req *GetAccountRequest) mapValidatorErr2WpErr(err error) apperror.Errors {
	var wpErr apperror.Errors
	errs := err.(govalidator.Errors).Errors()
	for _, e := range errs {
		if gErr, ok := e.(govalidator.Errors); ok {
			wpErr = append(wpErr, req.mapValidatorErr2WpErr(gErr)...)
			continue
		}
		gErr := e.(govalidator.Error)
		wpErr = append(wpErr, errors.NewErrFieldMissingorInvalid(gErr.Name).AppError)
	}
	return wpErr
}

type CreateTransactionRequest struct {
	AccountID       uuid.UUID       `json:"accountID" valid:"required"`
	OperationTypeID int             `json:"operationTypeID"  valid:"required"`
	Amount          decimal.Decimal `json:"amount"  valid:"required"`
	IsPositive      bool            `json:"-"` //injected
}

func (req *CreateTransactionRequest) Validate() apperror.Errors {
	var valErrors apperror.Errors
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		valErrors = req.mapValidatorErr2WpErr(err)
	}

	if req.OperationTypeID > 0 {
		isValid := false
		for _, op := range model.DefaultEnum.OperationTypes {
			if op.OperationTypeID == req.OperationTypeID {
				isValid = true
				req.IsPositive = op.IsPositive
				break
			}
		}
		if !isValid {
			valErrors = append(valErrors, errors.NewErrFieldMissingorInvalid("operationTypeID").AppError)
		}
	}
	if req.Amount.IsNegative() {
		valErrors = append(valErrors, errors.NewErrFieldMissingorInvalid("amount").AppError)
	}

	if len(valErrors) > 0 {
		return valErrors
	}
	return nil
}

func (req *CreateTransactionRequest) mapValidatorErr2WpErr(err error) apperror.Errors {
	var wpErr apperror.Errors
	errs := err.(govalidator.Errors).Errors()
	for _, e := range errs {
		if gErr, ok := e.(govalidator.Errors); ok {
			wpErr = append(wpErr, req.mapValidatorErr2WpErr(gErr)...)
			continue
		}
		gErr := e.(govalidator.Error)
		wpErr = append(wpErr, errors.NewErrFieldMissingorInvalid(gErr.Name).AppError)
	}
	return wpErr
}
