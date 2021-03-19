package staker

import (
	"context"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
)

type InboxReader struct {
	// Only in run thread
	bridge            *ethbridge.BridgeWatcher
	db                core.ArbCore
	firstMessageBlock *big.Int
	caughtUp          bool
	caughtUpChan      chan bool
	caughtUpTarget    *big.Int

	// Only in main thread
	running    bool
	cancelFunc context.CancelFunc
	completed  chan bool
}

func NewInboxReader(ctx context.Context, bridge *ethbridge.BridgeWatcher, db core.ArbCore) (*InboxReader, error) {
	firstMessageBlock, err := bridge.LookupMessageBlock(ctx, big.NewInt(0))
	if err != nil {
		return nil, err
	}
	return &InboxReader{
		bridge:            bridge,
		db:                db,
		firstMessageBlock: firstMessageBlock.Height.AsInt(),
		completed:         make(chan bool, 1),
		caughtUpChan:      make(chan bool, 1),
	}, nil
}

func (ir *InboxReader) Start(parentCtx context.Context) {
	ctx, cancelFunc := context.WithCancel(parentCtx)
	go func() {
		defer func() {
			ir.completed <- true
		}()
		for {
			err := ir.getMessages(ctx)
			if err == nil {
				break
			}
			logger.Warn().Stack().Err(err).Msg("Failed to read inbox messages")
			<-time.After(time.Second * 2)
		}
	}()
	ir.cancelFunc = cancelFunc
	ir.running = true
}

func (ir *InboxReader) Stop() {
	ir.cancelFunc()
	<-ir.completed
	ir.running = false
}

func (ir *InboxReader) IsRunning() bool {
	return ir.running
}

// May only be called once
func (ir *InboxReader) WaitToCatchUp() {
	<-ir.caughtUpChan
}

func (ir *InboxReader) getMessages(ctx context.Context) error {
	from, err := ir.getNextBlockToRead()
	if err != nil {
		return err
	}
	reorging := false
	blocksToFetch := uint64(100)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		currentHeight, err := ir.bridge.CurrentBlockHeight(ctx)
		if err != nil {
			return err
		}
		for {
			if !ir.caughtUp && ir.caughtUpTarget != nil {
				arbCorePosition := ir.db.MachineMessagesRead()
				if arbCorePosition.Cmp(ir.caughtUpTarget) >= 0 {
					ir.caughtUp = true
					ir.caughtUpChan <- true
				}
			}
			if from.Cmp(currentHeight) >= 0 {
				break
			}
			to := new(big.Int).Add(from, new(big.Int).SetUint64(blocksToFetch))
			if to.Cmp(currentHeight) > 0 {
				to = currentHeight
			}
			newMessages, err := ir.bridge.LookupMessagesInRange(ctx, from, to)
			if err != nil {
				return err
			}
			if ir.caughtUpTarget == nil && to.Cmp(currentHeight) == 0 {
				if len(newMessages) > 0 {
					ir.caughtUpTarget = newMessages[len(newMessages)-1].Message.InboxSeqNum
				} else {
					dbMessageCount, err := ir.db.GetMessageCount()
					if err != nil {
						return err
					}
					ir.caughtUpTarget = dbMessageCount
				}
			}
			if len(newMessages) < 40 {
				blocksToFetch += 20
			} else if len(newMessages) > 90 {
				blocksToFetch /= 2
			}
			if blocksToFetch == 0 {
				blocksToFetch++
			}
			logger.Debug().
				Str("from", from.String()).
				Str("to", to.String()).
				Int("count", len(newMessages)).
				Msg("Looking up messages")
			if len(newMessages) == 0 {
				if reorging {
					from, err = ir.getPrevBlockForReorg(from)
					if err != nil {
						return err
					}
				} else {
					from = to
				}
			} else {
				success, err := ir.addMessages(newMessages)
				if err != nil {
					return err
				}
				reorging = !success
				if !success {
					from, err = ir.getPrevBlockForReorg(from)
					if err != nil {
						return err
					}
				} else {
					from = to
				}
			}
		}
		<-time.After(time.Second * 1)
	}
}

func (ir *InboxReader) getNextBlockToRead() (*big.Int, error) {
	messageCount, err := ir.db.GetMessageCount()
	if err != nil {
		return nil, err
	}
	if messageCount.Cmp(big.NewInt(0)) == 0 {
		return ir.firstMessageBlock, nil
	}
	msg, err := core.GetSingleMessage(ir.db, new(big.Int).Sub(messageCount, big.NewInt(1)))
	if err != nil {
		return nil, err
	}
	return msg.ChainTime.BlockNum.AsInt(), nil
}

func (ir *InboxReader) getPrevBlockForReorg(from *big.Int) (*big.Int, error) {
	if from.Cmp(ir.firstMessageBlock) == 0 {
		return nil, errors.New("can't get older messages")
	}
	newFrom := new(big.Int).Sub(from, big.NewInt(10))
	if newFrom.Cmp(ir.firstMessageBlock) < 0 {
		newFrom = ir.firstMessageBlock
	}
	return newFrom, nil
}

func (ir *InboxReader) addMessages(newMessages []*ethbridge.DeliveredInboxMessage) (bool, error) {
	if len(newMessages) == 0 {
		return false, errors.New("must have messages to add")
	}

	messages := make([]inbox.InboxMessage, 0, len(newMessages))
	for _, msg := range newMessages {
		messages = append(messages, msg.Message)
	}
	return core.DeliverMessagesAndWait(
		ir.db,
		messages,
		newMessages[0].BeforeInboxAcc,
		true,
	)
}
