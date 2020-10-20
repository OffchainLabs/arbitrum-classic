package main

import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	file := os.Args[1]
	log.Println("Running test:", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	if err := checkTest(data); err != nil {
		log.Println("Test failed:", err)
	} else {
		log.Println("Test passed")
	}
}

func checkTest(data []byte) error {
	inboxMessages, avmLogs, avmSends, err := inbox.LoadTestVector(data)
	if err != nil {
		return err
	}

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		return err
	}

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(100000000000, inboxMessages, 0)

	calcLogs := assertion.ParseLogs()
	calcSends := assertion.ParseOutMessages()

	commonLogCount := len(avmLogs)
	if len(calcLogs) < commonLogCount {
		commonLogCount = len(calcLogs)
	}

	commonSendCount := len(avmSends)
	if len(calcSends) < commonSendCount {
		commonSendCount = len(calcSends)
	}

	for i := 0; i < commonLogCount; i++ {
		calcRes, err := evm.NewResultFromValue(calcLogs[i])
		if err != nil {
			return err
		}
		res, err := evm.NewResultFromValue(avmLogs[i])
		if err != nil {
			return err
		}
		if !value.Eq(calcRes.AsValue(), res.AsValue()) {
			log.Println("Calculated:", calcRes)
			log.Println("Correct:", res)
			return errors.New("wrong log")
		}
	}

	for i := 0; i < commonSendCount; i++ {
		if !value.Eq(calcSends[i], avmSends[i]) {
			return errors.New("wrong send")
		}
	}

	if len(calcLogs) != len(avmLogs) {
		return errors.New("wrong log count")
	}
	if len(calcSends) != len(avmSends) {
		return errors.New("wrong send count")
	}
	return nil
}
