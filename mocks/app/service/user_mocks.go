// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/personal/user-manager-backend/app/service (interfaces: UserService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	dto "github.com/personal/user-manager-backend/app/dto"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(arg0 *gin.Context, arg1 *dto.CreateUserDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(arg0 *gin.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), arg0, arg1)
}

// FindCompanyByID mocks base method.
func (m *MockUserService) FindCompanyByID(arg0 *gin.Context, arg1 int) (*dto.CompanyDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCompanyByID", arg0, arg1)
	ret0, _ := ret[0].(*dto.CompanyDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCompanyByID indicates an expected call of FindCompanyByID.
func (mr *MockUserServiceMockRecorder) FindCompanyByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCompanyByID", reflect.TypeOf((*MockUserService)(nil).FindCompanyByID), arg0, arg1)
}

// FindUserByEmail mocks base method.
func (m *MockUserService) FindUserByEmail(arg0 *gin.Context, arg1 string) (*dto.UserDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*dto.UserDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserServiceMockRecorder) FindUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserService)(nil).FindUserByEmail), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockUserService) GetUser(arg0 *gin.Context, arg1 int) (*dto.UserDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*dto.UserDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockUserService) GetUsers(arg0 *gin.Context, arg1, arg2 int) (*dto.UserDtoPaginated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.UserDtoPaginated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceMockRecorder) GetUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), arg0, arg1, arg2)
}

// Search mocks base method.
func (m *MockUserService) Search(arg0 *gin.Context, arg1 string, arg2, arg3 int) (*dto.UserDtoPaginated, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*dto.UserDtoPaginated)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockUserServiceMockRecorder) Search(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUserService)(nil).Search), arg0, arg1, arg2, arg3)
}

// UpdateUser mocks base method.
func (m *MockUserService) UpdateUser(arg0 *gin.Context, arg1 *dto.CreateUserDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserService)(nil).UpdateUser), arg0, arg1)
}
