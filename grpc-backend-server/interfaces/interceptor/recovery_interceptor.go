package interceptor

import (
	"context"
	"git.paylabo.com/c002/harp/backend-purchase/common/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		logger := logging.GetLoggerFromContext(ctx)
		defer func() {
			if recovered := recover(); recovered != nil {
				logger.Errorw("panic occurred", "error", recovered)
				err = status.Error(codes.Internal, "InternalServerError")
			}
		}()
		reply, err = handler(ctx, req)
		return
	}
}
