// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockmonitorInterface is a mock of monitorInterface interface
type MockmonitorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockmonitorInterfaceMockRecorder
}

// MockmonitorInterfaceMockRecorder is the mock recorder for MockmonitorInterface
type MockmonitorInterfaceMockRecorder struct {
	mock *MockmonitorInterface
}

// NewMockmonitorInterface creates a new mock instance
func NewMockmonitorInterface(ctrl *gomock.Controller) *MockmonitorInterface {
	mock := &MockmonitorInterface{ctrl: ctrl}
	mock.recorder = &MockmonitorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmonitorInterface) EXPECT() *MockmonitorInterfaceMockRecorder {
	return m.recorder
}

// WatchProemetheus mocks base method
func (m *MockmonitorInterface) WatchProemetheus() (chan []envoyTraffic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchProemetheus")
	ret0, _ := ret[0].(chan []envoyTraffic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchProemetheus indicates an expected call of WatchProemetheus
func (mr *MockmonitorInterfaceMockRecorder) WatchProemetheus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchProemetheus", reflect.TypeOf((*MockmonitorInterface)(nil).WatchProemetheus))
}

// Sacla mocks base method
func (m *MockmonitorInterface) Sacla(name string, num float64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sacla", name, num)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sacla indicates an expected call of Sacla
func (mr *MockmonitorInterfaceMockRecorder) Sacla(name, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sacla", reflect.TypeOf((*MockmonitorInterface)(nil).Sacla), name, num)
}

// WaitScala mocks base method
func (m *MockmonitorInterface) WaitScala(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitScala", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitScala indicates an expected call of WaitScala
func (mr *MockmonitorInterfaceMockRecorder) WaitScala(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitScala", reflect.TypeOf((*MockmonitorInterface)(nil).WaitScala), name)
}

// InvokeDeployService mocks base method
func (m *MockmonitorInterface) InvokeDeployService(name string, num float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeDeployService", name, num)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeDeployService indicates an expected call of InvokeDeployService
func (mr *MockmonitorInterfaceMockRecorder) InvokeDeployService(name, num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeDeployService", reflect.TypeOf((*MockmonitorInterface)(nil).InvokeDeployService), name, num)
}

// InitPloy mocks base method
func (m *MockmonitorInterface) InitPloy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitPloy")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitPloy indicates an expected call of InitPloy
func (mr *MockmonitorInterfaceMockRecorder) InitPloy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitPloy", reflect.TypeOf((*MockmonitorInterface)(nil).InitPloy))
}

// UpdatePloy mocks base method
func (m *MockmonitorInterface) UpdatePloy(arg0 string, arg1 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePloy", arg0, arg1)
}

// UpdatePloy indicates an expected call of UpdatePloy
func (mr *MockmonitorInterfaceMockRecorder) UpdatePloy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePloy", reflect.TypeOf((*MockmonitorInterface)(nil).UpdatePloy), arg0, arg1)
}

// GetPloy mocks base method
func (m *MockmonitorInterface) GetPloy() map[string]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPloy")
	ret0, _ := ret[0].(map[string]int)
	return ret0
}

// GetPloy indicates an expected call of GetPloy
func (mr *MockmonitorInterfaceMockRecorder) GetPloy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPloy", reflect.TypeOf((*MockmonitorInterface)(nil).GetPloy))
}

// NoticeProxyService mocks base method
func (m *MockmonitorInterface) NoticeProxyService(name, endpoint string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoticeProxyService", name, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// NoticeProxyService indicates an expected call of NoticeProxyService
func (mr *MockmonitorInterfaceMockRecorder) NoticeProxyService(name, endpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoticeProxyService", reflect.TypeOf((*MockmonitorInterface)(nil).NoticeProxyService), name, endpoint)
}

// NeedScala mocks base method
func (m *MockmonitorInterface) NeedScala(Traffic envoyTraffic) (bool, float64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedScala", Traffic)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// NeedScala indicates an expected call of NeedScala
func (mr *MockmonitorInterfaceMockRecorder) NeedScala(Traffic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedScala", reflect.TypeOf((*MockmonitorInterface)(nil).NeedScala), Traffic)
}

// DisableService mocks base method
func (m *MockmonitorInterface) DisableService(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableService", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableService indicates an expected call of DisableService
func (mr *MockmonitorInterfaceMockRecorder) DisableService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableService", reflect.TypeOf((*MockmonitorInterface)(nil).DisableService), arg0)
}

// MockprometheusInterface is a mock of prometheusInterface interface
type MockprometheusInterface struct {
	ctrl     *gomock.Controller
	recorder *MockprometheusInterfaceMockRecorder
}

// MockprometheusInterfaceMockRecorder is the mock recorder for MockprometheusInterface
type MockprometheusInterfaceMockRecorder struct {
	mock *MockprometheusInterface
}

// NewMockprometheusInterface creates a new mock instance
func NewMockprometheusInterface(ctrl *gomock.Controller) *MockprometheusInterface {
	mock := &MockprometheusInterface{ctrl: ctrl}
	mock.recorder = &MockprometheusInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockprometheusInterface) EXPECT() *MockprometheusInterfaceMockRecorder {
	return m.recorder
}

// QueryRange mocks base method
func (m *MockprometheusInterface) QueryRange(query string, step Step, stepVal int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRange", query, step, stepVal)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRange indicates an expected call of QueryRange
func (mr *MockprometheusInterfaceMockRecorder) QueryRange(query, step, stepVal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRange", reflect.TypeOf((*MockprometheusInterface)(nil).QueryRange), query, step, stepVal)
}

// QueryAllCluster mocks base method
func (m *MockprometheusInterface) QueryAllCluster() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAllCluster")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAllCluster indicates an expected call of QueryAllCluster
func (mr *MockprometheusInterfaceMockRecorder) QueryAllCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAllCluster", reflect.TypeOf((*MockprometheusInterface)(nil).QueryAllCluster))
}
