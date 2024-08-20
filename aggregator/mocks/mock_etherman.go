// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/etherman/types"
)

// Etherman is an autogenerated mock type for the etherman type
type Etherman struct {
	mock.Mock
}

// BuildTrustedVerifyBatchesTxData provides a mock function with given fields: lastVerifiedBatch, newVerifiedBatch, inputs, beneficiary
func (_m *Etherman) BuildTrustedVerifyBatchesTxData(lastVerifiedBatch uint64, newVerifiedBatch uint64, inputs *types.FinalProofInputs, beneficiary common.Address) (*common.Address, []byte, error) {
	ret := _m.Called(lastVerifiedBatch, newVerifiedBatch, inputs, beneficiary)

	if len(ret) == 0 {
		panic("no return value specified for BuildTrustedVerifyBatchesTxData")
	}

	var r0 *common.Address
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(uint64, uint64, *types.FinalProofInputs, common.Address) (*common.Address, []byte, error)); ok {
		return rf(lastVerifiedBatch, newVerifiedBatch, inputs, beneficiary)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64, *types.FinalProofInputs, common.Address) *common.Address); ok {
		r0 = rf(lastVerifiedBatch, newVerifiedBatch, inputs, beneficiary)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64, *types.FinalProofInputs, common.Address) []byte); ok {
		r1 = rf(lastVerifiedBatch, newVerifiedBatch, inputs, beneficiary)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(uint64, uint64, *types.FinalProofInputs, common.Address) error); ok {
		r2 = rf(lastVerifiedBatch, newVerifiedBatch, inputs, beneficiary)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLatestBlockHeader provides a mock function with given fields: ctx
func (_m *Etherman) GetLatestBlockHeader(ctx context.Context) (*coretypes.Header, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBlockHeader")
	}

	var r0 *coretypes.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*coretypes.Header, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *coretypes.Header); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestVerifiedBatchNum provides a mock function with given fields:
func (_m *Etherman) GetLatestVerifiedBatchNum() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLatestVerifiedBatchNum")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRollupId provides a mock function with given fields:
func (_m *Etherman) GetRollupId() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRollupId")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// NewEtherman creates a new instance of Etherman. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEtherman(t interface {
	mock.TestingT
	Cleanup(func())
}) *Etherman {
	mock := &Etherman{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
