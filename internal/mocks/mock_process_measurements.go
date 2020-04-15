// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/process_measurements.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	reflect "reflect"
)

// MockProcessMeasurementLister is a mock of ProcessMeasurementLister interface
type MockProcessMeasurementLister struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMeasurementListerMockRecorder
}

// MockProcessMeasurementListerMockRecorder is the mock recorder for MockProcessMeasurementLister
type MockProcessMeasurementListerMockRecorder struct {
	mock *MockProcessMeasurementLister
}

// NewMockProcessMeasurementLister creates a new mock instance
func NewMockProcessMeasurementLister(ctrl *gomock.Controller) *MockProcessMeasurementLister {
	mock := &MockProcessMeasurementLister{ctrl: ctrl}
	mock.recorder = &MockProcessMeasurementListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessMeasurementLister) EXPECT() *MockProcessMeasurementListerMockRecorder {
	return m.recorder
}

// ProcessMeasurements mocks base method
func (m *MockProcessMeasurementLister) ProcessMeasurements(arg0, arg1 string, arg2 int, arg3 *mongodbatlas.ProcessMeasurementListOptions) (*mongodbatlas.ProcessMeasurements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMeasurements", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mongodbatlas.ProcessMeasurements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessMeasurements indicates an expected call of ProcessMeasurements
func (mr *MockProcessMeasurementListerMockRecorder) ProcessMeasurements(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMeasurements", reflect.TypeOf((*MockProcessMeasurementLister)(nil).ProcessMeasurements), arg0, arg1, arg2, arg3)
}
