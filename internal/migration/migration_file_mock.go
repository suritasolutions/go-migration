// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/migration/migration_file_interface.go
//
// Generated by this command:
//
//	mockgen -source=./internal/migration/migration_file_interface.go -destination=./internal/migration/migration_file_mock.go -package=migration
//

// Package migration is a generated GoMock package.
package migration

import (
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMigrationFile is a mock of MigrationFile interface.
type MockMigrationFile struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationFileMockRecorder
}

// MockMigrationFileMockRecorder is the mock recorder for MockMigrationFile.
type MockMigrationFileMockRecorder struct {
	mock *MockMigrationFile
}

// NewMockMigrationFile creates a new mock instance.
func NewMockMigrationFile(ctrl *gomock.Controller) *MockMigrationFile {
	mock := &MockMigrationFile{ctrl: ctrl}
	mock.recorder = &MockMigrationFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrationFile) EXPECT() *MockMigrationFileMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMigrationFile) Create(database, path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", database, path)
}

// Create indicates an expected call of Create.
func (mr *MockMigrationFileMockRecorder) Create(database, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMigrationFile)(nil).Create), database, path)
}

// DBMigrationFolderExists mocks base method.
func (m *MockMigrationFile) DBMigrationFolderExists(folder string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBMigrationFolderExists", folder)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DBMigrationFolderExists indicates an expected call of DBMigrationFolderExists.
func (mr *MockMigrationFileMockRecorder) DBMigrationFolderExists(folder any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBMigrationFolderExists", reflect.TypeOf((*MockMigrationFile)(nil).DBMigrationFolderExists), folder)
}

// GetMigrationFileContent mocks base method.
func (m *MockMigrationFile) GetMigrationFileContent(file os.DirEntry) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigrationFileContent", file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMigrationFileContent indicates an expected call of GetMigrationFileContent.
func (mr *MockMigrationFileMockRecorder) GetMigrationFileContent(file any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrationFileContent", reflect.TypeOf((*MockMigrationFile)(nil).GetMigrationFileContent), file)
}

// GetMigrationSQLFiles mocks base method.
func (m *MockMigrationFile) GetMigrationSQLFiles() []os.DirEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigrationSQLFiles")
	ret0, _ := ret[0].([]os.DirEntry)
	return ret0
}

// GetMigrationSQLFiles indicates an expected call of GetMigrationSQLFiles.
func (mr *MockMigrationFileMockRecorder) GetMigrationSQLFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrationSQLFiles", reflect.TypeOf((*MockMigrationFile)(nil).GetMigrationSQLFiles))
}
