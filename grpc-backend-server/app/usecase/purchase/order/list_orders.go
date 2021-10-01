package order

import (
	"context"

	"git.paylabo.com/c002/harp/backend-purchase/domain/model"
	"git.paylabo.com/c002/harp/backend-purchase/domain/repository"
)

type ListOrdersParams struct {
	repository.ListOrdersParams
}

func (l *ListOrdersParams) validate() error {
	return nil
}

type ListOrders func(ctx context.Context, params ListOrdersParams) ([]*model.Order, error)

func NewListOrders(OrderRepository repository.OrderRepository) ListOrders {
	return func(ctx context.Context, params ListOrdersParams) ([]*model.Order, error) {
		if err := params.validate(); err != nil {
			return nil, err
		}
		limit := 10
		if params.Limit == 5 || params.Limit == 10 || params.Limit == 15 {
			limit = params.Limit
		}
		if params.Limit == 0 {
			limit = 0
		}
		offset := 0
		if params.Offset >= 0 {
			offset = params.Offset
		}
		return OrderRepository.ListOrders(ctx, repository.ListOrdersParams{
			Limit:   limit,
			Offset:  offset,
			OrderBy: params.OrderBy,
			FilterOrdersParams: repository.FilterOrdersParams{
				IDs: params.IDs,
			},
		})
	}
}
