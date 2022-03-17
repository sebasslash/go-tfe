// Code generated by MockGen. DO NOT EDIT.
// Source: configuration_version.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockConfigurationVersions is a mock of ConfigurationVersions interface.
type MockConfigurationVersions struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationVersionsMockRecorder
}

// MockConfigurationVersionsMockRecorder is the mock recorder for MockConfigurationVersions.
type MockConfigurationVersionsMockRecorder struct {
	mock *MockConfigurationVersions
}

// NewMockConfigurationVersions creates a new mock instance.
func NewMockConfigurationVersions(ctrl *gomock.Controller) *MockConfigurationVersions {
	mock := &MockConfigurationVersions{ctrl: ctrl}
	mock.recorder = &MockConfigurationVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationVersions) EXPECT() *MockConfigurationVersionsMockRecorder {
	return m.recorder
}

// Archive mocks base method.
func (m *MockConfigurationVersions) Archive(ctx context.Context, cvID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Archive", ctx, cvID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Archive indicates an expected call of Archive.
func (mr *MockConfigurationVersionsMockRecorder) Archive(ctx, cvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Archive", reflect.TypeOf((*MockConfigurationVersions)(nil).Archive), ctx, cvID)
}

// Create mocks base method.
func (m *MockConfigurationVersions) Create(ctx context.Context, workspaceID string, options tfe.ConfigurationVersionCreateOptions) (*tfe.ConfigurationVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.ConfigurationVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockConfigurationVersionsMockRecorder) Create(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConfigurationVersions)(nil).Create), ctx, workspaceID, options)
}

// Download mocks base method.
func (m *MockConfigurationVersions) Download(ctx context.Context, cvID string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, cvID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockConfigurationVersionsMockRecorder) Download(ctx, cvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockConfigurationVersions)(nil).Download), ctx, cvID)
}

// List mocks base method.
func (m *MockConfigurationVersions) List(ctx context.Context, workspaceID string, options *tfe.ConfigurationVersionListOptions) (*tfe.ConfigurationVersionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.ConfigurationVersionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockConfigurationVersionsMockRecorder) List(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockConfigurationVersions)(nil).List), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockConfigurationVersions) Read(ctx context.Context, cvID string) (*tfe.ConfigurationVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, cvID)
	ret0, _ := ret[0].(*tfe.ConfigurationVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockConfigurationVersionsMockRecorder) Read(ctx, cvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConfigurationVersions)(nil).Read), ctx, cvID)
}

// ReadWithOptions mocks base method.
func (m *MockConfigurationVersions) ReadWithOptions(ctx context.Context, cvID string, options *tfe.ConfigurationVersionReadOptions) (*tfe.ConfigurationVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, cvID, options)
	ret0, _ := ret[0].(*tfe.ConfigurationVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockConfigurationVersionsMockRecorder) ReadWithOptions(ctx, cvID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockConfigurationVersions)(nil).ReadWithOptions), ctx, cvID, options)
}

// Upload mocks base method.
func (m *MockConfigurationVersions) Upload(ctx context.Context, url, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, url, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockConfigurationVersionsMockRecorder) Upload(ctx, url, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockConfigurationVersions)(nil).Upload), ctx, url, path)
}
