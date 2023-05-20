// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	contextx "github.com/blackhorseya/irent/pkg/contextx"
	model "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIRepo is a mock of IRepo interface.
type MockIRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMockRecorder
}

// MockIRepoMockRecorder is the mock recorder for MockIRepo.
type MockIRepoMockRecorder struct {
	mock *MockIRepo
}

// NewMockIRepo creates a new mock instance.
func NewMockIRepo(ctrl *gomock.Controller) *MockIRepo {
	mock := &MockIRepo{ctrl: ctrl}
	mock.recorder = &MockIRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepo) EXPECT() *MockIRepoMockRecorder {
	return m.recorder
}

// FetchAvailableCars mocks base method.
func (m *MockIRepo) FetchAvailableCars(ctx contextx.Contextx) ([]*model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAvailableCars", ctx)
	ret0, _ := ret[0].([]*model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAvailableCars indicates an expected call of FetchAvailableCars.
func (mr *MockIRepoMockRecorder) FetchAvailableCars(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAvailableCars", reflect.TypeOf((*MockIRepo)(nil).FetchAvailableCars), ctx)
}

// ListCars mocks base method.
func (m *MockIRepo) ListCars(ctx contextx.Contextx) ([]*model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCars", ctx)
	ret0, _ := ret[0].([]*model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCars indicates an expected call of ListCars.
func (mr *MockIRepoMockRecorder) ListCars(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCars", reflect.TypeOf((*MockIRepo)(nil).ListCars), ctx)
}

// UpdateStatusAllCars mocks base method.
func (m *MockIRepo) UpdateStatusAllCars(ctx contextx.Contextx, status model.CarStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusAllCars", ctx, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusAllCars indicates an expected call of UpdateStatusAllCars.
func (mr *MockIRepoMockRecorder) UpdateStatusAllCars(ctx, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusAllCars", reflect.TypeOf((*MockIRepo)(nil).UpdateStatusAllCars), ctx, status)
}

// UpsertStatusCar mocks base method.
func (m *MockIRepo) UpsertStatusCar(ctx contextx.Contextx, target *model.Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertStatusCar", ctx, target)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertStatusCar indicates an expected call of UpsertStatusCar.
func (mr *MockIRepoMockRecorder) UpsertStatusCar(ctx, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStatusCar", reflect.TypeOf((*MockIRepo)(nil).UpsertStatusCar), ctx, target)
}