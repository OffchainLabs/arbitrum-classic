package machineobserver

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func setupRollup(ctx context.Context, client ethutils.EthClient, auth *bind.TransactOpts) (common.Address, common.Address, error) {
	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		StakeToken:              common.Address{},
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	factoryAddr, err := ethbridge.DeployRollupFactory(auth, client)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	arbClient := ethbridge.NewEthAuthClient(client, auth)

	factory, err := arbClient.NewArbFactory(common.NewAddressFromEth(factoryAddr))
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	rollupAddress, _, err := factory.CreateRollup(
		ctx,
		mach.Hash(),
		config,
		common.Address{},
	)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	inboxAddress, err := factory.GlobalInboxAddress()
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	return rollupAddress, inboxAddress, err
}

// TestReorg launches an Observer against a aggressively reorging chain in order
// to test it's ability to handle reorgs
func TestReorg(t *testing.T) {
	clnt, pks := test.SimulatedBackend()
	l1Client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}
	go func() {
		t := time.NewTicker(time.Millisecond * 10)
		for range t.C {
			l1Client.Commit()
		}
	}()

	common.SetDurationPerBlock(time.Millisecond * 10)

	rollupAddress, inboxAddress, err := setupRollup(context.Background(), l1Client, bind.NewKeyedTransactor(pks[0]))
	if err != nil {
		t.Fatal(err)
	}

	dbPath := "dbPath"

	if err := os.RemoveAll(dbPath); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dbPath); err != nil {
			t.Fatal(err)
		}
	}()

	auth := bind.NewKeyedTransactor(pks[0])

	authClient := ethbridge.NewEthAuthClient(l1Client, auth)
	inboxConn, err := authClient.NewGlobalInbox(inboxAddress, rollupAddress)
	if err != nil {
		t.Fatal(err)
	}

	dest := common.RandAddress()

	if err := inboxConn.DepositEthMessage(context.Background(), common.NewAddressFromEth(auth.From), big.NewInt(100000)); err != nil {
		t.Fatal(err)
	}

	errChan := make(chan error, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			<-time.After(time.Second)
			tx := types.NewTransaction(uint64(i), dest.ToEthAddress(), big.NewInt(1), 100000000, big.NewInt(0), nil)
			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(message.ChainAddressToID(rollupAddress)), pks[0])
			if err != nil {
				t.Fatal(err)
			}
			msg, err := message.NewL2Message(message.SignedTransaction{Tx: signedTx})
			if err != nil {
				errChan <- err
				return
			}

			ev, err := inboxConn.SendL2Message(context.Background(), msg.AsData())
			if err != nil {
				errChan <- err
				return
			}

			t.Log("Sent tx", ev.Message.InboxSeqNum)
		}
	}()

	_, err = RunObserver(
		context.Background(),
		rollupAddress,
		arbbridge.NewStressTestClient(ethbridge.NewEthClient(l1Client), time.Second),
		arbos.Path(),
		dbPath,
	)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-time.After(time.Second * 20):
		break
	case err := <-errChan:
		t.Fatal(err)
	}
}
