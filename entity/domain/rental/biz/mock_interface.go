// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	agg "github.com/blackhorseya/sion/entity/domain/rental/agg"
	model "github.com/blackhorseya/sion/entity/domain/rental/model"
	contextx "github.com/blackhorseya/sion/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIRentalBiz is a mock of IRentalBiz interface.
type MockIRentalBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIRentalBizMockRecorder
}

// MockIRentalBizMockRecorder is the mock recorder for MockIRentalBiz.
type MockIRentalBizMockRecorder struct {
	mock *MockIRentalBiz
}

// NewMockIRentalBiz creates a new mock instance.
func NewMockIRentalBiz(ctrl *gomock.Controller) *MockIRentalBiz {
	mock := &MockIRentalBiz{ctrl: ctrl}
	mock.recorder = &MockIRentalBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRentalBiz) EXPECT() *MockIRentalBizMockRecorder {
	return m.recorder
}

// ListByLocation mocks base method.
func (m *MockIRentalBiz) ListByLocation(ctx contextx.Contextx, location *model.Location, opts ListByLocationOptions) ([]*agg.Asset, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByLocation", ctx, location, opts)
	ret0, _ := ret[0].([]*agg.Asset)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByLocation indicates an expected call of ListByLocation.
func (mr *MockIRentalBizMockRecorder) ListByLocation(ctx, location, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByLocation", reflect.TypeOf((*MockIRentalBiz)(nil).ListByLocation), ctx, location, opts)
}
