package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"io/ioutil"
	golog "log"
	"math/big"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"

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

var logger zerolog.Logger

type App struct {
	inboxMessages []inbox.InboxMessage
	avmLogs       []value.Value
	avmSends      [][]byte

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
	assertion, _, _ := mach.ExecuteAssertion(100000000000, true, inboxMessages, true)

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
		logger.Warn().Stack().Err(err).Msg("logs incorrect")
	} else {
		logger.Info().Msg("logs verified")
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
			logger.Warn().Str("calculated", calcRes.AsValue().String()).Str("correct", res.AsValue().String()).Msg("wrong log")
		}
	}

	for i := 0; i < commonSendCount; i++ {
		if !bytes.Equal(calcSends[i], a.avmSends[i]) {
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
		logger.Info().Hex("Tx", txRes.IncomingRequest.MessageID.Bytes()).Hex("address", contractAddress.Bytes()).Msg("contract created")
	}
	return nil
}

func (a *App) getCode(account common.Address) error {
	code, err := a.snap.GetCode(account)
	if err != nil {
		return err
	}
	logger.Info().Hex("code", code)
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
			logger.Warn().Stack().Err(err).Msg("failed getting code")
		}
	case "constructors":
		if err := a.constructors(); err != nil {
			logger.Warn().Stack().Err(err).Msg("failed getting constructors")
		}
	}
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-replay-test").Logger()

	file := os.Args[1]
	logger.Info().Str("file", file).Msg("Running test")
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
