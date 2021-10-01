package interceptor

import (
	"context"
	"git.paylabo.com/c002/harp/backend-purchase/common/logging"
	"google.golang.org/grpc"
)

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := logging.GetLoggerFromContext(ctx)
		logger.Infow("request ==>", "request", req)
		res, err := handler(ctx, req)
		if res != nil {
			logger.Infow("<== response", "response", res)
		}
		if err != nil {
			logger.Infow("<== response", "response_error", err)
		}
		return res, err
	}
}
