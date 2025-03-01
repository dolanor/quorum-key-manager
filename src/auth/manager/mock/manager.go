// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	types "github.com/consensys/quorum-key-manager/src/auth/types"
	gomock "github.com/golang/mock/gomock"
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

// Group mocks base method.
func (m *MockManager) Group(ctx context.Context, name string) (*types.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", ctx, name)
	ret0, _ := ret[0].(*types.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockManagerMockRecorder) Group(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockManager)(nil).Group), ctx, name)
}

// Groups mocks base method.
func (m *MockManager) Groups(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Groups", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Groups indicates an expected call of Groups.
func (mr *MockManagerMockRecorder) Groups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Groups", reflect.TypeOf((*MockManager)(nil).Groups), arg0)
}

// Policies mocks base method.
func (m *MockManager) Policies(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policies", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Policies indicates an expected call of Policies.
func (mr *MockManagerMockRecorder) Policies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policies", reflect.TypeOf((*MockManager)(nil).Policies), arg0)
}

// Policy mocks base method.
func (m *MockManager) Policy(ctx context.Context, name string) (*types.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policy", ctx, name)
	ret0, _ := ret[0].(*types.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Policy indicates an expected call of Policy.
func (mr *MockManagerMockRecorder) Policy(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policy", reflect.TypeOf((*MockManager)(nil).Policy), ctx, name)
}
