// Code generated by MockGen. DO NOT EDIT.
// Source: admin_service.go
//
// Generated by this command:
//
//	mockgen -source=admin_service.go -destination=mock_admin_service.go -package=services
//

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	entities "github.com/Magetan-Boyz/Backend/internal/domain/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminService is a mock of AdminService interface.
type MockAdminService struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceMockRecorder
}

// MockAdminServiceMockRecorder is the mock recorder for MockAdminService.
type MockAdminServiceMockRecorder struct {
	mock *MockAdminService
}

// NewMockAdminService creates a new mock instance.
func NewMockAdminService(ctrl *gomock.Controller) *MockAdminService {
	mock := &MockAdminService{ctrl: ctrl}
	mock.recorder = &MockAdminServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminService) EXPECT() *MockAdminServiceMockRecorder {
	return m.recorder
}

// CreateSubject mocks base method.
func (m *MockAdminService) CreateSubject(subject *entities.Subject) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubject", subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubject indicates an expected call of CreateSubject.
func (mr *MockAdminServiceMockRecorder) CreateSubject(subject any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubject", reflect.TypeOf((*MockAdminService)(nil).CreateSubject), subject)
}

// CreateTeacher mocks base method.
func (m *MockAdminService) CreateTeacher(teacher *entities.Teacher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeacher", teacher)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeacher indicates an expected call of CreateTeacher.
func (mr *MockAdminServiceMockRecorder) CreateTeacher(teacher any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeacher", reflect.TypeOf((*MockAdminService)(nil).CreateTeacher), teacher)
}
