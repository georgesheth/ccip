// Code generated by mockery v2.28.1. DO NOT EDIT.

package ccipdata

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	prices "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"

	time "time"
)

// MockCommitStoreReader is an autogenerated mock type for the CommitStoreReader type
type MockCommitStoreReader struct {
	mock.Mock
}

// ChangeConfig provides a mock function with given fields: onchainConfig, offchainConfig
func (_m *MockCommitStoreReader) ChangeConfig(onchainConfig []byte, offchainConfig []byte) (common.Address, error) {
	ret := _m.Called(onchainConfig, offchainConfig)

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) (common.Address, error)); ok {
		return rf(onchainConfig, offchainConfig)
	}
	if rf, ok := ret.Get(0).(func([]byte, []byte) common.Address); ok {
		r0 = rf(onchainConfig, offchainConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, []byte) error); ok {
		r1 = rf(onchainConfig, offchainConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields: qopts
func (_m *MockCommitStoreReader) Close(qopts ...pg.QOpt) error {
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

// DecodeCommitReport provides a mock function with given fields: report
func (_m *MockCommitStoreReader) DecodeCommitReport(report []byte) (CommitStoreReport, error) {
	ret := _m.Called(report)

	var r0 CommitStoreReport
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (CommitStoreReport, error)); ok {
		return rf(report)
	}
	if rf, ok := ret.Get(0).(func([]byte) CommitStoreReport); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Get(0).(CommitStoreReport)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeCommitReport provides a mock function with given fields: report
func (_m *MockCommitStoreReader) EncodeCommitReport(report CommitStoreReport) ([]byte, error) {
	ret := _m.Called(report)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(CommitStoreReport) ([]byte, error)); ok {
		return rf(report)
	}
	if rf, ok := ret.Get(0).(func(CommitStoreReport) []byte); ok {
		r0 = rf(report)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(CommitStoreReport) error); ok {
		r1 = rf(report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasPriceEstimator provides a mock function with given fields:
func (_m *MockCommitStoreReader) GasPriceEstimator() prices.GasPriceEstimatorCommit {
	ret := _m.Called()

	var r0 prices.GasPriceEstimatorCommit
	if rf, ok := ret.Get(0).(func() prices.GasPriceEstimatorCommit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(prices.GasPriceEstimatorCommit)
		}
	}

	return r0
}

// GetAcceptedCommitReportsGteSeqNum provides a mock function with given fields: ctx, seqNum, confs
func (_m *MockCommitStoreReader) GetAcceptedCommitReportsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[CommitStoreReport], error) {
	ret := _m.Called(ctx, seqNum, confs)

	var r0 []Event[CommitStoreReport]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int) ([]Event[CommitStoreReport], error)); ok {
		return rf(ctx, seqNum, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int) []Event[CommitStoreReport]); ok {
		r0 = rf(ctx, seqNum, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Event[CommitStoreReport])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, int) error); ok {
		r1 = rf(ctx, seqNum, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcceptedCommitReportsGteTimestamp provides a mock function with given fields: ctx, ts, confs
func (_m *MockCommitStoreReader) GetAcceptedCommitReportsGteTimestamp(ctx context.Context, ts time.Time, confs int) ([]Event[CommitStoreReport], error) {
	ret := _m.Called(ctx, ts, confs)

	var r0 []Event[CommitStoreReport]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, int) ([]Event[CommitStoreReport], error)); ok {
		return rf(ctx, ts, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, int) []Event[CommitStoreReport]); ok {
		r0 = rf(ctx, ts, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Event[CommitStoreReport])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, int) error); ok {
		r1 = rf(ctx, ts, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpectedNextSequenceNumber provides a mock function with given fields: _a0
func (_m *MockCommitStoreReader) GetExpectedNextSequenceNumber(_a0 context.Context) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestPriceEpochAndRound provides a mock function with given fields: _a0
func (_m *MockCommitStoreReader) GetLatestPriceEpochAndRound(_a0 context.Context) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBlessed provides a mock function with given fields: ctx, root
func (_m *MockCommitStoreReader) IsBlessed(ctx context.Context, root [32]byte) (bool, error) {
	ret := _m.Called(ctx, root)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) (bool, error)); ok {
		return rf(ctx, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) bool); ok {
		r0 = rf(ctx, root)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, [32]byte) error); ok {
		r1 = rf(ctx, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDown provides a mock function with given fields: ctx
func (_m *MockCommitStoreReader) IsDown(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OffchainConfig provides a mock function with given fields:
func (_m *MockCommitStoreReader) OffchainConfig() CommitOffchainConfig {
	ret := _m.Called()

	var r0 CommitOffchainConfig
	if rf, ok := ret.Get(0).(func() CommitOffchainConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(CommitOffchainConfig)
	}

	return r0
}

// VerifyExecutionReport provides a mock function with given fields: ctx, report
func (_m *MockCommitStoreReader) VerifyExecutionReport(ctx context.Context, report ExecReport) (bool, error) {
	ret := _m.Called(ctx, report)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ExecReport) (bool, error)); ok {
		return rf(ctx, report)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ExecReport) bool); ok {
		r0 = rf(ctx, report)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ExecReport) error); ok {
		r1 = rf(ctx, report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockCommitStoreReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCommitStoreReader creates a new instance of MockCommitStoreReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCommitStoreReader(t mockConstructorTestingTNewMockCommitStoreReader) *MockCommitStoreReader {
	mock := &MockCommitStoreReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
