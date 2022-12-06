// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// RegisterManagedInstance provides a mock function with given fields: activationCode, activationID, publicKey, publicKeyType, fingerprint
func (_m *IClient) RegisterManagedInstance(activationCode string, activationID string, publicKey string, publicKeyType string, fingerprint string) (string, error) {
	ret := _m.Called(activationCode, activationID, publicKey, publicKeyType, fingerprint)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) string); ok {
		r0 = rf(activationCode, activationID, publicKey, publicKeyType, fingerprint)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(activationCode, activationID, publicKey, publicKeyType, fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIClient creates a new instance of IClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIClient(t testing.TB) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}