package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var nodeCreatedID ethcommon.Hash

func init() {
	parsedRollup, err := abi.JSON(strings.NewReader(ethbridgecontracts.RollupABI))
	if err != nil {
		panic(err)
	}
	nodeCreatedID = parsedRollup.Events["NodeCreated"].ID
}

type StakerInfo struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge *common.Address
}

type NodeInfo struct {
	NodeNum   NodeID
	Assertion *Assertion
}

type Assertion struct {
	BeforeProposedBlock *big.Int
	BeforeStepsRun      *big.Int
	BeforeMachineHash   common.Hash
	BeforeInboxHash     common.Hash
	BeforeInboxCount    *big.Int
	BeforeSendCount     *big.Int
	BeforeLogCount      *big.Int
	BeforeInboxMaxCount *big.Int
	StepsExecuted       *big.Int
	InboxDelta          common.Hash
	InboxMessagesRead   *big.Int
	GasUsed             *big.Int
	SendAcc             common.Hash
	SendCount           *big.Int
	LogAcc              common.Hash
	LogCount            *big.Int
	AfterInboxHash      common.Hash
	AfterMachineHash    common.Hash
}

func NewAssertionFromFields(a [7][32]byte, b [11]*big.Int) *Assertion {
	return &Assertion{
		BeforeProposedBlock: b[0],
		BeforeStepsRun:      b[1],
		BeforeMachineHash:   a[0],
		BeforeInboxHash:     a[1],
		BeforeInboxCount:    b[2],
		BeforeSendCount:     b[3],
		BeforeLogCount:      b[4],
		BeforeInboxMaxCount: b[5],
		StepsExecuted:       b[6],
		InboxDelta:          a[2],
		InboxMessagesRead:   b[7],
		GasUsed:             b[8],
		SendAcc:             a[3],
		SendCount:           b[9],
		LogAcc:              a[4],
		LogCount:            b[10],
		AfterInboxHash:      a[5],
		AfterMachineHash:    a[6],
	}
}

func (a *Assertion) BytesFields() [7][32]byte {
	return [7][32]byte{
		a.BeforeMachineHash,
		a.BeforeInboxHash,
		a.InboxDelta,
		a.SendAcc,
		a.LogAcc,
		a.AfterInboxHash,
		a.AfterMachineHash,
	}
}

func (a *Assertion) IntFields() [11]*big.Int {
	return [11]*big.Int{
		a.BeforeProposedBlock,
		a.BeforeStepsRun,
		a.BeforeInboxCount,
		a.BeforeSendCount,
		a.BeforeLogCount,
		a.BeforeInboxMaxCount,
		a.StepsExecuted,
		a.InboxMessagesRead,
		a.GasUsed,
		a.SendCount,
		a.LogCount,
	}
}

func bisectionChunkHash(
	length *big.Int,
	segmentEnd *big.Int,
	startHash common.Hash,
	endHash common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(length),
		hashing.Uint256(segmentEnd),
		hashing.Bytes32(startHash),
		hashing.Bytes32(endHash),
	)
}

func inboxDeltaHash(inboxAcc, deltaAcc common.Hash) common.Hash {
	return hashing.SoliditySHA3(hashing.Bytes32(inboxAcc), hashing.Bytes32(deltaAcc))
}

func assertionHash(
	inboxDelta common.Hash,
	gasUsed *big.Int,
	outputAcc common.Hash,
	machineState common.Hash,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(inboxDelta),
		hashing.Uint256(gasUsed),
		hashing.Bytes32(outputAcc),
		hashing.Bytes32(machineState),
	)
}

func outputAccHash(
	sendAcc common.Hash,
	sendCount *big.Int,
	logAcc common.Hash,
	logCount *big.Int,
) common.Hash {
	return hashing.SoliditySHA3(
		hashing.Bytes32(sendAcc),
		hashing.Uint256(sendCount),
		hashing.Bytes32(logAcc),
		hashing.Uint256(logCount),
	)
}

