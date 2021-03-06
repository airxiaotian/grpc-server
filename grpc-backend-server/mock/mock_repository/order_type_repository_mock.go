// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/order_type_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "git.paylabo.com/c002/harp/backend-purchase/domain/model"
	repository "git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderTypeRepository is a mock of OrderTypeRepository interface
type MockOrderTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderTypeRepositoryMockRecorder
}

// MockOrderTypeRepositoryMockRecorder is the mock recorder for MockOrderTypeRepository
type MockOrderTypeRepositoryMockRecorder struct {
	mock *MockOrderTypeRepository
}

// NewMockOrderTypeRepository creates a new mock instance
func NewMockOrderTypeRepository(ctrl *gomock.Controller) *MockOrderTypeRepository {
	mock := &MockOrderTypeRepository{ctrl: ctrl}
	mock.recorder = &MockOrderTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderTypeRepository) EXPECT() *MockOrderTypeRepositoryMockRecorder {
	return m.recorder
}

// ListOrderTypes mocks base method
func (m *MockOrderTypeRepository) ListOrderTypes(ctx context.Context, params repository.ListOrderTypesParams) ([]*model.OrderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrderTypes", ctx, params)
	ret0, _ := ret[0].([]*model.OrderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrderTypes indicates an expected call of ListOrderTypes
func (mr *MockOrderTypeRepositoryMockRecorder) ListOrderTypes(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrderTypes", reflect.TypeOf((*MockOrderTypeRepository)(nil).ListOrderTypes), ctx, params)
}

// GetOrderType mocks base method
func (m *MockOrderTypeRepository) GetOrderType(ctx context.Context, typeValue string) (*model.OrderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderType", ctx, typeValue)
	ret0, _ := ret[0].(*model.OrderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderType indicates an expected call of GetOrderType
func (mr *MockOrderTypeRepositoryMockRecorder) GetOrderType(ctx, typeValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderType", reflect.TypeOf((*MockOrderTypeRepository)(nil).GetOrderType), ctx, typeValue)
}
