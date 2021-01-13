package ethbridge

import (
	"context"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type RawTransaction struct {
	Data   []byte
	Dest   ethcommon.Address
	Amount *big.Int
}

type Rollup struct {
	*RollupWatcher
}

func NewRollup(address ethcommon.Address, client ethutils.EthClient) (*Rollup, error) {
	watcher, err := NewRollupWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		RollupWatcher: watcher,
	}, nil
}

func (r *Rollup) buildTx(data []byte, amount *big.Int) *RawTransaction {
	return &RawTransaction{
		Data:   data,
		Dest:   r.address,
		Amount: amount,
	}
}

func (r *Rollup) buildSimpleTx(data []byte) *RawTransaction {
	return r.buildTx(data, big.NewInt(0))
}

func (r *Rollup) RejectNextNode(node *big.Int, staker common.Address) (*RawTransaction, error) {
	data, err := rollupABI.Pack("rejectNextNode", node, staker.ToEthAddress())
	return r.buildSimpleTx(data), err
}

func (r *Rollup) ConfirmNextNode(logAcc common.Hash, sends [][]byte) (*RawTransaction, error) {
	var sendsData []byte
	sendLengths := make([]*big.Int, 0, len(sends))
	for _, msg := range sends {
		sendsData = append(sendsData, msg...)
		sendLengths = append(sendLengths, new(big.Int).SetInt64(int64(len(msg))))
	}
	data, err := rollupABI.Pack("confirmNextNode", logAcc, sendsData, sendLengths)
	return r.buildSimpleTx(data), err
}

func (r *Rollup) NewStake(amount *big.Int) (*RawTransaction, error) {
	data, err := rollupABI.Pack("newStake")
	return r.buildTx(data, amount), err
}

func (r *Rollup) StakeOnExistingNode(block *common.BlockId, node core.NodeID) (*RawTransaction, error) {
	data, err := rollupABI.Pack("stakeOnExistingNode", block.HeaderHash.ToEthHash(), block.Height.AsInt(), node)
	return r.buildSimpleTx(data), err
}

func (r *Rollup) StakeOnNewNode(
	block *common.BlockId,
	node core.NodeID,
	assertion *core.Assertion,
) (*RawTransaction, error) {
	data, err := rollupABI.Pack(
		"stakeOnNewNode",
		block.HeaderHash.ToEthHash(),
		block.Height.AsInt(),
		node,
		assertion.BytesFields(),
		assertion.IntFields(),
	)
	return r.buildSimpleTx(data), err
}

func (r *Rollup) ReturnOldDeposit(staker common.Address) (*RawTransaction, error) {
	data, err := rollupABI.Pack("returnOldDeposit", staker.ToEthAddress())
	return r.buildSimpleTx(data), err
}

func (r *Rollup) AddToDeposit(address common.Address, amount *big.Int) (*RawTransaction, error) {
	data, err := rollupABI.Pack("addToDeposit", address.ToEthAddress())
	return r.buildTx(data, amount), err
}

func (r *Rollup) ReduceDeposit(amount *big.Int) (*RawTransaction, error) {
	data, err := rollupABI.Pack("reduceDeposit", amount)
	return r.buildSimpleTx(data), err
}

func (r *Rollup) CreateChallenge(
	ctx context.Context,
	staker1 common.Address,
	node1 core.NodeID,
	staker2 common.Address,
	node2 core.NodeID,
	assertion *core.Assertion,
	inboxMaxHash common.Hash,
	inboxMaxCount *big.Int,
) (*RawTransaction, error) {
	speedLimit, err := r.ArbGasSpeedLimitPerBlock(ctx)
	if err != nil {
		return nil, err
	}
	data, err := rollupABI.Pack(
		"createChallenge",
		[2]ethcommon.Address{staker1.ToEthAddress(), staker2.ToEthAddress()},
		[2]*big.Int{node1, node2},
		[3][32]byte{
			assertion.InboxConsistencyHash(inboxMaxHash, inboxMaxCount),
			assertion.InboxDeltaHash(),
			assertion.ExecutionHash(),
		},
		assertion.CheckTime(speedLimit),
	)
	return r.buildSimpleTx(data), err
}

func (r *Rollup) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*RawTransaction, error) {
	data, err := rollupABI.Pack("removeOldZombies", zombieNum, maxNodes)
	return r.buildSimpleTx(data), err
}

func (r *Rollup) RemoveOldZombies(startIndex *big.Int) (*RawTransaction, error) {
	data, err := rollupABI.Pack("removeOldZombies", startIndex)
	return r.buildSimpleTx(data), err
}
