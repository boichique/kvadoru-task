package logger

import (
	"context"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptor that adds a logger to the context
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		attrs := []any{slog.String("method", info.FullMethod)} // Create a new logger with the method name as an attribute
		logger := slog.Default().With(attrs...)

		ctx = WithLogger(ctx, logger) // Add the logger to the context
		return handler(ctx, req)      // Invoke the handler with the updated context and the incoming request, and get the response and error
	}
}
