package main

import (
	"errors"
	"fmt"
	zerolog "github.com/rs/zerolog/log"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type App struct {
	inboxMessages []inbox.InboxMessage
	avmLogs       []value.Value
	avmSends      []value.Value

	snap      *snapshot.Snapshot
	assertion *protocol.ExecutionAssertion
}

func newApp(data []byte) (*App, error) {
	inboxMessages, avmLogs, avmSends, err := inbox.LoadTestVector(data)
	if err != nil {
		panic(err)
	}

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		return nil, err
	}

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(100000000000, inboxMessages, 0)

	chainId := message.ChainAddressToID(inboxMessages[0].Sender)

	snap := snapshot.NewSnapshot(mach, inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}, chainId, big.NewInt(int64(len(inboxMessages))+1))
	a := &App{
		inboxMessages: inboxMessages,
		avmLogs:       avmLogs,
		avmSends:      avmSends,
		snap:          snap,
		assertion:     assertion,
	}

	if err := a.verify(); err != nil {
		log.Println("Logs incorrect:", err)
	} else {
		log.Println("Logs verified")
	}
	return a, nil
}

func (a *App) verify() error {
	calcLogs := a.assertion.ParseLogs()
	calcSends := a.assertion.ParseOutMessages()

	commonLogCount := len(a.avmLogs)
	if len(calcLogs) < commonLogCount {
		commonLogCount = len(calcLogs)
	}

	commonSendCount := len(a.avmSends)
	if len(calcSends) < commonSendCount {
		commonSendCount = len(calcSends)
	}

	for i := 0; i < commonLogCount; i++ {
		calcRes, err := evm.NewResultFromValue(calcLogs[i])
		if err != nil {
			return err
		}
		res, err := evm.NewResultFromValue(a.avmLogs[i])
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
		if !value.Eq(calcSends[i], a.avmSends[i]) {
			return errors.New("wrong send")
		}
	}

	if len(calcLogs) != len(a.avmLogs) {
		return errors.New("wrong log count")
	}
	if len(calcSends) != len(a.avmSends) {
		return errors.New("wrong send count")
	}
	return nil
}

func (a *App) constructors() error {
	for _, avmLog := range a.avmLogs {
		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			return err
		}

		txRes, ok := res.(*evm.TxResult)
		if !ok {
			continue
		}

		req := txRes.IncomingRequest
		nested, err := message.NestedMessage(req.Data, req.Kind)
		if err != nil {
			return err
		}
		l2, ok := nested.(message.L2Message)
		if !ok {
			continue
		}
		msg, err := l2.AbstractMessage()
		if err != nil {
			return err
		}
		abstractTx, ok := msg.(message.AbstractTransaction)
		if !ok {
			continue
		}
		zeroAddr := common.Address{}
		if abstractTx.Destination() != zeroAddr {
			continue
		}

		if len(txRes.ReturnData) != 32 {
			return errors.New("unexpected constructor result length")
		}
		var contractAddress common.Address
		copy(contractAddress[:], txRes.ReturnData[12:])
		log.Println("Tx", txRes.IncomingRequest.MessageID, "created", contractAddress)
	}
	return nil
}

func (a *App) getCode(account common.Address) error {
	code, err := a.snap.GetCode(account)
	if err != nil {
		return err
	}
	log.Println(hexutil.Encode(code))
	return nil
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "code"},
		{Text: "constructors"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (a *App) Execute(line string) {
	fmt.Println("Got: ", line)
	line = strings.TrimSpace(line)

	blocks := strings.Split(line, " ")
	switch blocks[0] {
	case "code":
		if err := a.getCode(common.HexToAddress(blocks[1])); err != nil {
			log.Println("failed getting code", err)
		}
	case "constructors":
		if err := a.constructors(); err != nil {
			log.Println("failed getting constructors", err)
		}
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	zerolog.Logger = zerolog.With().Caller().Logger()
	file := os.Args[1]
	log.Println("Running test:", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	a, err := newApp(data)
	if err != nil {
		panic(err)
	}
	handleExit := prompt.OptionSetExitCheckerOnInput(func(line string, breakline bool) bool { return breakline && line == "exit" })
	p := prompt.New(a.Execute, completer, handleExit)

	p.Run()
}
