package repository

import (
	"context"
	"errors"

	"git.paylabo.com/c002/harp/backend-purchase/app/usecase/utils"
	"git.paylabo.com/c002/harp/backend-purchase/domain/model"
	"git.paylabo.com/c002/harp/backend-purchase/domain/validator"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type OrderRepository interface {
	ListUnacceptedOrdersThisMonth(ctx context.Context, params ListOrdersParams) ([]*model.Order, error)
	ListOrders(ctx context.Context, params ListOrdersParams) ([]*model.Order, error)
	GetOrder(ctx context.Context, id string) (*model.Order, error)
	CountOrders(ctx context.Context, params CountOrdersParams) (int64, error)
	CreateOrder(ctx context.Context, params CreateOrderParams) (CreateOrderReturning, error)
	UpdateOrder(ctx context.Context, params UpdateOrderParams) (int64, error)
	UpdateOrderProjectCostInfo(ctx context.Context, params UpdateOrderProjectCostInfoParams) (int64, error)
	DeleteOrder(ctx context.Context, params DeleteOrderParams) (int64, error)
	GetOrderRequesterAggregate(ctx context.Context, params GetOrderAggregateParams) ([]*model.Order, error)
	GetOrderSupplierAggregate(ctx context.Context, params GetOrderAggregateParams) ([]*model.Order, error)
	CountOrdersWithGroupBy(ctx context.Context, params CountOrdersWithGroupByParams) ([]*model.OrdersGroupBy, error)
	SumNearestTwoMonthsAmount(ctx context.Context, params SumNearestTwoMonthsAmountParams) ([]*model.NearestTwoMonthsAmount, error)
}

type FilterOrdersParams struct {
	IDs                    []int32
	SuppliersIDs           []int32
	Subject                string
	RequestOrganizationIDs []string
	RequestBys             []string
	OrderCaseCds           []int32
	OrderStatuses          []int32
	ProjectsIDs            []string
	ProjectCostIDs         []string
}

type CreateOrderReturning struct {
	AffectedRows int64
	Order        *model.Order
}

type CountOrdersParams struct {
	FilterOrdersParams
}

type SumOrdersParams struct {
	IDs []int32
}

type CreateOrderParams struct {
	OrderNo                    string
	SuppliersID                string
	CompanyGroupType           string
	Subject                    string
	RequestOrganizationID      string
	RequestDate                *timestamp.Timestamp
	RequestBy                  string
	ApprovalFile               string
	DerivationSourceOrderID    *string
	Remarks                    string
	SuperiorApprovalDate       *timestamp.Timestamp
	PurchasingDeptApprovalDate *timestamp.Timestamp
	OrderIssueDate             *timestamp.Timestamp
	FinalAcceptanceDate        *timestamp.Timestamp
	AcceptanceCompletedDate    *timestamp.Timestamp
	CancelDate                 *timestamp.Timestamp
	OrderCaseCd                string
	OrderStatus                string
	JiraNo                     string
	QuotationsID               *string
	OrderApprovalStaffsID      string
}

type UpdateOrderParams struct {
	ID string
	CreateOrderParams
}

type UpdateOrderProjectCostInfoParams struct {
	ID            string
	ProjectsID    string
	ProjectCostID string
	CostTypesID   string
}

type DeleteOrderParams struct {
	ID string
}

type ListOrdersParams struct {
	Limit   int
	Offset  int
	OrderBy *ListOrdersOrderBy
	FilterOrdersParams
}
type ListOrdersOrderBy struct {
	Id                  utils.SortEnum
	Subject             utils.SortEnum
	OrderCaseCd         utils.SortEnum
	OrderStatus         utils.SortEnum
	SuppliersId         utils.SortEnum
	RequestDate         utils.SortEnum
	FinalAcceptanceDate utils.SortEnum
	UpdatedAt           utils.SortEnum
}
type GetOrderAggregateParams struct {
}

func ExistingOrderRule(OrderRepository OrderRepository) func(interface{}) error {
	return func(value interface{}) error {
		if OrderRepository == nil {
			return errors.New(validator.NotExistingMessage)
		}
		id, ok := value.(string)
		if !ok {
			return errors.New(validator.NotExistingMessage)
		}
		if _, err := OrderRepository.GetOrder(context.Background(), id); err != nil {
			return errors.New(validator.NotExistingMessage)
		}
		return nil
	}
}

type CountOrdersWithGroupByParams struct {
	RequestOrganizationID string
	RequestBy             string
	RecentMonth           int64
	GroupBy               string
}

type SumNearestTwoMonthsAmountParams struct {
	RequestOrganizationID string
}
