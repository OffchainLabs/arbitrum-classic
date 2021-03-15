package core

import (
	"context"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "core").Logger()

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
		var deletedLogs []value.Value
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			// Loop until new logs retrieved, may get deleted logs if reorg happened
			firstIndex, logs, deletedLogs, err = lr.cursor.LogsCursorGetLogs(lr.cursorIndex)
			if err != nil {
				return err
			}

			if len(logs) != 0 || len(deletedLogs) != 0 {
				// Retrieved logs successfully
				break
			}

			err = lr.cursor.LogsCursorCheckError(lr.cursorIndex)
			if err != nil {
				return err
			}

			// No new logs or errors so give some time for new logs to be added
			time.Sleep(lr.sleepTime)
		}

		logger.Debug().
			Uint64("cursorIndex", lr.cursorIndex.Uint64()).
			Str("firstIndex", bigIntAsString(firstIndex)).
			Int("log count", len(logs)).
			Int("deletedLog count", len(deletedLogs)).
			Msg("logs received from log cursor")

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

		for {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			status, err := lr.cursor.LogsCursorConfirmReceived(lr.cursorIndex)
			if err != nil {
				return err
			}
			if status {
				// Successfully confirmed receipt of logs
				logger.Debug().Uint64("cursorIndex", lr.cursorIndex.Uint64()).Msg("confirmed receipt of logs")
				break
			}

			// Reorg may have happened since previous call to GetLogs.
			// Post-retrieve reorg of logscursor will only include extra deleted logs, won't add any new logs
			_, _, newDeletedLogs, err := lr.cursor.LogsCursorGetLogs(lr.cursorIndex)
			if err != nil {
				return err
			}

			if len(newDeletedLogs) > 0 {
				// Got deleted logs successfully
				err = lr.consumer.DeleteLogs(newDeletedLogs)
				if err != nil {
					return err
				}
			}

			err = lr.cursor.LogsCursorCheckError(lr.cursorIndex)
			if err != nil {
				return err
			}
		}
	}
}
