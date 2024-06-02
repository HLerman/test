// Code generated by MockGen. DO NOT EDIT.
// Source: endpoint.go
//
// Generated by this command:
//
//	mockgen -source=endpoint.go -package transport -destination=mock.go
//

// Package transport is a generated GoMock package.
package transport

import (
	context "context"
	reflect "reflect"

	api "github.com/HLerman/test"
	gomock "go.uber.org/mock/gomock"
)

// MockInvoiceService is a mock of InvoiceService interface.
type MockInvoiceService struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceServiceMockRecorder
}

// MockInvoiceServiceMockRecorder is the mock recorder for MockInvoiceService.
type MockInvoiceServiceMockRecorder struct {
	mock *MockInvoiceService
}

// NewMockInvoiceService creates a new mock instance.
func NewMockInvoiceService(ctrl *gomock.Controller) *MockInvoiceService {
	mock := &MockInvoiceService{ctrl: ctrl}
	mock.recorder = &MockInvoiceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceService) EXPECT() *MockInvoiceServiceMockRecorder {
	return m.recorder
}

// AddInvoice mocks base method.
func (m *MockInvoiceService) AddInvoice(ctx context.Context, invoice api.PostInvoiceJSONRequestBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInvoice", ctx, invoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInvoice indicates an expected call of AddInvoice.
func (mr *MockInvoiceServiceMockRecorder) AddInvoice(ctx, invoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInvoice", reflect.TypeOf((*MockInvoiceService)(nil).AddInvoice), ctx, invoice)
}

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// AddTransaction mocks base method.
func (m *MockTransactionService) AddTransaction(ctx context.Context, transaction api.PostTransactionJSONRequestBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", ctx, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction.
func (mr *MockTransactionServiceMockRecorder) AddTransaction(ctx, transaction any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockTransactionService)(nil).AddTransaction), ctx, transaction)
}

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

// GetUsers mocks base method.
func (m *MockUserService) GetUsers(ctx context.Context) (api.UserArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].(api.UserArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceMockRecorder) GetUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), ctx)
}