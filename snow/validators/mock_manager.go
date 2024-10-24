// Code generated by MockGen. DO NOT EDIT.
// Source: snow/validators/manager.go
//
// Generated by this command:
//
//	mockgen -source=snow/validators/manager.go -destination=snow/validators/mock_manager.go -package=validators -exclude_interfaces=SetCallbackListener
//

// Package validators is a generated GoMock package.
package validators

import (
	reflect "reflect"

	ids "github.com/shubham1dubay/metalgo/ids"
	bls "github.com/shubham1dubay/metalgo/utils/crypto/bls"
	set "github.com/shubham1dubay/metalgo/utils/set"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddStaker mocks base method.
func (m *MockManager) AddStaker(subnetID ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStaker", subnetID, nodeID, pk, txID, weight)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStaker indicates an expected call of AddStaker.
func (mr *MockManagerMockRecorder) AddStaker(subnetID, nodeID, pk, txID, weight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStaker", reflect.TypeOf((*MockManager)(nil).AddStaker), subnetID, nodeID, pk, txID, weight)
}

// AddWeight mocks base method.
func (m *MockManager) AddWeight(subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWeight", subnetID, nodeID, weight)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWeight indicates an expected call of AddWeight.
func (mr *MockManagerMockRecorder) AddWeight(subnetID, nodeID, weight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWeight", reflect.TypeOf((*MockManager)(nil).AddWeight), subnetID, nodeID, weight)
}

// Count mocks base method.
func (m *MockManager) Count(subnetID ids.ID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", subnetID)
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockManagerMockRecorder) Count(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockManager)(nil).Count), subnetID)
}

// GetMap mocks base method.
func (m *MockManager) GetMap(subnetID ids.ID) map[ids.NodeID]*GetValidatorOutput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMap", subnetID)
	ret0, _ := ret[0].(map[ids.NodeID]*GetValidatorOutput)
	return ret0
}

// GetMap indicates an expected call of GetMap.
func (mr *MockManagerMockRecorder) GetMap(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMap", reflect.TypeOf((*MockManager)(nil).GetMap), subnetID)
}

// GetValidator mocks base method.
func (m *MockManager) GetValidator(subnetID ids.ID, nodeID ids.NodeID) (*Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", subnetID, nodeID)
	ret0, _ := ret[0].(*Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockManagerMockRecorder) GetValidator(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockManager)(nil).GetValidator), subnetID, nodeID)
}

// GetValidatorIDs mocks base method.
func (m *MockManager) GetValidatorIDs(subnetID ids.ID) []ids.NodeID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorIDs", subnetID)
	ret0, _ := ret[0].([]ids.NodeID)
	return ret0
}

// GetValidatorIDs indicates an expected call of GetValidatorIDs.
func (mr *MockManagerMockRecorder) GetValidatorIDs(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorIDs", reflect.TypeOf((*MockManager)(nil).GetValidatorIDs), subnetID)
}

// GetWeight mocks base method.
func (m *MockManager) GetWeight(subnetID ids.ID, nodeID ids.NodeID) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeight", subnetID, nodeID)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetWeight indicates an expected call of GetWeight.
func (mr *MockManagerMockRecorder) GetWeight(subnetID, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeight", reflect.TypeOf((*MockManager)(nil).GetWeight), subnetID, nodeID)
}

// RegisterCallbackListener mocks base method.
func (m *MockManager) RegisterCallbackListener(subnetID ids.ID, listener SetCallbackListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterCallbackListener", subnetID, listener)
}

// RegisterCallbackListener indicates an expected call of RegisterCallbackListener.
func (mr *MockManagerMockRecorder) RegisterCallbackListener(subnetID, listener any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCallbackListener", reflect.TypeOf((*MockManager)(nil).RegisterCallbackListener), subnetID, listener)
}

// RemoveWeight mocks base method.
func (m *MockManager) RemoveWeight(subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWeight", subnetID, nodeID, weight)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWeight indicates an expected call of RemoveWeight.
func (mr *MockManagerMockRecorder) RemoveWeight(subnetID, nodeID, weight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWeight", reflect.TypeOf((*MockManager)(nil).RemoveWeight), subnetID, nodeID, weight)
}

// Sample mocks base method.
func (m *MockManager) Sample(subnetID ids.ID, size int) ([]ids.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sample", subnetID, size)
	ret0, _ := ret[0].([]ids.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sample indicates an expected call of Sample.
func (mr *MockManagerMockRecorder) Sample(subnetID, size any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sample", reflect.TypeOf((*MockManager)(nil).Sample), subnetID, size)
}

// String mocks base method.
func (m *MockManager) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockManagerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockManager)(nil).String))
}

// SubsetWeight mocks base method.
func (m *MockManager) SubsetWeight(subnetID ids.ID, validatorIDs set.Set[ids.NodeID]) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubsetWeight", subnetID, validatorIDs)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubsetWeight indicates an expected call of SubsetWeight.
func (mr *MockManagerMockRecorder) SubsetWeight(subnetID, validatorIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubsetWeight", reflect.TypeOf((*MockManager)(nil).SubsetWeight), subnetID, validatorIDs)
}

// TotalWeight mocks base method.
func (m *MockManager) TotalWeight(subnetID ids.ID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalWeight", subnetID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalWeight indicates an expected call of TotalWeight.
func (mr *MockManagerMockRecorder) TotalWeight(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalWeight", reflect.TypeOf((*MockManager)(nil).TotalWeight), subnetID)
}
