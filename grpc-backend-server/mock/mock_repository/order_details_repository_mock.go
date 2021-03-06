// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/order_details_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "git.paylabo.com/c002/harp/backend-purchase/domain/model"
	repository "git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderDetailsRepository is a mock of OrderDetailsRepository interface
type MockOrderDetailsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderDetailsRepositoryMockRecorder
}

// MockOrderDetailsRepositoryMockRecorder is the mock recorder for MockOrderDetailsRepository
type MockOrderDetailsRepositoryMockRecorder struct {
	mock *MockOrderDetailsRepository
}

// NewMockOrderDetailsRepository creates a new mock instance
func NewMockOrderDetailsRepository(ctrl *gomock.Controller) *MockOrderDetailsRepository {
	mock := &MockOrderDetailsRepository{ctrl: ctrl}
	mock.recorder = &MockOrderDetailsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderDetailsRepository) EXPECT() *MockOrderDetailsRepositoryMockRecorder {
	return m.recorder
}

// CountOrderDetails mocks base method
func (m *MockOrderDetailsRepository) CountOrderDetails(ctx context.Context, params repository.FilterOrderDetailsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountOrderDetails", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountOrderDetails indicates an expected call of CountOrderDetails
func (mr *MockOrderDetailsRepositoryMockRecorder) CountOrderDetails(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountOrderDetails", reflect.TypeOf((*MockOrderDetailsRepository)(nil).CountOrderDetails), ctx, params)
}

// GetOrderDetail mocks base method
func (m *MockOrderDetailsRepository) GetOrderDetail(ctx context.Context, params repository.GetOrderDetailParams) (*model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDetail", ctx, params)
	ret0, _ := ret[0].(*model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDetail indicates an expected call of GetOrderDetail
func (mr *MockOrderDetailsRepositoryMockRecorder) GetOrderDetail(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDetail", reflect.TypeOf((*MockOrderDetailsRepository)(nil).GetOrderDetail), ctx, params)
}

// ListOrderDetails mocks base method
func (m *MockOrderDetailsRepository) ListOrderDetails(ctx context.Context, params repository.ListOrderDetailsParams) ([]*model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrderDetails", ctx, params)
	ret0, _ := ret[0].([]*model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrderDetails indicates an expected call of ListOrderDetails
func (mr *MockOrderDetailsRepositoryMockRecorder) ListOrderDetails(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrderDetails", reflect.TypeOf((*MockOrderDetailsRepository)(nil).ListOrderDetails), ctx, params)
}

// SumOrderDetailsOrderQuantity mocks base method
func (m *MockOrderDetailsRepository) SumOrderDetailsOrderQuantity(ctx context.Context, params repository.FilterOrderDetailsParams) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumOrderDetailsOrderQuantity", ctx, params)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumOrderDetailsOrderQuantity indicates an expected call of SumOrderDetailsOrderQuantity
func (mr *MockOrderDetailsRepositoryMockRecorder) SumOrderDetailsOrderQuantity(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumOrderDetailsOrderQuantity", reflect.TypeOf((*MockOrderDetailsRepository)(nil).SumOrderDetailsOrderQuantity), ctx, params)
}

// SumOrderDetailsCancelQuantity mocks base method
func (m *MockOrderDetailsRepository) SumOrderDetailsCancelQuantity(ctx context.Context, params repository.FilterOrderDetailsParams) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SumOrderDetailsCancelQuantity", ctx, params)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SumOrderDetailsCancelQuantity indicates an expected call of SumOrderDetailsCancelQuantity
func (mr *MockOrderDetailsRepositoryMockRecorder) SumOrderDetailsCancelQuantity(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SumOrderDetailsCancelQuantity", reflect.TypeOf((*MockOrderDetailsRepository)(nil).SumOrderDetailsCancelQuantity), ctx, params)
}

// CreateOrderDetail mocks base method
func (m *MockOrderDetailsRepository) CreateOrderDetail(ctx context.Context, params repository.CreateOrderDetailParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderDetail", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderDetail indicates an expected call of CreateOrderDetail
func (mr *MockOrderDetailsRepositoryMockRecorder) CreateOrderDetail(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderDetail", reflect.TypeOf((*MockOrderDetailsRepository)(nil).CreateOrderDetail), ctx, params)
}

// CreateOrderDetails mocks base method
func (m *MockOrderDetailsRepository) CreateOrderDetails(ctx context.Context, params repository.CreateOrderDetailsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderDetails", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderDetails indicates an expected call of CreateOrderDetails
func (mr *MockOrderDetailsRepositoryMockRecorder) CreateOrderDetails(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderDetails", reflect.TypeOf((*MockOrderDetailsRepository)(nil).CreateOrderDetails), ctx, params)
}

// UpdateOrderDetail mocks base method
func (m *MockOrderDetailsRepository) UpdateOrderDetail(ctx context.Context, params repository.UpdateOrderDetailParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderDetail", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderDetail indicates an expected call of UpdateOrderDetail
func (mr *MockOrderDetailsRepositoryMockRecorder) UpdateOrderDetail(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderDetail", reflect.TypeOf((*MockOrderDetailsRepository)(nil).UpdateOrderDetail), ctx, params)
}

// DeleteOrderDetail mocks base method
func (m *MockOrderDetailsRepository) DeleteOrderDetail(ctx context.Context, params repository.DeleteOrderDetailParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderDetail", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrderDetail indicates an expected call of DeleteOrderDetail
func (mr *MockOrderDetailsRepositoryMockRecorder) DeleteOrderDetail(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderDetail", reflect.TypeOf((*MockOrderDetailsRepository)(nil).DeleteOrderDetail), ctx, params)
}

// DeleteOrderDetails mocks base method
func (m *MockOrderDetailsRepository) DeleteOrderDetails(ctx context.Context, params repository.DeleteOrderDetailsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderDetails", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrderDetails indicates an expected call of DeleteOrderDetails
func (mr *MockOrderDetailsRepositoryMockRecorder) DeleteOrderDetails(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderDetails", reflect.TypeOf((*MockOrderDetailsRepository)(nil).DeleteOrderDetails), ctx, params)
}
