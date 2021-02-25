package challenge

import (
	"io/ioutil"
	"math"
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

func runExecutionTest(t *testing.T, messages []inbox.InboxMessage, gas *big.Int, faultConfig faultConfig) {
	tmpDir, err := ioutil.TempDir("", "arbitrum")
	test.FailIfError(t, err)
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()
	storage, err := cmachine.NewArbStorage(tmpDir)
	test.FailIfError(t, err)
	defer storage.CloseArbStorage()

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

	challengedNode := initializeChallengeData(t, arbCore, gas)

	faultyCore := newFaultyCore(arbCore, faultConfig)

	time := big.NewInt(100)
	executeChallenge(
		t,
		challengedNode,
		time,
		time,
		arbCore,
		faultyCore,
	)
}

func TestChallengeToOSP(t *testing.T) {
	runExecutionTest(t, []inbox.InboxMessage{}, big.NewInt(200000), faultConfig{distortMachineAtGas: big.NewInt(100000)})
}

func TestChallengeToOSPWithMessage(t *testing.T) {
	owner := common.RandAddress()
	chain := common.RandAddress()
	initMsg := message.NewInboxMessage(
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
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	)
	runExecutionTest(t, []inbox.InboxMessage{initMsg}, big.NewInt(math.MaxInt64), faultConfig{distortMachineAtGas: big.NewInt(1500000)})
}