func (a *Assertion) InboxConsistencyHash(inboxTopHash common.Hash, inboxTopCount *big.Int) common.Hash {
	messagesAfterCount := new(big.Int).Sub(inboxTopCount, a.BeforeInboxCount)
	messagesAfterCount = messagesAfterCount.Sub(messagesAfterCount, a.InboxMessagesRead)
	return bisectionChunkHash(messagesAfterCount, messagesAfterCount, inboxTopHash, a.AfterInboxHash)
}

func (a *Assertion) InboxDeltaHash() common.Hash {
	return bisectionChunkHash(
		a.InboxMessagesRead,
		a.InboxMessagesRead,
		inboxDeltaHash(a.AfterInboxHash, common.Hash{}),
		inboxDeltaHash(a.BeforeInboxHash, a.InboxDelta),
	)
}

func (a *Assertion) ExecutionHash() common.Hash {
	return bisectionChunkHash(
		a.GasUsed,
		a.GasUsed,
		assertionHash(
			a.InboxDelta,
			big.NewInt(0),
			outputAccHash(common.Hash{}, big.NewInt(0), common.Hash{}, big.NewInt(0)),
			a.BeforeMachineHash,
		),
		assertionHash(
			common.Hash{},
			a.GasUsed,
			outputAccHash(
				a.SendAcc,
				a.SendCount,
				a.LogAcc,
				a.LogCount,
			),
			a.AfterMachineHash,
		),
	)
}

func (a *Assertion) CheckTime(arbGasSpeedLimitPerBlock *big.Int) *big.Int {
	return new(big.Int).Div(a.GasUsed, arbGasSpeedLimitPerBlock)
}

type NodeID *big.Int

type RollupWatcher struct {
	con     *ethbridgecontracts.Rollup
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewRollupWatcher(address ethcommon.Address, client ethutils.EthClient) (*RollupWatcher, error) {
	con, err := ethbridgecontracts.NewRollup(address, client)
	if err != nil {
		return nil, err
	}

	return &RollupWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (r *RollupWatcher) LookupNodes(ctx context.Context, nodes []*big.Int) ([]*NodeInfo, error) {
	var nodeQuery []ethcommon.Hash
	for _, node := range nodes {
		var nd ethcommon.Hash
		copy(nd[:], math.U256Bytes(node))
		nodeQuery = append(nodeQuery, nd)
	}
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{r.address},
		Topics:    [][]ethcommon.Hash{{nodeCreatedID}, nodeQuery},
	}
	logs, err := r.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	infos := make([]*NodeInfo, 0, len(logs))
	for _, ethLog := range logs {
		parsedLog, err := r.con.ParseNodeCreated(ethLog)
		if err != nil {
			return nil, err
		}
		infos = append(infos, &NodeInfo{
			NodeNum:   parsedLog.NodeNum,
			Assertion: NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
		})
	}
	return infos, nil
}

func (r *RollupWatcher) StakerInfo(ctx context.Context, staker common.Address) (*StakerInfo, error) {
	info, err := r.con.StakerMap(&bind.CallOpts{Context: ctx}, staker.ToEthAddress())
	if err != nil {
		return nil, err
	}
	if !info.IsStaked {
		return nil, nil
	}
	stakerInfo := &StakerInfo{
		Index:            info.Index,
		LatestStakedNode: info.LatestStakedNode,
		AmountStaked:     info.AmountStaked,
	}
	emptyAddress := ethcommon.Address{}
	if info.CurrentChallenge != emptyAddress {
		chal := common.NewAddressFromEth(info.CurrentChallenge)
		stakerInfo.CurrentChallenge = &chal
	}
	return stakerInfo, nil
}

func (r *RollupWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	return r.con.StakerCount(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) CurrentRequiredStake(ctx context.Context) (*big.Int, error) {
	return r.con.CurrentRequiredStake(&bind.CallOpts{Context: ctx})
}

func (r *Rollup) LatestConfirmedNode(ctx context.Context) (*big.Int, error) {
	return r.con.LatestConfirmed(&bind.CallOpts{Context: ctx})
}

func (r *Rollup) FirstUnresolvedNode(ctx context.Context) (*big.Int, error) {
	return r.con.FirstUnresolvedNode(&bind.CallOpts{Context: ctx})
}

func (r *Rollup) ChallengePeriodBlocks(ctx context.Context) (*big.Int, error) {
	return r.con.ChallengePeriodBlocks(&bind.CallOpts{Context: ctx})
}

func (r *RollupWatcher) GetNode(ctx context.Context, node NodeID) (*NodeWatcher, error) {
	nodeAddress, err := r.con.Nodes(&bind.CallOpts{Context: ctx}, node)
	if err != nil {
		return nil, err
	}
	return NewNodeWatcher(nodeAddress, r.client)
}

type Rollup struct {
	*RollupWatcher
	auth *TransactAuth
}

func NewRollup(address ethcommon.Address, client ethutils.EthClient, auth *TransactAuth) (*Rollup, error) {
	watcher, err := NewRollupWatcher(address, client)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		RollupWatcher: watcher,
		auth:          auth,
	}, nil
}

func (r *Rollup) RejectNextNode(ctx context.Context, node NodeID, staker common.Address) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RejectNextNode(auth, node, staker.ToEthAddress())
	})
}

