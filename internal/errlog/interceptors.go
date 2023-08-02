package errlog

import (
	"context"

	"github.com/boichique/kvadoru_task/internal/logger"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns a new unary server interceptor that logs errors
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		res, err := handler(ctx, req) // Invoke the handler with the incoming request and get the response and error
		if err != nil {
			logger := logger.FromContext(ctx)
			logger.Error("error", "handler error", err)
		}

		return res, err // Return the response and error to the client
	}
}
