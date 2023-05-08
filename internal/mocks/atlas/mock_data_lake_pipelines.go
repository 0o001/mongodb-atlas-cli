// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store/atlas (interfaces: PipelinesLister,PipelinesDescriber,PipelinesCreator,PipelinesUpdater,PipelinesDeleter,PipelineAvailableSnapshotsLister,PipelineAvailableSchedulesLister,PipelinesTriggerer,PipelinesPauser,PipelinesResumer)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockPipelinesLister is a mock of PipelinesLister interface.
type MockPipelinesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesListerMockRecorder
}

// MockPipelinesListerMockRecorder is the mock recorder for MockPipelinesLister.
type MockPipelinesListerMockRecorder struct {
	mock *MockPipelinesLister
}

// NewMockPipelinesLister creates a new mock instance.
func NewMockPipelinesLister(ctrl *gomock.Controller) *MockPipelinesLister {
	mock := &MockPipelinesLister{ctrl: ctrl}
	mock.recorder = &MockPipelinesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesLister) EXPECT() *MockPipelinesListerMockRecorder {
	return m.recorder
}

// Pipelines mocks base method.
func (m *MockPipelinesLister) Pipelines(arg0 string) ([]mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipelines", arg0)
	ret0, _ := ret[0].([]mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipelines indicates an expected call of Pipelines.
func (mr *MockPipelinesListerMockRecorder) Pipelines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipelines", reflect.TypeOf((*MockPipelinesLister)(nil).Pipelines), arg0)
}

// MockPipelinesDescriber is a mock of PipelinesDescriber interface.
type MockPipelinesDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesDescriberMockRecorder
}

// MockPipelinesDescriberMockRecorder is the mock recorder for MockPipelinesDescriber.
type MockPipelinesDescriberMockRecorder struct {
	mock *MockPipelinesDescriber
}

// NewMockPipelinesDescriber creates a new mock instance.
func NewMockPipelinesDescriber(ctrl *gomock.Controller) *MockPipelinesDescriber {
	mock := &MockPipelinesDescriber{ctrl: ctrl}
	mock.recorder = &MockPipelinesDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesDescriber) EXPECT() *MockPipelinesDescriberMockRecorder {
	return m.recorder
}

// Pipeline mocks base method.
func (m *MockPipelinesDescriber) Pipeline(arg0, arg1 string) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipeline", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipeline indicates an expected call of Pipeline.
func (mr *MockPipelinesDescriberMockRecorder) Pipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipeline", reflect.TypeOf((*MockPipelinesDescriber)(nil).Pipeline), arg0, arg1)
}

// MockPipelinesCreator is a mock of PipelinesCreator interface.
type MockPipelinesCreator struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesCreatorMockRecorder
}

// MockPipelinesCreatorMockRecorder is the mock recorder for MockPipelinesCreator.
type MockPipelinesCreatorMockRecorder struct {
	mock *MockPipelinesCreator
}

// NewMockPipelinesCreator creates a new mock instance.
func NewMockPipelinesCreator(ctrl *gomock.Controller) *MockPipelinesCreator {
	mock := &MockPipelinesCreator{ctrl: ctrl}
	mock.recorder = &MockPipelinesCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesCreator) EXPECT() *MockPipelinesCreatorMockRecorder {
	return m.recorder
}

// CreatePipeline mocks base method.
func (m *MockPipelinesCreator) CreatePipeline(arg0 string, arg1 mongodbatlasv2.IngestionPipeline) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockPipelinesCreatorMockRecorder) CreatePipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockPipelinesCreator)(nil).CreatePipeline), arg0, arg1)
}

// MockPipelinesUpdater is a mock of PipelinesUpdater interface.
type MockPipelinesUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesUpdaterMockRecorder
}

// MockPipelinesUpdaterMockRecorder is the mock recorder for MockPipelinesUpdater.
type MockPipelinesUpdaterMockRecorder struct {
	mock *MockPipelinesUpdater
}

// NewMockPipelinesUpdater creates a new mock instance.
func NewMockPipelinesUpdater(ctrl *gomock.Controller) *MockPipelinesUpdater {
	mock := &MockPipelinesUpdater{ctrl: ctrl}
	mock.recorder = &MockPipelinesUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesUpdater) EXPECT() *MockPipelinesUpdaterMockRecorder {
	return m.recorder
}

