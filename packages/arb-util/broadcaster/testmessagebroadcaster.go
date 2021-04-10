package broadcaster

import (
	"fmt"
	"sync"
	"time"
)

type TestMessageBroadcaster struct {
	clientManager            *ClientManager
	startWorkerMutex         *sync.Mutex
	messageBroadcasterWorker *time.Ticker
	intervalDuration         time.Duration
	workerStarted            bool
}

func NewTestMessageBroadcaster() *TestMessageBroadcaster {
	tmb := &TestMessageBroadcaster{}
	tmb.startWorkerMutex = &sync.Mutex{}
	tmb.intervalDuration = time.Duration(2) * time.Second
	tmb.workerStarted = false
	return tmb
}

func (tmb *TestMessageBroadcaster) setClientManager(clientManager *ClientManager) {
	tmb.clientManager = clientManager
}

func (tmb *TestMessageBroadcaster) startWorker() {
	tmb.startWorkerMutex.Lock()
	defer tmb.startWorkerMutex.Unlock()
	if tmb.workerStarted {
		return
	}

	ticker := time.NewTicker(tmb.intervalDuration)
	go func() {
		for t := range ticker.C {
			err := tmb.clientManager.Broadcast("new message", Object{
				"message": fmt.Sprintf("message thing: %v", t),
			})
			if err != nil {
				logger.Warn().Err(err).Msg("error sending broadcast")
			}
		}
	}()

	tmb.messageBroadcasterWorker = ticker
	tmb.workerStarted = true
}

func (tmb *TestMessageBroadcaster) stopWorker() {
	if tmb.messageBroadcasterWorker != nil {
		tmb.messageBroadcasterWorker.Stop()
		tmb.workerStarted = false
	}
}
