/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network/networkapi (interfaces: RouteTablesClientAPI)

// Package mock_routetables is a generated GoMock package.
package mock_routetables

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/network/mgmt/network"
	gomock "github.com/golang/mock/gomock"
)

// MockRouteTablesClientAPI is a mock of RouteTablesClientAPI interface
type MockRouteTablesClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTablesClientAPIMockRecorder
}

// MockRouteTablesClientAPIMockRecorder is the mock recorder for MockRouteTablesClientAPI
type MockRouteTablesClientAPIMockRecorder struct {
	mock *MockRouteTablesClientAPI
}

// NewMockRouteTablesClientAPI creates a new mock instance
func NewMockRouteTablesClientAPI(ctrl *gomock.Controller) *MockRouteTablesClientAPI {
	mock := &MockRouteTablesClientAPI{ctrl: ctrl}
	mock.recorder = &MockRouteTablesClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteTablesClientAPI) EXPECT() *MockRouteTablesClientAPIMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockRouteTablesClientAPI) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 network.RouteTable) (network.RouteTablesCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.RouteTablesCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockRouteTablesClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockRouteTablesClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockRouteTablesClientAPI) Delete(arg0 context.Context, arg1, arg2 string) (network.RouteTablesDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.RouteTablesDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockRouteTablesClientAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRouteTablesClientAPI)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockRouteTablesClientAPI) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.RouteTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.RouteTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRouteTablesClientAPIMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRouteTablesClientAPI)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockRouteTablesClientAPI) List(arg0 context.Context, arg1 string) (network.RouteTableListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.RouteTableListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRouteTablesClientAPIMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRouteTablesClientAPI)(nil).List), arg0, arg1)
}

// ListAll mocks base method
func (m *MockRouteTablesClientAPI) ListAll(arg0 context.Context) (network.RouteTableListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.RouteTableListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockRouteTablesClientAPIMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockRouteTablesClientAPI)(nil).ListAll), arg0)
}

// UpdateTags mocks base method
func (m *MockRouteTablesClientAPI) UpdateTags(arg0 context.Context, arg1, arg2 string, arg3 network.TagsObject) (network.RouteTablesUpdateTagsFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTags", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.RouteTablesUpdateTagsFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTags indicates an expected call of UpdateTags
func (mr *MockRouteTablesClientAPIMockRecorder) UpdateTags(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTags", reflect.TypeOf((*MockRouteTablesClientAPI)(nil).UpdateTags), arg0, arg1, arg2, arg3)
}
