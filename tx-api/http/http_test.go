package httptransport

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"tx-api/dto"
	"tx-api/endpoint"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// ---------- decodeCreateAccountRequest ----------
func TestDecodeCreateAccountRequest_Success(t *testing.T) {
	reqBody := `{"documentNumber":"12345"}`
	r := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBufferString(reqBody))

	decoded, err := decodeCreateAccountRequest(context.Background(), r)

	assert.NoError(t, err)
	req := decoded.(endpoint.CreateAccountRequest)
	assert.Equal(t, "12345", req.DocumentNumber)
}

func TestDecodeCreateAccountRequest_InvalidJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBufferString("{bad json"))

	decoded, err := decodeCreateAccountRequest(context.Background(), r)

	assert.Error(t, err)
	assert.Nil(t, decoded)
}

// ---------- decodeCreateTransactionRequest ----------
func TestDecodeCreateTransactionRequest_Success(t *testing.T) {
	reqBody := `{"accountID":"` + uuid.New().String() + `","operationTypeID":1,"amount":"100.00"}`
	r := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(reqBody))

	decoded, err := decodeCreateTransactionRequest(context.Background(), r)

	assert.NoError(t, err)
	req := decoded.(endpoint.CreateTransactionRequest)
	assert.Equal(t, 1, req.OperationTypeID)
}

func TestDecodeCreateTransactionRequest_InvalidJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString("{bad json"))

	decoded, err := decodeCreateTransactionRequest(context.Background(), r)

	assert.Error(t, err)
	assert.Nil(t, decoded)
}

// ---------- decodeGetAccountRequest ----------
func TestDecodeGetAccountRequest_Success(t *testing.T) {
	id := uuid.New()
	r := httptest.NewRequest(http.MethodGet, "/accounts/"+id.String(), nil)
	r = mux.SetURLVars(r, map[string]string{"accountId": id.String()})

	decoded, err := decodeGetAccountRequest(context.Background(), r)

	assert.NoError(t, err)
	req := decoded.(endpoint.GetAccountRequest)
	assert.Equal(t, id, req.AccountID)
}

func TestDecodeGetAccountRequest_InvalidUUID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/accounts/invalid-uuid", nil)
	r = mux.SetURLVars(r, map[string]string{"accountId": "invalid-uuid"})

	decoded, err := decodeGetAccountRequest(context.Background(), r)

	assert.Error(t, err)
	assert.Nil(t, decoded)
}

func TestDecodeGetAccountRequest_MissingParam(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/accounts", nil)
	// no vars set

	decoded, err := decodeGetAccountRequest(context.Background(), r)

	assert.Error(t, err)
	assert.Nil(t, decoded)
}

// ---------- NewHTTPHandler ----------
func TestNewHTTPHandler_Routes(t *testing.T) {
	// fake endpoints
	okEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{"status": "ok"}, nil
	}

	endpoints := endpoint.Set{
		CreateAccountEndpoint:     okEndpoint,
		GetAccountEndpoint:        okEndpoint,
		CreateTransactionEndpoint: okEndpoint,
	}

	handler := NewHTTPHandler(endpoints)

	// test POST /accounts
	w := httptest.NewRecorder()
	body, _ := json.Marshal(endpoint.CreateAccountRequest{CreateAccountRequest: dto.CreateAccountRequest{DocumentNumber: "12345"}})
	r := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(body))
	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	// test GET /accounts/{accountId}
	id := uuid.New()
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/accounts/"+id.String(), nil)
	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	// test POST /transactions
	w = httptest.NewRecorder()
	body, _ = json.Marshal(endpoint.CreateTransactionRequest{CreateTransactionRequest: dto.CreateTransactionRequest{AccountID: id, OperationTypeID: 1}})
	r = httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
