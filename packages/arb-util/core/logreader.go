package core

import (
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

var logger = log.With().Caller().Str("component", "logreader").Logger()

type LogReader struct {
	consumer    LogConsumer
	cursor      LogsCursor
	cursorIndex *big.Int
	maxCount    *big.Int

	// Only in main thread
	running    bool
	cancelFunc context.CancelFunc
}

func NewLogReader(consumer LogConsumer, cursor LogsCursor, cursorIndex *big.Int, maxCount *big.Int) (*LogReader, error) {
	return &LogReader{
		consumer:    consumer,
		cursor:      cursor,
		cursorIndex: cursorIndex,
		maxCount:    maxCount,
	}, nil
}

func (lr *LogReader) Start(parentCtx context.Context) <-chan error {
	errChan := make(chan error, 1)
	ctx, cancelFunc := context.WithCancel(parentCtx)
	go func() {
		defer close(errChan)
		errChan <- lr.getLogs(ctx)
	}()
	lr.cancelFunc = cancelFunc
	lr.running = true
	return errChan
}

func (lr *LogReader) Stop() {
	lr.cancelFunc()
	lr.running = false
}

func (lr *LogReader) IsRunning() bool {
	return lr.running
}

func (lr *LogReader) getLogs(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		err := lr.cursor.LogsCursorRequest(lr.cursorIndex, lr.maxCount)
		if err != nil {
			return err
		}

		var firstIndex *big.Int
		var logs []value.Value
		var firstDeletedIndex *big.Int
		var deletedLogs []value.Value
		for {
			// Loop until new logs retrieved, may get deleted logs if reorg happened
			// Cannot retrieve new logs until deleted logs have been retrieved
			firstIndex, logs, err = lr.cursor.LogsCursorGetLogs(lr.cursorIndex)
			if err != nil {
				return err
			}
			if logs != nil || deletedLogs != nil {
				// Retrieved logs successfully
				break
			}

			// No new logs yet, check if deleted logs
			firstDeletedIndex, deletedLogs, err = lr.cursor.LogsCursorGetDeletedLogs(lr.cursorIndex)
			if err != nil {
				return err
			}
			if deletedLogs != nil {
				// Got deleted logs successfully, retry loop to get any new logs without waiting
				continue
			}

			// No new logs or deleted logs so give some time for new logs to be added
			time.Sleep(1 * time.Second)
		}

		currentLogCount, err := lr.consumer.CurrentLogCount()
		if err != nil {
			return err
		}

		currentLogIndex := new(big.Int).Sub(currentLogCount, big.NewInt(1))

		if len(deletedLogs) > 0 && firstDeletedIndex.Cmp(currentLogIndex) <= 0 {
			// Existing logs to delete
			deletedCount := new(big.Int).Sub(currentLogCount, firstDeletedIndex)
			if deletedCount.Cmp(big.NewInt(int64(len(deletedLogs)))) != 0 {
				logger.Warn().
					Uint64("currentLogCount", currentLogCount.Uint64()).
					Uint64("firstDeletedIndex", firstDeletedIndex.Uint64()).
					Int("deletedLogs count", len(deletedLogs)).
					Msg("more deleted logs sent than we previously received")
			}
			if err = lr.consumer.DeleteLogs(deletedLogs[:deletedCount.Uint64()]); err != nil {
				return err
			}

			currentLogCount = firstDeletedIndex
			err = lr.consumer.UpdateCurrentLogCount(currentLogCount)
			if err != nil {
				return err
			}
		}

		if len(logs) > 0 {
			cmp := firstIndex.Cmp(currentLogCount)
			if cmp == 1 {
				// Some logs were skipped, pull these manually
				// TODO
				//skippedCount := new(big.Int).Sub(firstIndex, currentLogIndex)
				//result, err := lr.cursor
			}

			if err = lr.consumer.AddLogs(logs); err != nil {
				return err
			}
		}

		for {
			status, err := lr.cursor.LogsCursorConfirmReceived(lr.cursorIndex)
			if err != nil {
				return err
			}
			if status {
				// Successfully confirmed receipt of logs
				break
			}

			// Reorg happened since previous call to GetLogs.  Post-retrieve reorg of logscursor will only include
			// extra deleted logs, won't add any new logs
			firstNewDeletedIndex, newDeletedLogs, err := lr.cursor.LogsCursorGetDeletedLogs(lr.cursorIndex)
			if err != nil {
				return err
			}
			if newDeletedLogs == nil {
				return errors.New("missing expected deleted logs")
			}

			// Got deleted logs successfully
			if len(newDeletedLogs) > 0 {
				err = lr.consumer.DeleteLogs(newDeletedLogs)
				if err != nil {
					return err
				}
			} else {
				// Logs not provided to delete
				// TODO
				_ = firstNewDeletedIndex
			}
		}
	}
}
