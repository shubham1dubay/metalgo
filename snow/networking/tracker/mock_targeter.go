// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shubham1dubay/metalgo/snow/networking/tracker (interfaces: Targeter)
//
// Generated by this command:
//
//	mockgen -package=tracker -destination=snow/networking/tracker/mock_targeter.go github.com/shubham1dubay/metalgo/snow/networking/tracker Targeter
//

// Package tracker is a generated GoMock package.
package tracker

import (
	reflect "reflect"

	ids "github.com/shubham1dubay/metalgo/ids"
	gomock "go.uber.org/mock/gomock"
)

// MockTargeter is a mock of Targeter interface.
type MockTargeter struct {
	ctrl     *gomock.Controller
	recorder *MockTargeterMockRecorder
}

// MockTargeterMockRecorder is the mock recorder for MockTargeter.
type MockTargeterMockRecorder struct {
	mock *MockTargeter
}

// NewMockTargeter creates a new mock instance.
func NewMockTargeter(ctrl *gomock.Controller) *MockTargeter {
	mock := &MockTargeter{ctrl: ctrl}
	mock.recorder = &MockTargeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargeter) EXPECT() *MockTargeterMockRecorder {
	return m.recorder
}

// TargetUsage mocks base method.
func (m *MockTargeter) TargetUsage(arg0 ids.NodeID) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetUsage", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// TargetUsage indicates an expected call of TargetUsage.
func (mr *MockTargeterMockRecorder) TargetUsage(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetUsage", reflect.TypeOf((*MockTargeter)(nil).TargetUsage), arg0)
}
