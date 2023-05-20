// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	contextx "github.com/blackhorseya/irent/pkg/contextx"
	model "github.com/blackhorseya/irent/pkg/entity/domain/account/model"
	model0 "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	model1 "github.com/blackhorseya/irent/pkg/entity/domain/rental/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIBiz is a mock of IBiz interface.
type MockIBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIBizMockRecorder
}

// MockIBizMockRecorder is the mock recorder for MockIBiz.
type MockIBizMockRecorder struct {
	mock *MockIBiz
}

// NewMockIBiz creates a new mock instance.
func NewMockIBiz(ctrl *gomock.Controller) *MockIBiz {
	mock := &MockIBiz{ctrl: ctrl}
	mock.recorder = &MockIBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBiz) EXPECT() *MockIBizMockRecorder {
	return m.recorder
}

// BookRental mocks base method.
func (m *MockIBiz) BookRental(ctx contextx.Contextx, from *model.Profile, target *model1.Car) (*model0.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookRental", ctx, from, target)
	ret0, _ := ret[0].(*model0.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookRental indicates an expected call of BookRental.
func (mr *MockIBizMockRecorder) BookRental(ctx, from, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookRental", reflect.TypeOf((*MockIBiz)(nil).BookRental), ctx, from, target)
}

// CancelLease mocks base method.
func (m *MockIBiz) CancelLease(ctx contextx.Contextx, from *model.Profile, target *model0.Lease) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLease", ctx, from, target)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelLease indicates an expected call of CancelLease.
func (mr *MockIBizMockRecorder) CancelLease(ctx, from, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLease", reflect.TypeOf((*MockIBiz)(nil).CancelLease), ctx, from, target)
}

// GetArrears mocks base method.
func (m *MockIBiz) GetArrears(ctx contextx.Contextx, from, target *model.Profile) (*model0.Arrears, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArrears", ctx, from, target)
	ret0, _ := ret[0].(*model0.Arrears)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArrears indicates an expected call of GetArrears.
func (mr *MockIBizMockRecorder) GetArrears(ctx, from, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArrears", reflect.TypeOf((*MockIBiz)(nil).GetArrears), ctx, from, target)
}

// ListLease mocks base method.
func (m *MockIBiz) ListLease(ctx contextx.Contextx, from *model.Profile) ([]*model0.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLease", ctx, from)
	ret0, _ := ret[0].([]*model0.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLease indicates an expected call of ListLease.
func (mr *MockIBizMockRecorder) ListLease(ctx, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLease", reflect.TypeOf((*MockIBiz)(nil).ListLease), ctx, from)
}

// Liveness mocks base method.
func (m *MockIBiz) Liveness(ctx contextx.Contextx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Liveness", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Liveness indicates an expected call of Liveness.
func (mr *MockIBizMockRecorder) Liveness(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Liveness", reflect.TypeOf((*MockIBiz)(nil).Liveness), ctx)
}

// Readiness mocks base method.
func (m *MockIBiz) Readiness(ctx contextx.Contextx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readiness", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Readiness indicates an expected call of Readiness.
func (mr *MockIBizMockRecorder) Readiness(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readiness", reflect.TypeOf((*MockIBiz)(nil).Readiness), ctx)
}
