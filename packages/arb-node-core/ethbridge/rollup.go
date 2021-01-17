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

func (r *Rollup) buildTx(name string, amount *big.Int, args ...interface{}) (*RawTransaction, error) {
	data, err := rollupABI.Pack(name, args...)
	return &RawTransaction{
		Data:   data,
		Dest:   r.address,
		Amount: amount,
	}, err
}

func (r *Rollup) buildSimpleTx(name string, args ...interface{}) (*RawTransaction, error) {
	return r.buildTx(name, big.NewInt(0), args...)
}

func (r *Rollup) RejectNextNode(node *big.Int, staker common.Address) (*RawTransaction, error) {
	return r.buildSimpleTx("rejectNextNode", node, staker.ToEthAddress())
}

func (r *Rollup) ConfirmNextNode(logAcc common.Hash, sends [][]byte) (*RawTransaction, error) {
	var sendsData []byte
	sendLengths := make([]*big.Int, 0, len(sends))
	for _, msg := range sends {
		sendsData = append(sendsData, msg...)
		sendLengths = append(sendLengths, new(big.Int).SetInt64(int64(len(msg))))
	}
	return r.buildSimpleTx("confirmNextNode", logAcc, sendsData, sendLengths)
}

func (r *Rollup) NewStake(ctx context.Context, amount *big.Int) (*RawTransaction, error) {
	tokenType, err := r.StakeToken(ctx)
	if err != nil {
		return nil, err
	}
	emptyAddress := common.Address{}
	if tokenType != emptyAddress {
		return r.buildSimpleTx("newStake", amount)
	} else {
		return r.buildTx("newStake", amount, big.NewInt(0))
	}
}

func (r *Rollup) StakeOnExistingNode(block *common.BlockId, node core.NodeID) (*RawTransaction, error) {
	return r.buildSimpleTx("stakeOnExistingNode", block.HeaderHash.ToEthHash(), block.Height.AsInt(), node)
}

func (r *Rollup) StakeOnNewNode(
	block *common.BlockId,
	node core.NodeID,
	assertion *core.Assertion,
) (*RawTransaction, error) {
	return r.buildSimpleTx(
		"stakeOnNewNode",
		block.HeaderHash.ToEthHash(),
		block.Height.AsInt(),
		node,
		assertion.BytesFields(),
		assertion.IntFields(),
	)
}

func (r *Rollup) ReturnOldDeposit(staker common.Address) (*RawTransaction, error) {
	return r.buildSimpleTx("returnOldDeposit", staker.ToEthAddress())
}

func (r *Rollup) AddToDeposit(ctx context.Context, address common.Address, amount *big.Int) (*RawTransaction, error) {
	tokenType, err := r.StakeToken(ctx)
	if err != nil {
		return nil, err
	}
	emptyAddress := common.Address{}
	if tokenType != emptyAddress {
		return r.buildTx("addToDeposit", big.NewInt(0), address, amount)
	} else {
		return r.buildTx("addToDeposit", amount, address, big.NewInt(0))
	}
}

func (r *Rollup) ReduceDeposit(amount *big.Int, destination common.Address) (*RawTransaction, error) {
	return r.buildSimpleTx("reduceDeposit", amount, destination.ToEthAddress())
}

func (r *Rollup) CreateChallenge(
	staker1 common.Address,
	node1 core.NodeID,
	staker2 common.Address,
	node2 core.NodeID,
	assertion *core.Assertion,
	inboxMaxHash common.Hash,
	inboxMaxCount *big.Int,
) (*RawTransaction, error) {
	return r.buildSimpleTx(
		"createChallenge",
		[2]ethcommon.Address{staker1.ToEthAddress(), staker2.ToEthAddress()},
		[2]*big.Int{node1, node2},
		[3][32]byte{
			assertion.InboxConsistencyHash(inboxMaxHash, inboxMaxCount),
			assertion.InboxDeltaHash(),
			assertion.ExecutionHash(),
		},
		assertion.GasUsed(),
	)
}

func (r *Rollup) RemoveZombie(zombieNum *big.Int, maxNodes *big.Int) (*RawTransaction, error) {
	return r.buildSimpleTx("removeZombie", zombieNum, maxNodes)
}

func (r *Rollup) RemoveOldZombies(startIndex *big.Int) (*RawTransaction, error) {
	return r.buildSimpleTx("removeOldZombies", startIndex)
}
