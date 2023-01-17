package state

import (
	"errors"
	"math/big"

	"github.com/0xPolygonHermez/zkevm-node/state/runtime/instrumentation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var ()

// TouchedAddress represents affected address after executor processing one or multiple txs
type TouchedAddress struct {
	Address common.Address
	Nonce   *uint64
	Balance *big.Int
}

// ProcessRequest represents the request of a batch process.
type ProcessRequest struct {
	BatchNumber     uint64
	GlobalExitRoot  common.Hash
	OldStateRoot    common.Hash
	OldAccInputHash common.Hash
	Transactions    []byte
	Coinbase        common.Address
	Timestamp       uint64
	Caller          CallerLabel
}

// ProcessBatchResponse represents the response of a batch process.
type ProcessBatchResponse struct {
	NewStateRoot     common.Hash
	NewAccInputHash  common.Hash
	NewLocalExitRoot common.Hash
	NewBatchNumber   uint64
	UsedZkCounters   ZKCounters
	Responses        []*ProcessTransactionResponse
	Error            error
	IsBatchProcessed bool
	TouchedAddresses map[common.Address]*TouchedAddress
}

// ProcessTransactionResponse represents the response of a tx process.
type ProcessTransactionResponse struct {
	// TxHash is the hash of the transaction
	TxHash common.Hash
	// Type indicates legacy transaction
	// It will be always 0 (legacy) in the executor
	Type uint32
	// ReturnValue is the returned data from the runtime (function result or data supplied with revert opcode)
	ReturnValue []byte
	// GasLeft is the total gas left as result of execution
	GasLeft uint64
	// GasUsed is the total gas used as result of execution or gas estimation
	GasUsed uint64
	// GasRefunded is the total gas refunded as result of execution
	GasRefunded uint64
	// Error represents any error encountered during the execution
	Error error
	// CreateAddress is the new SC Address in case of SC creation
	CreateAddress common.Address
	// StateRoot is the State Root
	StateRoot common.Hash
	// Logs emitted by LOG opcode
	Logs []*types.Log
	// IsProcessed indicates if this tx didn't fit into the batch
	IsProcessed bool
	// Tx is the whole transaction object
	Tx types.Transaction
	// ExecutionTrace contains the traces produced in the execution
	ExecutionTrace []instrumentation.StructLog
	// CallTrace contains the call trace.
	CallTrace instrumentation.ExecutorTrace
}

// ZKCounters counters for the tx
type ZKCounters struct {
	CumulativeGasUsed    uint64
	UsedKeccakHashes     uint32
	UsedPoseidonHashes   uint32
	UsedPoseidonPaddings uint32
	UsedMemAligns        uint32
	UsedArithmetics      uint32
	UsedBinaries         uint32
	UsedSteps            uint32
}

// SumUp sum ups zk counters with passed tx zk counters
func (z *ZKCounters) SumUp(other ZKCounters) {
	z.CumulativeGasUsed += other.CumulativeGasUsed
	z.UsedKeccakHashes += other.UsedKeccakHashes
	z.UsedPoseidonHashes += other.UsedPoseidonHashes
	z.UsedPoseidonPaddings += other.UsedPoseidonPaddings
	z.UsedMemAligns += other.UsedMemAligns
	z.UsedArithmetics += other.UsedArithmetics
	z.UsedBinaries += other.UsedBinaries
	z.UsedSteps += other.UsedSteps
}

// Sub subtract zk counters with passed zk counters (not safe)
func (z *ZKCounters) Sub(other ZKCounters) error {
	// ZKCounters
	if other.CumulativeGasUsed > z.CumulativeGasUsed {
		return GetZKCounterError("CumulativeGasUsed")
	}
	if other.UsedKeccakHashes > z.UsedKeccakHashes {
		return GetZKCounterError("UsedKeccakHashes")
	}
	if other.UsedPoseidonHashes > z.UsedPoseidonHashes {
		return GetZKCounterError("UsedPoseidonHashes")
	}
	if other.UsedPoseidonPaddings > z.UsedPoseidonPaddings {
		return errors.New("underflow ZKCounter: UsedPoseidonPaddings")
	}
	if other.UsedMemAligns > z.UsedMemAligns {
		return GetZKCounterError("UsedMemAligns")
	}
	if other.UsedArithmetics > z.UsedArithmetics {
		return GetZKCounterError("UsedArithmetics")
	}
	if other.UsedBinaries > z.UsedBinaries {
		return GetZKCounterError("UsedBinaries")
	}
	if other.UsedSteps > z.UsedSteps {
		return GetZKCounterError("UsedSteps")
	}

	z.CumulativeGasUsed -= other.CumulativeGasUsed
	z.UsedKeccakHashes -= other.UsedKeccakHashes
	z.UsedPoseidonHashes -= other.UsedPoseidonHashes
	z.UsedPoseidonPaddings -= other.UsedPoseidonPaddings
	z.UsedMemAligns -= other.UsedMemAligns
	z.UsedArithmetics -= other.UsedArithmetics
	z.UsedBinaries -= other.UsedBinaries
	z.UsedSteps -= other.UsedSteps

	return nil
}
