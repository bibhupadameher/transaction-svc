package kit

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// fakeFailer implements kitendpoint.Failer
type fakeFailer struct {
	err error
}

func (f fakeFailer) Failed() error {
	return f.err
}

func TestGenericEncodeResponse_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := map[string]string{"hello": "world"}

	err := GenericEncodeResponse(context.Background(), rr, resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Errorf("expected json content-type, got %q", ct)
	}

	var got map[string]string
	if decodeErr := json.NewDecoder(rr.Body).Decode(&got); decodeErr != nil {
		t.Fatalf("decode failed: %v", decodeErr)
	}
	if got["hello"] != "world" {
		t.Errorf("expected hello=world, got %v", got)
	}
}

func TestGenericEncodeResponse_WithFailer(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := fakeFailer{err: errors.New("fail error")}

	err := GenericEncodeResponse(context.Background(), rr, resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rr.Code)
	}

}
