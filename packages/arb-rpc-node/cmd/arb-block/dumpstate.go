package main

import (
	"os"
	"path"

	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

func dumpArbState(arbcore core.ArbCore, height uint64, dirname string) error {
	cursor, err := arbcore.GetExecutionCursorAtEndOfBlock(height, true)
	if err != nil {
		return err
	}
	logger.Info().Uint64("block", height).Msg("taking machine")
	machine, err := arbcore.TakeMachine(cursor)
	if err != nil {
		return err
	}
	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	return arbcore.DumpAccounts(machine, path.Join(dirname, "accounts.json"))
}
