// Code generated by MockGen. DO NOT EDIT.
// Source: shim.go

// Package mock_host_query_reset is a generated GoMock package.
package mock_host_query_reset

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockShim is a mock of Shim interface.
type MockShim struct {
	ctrl     *gomock.Controller
	recorder *MockShimMockRecorder
}

// MockShimMockRecorder is the mock recorder for MockShim.
type MockShimMockRecorder struct {
	mock *MockShim
}

// NewMockShim creates a new mock instance.
func NewMockShim(ctrl *gomock.Controller) *MockShim {
	mock := &MockShim{ctrl: ctrl}
	mock.recorder = &MockShimMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShim) EXPECT() *MockShimMockRecorder {
	return m.recorder
}

// Reset mocks base method.
func (m *MockShim) Reset(firstQuery, queryCount int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", firstQuery, queryCount)
}

// Reset indicates an expected call of Reset.
func (mr *MockShimMockRecorder) Reset(firstQuery, queryCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockShim)(nil).Reset), firstQuery, queryCount)
}