func (r *Rollup) RejectNextNodeOutOfOrder(ctx context.Context) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RejectNextNodeOutOfOrder(auth)
	})
}

func (r *Rollup) ConfirmNextNode(
	ctx context.Context,
	logAcc common.Hash,
	sends [][]byte,
) (*types.Transaction, error) {
	var sendsData []byte
	sendLengths := make([]*big.Int, 0, len(sends))
	for _, msg := range sends {
		sendsData = append(sendsData, msg...)
		sendLengths = append(sendLengths, new(big.Int).SetInt64(int64(len(msg))))
	}
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ConfirmNextNode(auth, logAcc, sendsData, sendLengths)
	})
}

func (r *Rollup) NewStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.NewStakeOnExistingNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
		)
	})
}

func (r *Rollup) AddStakeOnExistingNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.AddStakeOnExistingNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
		)
	})
}

func (r *Rollup) NewStakeOnNewNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
	prev NodeID,
	assertion *Assertion,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.NewStakeOnNewNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
			prev,
			assertion.BytesFields(),
			assertion.IntFields(),
		)
	})
}

func (r *Rollup) AddStakeOnNewNode(
	ctx context.Context,
	block *common.BlockId,
	node NodeID,
	assertion *Assertion,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.AddStakeOnNewNode(
			auth,
			block.HeaderHash.ToEthHash(),
			block.Height.AsInt(),
			node,
			assertion.BytesFields(),
			assertion.IntFields(),
		)
	})
}

func (r *Rollup) ReturnOldDeposit(ctx context.Context, staker common.Address) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ReturnOldDeposit(auth, staker.ToEthAddress())
	})
}

func (r *Rollup) AddToDeposit(ctx context.Context, address common.Address, amount *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		auth.Value = amount
		return r.con.AddToDeposit(auth, address.ToEthAddress())
	})
}

func (r *Rollup) ReduceDeposit(ctx context.Context, amount *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.ReduceDeposit(auth, amount)
	})
}

func (r *Rollup) CreateChallenge(
	ctx context.Context,
	staker1 common.Address,
	node1 NodeID,
	staker2 common.Address,
	node2 NodeID,
	assertion *Assertion,
	inboxMaxHash common.Hash,
	inboxMaxCount *big.Int,
	arbGasSpeedLimitPerBlock *big.Int,
) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.CreateChallenge(
			auth,
			staker1.ToEthAddress(),
			node1,
			staker2.ToEthAddress(),
			node2,
			assertion.InboxConsistencyHash(inboxMaxHash, inboxMaxCount),
			assertion.InboxDeltaHash(),
			assertion.ExecutionHash(),
			assertion.CheckTime(arbGasSpeedLimitPerBlock),
		)
	})
}

func (r *Rollup) RemoveZombie(ctx context.Context, zombieNum *big.Int, maxNodes *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RemoveZombie(auth, zombieNum, maxNodes)
	})
}

func (r *Rollup) RemoveOldZombies(ctx context.Context, startIndex *big.Int) (*types.Transaction, error) {
	return r.auth.makeTx(ctx, func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.con.RemoveOldZombies(auth, startIndex)
	})
}
