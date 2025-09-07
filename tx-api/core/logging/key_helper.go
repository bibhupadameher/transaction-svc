package logger

import "context"

type ctxKey string

const (
	traceIDKey ctxKey = "trace_id"
)

func WithTrace(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, traceIDKey, traceID)

	return ctx
}

func GetTrace(ctx context.Context) (traceID string) {
	if v, ok := ctx.Value(traceIDKey).(string); ok {
		traceID = v
	}

	return
}
