package core

import (
	"context"
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
	sleepTime   time.Duration

	// Only in main thread
	running    bool
	cancelFunc context.CancelFunc
	completed  chan bool
}

func NewLogReader(consumer LogConsumer, cursor LogsCursor, cursorIndex *big.Int, maxCount *big.Int, sleepTime time.Duration) *LogReader {
	return &LogReader{
		consumer:    consumer,
		cursor:      cursor,
		cursorIndex: cursorIndex,
		maxCount:    maxCount,
		sleepTime:   sleepTime,
		completed:   make(chan bool, 1),
	}
}

func (lr *LogReader) Start(parentCtx context.Context) <-chan error {
	errChan := make(chan error, 1)
	ctx, cancelFunc := context.WithCancel(parentCtx)
	go func() {
		defer close(errChan)
		errChan <- lr.getLogs(ctx)
		lr.completed <- true
	}()
	lr.cancelFunc = cancelFunc
	lr.running = true
	return errChan
}

func (lr *LogReader) Stop() {
	lr.cancelFunc()
	<-lr.completed
	lr.running = false
}

func (lr *LogReader) IsRunning() bool {
	return lr.running
}

func bigIntAsString(val *big.Int) string {
	if val == nil {
		return "nil"
	}
	return val.String()
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
			if len(logs) > 0 || len(deletedLogs) > 0 {
				// Retrieved logs successfully
				break
			}

			// No new logs yet, check if deleted logs
			firstDeletedIndex, deletedLogs, err = lr.cursor.LogsCursorGetDeletedLogs(lr.cursorIndex)
			if err != nil {
				return err
			}
			if len(deletedLogs) > 0 {
				// Got deleted logs successfully, retry loop to get any new logs without waiting
				continue
			}

			// No new logs or deleted logs so give some time for new logs to be added
			time.Sleep(lr.sleepTime)
		}

		if len(logs) > 0 || len(deletedLogs) > 0 {
			logger.Info().
				Str("firstDeletedIndex", bigIntAsString(firstDeletedIndex)).
				Int("deletedLog count", len(deletedLogs)).
				Str("firstIndex", bigIntAsString(firstIndex)).
				Int("log count", len(logs)).
				Msg("logs received from log cursor")
		}

		if len(deletedLogs) > 0 {
			// Existing logs to delete
			if err = lr.consumer.DeleteLogs(deletedLogs); err != nil {
				return err
			}
		}

		if len(logs) > 0 {
			if err = lr.consumer.AddLogs(firstIndex, logs); err != nil {
				return err
			}
		}

		if len(logs) > 0 || len(deletedLogs) > 0 {
			logger.Info().Uint64("cursorIndex", lr.cursorIndex.Uint64()).Msg("confirming receipt of logs")
			for {
				status, err := lr.cursor.LogsCursorConfirmReceived(lr.cursorIndex)
				if err != nil {
					return err
				}
				if status {
					// Successfully confirmed receipt of logs
					logger.Info().Uint64("cursorIndex", lr.cursorIndex.Uint64()).Msg("confirmed receipt of logs")
					break
				}

				// Reorg happened since previous call to GetLogs.  Post-retrieve reorg of logscursor will only include
				// extra deleted logs, won't add any new logs
				_, newDeletedLogs, err := lr.cursor.LogsCursorGetDeletedLogs(lr.cursorIndex)
				if err != nil {
					return err
				}

				// Got deleted logs successfully
				if len(newDeletedLogs) > 0 {
					err = lr.consumer.DeleteLogs(newDeletedLogs)
					if err != nil {
						return err
					}
				}
			}
		}
	}
}
