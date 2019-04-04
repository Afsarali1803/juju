// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/model (interfaces: CheckoutCommandAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCheckoutCommandAPI is a mock of CheckoutCommandAPI interface
type MockCheckoutCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCheckoutCommandAPIMockRecorder
}

// MockCheckoutCommandAPIMockRecorder is the mock recorder for MockCheckoutCommandAPI
type MockCheckoutCommandAPIMockRecorder struct {
	mock *MockCheckoutCommandAPI
}

// NewMockCheckoutCommandAPI creates a new mock instance
func NewMockCheckoutCommandAPI(ctrl *gomock.Controller) *MockCheckoutCommandAPI {
	mock := &MockCheckoutCommandAPI{ctrl: ctrl}
	mock.recorder = &MockCheckoutCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckoutCommandAPI) EXPECT() *MockCheckoutCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCheckoutCommandAPI) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCheckoutCommandAPIMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCheckoutCommandAPI)(nil).Close))
}

// HasActiveBranch mocks base method
func (m *MockCheckoutCommandAPI) HasActiveBranch(arg0, arg1 string) (bool, error) {
	ret := m.ctrl.Call(m, "HasActiveBranch", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasActiveBranch indicates an expected call of HasActiveBranch
func (mr *MockCheckoutCommandAPIMockRecorder) HasActiveBranch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasActiveBranch", reflect.TypeOf((*MockCheckoutCommandAPI)(nil).HasActiveBranch), arg0, arg1)
}
