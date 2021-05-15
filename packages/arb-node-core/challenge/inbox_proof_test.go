package challenge

import (
	"context"
	"math/big"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func TestInboxProof(t *testing.T) {
	testDir, err := gotest.OpCodeTestDir()
	test.FailIfError(t, err)
	mexe := filepath.Join(testDir, "inbox.mexe")
	backend, pks := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}

	auth, err := bind.NewKeyedTransactorWithChainID(pks[0], big.NewInt(1337))
	test.FailIfError(t, err)
	sequencer := auth.From
	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)
	rollup := common.RandAddress()

	_, _, osp1, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	test.FailIfError(t, err)
	delayedBridgeAddr, _, delayedBridge, err := ethbridgecontracts.DeployBridge(auth, client)
	test.FailIfError(t, err)
	sequencerAddr, _, sequencerCon, err := ethbridgecontracts.DeploySequencerInbox(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.Initialize(auth)
	test.FailIfError(t, err)
	_, err = sequencerCon.Initialize(auth, delayedBridgeAddr, sequencer, maxDelayBlocks, maxDelaySeconds)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.SetInbox(auth, auth.From, true)
	test.FailIfError(t, err)
	client.Commit()

	initMsg, err := message.NewInitMessage(protocol.NewRandomChainParams(), common.RandAddress(), nil)
	test.FailIfError(t, err)

	chainTime := inbox.NewRandomChainTime()

	arbCore, cancel := monitor.PrepareArbCoreWithMexe(t, mexe)
	defer cancel()

	delayedMsg := message.NewInboxMessage(initMsg, rollup, big.NewInt(0), big.NewInt(0), chainTime)
	tx, err := delayedBridge.DeliverMessageToInbox(auth, uint8(delayedMsg.Kind), delayedMsg.Sender.ToEthAddress(), hashing.SoliditySHA3(delayedMsg.Data))
	test.FailIfError(t, err)
	client.Commit()
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)
	ev, err := delayedBridge.ParseMessageDelivered(*receipt.Logs[0])
	test.FailIfError(t, err)
	header, err := client.HeaderByHash(context.Background(), ev.Raw.BlockHash)
	test.FailIfError(t, err)
	delayedMsg.ChainTime = inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(int64(ev.Raw.BlockNumber)),
		Timestamp: new(big.Int).SetUint64(header.Time),
	}
	delayedMsg.GasPrice = tx.GasPrice()
	delayed := inbox.NewDelayedMessage(common.Hash{}, delayedMsg)
	delayedAcc, err := delayedBridge.InboxAccs(&bind.CallOpts{}, big.NewInt(0))
	test.FailIfError(t, err)
	if delayedAcc != delayed.DelayedAccumulator {
		t.Fatal("bad delayed acc", delayedAcc, delayed.DelayedAccumulator)
	}

	delayedItem := inbox.NewDelayedItem(big.NewInt(0), big.NewInt(1), common.Hash{}, big.NewInt(0), delayedAcc)
	latest, err := client.HeaderByNumber(context.Background(), nil)
	test.FailIfError(t, err)

	endOfBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(1),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(latest.Number),
			Timestamp: new(big.Int).SetUint64(latest.Time),
		},
	)
	endBlockBatchItem := inbox.NewSequencerItem(big.NewInt(1), endOfBlockMessage, delayedItem.Accumulator)

	_, err = sequencerCon.AddSequencerL2BatchFromOrigin(
		auth,
		nil,
		nil,
		latest.Number,
		new(big.Int).SetUint64(latest.Time),
		big.NewInt(1),
		endBlockBatchItem.Accumulator,
	)
	test.FailIfError(t, err)
	client.Commit()

	err = core.DeliverMessagesAndWait(arbCore.Core, common.Hash{}, []inbox.SequencerBatchItem{delayedItem, endBlockBatchItem}, []inbox.DelayedMessage{delayed}, nil)
	test.FailIfError(t, err)

	var cursors []core.ExecutionCursor
	cursor, err := arbCore.Core.GetExecutionCursor(big.NewInt(0))
	test.FailIfError(t, err)
	cursors = append(cursors, cursor.Clone())
	for {
		err = arbCore.Core.AdvanceExecutionCursor(cursor, big.NewInt(1), true)
		test.FailIfError(t, err)
		if cursor.TotalGasConsumed().Cmp(cursors[len(cursors)-1].TotalGasConsumed()) == 0 {
			break
		}
		cursors = append(cursors, cursor.Clone())
	}

	lastCursor := cursors[len(cursors)-1]
	lastMach, err := arbCore.Core.TakeMachine(lastCursor.Clone())
	test.FailIfError(t, err)
	t.Log(lastMach)

	t.Log("Generated", len(cursors), "curors")

	for i := 0; i < len(cursors)-1; i++ {
		beforeCursor := cursors[i]
		afterCursor := cursors[i+1]

		mach, err := arbCore.Core.TakeMachine(beforeCursor.Clone())
		test.FailIfError(t, err)
		proof, bproof, err := mach.MarshalForProof()
		test.FailIfError(t, err)

		op := proof[0]

		sequencerInboxWatcher, err := ethbridge.NewSequencerInboxWatcher(sequencerAddr, client)
		test.FailIfError(t, err)
		if op == 0x72 {
			// INBOX proving
			seqNum := beforeCursor.TotalMessagesRead()
			batch, err := LookupBatchContaining(context.Background(), arbCore.Core, sequencerInboxWatcher, seqNum)
			test.FailIfError(t, err)
			if batch == nil {
				t.Fatal("Failed to lookup batch containing message")
			}
			inboxProof, err := arbCore.Core.GenInboxProof(seqNum, batch.GetBatchIndex(), batch.GetAfterCount())
			test.FailIfError(t, err)
			proof = append(proof, inboxProof...)
		}

		t.Log("Op", op)
		ret, err := osp1.ExecuteStep(
			&bind.CallOpts{},
			[2]ethcommon.Address{sequencerAddr, delayedBridgeAddr},
			beforeCursor.TotalMessagesRead(),
			[2][32]byte{
				beforeCursor.SendAcc(),
				beforeCursor.LogAcc(),
			},
			proof,
			bproof,
		)
		if err != nil {
			t.Error(err)
			continue
		}

		beforeMachineHash := ret.Fields[0]
		afterMachineHash := ret.Fields[1]
		if beforeMachineHash != beforeCursor.MachineHash() {
			t.Error("wrong before machine hash")
		}
		if afterMachineHash != afterCursor.MachineHash() {
			t.Error("wrong after machine hash")
		}
	}
}
