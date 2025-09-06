package logger

import "context"

type ctxKey string

const (
	traceIDKey ctxKey = "trace_id"
	userIDKey  ctxKey = "user_id"
)

// WithTraceAndUser injects trace ID and user ID into context.
func WithTraceAndUser(ctx context.Context, traceID, userID string) context.Context {
	ctx = context.WithValue(ctx, traceIDKey, traceID)
	ctx = context.WithValue(ctx, userIDKey, userID)
	return ctx
}

// GetTraceAndUser extracts trace ID and user ID from context.
func GetTraceAndUser(ctx context.Context) (traceID, userID string) {
	if v, ok := ctx.Value(traceIDKey).(string); ok {
		traceID = v
	}
	if v, ok := ctx.Value(userIDKey).(string); ok {
		userID = v
	}
	return
}
