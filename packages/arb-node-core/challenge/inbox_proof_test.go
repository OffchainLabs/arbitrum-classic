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
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func TestInboxProof(t *testing.T) {
	testDir, err := gotest.OpCodeTestDir()
	test.FailIfError(t, err)
	mexe := filepath.Join(testDir, "inbox.mexe")
	backend, auths := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}

	auth := auths[0]
	sequencer := auth.From
	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)

	_, _, osp1, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	test.FailIfError(t, err)
	delayedBridgeAddr, _, delayedBridge, err := ethbridgecontracts.DeployBridge(auth, client)
	test.FailIfError(t, err)
	sequencerAddr, _, sequencerCon, err := ethbridgecontracts.DeploySequencerInbox(auth, client)
	test.FailIfError(t, err)
	rollupAddr, _, rollup, err := ethbridgetestcontracts.DeployRollupMock(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	sequencerInboxWatcher, err := ethbridge.NewSequencerInboxWatcher(sequencerAddr, client)
	test.FailIfError(t, err)

	_, err = rollup.SetMock(auth, maxDelayBlocks, maxDelaySeconds)
	test.FailIfError(t, err)
	_, err = delayedBridge.Initialize(auth)
	test.FailIfError(t, err)
	_, err = sequencerCon.Initialize(auth, delayedBridgeAddr, sequencer, rollupAddr)
	test.FailIfError(t, err)
	client.Commit()

	_, err = delayedBridge.SetInbox(auth, auth.From, true)
	test.FailIfError(t, err)
	client.Commit()

	initMsg, err := message.NewInitMessage(protocol.NewRandomChainParams(), common.RandAddress(), nil)
	test.FailIfError(t, err)

	arbCore, cancel := monitor.PrepareArbCoreWithMexe(t, mexe)
	defer cancel()

	addDelayed := func(prevAcc common.Hash, msg message.Message, sender common.Address, msgNum int64) ([32]byte, inbox.DelayedMessage) {
		t.Helper()
		tx, err := delayedBridge.DeliverMessageToInbox(auth, uint8(msg.Type()), sender.ToEthAddress(), hashing.SoliditySHA3(msg.AsData()))
		test.FailIfError(t, err)
		client.Commit()
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		test.FailIfError(t, err)
		ev, err := delayedBridge.ParseMessageDelivered(*receipt.Logs[0])
		test.FailIfError(t, err)
		header, err := client.HeaderByHash(context.Background(), ev.Raw.BlockHash)
		test.FailIfError(t, err)
		delayedMsg := message.NewInboxMessage(msg, sender, ev.MessageIndex, gasPrice(tx, header.BaseFee), inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(int64(ev.Raw.BlockNumber)),
			Timestamp: new(big.Int).SetUint64(header.Time),
		})
		delayed := inbox.NewDelayedMessage(prevAcc, delayedMsg)
		delayedAcc, err := delayedBridge.InboxAccs(&bind.CallOpts{}, big.NewInt(msgNum))
		test.FailIfError(t, err)
		if delayedAcc != delayed.DelayedAccumulator {
			t.Fatal("bad delayed acc", delayedAcc, delayed.DelayedAccumulator)
		}
		return delayedAcc, delayed
	}

	delayedAcc1, delayed1 := addDelayed(common.Hash{}, initMsg, common.NewAddressFromEth(rollupAddr), 0)
	delayedAcc2, delayed2 := addDelayed(delayedAcc1, message.NewSafeL2Message(message.NewRandomTransaction()), common.RandAddress(), 1)

	delayedItem1 := inbox.NewDelayedItem(big.NewInt(0), big.NewInt(1), common.Hash{}, big.NewInt(0), delayedAcc1)

	latest, err := client.HeaderByNumber(context.Background(), nil)
	test.FailIfError(t, err)

	endOfBlockMessage1 := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(1),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(latest.Number),
			Timestamp: new(big.Int).SetUint64(latest.Time),
		},
	)
	endBlockBatchItem1 := inbox.NewSequencerItem(big.NewInt(1), endOfBlockMessage1, delayedItem1.Accumulator)

	seqMessage := message.NewInboxMessage(
		message.NewSafeL2Message(message.NewRandomTransaction()),
		common.NewAddressFromEth(sequencer),
		big.NewInt(2),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(latest.Number),
			Timestamp: new(big.Int).SetUint64(latest.Time),
		},
	)
	seqMsgItem := inbox.NewSequencerItem(big.NewInt(1), seqMessage, endBlockBatchItem1.Accumulator)

	delayedItem2 := inbox.NewDelayedItem(big.NewInt(3), big.NewInt(2), seqMsgItem.Accumulator, big.NewInt(1), delayedAcc2)

	endOfBlockMessage2 := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(4),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(latest.Number),
			Timestamp: new(big.Int).SetUint64(latest.Time),
		},
	)
	endBlockBatchItem2 := inbox.NewSequencerItem(big.NewInt(2), endOfBlockMessage2, delayedItem2.Accumulator)

	delayedAcc1Int := new(big.Int).SetBytes(delayedAcc1[:])
	delayedAcc2Int := new(big.Int).SetBytes(delayedAcc2[:])
	batch2Metadata := []*big.Int{
		big.NewInt(0), latest.Number, new(big.Int).SetUint64(latest.Time), big.NewInt(1), delayedAcc1Int,
		big.NewInt(1), latest.Number, new(big.Int).SetUint64(latest.Time), big.NewInt(2), delayedAcc2Int,
	}
	_, err = sequencerCon.AddSequencerL2BatchFromOrigin(
		auth,
		seqMessage.Data,
		[]*big.Int{big.NewInt(int64(len(seqMessage.Data)))},
		batch2Metadata,
		endBlockBatchItem2.Accumulator,
	)
	test.FailIfError(t, err)
	client.Commit()

	err = core.DeliverMessagesAndWait(arbCore.Core, big.NewInt(0), common.Hash{}, []inbox.SequencerBatchItem{delayedItem1, endBlockBatchItem1}, []inbox.DelayedMessage{delayed1}, nil)
	test.FailIfError(t, err)

	err = core.DeliverMessagesAndWait(
		arbCore.Core,
		big.NewInt(2),
		endBlockBatchItem1.Accumulator,
		[]inbox.SequencerBatchItem{seqMsgItem, delayedItem2, endBlockBatchItem2},
		[]inbox.DelayedMessage{delayed2},
		nil,
	)
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

	t.Log("Generated", len(cursors), "curors")

	for i := 0; i < len(cursors)-1; i++ {
		beforeCursor := cursors[i]
		afterCursor := cursors[i+1]

		mach, err := arbCore.Core.TakeMachine(beforeCursor.Clone())
		test.FailIfError(t, err)
		proof, bproof, err := mach.MarshalForProof()
		test.FailIfError(t, err)

		op := proof[0]
		if op != 0x72 {
			continue
		}

		t.Log("Proving inbox opcode")

		seqNum := beforeCursor.TotalMessagesRead()
		batch, err := LookupBatchContaining(context.Background(), arbCore.Core, sequencerInboxWatcher, seqNum)
		test.FailIfError(t, err)
		if batch == nil {
			t.Fatal("Failed to lookup batch containing message")
		}
		inboxProof, err := arbCore.Core.GenInboxProof(seqNum, batch.GetBatchIndex(), batch.GetAfterCount())
		test.FailIfError(t, err)
		proof = append(proof, inboxProof...)

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
