package staker

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
	"math/big"
	"time"
)

type InboxReader struct {
	// Only in run thread
	bridge            ethbridge.BridgeWatcher
	db                core.ArbCore
	firstMessageBlock *big.Int

	// Only in main thread
	running    bool
	cancelFunc context.CancelFunc
}

func NewInboxReader(ctx context.Context, bridge ethbridge.BridgeWatcher, db core.ArbCore) (*InboxReader, error) {
	firstMessageBlock, err := bridge.LookupMessageBlock(ctx, big.NewInt(0))
	if err != nil {
		return nil, err
	}
	return &InboxReader{
		bridge:            bridge,
		db:                db,
		firstMessageBlock: firstMessageBlock.Height.AsInt(),
	}, nil
}

func (ir *InboxReader) Start(parentCtx context.Context) <-chan error {
	errChan := make(chan error, 1)
	ctx, cancelFunc := context.WithCancel(parentCtx)
	go func() {
		defer close(errChan)
		errChan <- ir.getMessages(ctx)
	}()
	ir.cancelFunc = cancelFunc
	ir.running = true
	return errChan
}

func (ir *InboxReader) Stop() {
	ir.cancelFunc()
	ir.running = false
}

func (ir *InboxReader) IsRunning() bool {
	return ir.running
}

func (ir *InboxReader) getMessages(ctx context.Context) error {
	from, err := ir.getNextBlockToRead()
	if err != nil {
		return err
	}
	reorging := false
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		for {
			currentHeight, err := ir.bridge.CurrentBlockHeight(ctx)
			if err != nil {
				return err
			}
			if from.Cmp(currentHeight) >= 0 {
				break
			}
			to := new(big.Int).Add(from, big.NewInt(10))
			if to.Cmp(currentHeight) > 0 {
				to = currentHeight
			}

			newMessages, err := ir.bridge.LookupMessagesInRange(ctx, from, to)
			if err != nil {
				return err
			}
			if len(newMessages) == 0 {
				if !reorging {
					from, err = ir.getPrevBlockForReorg(from)
					if err != nil {
						return err
					}
				} else {
					from = to
				}
			} else {
				needOlder, err := ir.addMessages(newMessages)
				if err != nil {
					return err
				}
				reorging = needOlder
				if needOlder {
					from, err = ir.getPrevBlockForReorg(from)
					if err != nil {
						return err
					}
				} else {
					from = to
				}
			}
		}
		<-time.After(time.Second * 2)
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
