// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/item_units_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "git.paylabo.com/c002/harp/backend-purchase/domain/model"
	repository "git.paylabo.com/c002/harp/backend-purchase/domain/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockItemUnitsRepository is a mock of ItemUnitsRepository interface
type MockItemUnitsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemUnitsRepositoryMockRecorder
}

// MockItemUnitsRepositoryMockRecorder is the mock recorder for MockItemUnitsRepository
type MockItemUnitsRepositoryMockRecorder struct {
	mock *MockItemUnitsRepository
}

// NewMockItemUnitsRepository creates a new mock instance
func NewMockItemUnitsRepository(ctrl *gomock.Controller) *MockItemUnitsRepository {
	mock := &MockItemUnitsRepository{ctrl: ctrl}
	mock.recorder = &MockItemUnitsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemUnitsRepository) EXPECT() *MockItemUnitsRepositoryMockRecorder {
	return m.recorder
}

// GetItemUnit mocks base method
func (m *MockItemUnitsRepository) GetItemUnit(ctx context.Context, param repository.GetItemUnitParams) (*model.ItemUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemUnit", ctx, param)
	ret0, _ := ret[0].(*model.ItemUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemUnit indicates an expected call of GetItemUnit
func (mr *MockItemUnitsRepositoryMockRecorder) GetItemUnit(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemUnit", reflect.TypeOf((*MockItemUnitsRepository)(nil).GetItemUnit), ctx, param)
}

// ListItemUnits mocks base method
func (m *MockItemUnitsRepository) ListItemUnits(ctx context.Context, param repository.ListItemUnitsParams) ([]*model.ItemUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItemUnits", ctx, param)
	ret0, _ := ret[0].([]*model.ItemUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItemUnits indicates an expected call of ListItemUnits
func (mr *MockItemUnitsRepositoryMockRecorder) ListItemUnits(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemUnits", reflect.TypeOf((*MockItemUnitsRepository)(nil).ListItemUnits), ctx, param)
}
