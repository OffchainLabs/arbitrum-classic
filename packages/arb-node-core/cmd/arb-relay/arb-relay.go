package arbrelay

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
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

	_ = messages
	/*
		go func() {
			for {
				select {
				case receivedMsgs := <-messages:
					for i := range receivedMsgs.Messages {
						m := receivedMsgs.Messages[i]
						ar.broadcaster.Broadcast(m.BeforeAccumulator, m.InboxMessage, m.Signature)
					}

				}
			}
		}()
	*/
}

func (ar *ArbRelay) Stop() {
	ar.broadcastClient.Close()
	ar.broadcaster.Stop()
}
