package interceptor

import (
	"context"
	"errors"
	"git.paylabo.com/c002/harp/backend-purchase/common/logging"
	"git.paylabo.com/c002/harp/backend-purchase/domain/domainerror"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandleInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger := logging.GetLoggerFromContext(ctx)
		res, err := handler(ctx, req)
		if err == nil {
			return res, nil
		}
		if domainerror.IsValidationError(err) {
			return nil, handleValidationError(err)
		}
		if domainerror.IsNotFoundError(err) {
			return nil, handleNotFound(err)
		}
		if domainerror.IsInternalServerError(err) {
			return nil, handleInternalServerError(err)
		}
		logger.Error("unknown error", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func handleValidationError(err error) error {
	var resourceError domainerror.ApplicationError
	errors.As(err, &resourceError)
	sts := status.New(codes.InvalidArgument, "invalid argument")
	violations := make([]*errdetails.BadRequest_FieldViolation, 0)
	for key, values := range resourceError.(*domainerror.ValidationError).GetFields() {
		for _, value := range values {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field:       key,
				Description: value,
			})
		}
	}
	stsWithDetails, _ := sts.WithDetails(&errdetails.BadRequest{FieldViolations: violations})
	return stsWithDetails.Err()
}

func handleNotFound(err error) error {
	return status.Error(codes.NotFound, err.Error())
}

func handleInternalServerError(err error) error {
	return status.Error(codes.Internal, err.Error())
}
