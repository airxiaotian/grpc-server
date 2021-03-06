// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/order_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	model "git.paylabo.com/c002/harp/backend-purchase/domain/model"
	repository "git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// ListUnacceptedOrdersThisMonth mocks base method
func (m *MockOrderRepository) ListUnacceptedOrdersThisMonth(ctx context.Context, params repository.ListOrdersParams) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnacceptedOrdersThisMonth", ctx, params)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnacceptedOrdersThisMonth indicates an expected call of ListUnacceptedOrdersThisMonth
func (mr *MockOrderRepositoryMockRecorder) ListUnacceptedOrdersThisMonth(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnacceptedOrdersThisMonth", reflect.TypeOf((*MockOrderRepository)(nil).ListUnacceptedOrdersThisMonth), ctx, params)
}

// ListOrders mocks base method
func (m *MockOrderRepository) ListOrders(ctx context.Context, params repository.ListOrdersParams) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", ctx, params)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders
func (mr *MockOrderRepositoryMockRecorder) ListOrders(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockOrderRepository)(nil).ListOrders), ctx, params)
}

// GetOrder mocks base method
func (m *MockOrderRepository) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, id)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder
func (mr *MockOrderRepositoryMockRecorder) GetOrder(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrderRepository)(nil).GetOrder), ctx, id)
}

// CountOrders mocks base method
func (m *MockOrderRepository) CountOrders(ctx context.Context, params repository.CountOrdersParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountOrders", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountOrders indicates an expected call of CountOrders
func (mr *MockOrderRepositoryMockRecorder) CountOrders(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountOrders", reflect.TypeOf((*MockOrderRepository)(nil).CountOrders), ctx, params)
}

// CreateOrder mocks base method
func (m *MockOrderRepository) CreateOrder(ctx context.Context, params repository.CreateOrderParams) (repository.CreateOrderReturning, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, params)
	ret0, _ := ret[0].(repository.CreateOrderReturning)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), ctx, params)
}

// UpdateOrder mocks base method
func (m *MockOrderRepository) UpdateOrder(ctx context.Context, params repository.UpdateOrderParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder
func (mr *MockOrderRepositoryMockRecorder) UpdateOrder(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOrderRepository)(nil).UpdateOrder), ctx, params)
}

// UpdateOrderProjectCostInfo mocks base method
func (m *MockOrderRepository) UpdateOrderProjectCostInfo(ctx context.Context, params repository.UpdateOrderProjectCostInfoParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderProjectCostInfo", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderProjectCostInfo indicates an expected call of UpdateOrderProjectCostInfo
func (mr *MockOrderRepositoryMockRecorder) UpdateOrderProjectCostInfo(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderProjectCostInfo", reflect.TypeOf((*MockOrderRepository)(nil).UpdateOrderProjectCostInfo), ctx, params)
}

// DeleteOrder mocks base method
func (m *MockOrderRepository) DeleteOrder(ctx context.Context, params repository.DeleteOrderParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrder indicates an expected call of DeleteOrder
func (mr *MockOrderRepositoryMockRecorder) DeleteOrder(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOrderRepository)(nil).DeleteOrder), ctx, params)
}

// GetOrderRequesterAggregate mocks base method
func (m *MockOrderRepository) GetOrderRequesterAggregate(ctx context.Context, params repository.GetOrderAggregateParams) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderRequesterAggregate", ctx, params)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderRequesterAggregate indicates an expected call of GetOrderRequesterAggregate
func (mr *MockOrderRepositoryMockRecorder) GetOrderRequesterAggregate(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderRequesterAggregate", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderRequesterAggregate), ctx, params)
}

// GetOrderSupplierAggregate mocks base method
func (m *MockOrderRepository) GetOrderSupplierAggregate(ctx context.Context, params repository.GetOrderAggregateParams) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderSupplierAggregate", ctx, params)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderSupplierAggregate indicates an expected call of GetOrderSupplierAggregate
func (mr *MockOrderRepositoryMockRecorder) GetOrderSupplierAggregate(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderSupplierAggregate", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderSupplierAggregate), ctx, params)
}

// CountOrdersWithGroupBy mocks base method
func (m *MockOrderRepository) CountOrdersWithGroupBy(ctx context.Context, params repository.CountOrdersWithGroupByParams) ([]*model.OrdersGroupBy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountOrdersWithGroupBy", ctx, params)
	ret0, _ := ret[0].([]*model.OrdersGroupBy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountOrdersWithGroupBy indicates an expected call of CountOrdersWithGroupBy
func (mr *MockOrderRepositoryMockRecorder) CountOrdersWithGroupBy(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountOrdersWithGroupBy", reflect.TypeOf((*MockOrderRepository)(nil).CountOrdersWithGroupBy), ctx, params)
}

// SumNearestTwoMonthsAmount mocks base method
func (m *MockOrderRepository) SumNearestTwoMonthsAmount(ctx context.Context, params repository.SumNearestTwoMonthsAmountParams) ([]*model.NearestTwoMonthsAmount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumNearestTwoMonthsAmount", ctx, params)
	ret0, _ := ret[0].([]*model.NearestTwoMonthsAmount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumNearestTwoMonthsAmount indicates an expected call of SumNearestTwoMonthsAmount
func (mr *MockOrderRepositoryMockRecorder) SumNearestTwoMonthsAmount(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumNearestTwoMonthsAmount", reflect.TypeOf((*MockOrderRepository)(nil).SumNearestTwoMonthsAmount), ctx, params)
}
