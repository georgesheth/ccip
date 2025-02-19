// Code generated by mockery v2.28.1. DO NOT EDIT.

package ccipdata

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

// MockUSDCReader is an autogenerated mock type for the USDCReader type
type MockUSDCReader struct {
	mock.Mock
}

// Close provides a mock function with given fields: qopts
func (_m *MockUSDCReader) Close(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLastUSDCMessagePriorToLogIndexInTx provides a mock function with given fields: ctx, logIndex, txHash
func (_m *MockUSDCReader) GetLastUSDCMessagePriorToLogIndexInTx(ctx context.Context, logIndex int64, txHash common.Hash) ([]byte, error) {
	ret := _m.Called(ctx, logIndex, txHash)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, common.Hash) ([]byte, error)); ok {
		return rf(ctx, logIndex, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, common.Hash) []byte); ok {
		r0 = rf(ctx, logIndex, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, common.Hash) error); ok {
		r1 = rf(ctx, logIndex, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockUSDCReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUSDCReader creates a new instance of MockUSDCReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUSDCReader(t mockConstructorTestingTNewMockUSDCReader) *MockUSDCReader {
	mock := &MockUSDCReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
