// Code generated by MockGen. DO NOT EDIT.
// Source: secrets.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entities "github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"
	secrets "github.com/ConsenSysQuorum/quorum-key-manager/src/store/secrets"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockStore) Info(arg0 context.Context) (*entities.StoreInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(*entities.StoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockStoreMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockStore)(nil).Info), arg0)
}

// Set mocks base method
func (m *MockStore) Set(ctx context.Context, id, value string, attr *entities.Attributes) (*entities.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, id, value, attr)
	ret0, _ := ret[0].(*entities.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockStoreMockRecorder) Set(ctx, id, value, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStore)(nil).Set), ctx, id, value, attr)
}

// Get mocks base method
func (m *MockStore) Get(ctx context.Context, id string, version int) (*entities.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id, version)
	ret0, _ := ret[0].(*entities.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(ctx, id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx, id, version)
}

// List mocks base method
func (m *MockStore) List(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockStoreMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), ctx)
}

// Refresh mocks base method
func (m *MockStore) Refresh(ctx context.Context, id string, expirationDate time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ctx, id, expirationDate)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockStoreMockRecorder) Refresh(ctx, id, expirationDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockStore)(nil).Refresh), ctx, id, expirationDate)
}

// Delete mocks base method
func (m *MockStore) Delete(ctx context.Context, id string, versions ...int) (*entities.Secret, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range versions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*entities.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(ctx, id interface{}, versions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, versions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), varargs...)
}

// GetDeleted mocks base method
func (m *MockStore) GetDeleted(ctx context.Context, id string) (*entities.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeleted", ctx, id)
	ret0, _ := ret[0].(*entities.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeleted indicates an expected call of GetDeleted
func (mr *MockStoreMockRecorder) GetDeleted(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeleted", reflect.TypeOf((*MockStore)(nil).GetDeleted), ctx, id)
}

// ListDeleted mocks base method
func (m *MockStore) ListDeleted(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeleted", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeleted indicates an expected call of ListDeleted
func (mr *MockStoreMockRecorder) ListDeleted(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeleted", reflect.TypeOf((*MockStore)(nil).ListDeleted), ctx)
}

// Undelete mocks base method
func (m *MockStore) Undelete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Undelete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Undelete indicates an expected call of Undelete
func (mr *MockStoreMockRecorder) Undelete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Undelete", reflect.TypeOf((*MockStore)(nil).Undelete), ctx, id)
}

// Destroy mocks base method
func (m *MockStore) Destroy(ctx context.Context, id string, versions ...int) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range versions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Destroy", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockStoreMockRecorder) Destroy(ctx, id interface{}, versions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, versions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockStore)(nil).Destroy), varargs...)
}

// MockInstrument is a mock of Instrument interface
type MockInstrument struct {
	ctrl     *gomock.Controller
	recorder *MockInstrumentMockRecorder
}

// MockInstrumentMockRecorder is the mock recorder for MockInstrument
type MockInstrumentMockRecorder struct {
	mock *MockInstrument
}

// NewMockInstrument creates a new mock instance
func NewMockInstrument(ctrl *gomock.Controller) *MockInstrument {
	mock := &MockInstrument{ctrl: ctrl}
	mock.recorder = &MockInstrumentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstrument) EXPECT() *MockInstrumentMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockInstrument) Apply(arg0 secrets.Store) secrets.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0)
	ret0, _ := ret[0].(secrets.Store)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockInstrumentMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockInstrument)(nil).Apply), arg0)
}
