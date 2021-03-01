package challenge

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func PrepareTestArbCore(t *testing.T, messages []inbox.InboxMessage) (core.ArbCore, func()) {
	tmpDir, err := ioutil.TempDir("", "arbitrum")
	test.FailIfError(t, err)
	storage, err := cmachine.NewArbStorage(tmpDir)
	if err != nil {
		os.RemoveAll(tmpDir)
	}
	test.FailIfError(t, err)
	shutdown := func() {
		storage.CloseArbStorage()
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}
	returning := false
	defer (func() {
		if !returning {
			shutdown()
		}
	})()

	err = storage.Initialize(arbos.Path())
	test.FailIfError(t, err)

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		t.Fatal("failed to start thread")
	}

	if len(messages) > 0 {
		_, err = core.DeliverMessagesAndWait(arbCore, messages, common.Hash{}, false)
		test.FailIfError(t, err)
	}
	for {
		if arbCore.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	returning = true
	return arbCore, shutdown
}

func MakeTestInitMsg() inbox.InboxMessage {
	owner := common.RandAddress()
	chain := common.RandAddress()
	return message.NewInboxMessage(
		message.Init{
			ChainParams: protocol.ChainParams{
				StakeRequirement:        big.NewInt(0),
				StakeToken:              common.Address{},
				GracePeriod:             common.NewTimeBlocks(big.NewInt(3)),
				MaxExecutionSteps:       0,
				ArbGasSpeedLimitPerTick: 0,
			},
			Owner:       owner,
			ExtraConfig: []byte{},
		},
		chain,
		big.NewInt(0),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	)
}

func distortHash(hash common.Hash) common.Hash {
	hash[0] = 0xde
	hash[1] = 0xad
	hash[2] = 0xbe
	hash[3] = 0xef
	return hash
}

type FaultConfig struct {
	DistortMachineAtGas *big.Int
	MessagesReadCap     *big.Int
	PhantomMessageAtGas *big.Int
	StallMachineAt      *big.Int
}

type FaultyExecutionCursor struct {
	config     FaultConfig
	phantomGas *big.Int
	core.ExecutionCursor
}

func (e FaultyExecutionCursor) Clone() core.ExecutionCursor {
	return FaultyExecutionCursor{
		config:          e.config,
		phantomGas:      new(big.Int).Set(e.phantomGas),
		ExecutionCursor: e.ExecutionCursor.Clone(),
	}
}

func (e FaultyExecutionCursor) MachineHash() common.Hash {
	hash := e.ExecutionCursor.MachineHash()
	if e.config.DistortMachineAtGas != nil && e.ExecutionCursor.TotalGasConsumed().Cmp(e.config.DistortMachineAtGas) >= 0 {
		hash = distortHash(hash)
	}
	return hash
}

func (e FaultyExecutionCursor) TotalMessagesRead() *big.Int {
	messages := e.ExecutionCursor.TotalMessagesRead()
	if e.config.PhantomMessageAtGas != nil && e.ExecutionCursor.TotalGasConsumed().Cmp(e.config.PhantomMessageAtGas) > 0 {
		messages = new(big.Int).Add(messages, big.NewInt(1))
	}
	if e.config.MessagesReadCap != nil && messages.Cmp(e.config.MessagesReadCap) > 0 {
		messages = new(big.Int).Set(e.config.MessagesReadCap)
	}
	return messages
}

func (e FaultyExecutionCursor) TotalGasConsumed() *big.Int {
	gas := e.ExecutionCursor.TotalGasConsumed()
	return new(big.Int).Add(gas, e.phantomGas)
}

type FaultyCore struct {
	config FaultConfig
	core.ArbCore
}

func NewFaultyCore(core core.ArbCore, config FaultConfig) FaultyCore {
	return FaultyCore{
		config:  config,
		ArbCore: core,
	}
}

func (c FaultyCore) GetExecutionCursor(totalGasUsed *big.Int) (core.ExecutionCursor, error) {
	cursor, err := c.ArbCore.GetExecutionCursor(totalGasUsed)
	if err != nil {
		return nil, err
	}
	return FaultyExecutionCursor{
		config:          c.config,
		phantomGas:      big.NewInt(0),
		ExecutionCursor: cursor,
	}, nil
}

func (c FaultyCore) AdvanceExecutionCursor(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool) error {
	faultyCursor := executionCursor.(FaultyExecutionCursor)
	targetGas := new(big.Int).Add(executionCursor.TotalGasConsumed(), maxGas)
	if c.config.StallMachineAt != nil && targetGas.Cmp(c.config.StallMachineAt) > 0 {
		phantomGas := new(big.Int).Sub(targetGas, c.config.StallMachineAt)
		maxGas = new(big.Int).Sub(targetGas, phantomGas)
		faultyCursor.phantomGas.Set(phantomGas)
		if maxGas.Cmp(big.NewInt(0)) <= 0 {
			return nil
		}
	}
	return c.ArbCore.AdvanceExecutionCursor(faultyCursor.ExecutionCursor, maxGas, goOverGas)
}