// UpdatePipeline mocks base method.
func (m *MockPipelinesUpdater) UpdatePipeline(arg0, arg1 string, arg2 mongodbatlasv2.IngestionPipeline) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipeline indicates an expected call of UpdatePipeline.
func (mr *MockPipelinesUpdaterMockRecorder) UpdatePipeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockPipelinesUpdater)(nil).UpdatePipeline), arg0, arg1, arg2)
}

// MockPipelinesDeleter is a mock of PipelinesDeleter interface.
type MockPipelinesDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesDeleterMockRecorder
}

// MockPipelinesDeleterMockRecorder is the mock recorder for MockPipelinesDeleter.
type MockPipelinesDeleterMockRecorder struct {
	mock *MockPipelinesDeleter
}

// NewMockPipelinesDeleter creates a new mock instance.
func NewMockPipelinesDeleter(ctrl *gomock.Controller) *MockPipelinesDeleter {
	mock := &MockPipelinesDeleter{ctrl: ctrl}
	mock.recorder = &MockPipelinesDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesDeleter) EXPECT() *MockPipelinesDeleterMockRecorder {
	return m.recorder
}

// DeletePipeline mocks base method.
func (m *MockPipelinesDeleter) DeletePipeline(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockPipelinesDeleterMockRecorder) DeletePipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockPipelinesDeleter)(nil).DeletePipeline), arg0, arg1)
}

// MockPipelineAvailableSnapshotsLister is a mock of PipelineAvailableSnapshotsLister interface.
type MockPipelineAvailableSnapshotsLister struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineAvailableSnapshotsListerMockRecorder
}

// MockPipelineAvailableSnapshotsListerMockRecorder is the mock recorder for MockPipelineAvailableSnapshotsLister.
type MockPipelineAvailableSnapshotsListerMockRecorder struct {
	mock *MockPipelineAvailableSnapshotsLister
}

// NewMockPipelineAvailableSnapshotsLister creates a new mock instance.
func NewMockPipelineAvailableSnapshotsLister(ctrl *gomock.Controller) *MockPipelineAvailableSnapshotsLister {
	mock := &MockPipelineAvailableSnapshotsLister{ctrl: ctrl}
	mock.recorder = &MockPipelineAvailableSnapshotsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineAvailableSnapshotsLister) EXPECT() *MockPipelineAvailableSnapshotsListerMockRecorder {
	return m.recorder
}

// PipelineAvailableSnapshots mocks base method.
func (m *MockPipelineAvailableSnapshotsLister) PipelineAvailableSnapshots(arg0, arg1 string, arg2 *time.Time, arg3 *mongodbatlas.ListOptions) (*mongodbatlasv2.PaginatedBackupSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineAvailableSnapshots", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mongodbatlasv2.PaginatedBackupSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineAvailableSnapshots indicates an expected call of PipelineAvailableSnapshots.
func (mr *MockPipelineAvailableSnapshotsListerMockRecorder) PipelineAvailableSnapshots(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineAvailableSnapshots", reflect.TypeOf((*MockPipelineAvailableSnapshotsLister)(nil).PipelineAvailableSnapshots), arg0, arg1, arg2, arg3)
}

// MockPipelineAvailableSchedulesLister is a mock of PipelineAvailableSchedulesLister interface.
type MockPipelineAvailableSchedulesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineAvailableSchedulesListerMockRecorder
}

// MockPipelineAvailableSchedulesListerMockRecorder is the mock recorder for MockPipelineAvailableSchedulesLister.
type MockPipelineAvailableSchedulesListerMockRecorder struct {
	mock *MockPipelineAvailableSchedulesLister
}

// NewMockPipelineAvailableSchedulesLister creates a new mock instance.
func NewMockPipelineAvailableSchedulesLister(ctrl *gomock.Controller) *MockPipelineAvailableSchedulesLister {
	mock := &MockPipelineAvailableSchedulesLister{ctrl: ctrl}
	mock.recorder = &MockPipelineAvailableSchedulesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineAvailableSchedulesLister) EXPECT() *MockPipelineAvailableSchedulesListerMockRecorder {
	return m.recorder
}

