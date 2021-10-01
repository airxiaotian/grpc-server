package interceptor

import (
	"context"
	"git.paylabo.com/c002/harp/backend-purchase/common/logging"
	"google.golang.org/grpc"
)

func LoggerSupplyInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := logging.NewLogger()
		logger = logging.With(logger, "method", info.FullMethod)
		ctx = logging.SetLoggerToContext(ctx, logger)
		return handler(ctx, req)
	}
}
