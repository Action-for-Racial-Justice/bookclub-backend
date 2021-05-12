// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers (interfaces: Handlers)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHandlers is a mock of Handlers interface.
type MockHandlers struct {
	ctrl     *gomock.Controller
	recorder *MockHandlersMockRecorder
}

// MockHandlersMockRecorder is the mock recorder for MockHandlers.
type MockHandlersMockRecorder struct {
	mock *MockHandlers
}

// NewMockHandlers creates a new mock instance.
func NewMockHandlers(ctrl *gomock.Controller) *MockHandlers {
	mock := &MockHandlers{ctrl: ctrl}
	mock.recorder = &MockHandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlers) EXPECT() *MockHandlersMockRecorder {
	return m.recorder
}

// CreateClub mocks base method.
func (m *MockHandlers) CreateClub(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateClub", arg0, arg1)
}

// CreateClub indicates an expected call of CreateClub.
func (mr *MockHandlersMockRecorder) CreateClub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClub", reflect.TypeOf((*MockHandlers)(nil).CreateClub), arg0, arg1)
}

// EndUserSession mocks base method.
func (m *MockHandlers) EndUserSession(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndUserSession", arg0, arg1)
}

// EndUserSession indicates an expected call of EndUserSession.
func (mr *MockHandlersMockRecorder) EndUserSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndUserSession", reflect.TypeOf((*MockHandlers)(nil).EndUserSession), arg0, arg1)
}

// GetBookData mocks base method.
func (m *MockHandlers) GetBookData(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBookData", arg0, arg1)
}

// GetBookData indicates an expected call of GetBookData.
func (mr *MockHandlersMockRecorder) GetBookData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookData", reflect.TypeOf((*MockHandlers)(nil).GetBookData), arg0, arg1)
}

// GetClubData mocks base method.
func (m *MockHandlers) GetClubData(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetClubData", arg0, arg1)
}

// GetClubData indicates an expected call of GetClubData.
func (mr *MockHandlersMockRecorder) GetClubData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClubData", reflect.TypeOf((*MockHandlers)(nil).GetClubData), arg0, arg1)
}

// GetClubs mocks base method.
func (m *MockHandlers) GetClubs(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetClubs", arg0, arg1)
}

// GetClubs indicates an expected call of GetClubs.
func (mr *MockHandlersMockRecorder) GetClubs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClubs", reflect.TypeOf((*MockHandlers)(nil).GetClubs), arg0, arg1)
}

// GetSSOToken mocks base method.
func (m *MockHandlers) GetSSOToken(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetSSOToken", arg0, arg1)
}

// GetSSOToken indicates an expected call of GetSSOToken.
func (mr *MockHandlersMockRecorder) GetSSOToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSOToken", reflect.TypeOf((*MockHandlers)(nil).GetSSOToken), arg0, arg1)
}

// GetUserClubs mocks base method.
func (m *MockHandlers) GetUserClubs(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserClubs", arg0, arg1)
}

// GetUserClubs indicates an expected call of GetUserClubs.
func (mr *MockHandlersMockRecorder) GetUserClubs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserClubs", reflect.TypeOf((*MockHandlers)(nil).GetUserClubs), arg0, arg1)
}

// GetUserData mocks base method.
func (m *MockHandlers) GetUserData(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserData", arg0, arg1)
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockHandlersMockRecorder) GetUserData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*MockHandlers)(nil).GetUserData), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockHandlers) HealthCheck(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HealthCheck", arg0, arg1)
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockHandlersMockRecorder) HealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockHandlers)(nil).HealthCheck), arg0, arg1)
}

// JoinClub mocks base method.
func (m *MockHandlers) JoinClub(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JoinClub", arg0, arg1)
}

// JoinClub indicates an expected call of JoinClub.
func (mr *MockHandlersMockRecorder) JoinClub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinClub", reflect.TypeOf((*MockHandlers)(nil).JoinClub), arg0, arg1)
}

// LeaveClub mocks base method.
func (m *MockHandlers) LeaveClub(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LeaveClub", arg0, arg1)
}

// LeaveClub indicates an expected call of LeaveClub.
func (mr *MockHandlersMockRecorder) LeaveClub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveClub", reflect.TypeOf((*MockHandlers)(nil).LeaveClub), arg0, arg1)
}

// ServeHTTP mocks base method.
func (m *MockHandlers) ServeHTTP(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTP", arg0, arg1)
}

// ServeHTTP indicates an expected call of ServeHTTP.
func (mr *MockHandlersMockRecorder) ServeHTTP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockHandlers)(nil).ServeHTTP), arg0, arg1)
}
