// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/marketplaceordering (interfaces: MarketPlaceAgreementsClient)

// Package mock_marketplaceordering is a generated GoMock package.
package mock_marketplaceordering

import (
	context "context"
	reflect "reflect"

	marketplaceordering "github.com/Azure/azure-sdk-for-go/services/marketplaceordering/mgmt/2015-06-01/marketplaceordering"
	gomock "github.com/golang/mock/gomock"
)

// MockMarketPlaceAgreementsClient is a mock of MarketPlaceAgreementsClient interface
type MockMarketPlaceAgreementsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMarketPlaceAgreementsClientMockRecorder
}

// MockMarketPlaceAgreementsClientMockRecorder is the mock recorder for MockMarketPlaceAgreementsClient
type MockMarketPlaceAgreementsClientMockRecorder struct {
	mock *MockMarketPlaceAgreementsClient
}

// NewMockMarketPlaceAgreementsClient creates a new mock instance
func NewMockMarketPlaceAgreementsClient(ctrl *gomock.Controller) *MockMarketPlaceAgreementsClient {
	mock := &MockMarketPlaceAgreementsClient{ctrl: ctrl}
	mock.recorder = &MockMarketPlaceAgreementsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketPlaceAgreementsClient) EXPECT() *MockMarketPlaceAgreementsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMarketPlaceAgreementsClient) Create(arg0 context.Context, arg1, arg2, arg3 string, arg4 marketplaceordering.AgreementTerms) (marketplaceordering.AgreementTerms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(marketplaceordering.AgreementTerms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMarketPlaceAgreementsClientMockRecorder) Create(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMarketPlaceAgreementsClient)(nil).Create), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockMarketPlaceAgreementsClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (marketplaceordering.AgreementTerms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(marketplaceordering.AgreementTerms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMarketPlaceAgreementsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMarketPlaceAgreementsClient)(nil).Get), arg0, arg1, arg2, arg3)
}