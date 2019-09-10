package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)

func StartEventListeners(fibonacci *Fibonacci) chan interface{} {
	ch := make(chan interface{}, 64)

	startGenerateFibListener(fibonacci, ch)
	//startSubmitYourDeckHashListener(poker, ch)
	//startEncryptAndProveCommandListener(poker, ch)
	//startDealFaceDownCardListener(poker, ch)
	//startDealFaceUpCardListener(poker, ch)
	//startCallForBetActionListener(poker, ch)
	//startPresentBestHandListener(poker, ch)
	//startDeclareWinningsListener(poker, ch)
	//startReportZKPListener(poker, ch)

	return ch
}

type ListenerError struct {
	ListenerName string
	Err          error
}

func startGenerateFibListener(fibonacci *Fibonacci, ch chan interface{}) {
	go func() {
		evCh := make(chan *FibonacciTestEvent, 2)
		start := uint64(0)
		watch := &bind.WatchOpts{
			Context: context.Background(),
			Start:   &start,
		}
		sub, err := fibonacci.WatchTestEvent(watch, evCh)
		if err != nil {
			log.Println(err)
			panic("listener startup error")
		}
		defer sub.Unsubscribe()
		errChan := sub.Err()
		defer log.Println("FibonacciTestEvent listener exiting")
		fmt.Println("test event listener")
		for {
			select {
			case ev, ok := <-evCh:
				fmt.Println("test event received")
				if ok {
					ch <- ev
				} else {
					return
				}
			case err, ok := <-errChan:
				fmt.Println("test event error received")
				if ok {
					log.Println("FibonacciTestEvent error:", err)
					ch <- &ListenerError{"FibonacciTestEvent listener error", err}
				} else {
					return
				}
			}
		}
	}()
}
