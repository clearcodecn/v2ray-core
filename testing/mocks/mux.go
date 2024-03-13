// Code generated by MockGen. DO NOT EDIT.
// Source: v2ray.com/core/common/mux (interfaces: ClientWorkerFactory)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	mux "github.com/clearcodecn/v2ray-core/common/mux"
)

// MuxClientWorkerFactory is a mock of ClientWorkerFactory interface
type MuxClientWorkerFactory struct {
	ctrl     *gomock.Controller
	recorder *MuxClientWorkerFactoryMockRecorder
}

// MuxClientWorkerFactoryMockRecorder is the mock recorder for MuxClientWorkerFactory
type MuxClientWorkerFactoryMockRecorder struct {
	mock *MuxClientWorkerFactory
}

// NewMuxClientWorkerFactory creates a new mock instance
func NewMuxClientWorkerFactory(ctrl *gomock.Controller) *MuxClientWorkerFactory {
	mock := &MuxClientWorkerFactory{ctrl: ctrl}
	mock.recorder = &MuxClientWorkerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MuxClientWorkerFactory) EXPECT() *MuxClientWorkerFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MuxClientWorkerFactory) Create() (*mux.ClientWorker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(*mux.ClientWorker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MuxClientWorkerFactoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MuxClientWorkerFactory)(nil).Create))
}
