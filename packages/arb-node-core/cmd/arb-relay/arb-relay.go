package arbrelay

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/rs/zerolog/log"
)

type ArbRelay struct {
	ArbitrumBroadcasterWebsocketUrl string
	broadcastClient                 *broadcastclient.BroadcastClient
	broadcaster                     *broadcaster.Broadcaster
}

var logger = log.With().Caller().Str("component", "broadcaster").Logger()

func NewArbRelay(websocketUrl string, rebroadcastSettings broadcaster.Settings) *ArbRelay {
	ar := &ArbRelay{}
	ar.ArbitrumBroadcasterWebsocketUrl = websocketUrl

	ar.broadcaster = broadcaster.NewBroadcaster(rebroadcastSettings)

	return ar
}

func (ar *ArbRelay) Start() {
	ar.broadcastClient = broadcastclient.NewBroadcastClient(ar.ArbitrumBroadcasterWebsocketUrl, nil)

	err := ar.broadcaster.Start()
	if err != nil {
		logger.Error().Err(err).Msg("broadcasted unable to start")
	}

	// connect returns
	messages, err := ar.broadcastClient.Connect()
	if err != nil {
		logger.Error().Err(err).Msg("broadcast client unable to connect")
	}

	accList := make(chan common.Hash)

	ar.broadcastClient.SetConfirmedAccumulatorListner(accList)

	go func() {
		for {
			select {
			case msg := <-messages:
				ar.broadcaster.Broadcast(msg.FeedItem.PrevAcc, msg.FeedItem.BatchItem, msg.Signature)
			case ca := <-accList:
				ar.broadcaster.ConfirmedAccumulator(ca)
			}
		}
	}()
}

func (ar *ArbRelay) Stop() {
	ar.broadcastClient.Close()
	ar.broadcaster.Stop()
}
