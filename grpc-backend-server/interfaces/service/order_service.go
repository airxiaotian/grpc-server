package service

import (
	"context"

	"git.paylabo.com/c002/harp/backend-purchase/app/usecase/purchase/order"
	"git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	pb "git.paylabo.com/c002/harp/backend-purchase/interfaces/proto/git.paylabo.com/c002/harp"
)

type OrderService struct {
	listOrders order.ListOrders
}

func NewOrderService(
	listOrders order.ListOrders,
) pb.OrderServer {
	return &OrderService{
		listOrders: listOrders,
	}
}

// ListOrders は発注一覧を取得する。
func (o *OrderService) ListOrders(ctx context.Context, request *pb.ListOrdersRequest) (*pb.OrdersResponse, error) {
	params := order.ListOrdersParams{
		ListOrdersParams: repository.ListOrdersParams{
			Limit:  int(request.Limit),
			Offset: int(request.Offset),
			FilterOrdersParams: repository.FilterOrdersParams{
				IDs:  request.Ids,
				name: request.SuppliersIds,
			},
		},
	}
	return &pb.OrdersResponse{}
}
