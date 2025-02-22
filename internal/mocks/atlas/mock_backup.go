// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: CompliancePolicyDescriber,CompliancePolicy)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

// MockCompliancePolicyDescriber is a mock of CompliancePolicyDescriber interface.
type MockCompliancePolicyDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyDescriberMockRecorder
}

// MockCompliancePolicyDescriberMockRecorder is the mock recorder for MockCompliancePolicyDescriber.
type MockCompliancePolicyDescriberMockRecorder struct {
	mock *MockCompliancePolicyDescriber
}

// NewMockCompliancePolicyDescriber creates a new mock instance.
func NewMockCompliancePolicyDescriber(ctrl *gomock.Controller) *MockCompliancePolicyDescriber {
	mock := &MockCompliancePolicyDescriber{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicyDescriber) EXPECT() *MockCompliancePolicyDescriberMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicyDescriber) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyDescriberMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicyDescriber)(nil).DescribeCompliancePolicy), arg0)
}

// MockCompliancePolicy is a mock of CompliancePolicy interface.
type MockCompliancePolicy struct {
	ctrl     *gomock.Controller
	recorder *MockCompliancePolicyMockRecorder
}

// MockCompliancePolicyMockRecorder is the mock recorder for MockCompliancePolicy.
type MockCompliancePolicyMockRecorder struct {
	mock *MockCompliancePolicy
}

// NewMockCompliancePolicy creates a new mock instance.
func NewMockCompliancePolicy(ctrl *gomock.Controller) *MockCompliancePolicy {
	mock := &MockCompliancePolicy{ctrl: ctrl}
	mock.recorder = &MockCompliancePolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompliancePolicy) EXPECT() *MockCompliancePolicyMockRecorder {
	return m.recorder
}

// DescribeCompliancePolicy mocks base method.
func (m *MockCompliancePolicy) DescribeCompliancePolicy(arg0 string) (*admin.DataProtectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCompliancePolicy", arg0)
	ret0, _ := ret[0].(*admin.DataProtectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCompliancePolicy indicates an expected call of DescribeCompliancePolicy.
func (mr *MockCompliancePolicyMockRecorder) DescribeCompliancePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCompliancePolicy", reflect.TypeOf((*MockCompliancePolicy)(nil).DescribeCompliancePolicy), arg0)
}

// UpdateCompliancePolicy mocks base method.
func (m *MockCompliancePolicy) UpdateCompliancePolicy(arg0 string, arg1 *admin.DataProtectionSettings) (*admin.DataProtectionSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompliancePolicy", arg0, arg1)
	ret0, _ := ret[0].(*admin.DataProtectionSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompliancePolicy indicates an expected call of UpdateCompliancePolicy.
func (mr *MockCompliancePolicyMockRecorder) UpdateCompliancePolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompliancePolicy", reflect.TypeOf((*MockCompliancePolicy)(nil).UpdateCompliancePolicy), arg0, arg1)
}
