package e2e

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/db"
	"github.com/0xPolygonHermez/zkevm-node/event"
	"github.com/0xPolygonHermez/zkevm-node/event/nileventstorage"
	"github.com/0xPolygonHermez/zkevm-node/hex"
	"github.com/0xPolygonHermez/zkevm-node/merkletree"
	"github.com/0xPolygonHermez/zkevm-node/sequencer/broadcast"
	"github.com/0xPolygonHermez/zkevm-node/sequencer/broadcast/pb"
	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/0xPolygonHermez/zkevm-node/state/runtime/executor"
	"github.com/0xPolygonHermez/zkevm-node/test/constants"
	"github.com/0xPolygonHermez/zkevm-node/test/dbutils"
	"github.com/0xPolygonHermez/zkevm-node/test/operations"
	"github.com/0xPolygonHermez/zkevm-node/test/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	serverAddress     = "localhost:61090"
	totalBatches      = 2
	totalTxsLastBatch = 5
	forcedBatchNumber = 18
)

var (
	stateDBCfg      = dbutils.NewStateConfigFromEnv()
	ger             = common.HexToHash("deadbeef")
	mainnetExitRoot = common.HexToHash("caffe")
	rollupExitRoot  = common.HexToHash("bead")
)

func TestBroadcast(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	initOrResetDB()
	ctx := context.Background()

	require.NoError(t, operations.StartComponent("network"))
	require.NoError(t, operations.StartComponent("broadcast"))
	defer func() {
		require.NoError(t, operations.StopComponent("network"))
		require.NoError(t, operations.StopComponent("broadcast"))
	}()
	st, err := initState()
	require.NoError(t, err)

	require.NoError(t, populateDB(ctx, st))

	client, conn, cancel, err := broadcast.NewClient(ctx, serverAddress)
	require.NoError(t, err)
	defer func() {
		cancel()
		require.NoError(t, conn.Close())
	}()

	lastBatch, err := client.GetLastBatch(ctx, &emptypb.Empty{})
	require.NoError(t, err)
	require.Equal(t, totalBatches, int(lastBatch.BatchNumber))

	batch, err := client.GetBatch(ctx, &pb.GetBatchRequest{
		BatchNumber: uint64(totalBatches),
	})
	require.NoError(t, err)
	require.Equal(t, totalBatches, int(batch.BatchNumber))

	require.Equal(t, totalTxsLastBatch, len(batch.Transactions))
	require.EqualValues(t, forcedBatchNumber, batch.ForcedBatchNumber)

	require.Equal(t, mainnetExitRoot.String(), batch.MainnetExitRoot)
	require.Equal(t, rollupExitRoot.String(), batch.RollupExitRoot)
}

func initState() (*state.State, error) {
	ctx := context.Background()
	initOrResetDB()
	sqlDB, err := db.NewSQLDB(stateDBCfg)
	if err != nil {
		return nil, err
	}
	stateDb := state.NewPostgresStorage(sqlDB)
	executorUri := testutils.GetEnv(constants.ENV_ZKPROVER_URI, "localhost:50071")
	merkleTreeUri := testutils.GetEnv(constants.ENV_MERKLETREE_URI, "localhost:50061")
	executorClient, _, _ := executor.NewExecutorClient(ctx, executor.Config{URI: executorUri})
	mtDBClient, _, _ := merkletree.NewMTDBServiceClient(ctx, merkletree.Config{URI: merkleTreeUri})
	stateTree := merkletree.NewStateTree(mtDBClient)

	eventStorage, err := nileventstorage.NewNilEventStorage()
	if err != nil {
		return nil, err
	}
	eventLog := event.NewEventLog(event.Config{}, eventStorage)

	return state.NewState(state.Config{}, stateDb, executorClient, stateTree, eventLog), nil
}

func populateDB(ctx context.Context, st *state.State) error {
	const blockNumber = 1

	var parentHash common.Hash
	var l2Block types.Block

	const addBlock = "INSERT INTO state.block (block_num, received_at, block_hash) VALUES ($1, $2, $3)"
	if _, err := st.PostgresStorage.Exec(ctx, addBlock, blockNumber, time.Now(), ""); err != nil {
		return err
	}

	const addForcedBatch = "INSERT INTO state.forced_batch (forced_batch_num, global_exit_root, raw_txs_data, coinbase, timestamp, block_num) VALUES ($1, $2, $3, $4, $5, $6)"
	if _, err := st.PostgresStorage.Exec(ctx, addForcedBatch, forcedBatchNumber, ger.String(), "", common.HexToAddress("").String(), time.Now(), blockNumber); err != nil {
		return err
	}

	const addBatch = "INSERT INTO state.batch (batch_num, global_exit_root, timestamp, coinbase, local_exit_root, state_root, forced_batch_num) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err := st.PostgresStorage.Exec(ctx, addBatch, 1, ger.String(), time.Now(), common.HexToAddress("").String(), common.Hash{}.String(), common.Hash{}.String(), nil); err != nil {
		return err
	}
	if _, err := st.PostgresStorage.Exec(ctx, addBatch, 2, ger.String(), time.Now(), common.HexToAddress("").String(), common.Hash{}.String(), common.Hash{}.String(), forcedBatchNumber); err != nil {
		return err
	}

	for i := 1; i <= totalTxsLastBatch; i++ {
		if i == 1 {
			parentHash = state.ZeroHash
		} else {
			parentHash = l2Block.Hash()
		}

		// Store L2 Genesis Block
		header := new(types.Header)
		header.Number = new(big.Int).SetUint64(uint64(i - 1))
		header.ParentHash = parentHash
		l2Block := types.NewBlockWithHeader(header)
		l2Block.ReceivedAt = time.Now()

		if err := st.PostgresStorage.AddL2Block(ctx, totalBatches, l2Block, []*types.Receipt{}, nil); err != nil {
			return err
		}

		tx := types.NewTransaction(uint64(i), common.HexToAddress("0x1"), big.NewInt(0), uint64(0), nil, nil)
		bData, _ := tx.MarshalBinary()
		encoded := hex.EncodeToHex(bData)
		const addTransaction = "INSERT INTO state.transaction (hash, encoded, l2_block_num) VALUES ($1, $2, $3)"
		if _, err := st.PostgresStorage.Exec(ctx, addTransaction, tx.Hash().String(), encoded, l2Block.Number().Uint64()); err != nil {
			return err
		}
	}

	const addExitRoots = "INSERT INTO state.exit_root (block_num, timestamp, global_exit_root, mainnet_exit_root, rollup_exit_root) VALUES ($1, $2, $3, $4, $5)"
	_, err := st.PostgresStorage.Exec(ctx, addExitRoots, blockNumber, time.Now(), ger, mainnetExitRoot, rollupExitRoot)
	return err
}

func initOrResetDB() {
	if err := dbutils.InitOrResetState(stateDBCfg); err != nil {
		panic(err)
	}
}
