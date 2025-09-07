package endpoint

import (
	"context"
	"time"

	logger "tx-api/core/logging"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func LoggingMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			start := time.Now()

			ctx = logger.WithTrace(ctx, uuid.New().String())
			traceID := logger.GetTrace(ctx)

			resp, err := next(ctx, req)

			elapsed := time.Since(start).String()
			//	traceID = logger.GetTrace(ctx)

			fields := []zap.Field{
				zap.String("trace_id", traceID),
				zap.String("elapsed", elapsed),
			}

			if err != nil {
				logger.Error(ctx, "request failed", fields...)
			} else {
				logger.Info(ctx, "request completed", fields...)
			}

			return resp, err
		}
	}
}
