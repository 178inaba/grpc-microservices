package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/178inaba/grpc-microservices/shared/metadata"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func XTraceID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		traceID := metadata.GetTraceIDFromContext(ctx)
		ctx = metadata.AddTraceIDToContext(ctx, traceID)
		return handler(ctx, req)
	}
}

func Logging(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		h, err := handler(ctx, req)
		if err != nil {
			logger.Error(info.FullMethod,
				zap.String("trace_id", metadata.GetTraceIDFromContext(ctx)),
				zap.Duration("elapsed_time", time.Since(start)),
				zap.String("status_code", status.Code(err).String()),
				zap.Error(err))
		} else {
			logger.Info(info.FullMethod,
				zap.String("trace_id", metadata.GetTraceIDFromContext(ctx)),
				zap.Duration("elapsed_time", time.Since(start)),
				zap.String("status_code", status.Code(err).String()))
		}

		return h, err
	}
}

func XUserID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		userID, err := metadata.SafeGetUserIDFromContext(ctx)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("safe get user id from context: %v", err))
		}

		ctx = metadata.AddUserIDToContext(ctx, userID)
		return handler(ctx, req)
	}
}
