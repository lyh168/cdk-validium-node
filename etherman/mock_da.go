// Code generated by mockery v2.43.2. DO NOT EDIT.

package etherman

import (
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// daMock is an autogenerated mock type for the dataAvailabilityProvider type
type daMock struct {
	mock.Mock
}

// GetBatchL2Data provides a mock function with given fields: batchNum, hash, dataAvailabilityMessage
func (_m *daMock) GetBatchL2Data(batchNum []uint64, hash []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	ret := _m.Called(batchNum, hash, dataAvailabilityMessage)

	if len(ret) == 0 {
		panic("no return value specified for GetBatchL2Data")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]uint64, []common.Hash, []byte) ([][]byte, error)); ok {
		return rf(batchNum, hash, dataAvailabilityMessage)
	}
	if rf, ok := ret.Get(0).(func([]uint64, []common.Hash, []byte) [][]byte); ok {
		r0 = rf(batchNum, hash, dataAvailabilityMessage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]uint64, []common.Hash, []byte) error); ok {
		r1 = rf(batchNum, hash, dataAvailabilityMessage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newDaMock creates a new instance of daMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newDaMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *daMock {
	mock := &daMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
