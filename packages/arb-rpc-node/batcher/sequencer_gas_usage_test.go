package batcher

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestSequencerGasUsage(t *testing.T) {
	clnt, auths := test.SimulatedBackend(t)
	auth := auths[0]
	sequencer := common.NewAddressFromEth(auth.From)
	delayedInboxAddr, _, delayedBridge, err := ethbridgecontracts.DeployBridge(auth, clnt)
	test.FailIfError(t, err)

	_, err = delayedBridge.Initialize(auth)
	test.FailIfError(t, err)

	evBridgeAddr, _, evBridge, err := ethbridgetestcontracts.DeployRollupEventBridge(auth, clnt)
	test.FailIfError(t, err)
	_, _, seqInbox, err := ethbridgecontracts.DeploySequencerInbox(auth, clnt)
	test.FailIfError(t, err)
	clnt.Commit()
	_, err = evBridge.Initialize(auth, delayedInboxAddr, auth.From)
	test.FailIfError(t, err)
	clnt.Commit()

	_, err = delayedBridge.SetInbox(auth, evBridgeAddr, true)
	test.FailIfError(t, err)

	_, err = seqInbox.Initialize(auth, delayedInboxAddr, auth.From, auth.From)
	test.FailIfError(t, err)

	clnt.Commit()

	_, err = seqInbox.SetMaxDelay(auth, big.NewInt(150), big.NewInt(9000))
	test.FailIfError(t, err)

	_, err = evBridge.RollupInitialized(
		auth,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		ethcommon.Address{},
		auth.From,
		nil,
	)
	test.FailIfError(t, err)
	clnt.Commit()

	latestHeader, err := clnt.HeaderByNumber(context.Background(), nil)
	test.FailIfError(t, err)

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(latestHeader.Number),
		Timestamp: new(big.Int).SetUint64(latestHeader.Time),
	}

	delayedAcc, err := delayedBridge.InboxAccs(&bind.CallOpts{}, big.NewInt(0))
	test.FailIfError(t, err)

	initBatchItem := inbox.NewDelayedItem(big.NewInt(0), big.NewInt(1), common.Hash{}, big.NewInt(0), delayedAcc)
	test.FailIfError(t, err)
	endBlockMsg := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(1),
		big.NewInt(0),
		chainTime,
	)
	endBlockBatchItem := inbox.NewSequencerItem(big.NewInt(1), endBlockMsg, initBatchItem.Accumulator)

	delayedAccInt := new(big.Int).SetBytes(delayedAcc[:])
	metadata := []*big.Int{big.NewInt(0), chainTime.BlockNum.AsInt(), chainTime.Timestamp, big.NewInt(1), delayedAccInt}
	_, err = seqInbox.AddSequencerL2BatchFromOrigin(
		auth,
		nil,
		nil,
		metadata,
		endBlockBatchItem.Accumulator,
	)
	test.FailIfError(t, err)
	clnt.Commit()

	prevAcc := endBlockBatchItem.Accumulator

	seq := big.NewInt(2)
	for _, totalCount := range []int{1, 10, 100, 500} {
		for _, dataSizePerTx := range []int{0, 1, 10, 100, 1000, 10000} {
			l2Msg := message.L2Message{Data: common.RandBytes(dataSizePerTx)}
			var transactionsData []byte
			var lengths []*big.Int
			for i := 0; i < totalCount; i++ {
				var msg inbox.InboxMessage
				if dataSizePerTx > 0 {
					msg = message.NewInboxMessage(l2Msg, sequencer, seq, big.NewInt(0), chainTime)
				} else {
					msg = message.NewInboxMessage(
						message.EndBlockMessage{},
						sequencer,
						seq,
						big.NewInt(0),
						chainTime,
					)
				}

				batchItem := inbox.NewSequencerItem(big.NewInt(1), msg, prevAcc)
				transactionsData = append(transactionsData, l2Msg.Data...)
				lengths = append(lengths, big.NewInt(int64(len(l2Msg.Data))))
				prevAcc = batchItem.Accumulator
				seq = new(big.Int).Add(seq, big.NewInt(1))
			}

			metadata := []*big.Int{big.NewInt(int64(totalCount)), chainTime.BlockNum.AsInt(), chainTime.Timestamp, big.NewInt(1), big.NewInt(0)}
			tx, err := seqInbox.AddSequencerL2BatchFromOrigin(
				auth,
				transactionsData,
				lengths,
				metadata,
				prevAcc,
			)
			test.FailIfError(t, err)
			clnt.Commit()
			receipt, err := clnt.TransactionReceipt(context.Background(), tx.Hash())
			test.FailIfError(t, err)
			t.Logf("%v, %v, %v", totalCount, dataSizePerTx, receipt.GasUsed)
		}
	}

}
