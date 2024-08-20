// Code generated by mockery v2.43.2. DO NOT EDIT.

package sequencesender

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethtxmanager "github.com/0xPolygonHermez/zkevm-node/ethtxmanager"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"
)

// EthTxManagerMock is an autogenerated mock type for the ethTxManager type
type EthTxManagerMock struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, owner, id, from, to, value, data, gasOffset, dbTx
func (_m *EthTxManagerMock) Add(ctx context.Context, owner string, id string, from common.Address, to *common.Address, value *big.Int, data []byte, gasOffset uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, owner, id, from, to, value, data, gasOffset, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, common.Address, *common.Address, *big.Int, []byte, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, owner, id, from, to, value, data, gasOffset, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessPendingMonitoredTxs provides a mock function with given fields: ctx, owner, failedResultHandler, dbTx
func (_m *EthTxManagerMock) ProcessPendingMonitoredTxs(ctx context.Context, owner string, failedResultHandler ethtxmanager.ResultHandler, dbTx pgx.Tx) {
	_m.Called(ctx, owner, failedResultHandler, dbTx)
}

// NewEthTxManagerMock creates a new instance of EthTxManagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthTxManagerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthTxManagerMock {
	mock := &EthTxManagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
