// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/model (interfaces: ShowGenerationCommandAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/juju/juju/core/model"
	reflect "reflect"
)

// MockShowGenerationCommandAPI is a mock of ShowGenerationCommandAPI interface
type MockShowGenerationCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockShowGenerationCommandAPIMockRecorder
}

// MockShowGenerationCommandAPIMockRecorder is the mock recorder for MockShowGenerationCommandAPI
type MockShowGenerationCommandAPIMockRecorder struct {
	mock *MockShowGenerationCommandAPI
}

// NewMockShowGenerationCommandAPI creates a new mock instance
func NewMockShowGenerationCommandAPI(ctrl *gomock.Controller) *MockShowGenerationCommandAPI {
	mock := &MockShowGenerationCommandAPI{ctrl: ctrl}
	mock.recorder = &MockShowGenerationCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShowGenerationCommandAPI) EXPECT() *MockShowGenerationCommandAPIMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockShowGenerationCommandAPI) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockShowGenerationCommandAPIMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockShowGenerationCommandAPI)(nil).Close))
}

// GenerationInfo mocks base method
func (m *MockShowGenerationCommandAPI) GenerationInfo(arg0 string) (map[model.GenerationVersion]model.Generation, error) {
	ret := m.ctrl.Call(m, "GenerationInfo", arg0)
	ret0, _ := ret[0].(map[model.GenerationVersion]model.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerationInfo indicates an expected call of GenerationInfo
func (mr *MockShowGenerationCommandAPIMockRecorder) GenerationInfo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerationInfo", reflect.TypeOf((*MockShowGenerationCommandAPI)(nil).GenerationInfo), arg0)
}