// PipelineAvailableSchedules mocks base method.
func (m *MockPipelineAvailableSchedulesLister) PipelineAvailableSchedules(arg0, arg1 string) ([]mongodbatlasv2.PolicyItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineAvailableSchedules", arg0, arg1)
	ret0, _ := ret[0].([]mongodbatlasv2.PolicyItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineAvailableSchedules indicates an expected call of PipelineAvailableSchedules.
func (mr *MockPipelineAvailableSchedulesListerMockRecorder) PipelineAvailableSchedules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineAvailableSchedules", reflect.TypeOf((*MockPipelineAvailableSchedulesLister)(nil).PipelineAvailableSchedules), arg0, arg1)
}

// MockPipelinesTriggerer is a mock of PipelinesTriggerer interface.
type MockPipelinesTriggerer struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesTriggererMockRecorder
}

// MockPipelinesTriggererMockRecorder is the mock recorder for MockPipelinesTriggerer.
type MockPipelinesTriggererMockRecorder struct {
	mock *MockPipelinesTriggerer
}

// NewMockPipelinesTriggerer creates a new mock instance.
func NewMockPipelinesTriggerer(ctrl *gomock.Controller) *MockPipelinesTriggerer {
	mock := &MockPipelinesTriggerer{ctrl: ctrl}
	mock.recorder = &MockPipelinesTriggererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesTriggerer) EXPECT() *MockPipelinesTriggererMockRecorder {
	return m.recorder
}

// PipelineTrigger mocks base method.
func (m *MockPipelinesTriggerer) PipelineTrigger(arg0, arg1 string) (*mongodbatlasv2.IngestionPipelineRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineTrigger", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipelineRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineTrigger indicates an expected call of PipelineTrigger.
func (mr *MockPipelinesTriggererMockRecorder) PipelineTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineTrigger", reflect.TypeOf((*MockPipelinesTriggerer)(nil).PipelineTrigger), arg0, arg1)
}

// MockPipelinesPauser is a mock of PipelinesPauser interface.
type MockPipelinesPauser struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesPauserMockRecorder
}

// MockPipelinesPauserMockRecorder is the mock recorder for MockPipelinesPauser.
type MockPipelinesPauserMockRecorder struct {
	mock *MockPipelinesPauser
}

// NewMockPipelinesPauser creates a new mock instance.
func NewMockPipelinesPauser(ctrl *gomock.Controller) *MockPipelinesPauser {
	mock := &MockPipelinesPauser{ctrl: ctrl}
	mock.recorder = &MockPipelinesPauserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesPauser) EXPECT() *MockPipelinesPauserMockRecorder {
	return m.recorder
}

// PipelinePause mocks base method.
func (m *MockPipelinesPauser) PipelinePause(arg0, arg1 string) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelinePause", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelinePause indicates an expected call of PipelinePause.
func (mr *MockPipelinesPauserMockRecorder) PipelinePause(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelinePause", reflect.TypeOf((*MockPipelinesPauser)(nil).PipelinePause), arg0, arg1)
}

// MockPipelinesResumer is a mock of PipelinesResumer interface.
type MockPipelinesResumer struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesResumerMockRecorder
}

// MockPipelinesResumerMockRecorder is the mock recorder for MockPipelinesResumer.
type MockPipelinesResumerMockRecorder struct {
	mock *MockPipelinesResumer
}

// NewMockPipelinesResumer creates a new mock instance.
func NewMockPipelinesResumer(ctrl *gomock.Controller) *MockPipelinesResumer {
	mock := &MockPipelinesResumer{ctrl: ctrl}
	mock.recorder = &MockPipelinesResumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesResumer) EXPECT() *MockPipelinesResumerMockRecorder {
	return m.recorder
}

// PipelineResume mocks base method.
func (m *MockPipelinesResumer) PipelineResume(arg0, arg1 string) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineResume", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineResume indicates an expected call of PipelineResume.
func (mr *MockPipelinesResumerMockRecorder) PipelineResume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineResume", reflect.TypeOf((*MockPipelinesResumer)(nil).PipelineResume), arg0, arg1)
}
