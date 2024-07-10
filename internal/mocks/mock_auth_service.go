// Code generated by MockGen. DO NOT EDIT.
// Source: auth_service.go
//
// Generated by this command:
//
//	mockgen -source=auth_service.go -destination=mock_auth_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	entities "github.com/Magetan-Boyz/Backend/internal/domain/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockAuthService) ChangePassword(userID, oldPassword, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", userID, oldPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockAuthServiceMockRecorder) ChangePassword(userID, oldPassword, newPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockAuthService)(nil).ChangePassword), userID, oldPassword, newPassword)
}

// CreateUserToken mocks base method.
func (m *MockAuthService) CreateUserToken(user *entities.User, role string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserToken", user, role)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserToken indicates an expected call of CreateUserToken.
func (mr *MockAuthServiceMockRecorder) CreateUserToken(user, role any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserToken", reflect.TypeOf((*MockAuthService)(nil).CreateUserToken), user, role)
}

// FindUserByToken mocks base method.
func (m *MockAuthService) FindUserByToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByToken indicates an expected call of FindUserByToken.
func (mr *MockAuthServiceMockRecorder) FindUserByToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByToken", reflect.TypeOf((*MockAuthService)(nil).FindUserByToken), token)
}

// GetRoleNameFromID mocks base method.
func (m *MockAuthService) GetRoleNameFromID(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleNameFromID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleNameFromID indicates an expected call of GetRoleNameFromID.
func (mr *MockAuthServiceMockRecorder) GetRoleNameFromID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleNameFromID", reflect.TypeOf((*MockAuthService)(nil).GetRoleNameFromID), id)
}

// GetUserByToken mocks base method.
func (m *MockAuthService) GetUserByToken(token string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByToken", token)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByToken indicates an expected call of GetUserByToken.
func (mr *MockAuthServiceMockRecorder) GetUserByToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByToken", reflect.TypeOf((*MockAuthService)(nil).GetUserByToken), token)
}

// LogIn mocks base method.
func (m *MockAuthService) LogIn(username, password string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", username, password)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogIn indicates an expected call of LogIn.
func (mr *MockAuthServiceMockRecorder) LogIn(username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockAuthService)(nil).LogIn), username, password)
}

// LogOut mocks base method.
func (m *MockAuthService) LogOut(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogOut", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogOut indicates an expected call of LogOut.
func (mr *MockAuthServiceMockRecorder) LogOut(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOut", reflect.TypeOf((*MockAuthService)(nil).LogOut), userID)
}
