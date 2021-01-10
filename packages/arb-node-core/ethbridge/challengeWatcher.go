package ethbridge

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"strings"
)

type ChallengeKind uint8

const (
	UNINITIALIZED ChallengeKind = iota
	INBOX_CONSISTENCY
	INBOX_DELTA
	EXECUTION
	STOPPED_SHORT
	NO_CHALLENGE
)

var bisectedID ethcommon.Hash

func init() {
	parsedChallenge, err := abi.JSON(strings.NewReader(ethbridgecontracts.ChallengeABI))
	if err != nil {
		panic(err)
	}
	bisectedID = parsedChallenge.Events["Bisected"].ID
}

type ChallengeTurn uint8

const (
	NONE ChallengeTurn = iota
	ASSERTER
	CHALLENGER
)

type ChallengeWatcher struct {
	con     *ethbridgecontracts.Challenge
	address ethcommon.Address
	client  ethutils.EthClient
}

func NewChallengeWatcher(address ethcommon.Address, client ethutils.EthClient) (*ChallengeWatcher, error) {
	con, err := ethbridgecontracts.NewChallenge(address, client)
	if err != nil {
		return nil, err
	}

	return &ChallengeWatcher{
		con:     con,
		address: address,
		client:  client,
	}, nil
}

func (c *ChallengeWatcher) ChallengedNodeNum(ctx context.Context) (NodeID, error) {
	return c.con.ChallengedNodeNum(&bind.CallOpts{Context: ctx})
}

func (c *ChallengeWatcher) Kind(ctx context.Context) (ChallengeKind, error) {
	rawKind, err := c.con.Kind(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}
	return ChallengeKind(rawKind), nil
}

func (c *ChallengeWatcher) Turn(ctx context.Context) (ChallengeTurn, error) {
	rawTurn, err := c.con.Turn(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}
	return ChallengeTurn(rawTurn), nil
}

func (c *ChallengeWatcher) Asserter(ctx context.Context) (common.Address, error) {
	asserter, err := c.con.Asserter(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}
	return common.NewAddressFromEth(asserter), nil
}

func (c *ChallengeWatcher) Challenger(ctx context.Context) (common.Address, error) {
	challenger, err := c.con.Challenger(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}
	return common.NewAddressFromEth(challenger), nil
}

func (c *ChallengeWatcher) CurrentResponder(ctx context.Context) (common.Address, error) {
	responder, err := c.con.CurrentResponder(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}
	return common.NewAddressFromEth(responder), nil
}

func (c *ChallengeWatcher) LookupNodes(ctx context.Context, challengeRoots []common.Hash) ([]*NodeInfo, error) {
	var query = ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []ethcommon.Address{c.address},
		Topics:    [][]ethcommon.Hash{{bisectedID}, common.NewEthHashesFromHashes(challengeRoots)},
	}
	logs, err := c.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	infos := make([]*NodeInfo, 0, len(logs))
	for _, ethLog := range logs {
		parsedLog, err := c.con.ParseBisected(ethLog)
		if err != nil {
			return nil, err
		}
		proposed := &common.BlockId{
			Height:     common.NewTimeBlocks(new(big.Int).SetUint64(ethLog.BlockNumber)),
			HeaderHash: common.NewHashFromEth(ethLog.BlockHash),
		}
		pa
		infos = append(infos, &NodeInfo{
			NodeNum:       parsedLog.NodeNum,
			BlockProposed: proposed,
			Assertion:     NewAssertionFromFields(parsedLog.AssertionBytes32Fields, parsedLog.AssertionIntFields),
			InboxMaxCount: parsedLog.InboxMaxCount,
		})
	}
	return infos, nil
}
