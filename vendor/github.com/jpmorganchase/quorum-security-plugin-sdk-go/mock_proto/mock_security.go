// Code generated by MockGen. DO NOT EDIT.
// Source: proto/security.pb.go

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	grpc "google.golang.org/grpc"
)

// MockTLSConfigurationSourceClient is a mock of TLSConfigurationSourceClient interface
type MockTLSConfigurationSourceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTLSConfigurationSourceClientMockRecorder
}

// MockTLSConfigurationSourceClientMockRecorder is the mock recorder for MockTLSConfigurationSourceClient
type MockTLSConfigurationSourceClientMockRecorder struct {
	mock *MockTLSConfigurationSourceClient
}

// NewMockTLSConfigurationSourceClient creates a new mock instance
func NewMockTLSConfigurationSourceClient(ctrl *gomock.Controller) *MockTLSConfigurationSourceClient {
	mock := &MockTLSConfigurationSourceClient{ctrl: ctrl}
	mock.recorder = &MockTLSConfigurationSourceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTLSConfigurationSourceClient) EXPECT() *MockTLSConfigurationSourceClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTLSConfigurationSourceClient) Get(ctx context.Context, in *proto.TLSConfiguration_Request, opts ...grpc.CallOption) (*proto.TLSConfiguration_Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*proto.TLSConfiguration_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTLSConfigurationSourceClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTLSConfigurationSourceClient)(nil).Get), varargs...)
}

// MockTLSConfigurationSourceServer is a mock of TLSConfigurationSourceServer interface
type MockTLSConfigurationSourceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTLSConfigurationSourceServerMockRecorder
}

// MockTLSConfigurationSourceServerMockRecorder is the mock recorder for MockTLSConfigurationSourceServer
type MockTLSConfigurationSourceServerMockRecorder struct {
	mock *MockTLSConfigurationSourceServer
}

// NewMockTLSConfigurationSourceServer creates a new mock instance
func NewMockTLSConfigurationSourceServer(ctrl *gomock.Controller) *MockTLSConfigurationSourceServer {
	mock := &MockTLSConfigurationSourceServer{ctrl: ctrl}
	mock.recorder = &MockTLSConfigurationSourceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTLSConfigurationSourceServer) EXPECT() *MockTLSConfigurationSourceServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTLSConfigurationSourceServer) Get(arg0 context.Context, arg1 *proto.TLSConfiguration_Request) (*proto.TLSConfiguration_Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*proto.TLSConfiguration_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTLSConfigurationSourceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTLSConfigurationSourceServer)(nil).Get), arg0, arg1)
}

// MockAuthenticationManagerClient is a mock of AuthenticationManagerClient interface
type MockAuthenticationManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationManagerClientMockRecorder
}

// MockAuthenticationManagerClientMockRecorder is the mock recorder for MockAuthenticationManagerClient
type MockAuthenticationManagerClientMockRecorder struct {
	mock *MockAuthenticationManagerClient
}

// NewMockAuthenticationManagerClient creates a new mock instance
func NewMockAuthenticationManagerClient(ctrl *gomock.Controller) *MockAuthenticationManagerClient {
	mock := &MockAuthenticationManagerClient{ctrl: ctrl}
	mock.recorder = &MockAuthenticationManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationManagerClient) EXPECT() *MockAuthenticationManagerClientMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockAuthenticationManagerClient) Authenticate(ctx context.Context, in *proto.AuthenticationToken, opts ...grpc.CallOption) (*proto.PreAuthenticatedAuthenticationToken, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Authenticate", varargs...)
	ret0, _ := ret[0].(*proto.PreAuthenticatedAuthenticationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAuthenticationManagerClientMockRecorder) Authenticate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationManagerClient)(nil).Authenticate), varargs...)
}

// MockAuthenticationManagerServer is a mock of AuthenticationManagerServer interface
type MockAuthenticationManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationManagerServerMockRecorder
}

// MockAuthenticationManagerServerMockRecorder is the mock recorder for MockAuthenticationManagerServer
type MockAuthenticationManagerServerMockRecorder struct {
	mock *MockAuthenticationManagerServer
}

// NewMockAuthenticationManagerServer creates a new mock instance
func NewMockAuthenticationManagerServer(ctrl *gomock.Controller) *MockAuthenticationManagerServer {
	mock := &MockAuthenticationManagerServer{ctrl: ctrl}
	mock.recorder = &MockAuthenticationManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationManagerServer) EXPECT() *MockAuthenticationManagerServerMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockAuthenticationManagerServer) Authenticate(arg0 context.Context, arg1 *proto.AuthenticationToken) (*proto.PreAuthenticatedAuthenticationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(*proto.PreAuthenticatedAuthenticationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAuthenticationManagerServerMockRecorder) Authenticate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationManagerServer)(nil).Authenticate), arg0, arg1)
}
