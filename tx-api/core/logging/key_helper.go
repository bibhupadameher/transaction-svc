package logger

import "context"

type ctxKey string

const (
	traceIDKey ctxKey = "trace_id"
)

// WithTraceAndUser injects trace ID and user ID into context.
func WithTrace(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, traceIDKey, traceID)

	return ctx
}

// GetTraceAndUser extracts trace ID and user ID from context.
func GetTrace(ctx context.Context) (traceID string) {
	if v, ok := ctx.Value(traceIDKey).(string); ok {
		traceID = v
	}

	return
}
