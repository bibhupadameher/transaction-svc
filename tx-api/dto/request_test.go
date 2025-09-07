package dto_test

import (
	"testing"
	"tx-api/dto"
	"tx-api/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		req       dto.CreateAccountRequest
		wantError bool
	}{
		{
			name:      "valid request",
			req:       dto.CreateAccountRequest{DocumentNumber: "12345"},
			wantError: false,
		},
		{
			name:      "missing document number",
			req:       dto.CreateAccountRequest{DocumentNumber: ""},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.req.Validate()
			if tt.wantError {
				assert.NotNil(t, errs)
				assert.NotEmpty(t, errs)
			} else {
				assert.Nil(t, errs)
			}
		})
	}
}

func TestGetAccountRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		req       dto.GetAccountRequest
		wantError bool
	}{
		{
			name:      "valid request",
			req:       dto.GetAccountRequest{AccountID: uuid.New()},
			wantError: false,
		},
		{
			name:      "missing account id",
			req:       dto.GetAccountRequest{},
			wantError: true,
		},
		{
			name:      "invalid uuid (zero value is fine, but simulate bad input)",
			req:       dto.GetAccountRequest{AccountID: uuid.Nil},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.req.Validate()
			if tt.wantError {
				assert.NotNil(t, errs)
				assert.NotEmpty(t, errs)
			} else {
				assert.Nil(t, errs)
			}
		})
	}
}

func TestCreateTransactionRequest_Validate(t *testing.T) {
	// inject model defaults for test
	model.DefaultEnum.OperationTypes = []model.OperationType{
		{OperationTypeID: 1, Description: "Payment", IsPositive: false},
		{OperationTypeID: 2, Description: "Deposit", IsPositive: true},
	}

	tests := []struct {
		name      string
		req       dto.CreateTransactionRequest
		wantError bool
	}{
		{
			name: "valid deposit transaction",
			req: dto.CreateTransactionRequest{
				AccountID:       uuid.New(),
				OperationTypeID: 2,
				Amount:          decimal.NewFromInt(100),
			},
			wantError: false,
		},
		{
			name: "invalid operation type id",
			req: dto.CreateTransactionRequest{
				AccountID:       uuid.New(),
				OperationTypeID: 99,
				Amount:          decimal.NewFromInt(50),
			},
			wantError: true,
		},
		{
			name: "negative amount",
			req: dto.CreateTransactionRequest{
				AccountID:       uuid.New(),
				OperationTypeID: 1,
				Amount:          decimal.NewFromInt(-10),
			},
			wantError: true,
		},
		{
			name: "missing account id",
			req: dto.CreateTransactionRequest{
				OperationTypeID: 1,
				Amount:          decimal.NewFromInt(20),
			},
			wantError: true,
		},
		{
			name: "missing amount",
			req: dto.CreateTransactionRequest{
				AccountID:       uuid.New(),
				OperationTypeID: 1,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.req.Validate()
			if tt.wantError {
				assert.NotNil(t, errs)
				assert.NotEmpty(t, errs)
			} else {
				assert.Nil(t, errs)
				// also check that IsPositive was set correctly
				if tt.req.OperationTypeID == 2 {
					assert.True(t, tt.req.IsPositive)
				}
			}
		})
	}
}
