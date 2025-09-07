package logger

import (
	"context"
	"testing"
)

func TestLoggingFunctions_NoPanic(t *testing.T) {
	_ = Init()
	ctx := WithTrace(context.Background(), "trace-xyz")

	Debug(ctx, "debug message")
	Info(ctx, "info message")
	Warn(ctx, "warn message")
	Error(ctx, "error message")
}

func TestGetLogger_ReturnsLogger(t *testing.T) {
	_ = Init()
	l := GetLogger()
	if l == nil {
		t.Fatalf("expected GetLogger() to return non-nil logger")
	}
}
