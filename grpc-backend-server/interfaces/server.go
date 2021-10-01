package interfaces

import (
	"git.paylabo.com/c002/harp/backend-purchase/app/usecase/purchase/order"
	"git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	"git.paylabo.com/c002/harp/backend-purchase/interfaces/interceptor"
	pb "git.paylabo.com/c002/harp/backend-purchase/interfaces/proto/git.paylabo.com/c002/harp"
	"git.paylabo.com/c002/harp/backend-purchase/interfaces/service"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	pbhelth "google.golang.org/grpc/health/grpc_health_v1"
)

type ServerParams struct {
	OrderRepository repository.OrderRepository
}

func NewServer(params ServerParams) *grpc.Server {
	options := make([]grpc.ServerOption, 0)
	options = append(options, grpc.UnaryInterceptor(
		middleware.ChainUnaryServer(
			interceptor.LoggerSupplyInterceptor(),
			interceptor.RecoveryInterceptor(),
			interceptor.LoggingInterceptor(),
			interceptor.ErrorHandleInterceptor(),
		)))
	server := grpc.NewServer(options...)

	registerOrdersServer(server, params)

	healthService := health.NewServer()
	pbhelth.RegisterHealthServer(server, healthService)
	return server
}

func registerOrdersServer(server *grpc.Server, params ServerParams) {
	orderService := service.NewOrderService(
		order.NewListOrders(params.OrderRepository),
	)
	pb.RegisterOrderServer(server, orderService)
}
