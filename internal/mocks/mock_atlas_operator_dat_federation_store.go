// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: AtlasOperatorDataFederationStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

// MockAtlasOperatorDataFederationStore is a mock of AtlasOperatorDataFederationStore interface.
type MockAtlasOperatorDataFederationStore struct {
	ctrl     *gomock.Controller
	recorder *MockAtlasOperatorDataFederationStoreMockRecorder
}

// MockAtlasOperatorDataFederationStoreMockRecorder is the mock recorder for MockAtlasOperatorDataFederationStore.
type MockAtlasOperatorDataFederationStoreMockRecorder struct {
	mock *MockAtlasOperatorDataFederationStore
}

// NewMockAtlasOperatorDataFederationStore creates a new mock instance.
func NewMockAtlasOperatorDataFederationStore(ctrl *gomock.Controller) *MockAtlasOperatorDataFederationStore {
	mock := &MockAtlasOperatorDataFederationStore{ctrl: ctrl}
	mock.recorder = &MockAtlasOperatorDataFederationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtlasOperatorDataFederationStore) EXPECT() *MockAtlasOperatorDataFederationStoreMockRecorder {
	return m.recorder
}

// CreateDataFederation mocks base method.
func (m *MockAtlasOperatorDataFederationStore) CreateDataFederation(arg0 string, arg1 *admin.DataLakeTenant) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataFederation indicates an expected call of CreateDataFederation.
func (mr *MockAtlasOperatorDataFederationStoreMockRecorder) CreateDataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataFederation", reflect.TypeOf((*MockAtlasOperatorDataFederationStore)(nil).CreateDataFederation), arg0, arg1)
}

// DataFederation mocks base method.
func (m *MockAtlasOperatorDataFederationStore) DataFederation(arg0, arg1 string) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederation", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederation indicates an expected call of DataFederation.
func (mr *MockAtlasOperatorDataFederationStoreMockRecorder) DataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederation", reflect.TypeOf((*MockAtlasOperatorDataFederationStore)(nil).DataFederation), arg0, arg1)
}

// DataFederationList mocks base method.
func (m *MockAtlasOperatorDataFederationStore) DataFederationList(arg0, arg1 string) ([]admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationList", arg0, arg1)
	ret0, _ := ret[0].([]admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationList indicates an expected call of DataFederationList.
func (mr *MockAtlasOperatorDataFederationStoreMockRecorder) DataFederationList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationList", reflect.TypeOf((*MockAtlasOperatorDataFederationStore)(nil).DataFederationList), arg0, arg1)
}

// DataFederationLogs mocks base method.
func (m *MockAtlasOperatorDataFederationStore) DataFederationLogs(arg0, arg1 string, arg2, arg3 int64) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataFederationLogs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataFederationLogs indicates an expected call of DataFederationLogs.
func (mr *MockAtlasOperatorDataFederationStoreMockRecorder) DataFederationLogs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFederationLogs", reflect.TypeOf((*MockAtlasOperatorDataFederationStore)(nil).DataFederationLogs), arg0, arg1, arg2, arg3)
}

// DeleteDataFederation mocks base method.
func (m *MockAtlasOperatorDataFederationStore) DeleteDataFederation(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataFederation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataFederation indicates an expected call of DeleteDataFederation.
func (mr *MockAtlasOperatorDataFederationStoreMockRecorder) DeleteDataFederation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataFederation", reflect.TypeOf((*MockAtlasOperatorDataFederationStore)(nil).DeleteDataFederation), arg0, arg1)
}

// UpdateDataFederation mocks base method.
func (m *MockAtlasOperatorDataFederationStore) UpdateDataFederation(arg0, arg1 string, arg2 *admin.DataLakeTenant) (*admin.DataLakeTenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDataFederation", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.DataLakeTenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDataFederation indicates an expected call of UpdateDataFederation.
func (mr *MockAtlasOperatorDataFederationStoreMockRecorder) UpdateDataFederation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDataFederation", reflect.TypeOf((*MockAtlasOperatorDataFederationStore)(nil).UpdateDataFederation), arg0, arg1, arg2)
}
