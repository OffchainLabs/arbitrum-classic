package challenge

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
)

func TestExecutionChallenge(t *testing.T) {
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

	for {
		if arbCore.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	challengedNode := initializeChallengeData(t, arbCore, big.NewInt(200000))

	faultyCore := newFaultyCore(arbCore, faultConfig{
		distortMachineAtGas: big.NewInt(100000),
	})

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
