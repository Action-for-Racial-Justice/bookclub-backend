// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Action-for-Racial-Justice/bookclub-backend/internal/requests (interfaces: IRequests)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIRequests is a mock of IRequests interface.
type MockIRequests struct {
	ctrl     *gomock.Controller
	recorder *MockIRequestsMockRecorder
}

// MockIRequestsMockRecorder is the mock recorder for MockIRequests.
type MockIRequestsMockRecorder struct {
	mock *MockIRequests
}

// NewMockIRequests creates a new mock instance.
func NewMockIRequests(ctrl *gomock.Controller) *MockIRequests {
	mock := &MockIRequests{ctrl: ctrl}
	mock.recorder = &MockIRequestsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRequests) EXPECT() *MockIRequestsMockRecorder {
	return m.recorder
}

// EndUserSession mocks base method.
func (m *MockIRequests) EndUserSession(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndUserSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EndUserSession indicates an expected call of EndUserSession.
func (mr *MockIRequestsMockRecorder) EndUserSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndUserSession", reflect.TypeOf((*MockIRequests)(nil).EndUserSession), arg0)
}

// GetLoginResponse mocks base method.
func (m *MockIRequests) GetLoginResponse(arg0 *models.UserLoginRequest) (*models.ArjAPILoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoginResponse", arg0)
	ret0, _ := ret[0].(*models.ArjAPILoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoginResponse indicates an expected call of GetLoginResponse.
func (mr *MockIRequestsMockRecorder) GetLoginResponse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginResponse", reflect.TypeOf((*MockIRequests)(nil).GetLoginResponse), arg0)
}

// GetUserData mocks base method.
func (m *MockIRequests) GetUserData(arg0 string) (*models.ArjAPIUserDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserData", arg0)
	ret0, _ := ret[0].(*models.ArjAPIUserDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockIRequestsMockRecorder) GetUserData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*MockIRequests)(nil).GetUserData), arg0)
}
